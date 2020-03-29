package telegram

import (
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
			tg.logger.LogDebug("Missing message data")
			continue
		case u.Message.Text != "":
			tg.logger.LogDebug("messageHandler triggered")
			messageHandler(tg, u)
		case u.Message.Sticker != nil:
			tg.logger.LogDebug("stickerHandler triggered")
			stickerHandler(tg, u)
		case u.Message.Document != nil:
			tg.logger.LogDebug("documentHandler triggered")
			documentHandler(tg, u.Message)
		default:
			tg.logger.LogDebug("triggered, but message type is currently unsupported")
			continue
		}
	}
}

/*
messageHandler handles the Message Telegram Object, which formats the
Telegram update into a simple string for IRC.
*/
func messageHandler(tg *Client, u tgbotapi.Update) {
	formatted := tg.Settings.Prefix + u.Message.From.UserName +
		tg.Settings.Suffix + u.Message.Text

	tg.sendToIrc(formatted)
}

/*
stickerHandler handles the Message.Sticker Telegram Object, which formats the
Telegram message into its base Emoji unicode character.
*/
func stickerHandler(tg *Client, u tgbotapi.Update) {
	formatted := tg.Settings.Prefix + u.Message.From.UserName +
		tg.Settings.Suffix + u.Message.Sticker.Emoji

	tg.sendToIrc(formatted)
}

/*
documentHandler receives a document object from Telegram, and sends
a notification to IRC.
*/
func documentHandler(tg *Client, u *tgbotapi.Message) {

	formatted := u.From.UserName + " shared a file (" + u.Document.MimeType + ")"

	if u.Caption != "" {
		formatted += " on Telegram with caption: " + "'" + u.Caption + "'."
	} else {
		formatted += " on Telegram with title: " + "'" + u.Document.FileName + "'."
	}

	tg.sendToIrc(formatted)
}
