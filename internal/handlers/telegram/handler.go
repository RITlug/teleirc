package telegram

import (
	"fmt"
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
	username := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.From)
	replyUsername := ""

	if tg.IRCSettings.NoForwardPrefix != "" && strings.HasPrefix(u.Message.Text, tg.IRCSettings.NoForwardPrefix) {
		return
	}

	// Don't forward messages to IRC that didn't come from the
	// chat we're bridging
	if u.Message.Chat.ID != tg.Settings.ChatID {
		return
	}

	if u.Message.ReplyToMessage != nil {
		replyUsername = GetUsername(tg.IRCSettings.ShowZWSP, u.Message.ReplyToMessage.From) + ":"
	}
	formatted := fmt.Sprintf("%s%s%s%s %s",
		tg.Settings.Prefix,
		username,
		tg.Settings.Suffix,
		replyUsername,
		// Trim unexpected trailing whitespace
		strings.Trim(u.Message.Text, " "))
	tg.sendToIrc(formatted)
}

/*
joinHandler handles when users join the Telegram group
*/
func joinHandler(tg *Client, users *[]tgbotapi.User) {
	if tg.IRCSettings.ShowJoinMessage {
		for _, user := range *users {
			username := GetFullUsername(tg.IRCSettings.ShowZWSP, &user)
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
		username := GetFullUsername(tg.IRCSettings.ShowZWSP, user)
		formatted := username + " has left the Telegram Group!"

		tg.sendToIrc(formatted)
	}
}

/*
stickerHandler handles the Message.Sticker Telegram Object, which formats the
Telegram message into its base Emoji unicode character.
*/
func stickerHandler(tg *Client, u tgbotapi.Update) {
	username := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.From)
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
	username := GetUsername(tg.IRCSettings.ShowZWSP, u.Message.From)
	formatted := username + " shared a photo on Telegram with caption: '" +
		u.Message.Caption + "'"

	tg.sendToIrc(formatted)
}

/*
documentHandler receives a document object from Telegram, and sends
a notification to IRC.
*/
func documentHandler(tg *Client, u *tgbotapi.Message) {
	username := GetUsername(tg.IRCSettings.ShowZWSP, u.From)
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
