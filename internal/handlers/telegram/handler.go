package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TODO: These format strings are currently unused, so commenting out for now
/*
const (
	joinFmt = "%s (@%s) has joined the Telegram Group!"
	partFmt = "%s (@%s) has left the Telegram Group."
)
*/

/*
Handler specifies a function that handles a Telegram update.
In this case, we take a Telegram client and update object,
where the specific Handler will "handle" the given event.
*/
type Handler = func(tg *Client, u tgbotapi.Update)

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
