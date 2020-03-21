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
