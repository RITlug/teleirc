package telegram

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
Handler specifies a function that handles a Telegram update.
In this case, we take a Telegram client and update object,
where the specific Handler will "handle" the given event.
*/
type Handler = func(tg *Client, u tgbotapi.Update)

/*
updateHandler takes in a Telegram Update channel, and determines
which handler to fire off
*/
func updateHandler(tg *Client, updates tgbotapi.UpdatesChannel) {
	for u := range updates {
		// Don't process any messages that didn't come from the
		// chat we're bridging
		if u.Message.Chat.ID != tg.Settings.ChatID {
			tg.logger.LogDebug("Ignored message from a telegram chat we're not bridging:", tg.Settings.ChatID)
			continue
		}

		switch {
		case u.Message == nil:
			tg.logger.LogError("Missing message data")
			continue
		case u.Message.NewChatMembers != nil:
			tg.logger.LogDebug("joinHandler triggered")
			joinHandler(tg, u.Message.NewChatMembers)
		case u.Message.LeftChatMember != nil:
			tg.logger.LogDebug("partHandler triggered")
			partHandler(tg, u.Message.LeftChatMember)
		case u.Message.Text != "":
			tg.logger.LogDebug("messageHandler triggered")
			messageHandler(tg, u)
		case u.Message.Sticker != nil:
			tg.logger.LogDebug("stickerHandler triggered")
			stickerHandler(tg, u)
		case u.Message.Document != nil:
			tg.logger.LogDebug("documentHandler triggered")
			documentHandler(tg, u.Message)
		case u.Message.Photo != nil:
			tg.logger.LogDebug("photoHandler triggered")
			photoHandler(tg, u)
		case u.Message.Location != nil:
			tg.logger.LogDebug("locationHandler triggered")
			locationHandler(tg, u.Message)
		default:
			tg.logger.LogWarning("Triggered, but message type is currently unsupported")
			tg.logger.LogWarning("Unhandled Update:", u)
			continue
		}
	}
}

/*
messageHandler handles the Message Telegram Object, which formats the
Telegram update into a simple string for IRC.
*/
func messageHandler(tg *Client, u tgbotapi.Update) {
	username := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.From, tg.Settings.PreferName)
	formatted := ""

	if tg.IRCSettings.NoForwardPrefix != "" && strings.HasPrefix(u.Message.Text, tg.IRCSettings.NoForwardPrefix) {
		return
	}

	// Telegram user replied to a message
	if u.Message.ReplyToMessage != nil {
		replyHandler(tg, u)
		return
	}

	formatted = fmt.Sprintf("%s%s%s %s",
		tg.Settings.Prefix,
		username,
		tg.Settings.Suffix,
		// Trim unexpected trailing whitespace
		strings.Trim(u.Message.Text, " "))

	tg.sendToIrc(formatted)
}

/*
replyHandler handles when users reply to a Telegram message
*/
func replyHandler(tg *Client, u tgbotapi.Update) {
	replyText := strings.Trim(u.Message.ReplyToMessage.Text, " ")
	username := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.From, tg.Settings.PreferName)
	replyUser := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.ReplyToMessage.From, tg.Settings.PreferName)

	// Only show a portion of the reply text
	if replyTextAsRunes := []rune(replyText); len(replyTextAsRunes) > tg.Settings.ReplyLength {
		replyText = string(replyTextAsRunes[:tg.Settings.ReplyLength]) + "…"
	}

	formatted := fmt.Sprintf("%s%s%s %sRe %s: %s%s %s",
		tg.Settings.Prefix,
		username,
		tg.Settings.Suffix,
		tg.Settings.ReplyPrefix,
		replyUser,
		replyText,
		tg.Settings.ReplySuffix,
		u.Message.Text)

	tg.sendToIrc(formatted)
}

/*
joinHandler handles when users join the Telegram group
*/
func joinHandler(tg *Client, users *[]tgbotapi.User) {
	if tg.IRCSettings.ShowJoinMessage {
		for _, user := range *users {
			user := user
			username := GetFullUsername(tg.IRCSettings.ShowZWSP, &user, tg.Settings.PreferName)
			formatted := username + " has joined the Telegram Group!"
			tg.sendToIrc(formatted)
		}
	}
}

/*
partHandler handles when users leave the Telegram group
*/
func partHandler(tg *Client, user *tgbotapi.User) {
	if tg.IRCSettings.ShowLeaveMessage {
		username := GetFullUsername(tg.IRCSettings.ShowZWSP, user, tg.Settings.PreferName)
		formatted := username + " has left the Telegram Group!"

		tg.sendToIrc(formatted)
	}
}

/*
stickerHandler handles the Message.Sticker Telegram Object, which formats the
Telegram message into its base Emoji unicode character.
*/
func stickerHandler(tg *Client, u tgbotapi.Update) {
	if !tg.IRCSettings.SendStickerEmoji {
		tg.logger.LogDebug("Skipped processing Message.Sticker. Reason: IRC_SEND_STICKER_EMOJI=false")
		return
	}

	username := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.From, tg.Settings.PreferName)
	formatted := fmt.Sprintf("%s%s%s %s",
		tg.Settings.Prefix,
		username,
		tg.Settings.Suffix,
		u.Message.Sticker.Emoji)
	tg.sendToIrc(formatted)
}

/*
photoHandler handles the Message.Photo Telegram object. Only acknowledges Photo
exists, and sends notification to IRC
*/
func photoHandler(tg *Client, u tgbotapi.Update) {
	if !tg.IRCSettings.SendPhoto {
		tg.logger.LogDebug("Skipped processing Message.Photo. Reason: IRC_SEND_PHOTO=false")
		return
	}

	link := uploadImage(tg, u)
	username := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.From, tg.Settings.PreferName)
	caption := u.Message.Caption
	if caption == "" {
		caption = "No caption provided."
	}

	// TeleIRC can fail to upload to Imgur
	if link == "" {
		tg.logger.LogError("Failed imgur photo upload for", username)
	} else {
		formatted := "'" + caption + "' uploaded by " + username + ": " + link
		tg.sendToIrc(formatted)
	}
}

/*
documentHandler receives a document object from Telegram, and sends
a notification to IRC.
*/
func documentHandler(tg *Client, u *tgbotapi.Message) {
	if !tg.IRCSettings.SendDocument {
		tg.logger.LogDebug("Skipped processing document object. Reason: IRC_SEND_DOCUMENT=false")
		return
	}

	username := GetUsername(tg.IRCSettings.ShowZWSP, u.From, tg.Settings.PreferName)
	formatted := username + " shared a file"
	if u.Document.MimeType != "" {
		formatted += " (" + u.Document.MimeType + ")"
	}

	if u.Caption != "" {
		formatted += " on Telegram with caption: " + "'" + u.Caption + "'."
	} else if u.Document.FileName != "" {
		formatted += " on Telegram with title: " + "'" + u.Document.FileName + "'."
	}

	tg.sendToIrc(formatted)
}

/*
locationHandler receivers a location object from Telegram, and sends
a notification to IRC.
*/
func locationHandler(tg *Client, u *tgbotapi.Message) {
	if !tg.IRCSettings.ShowLocationMessage {
		tg.logger.LogDebug("Skipped processing location object. Reason: IRC_SHOW_LOCATION_MESSAGE=false")
		return
	}

	username := GetUsername(tg.IRCSettings.ShowZWSP, u.From, tg.Settings.PreferName)
	formatted := username + " shared their location: ("

	// f means do not use an exponent.
	// -1 means use the smallest number of digits needed so parseFloat will return f exactly.
	// 64 to represent a standard 64 bit floating point number.
	formatted += strconv.FormatFloat(u.Location.Latitude, 'f', -1, 64)
	formatted += ", "
	formatted += strconv.FormatFloat(u.Location.Longitude, 'f', -1, 64)
	formatted += ")."

	tg.sendToIrc(formatted)
}
