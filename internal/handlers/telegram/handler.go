package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	joinFmt = "%s (@%s) has joined the Telegram Group!"
	partFmt = "%s (@%s) has left the Telegram Group."
)

/*
messageHandler handles the Message Telegram Object, which formats the
Telegram update into a simple string for IRC.
*/
func messageHandler(tg *Client, u tgbotapi.Update) {
	formatted := tg.Settings.Prefix + u.Message.From.UserName +
		tg.Settings.Suffix + u.Message.Text

	tg.sendToIrc(formatted)
}
