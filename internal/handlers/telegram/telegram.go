// Package Telegram handles all Telegram-side logic.
package telegram

import (
	"fmt"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ritlug/teleirc/internal"
)

/*
Client contains information for the Telegram bridge, including
the TelegramSettings needed to run the bot
*/
type Client struct {
	api      *tgbotapi.BotAPI
	Settings internal.TelegramSettings
}

/*
NewClient creates a new Telegram client
*/
func NewClient(settings internal.TelegramSettings) (Client, error) {
	bot, err := tgbotapi.NewBotAPI(settings.Token)
	if err != nil {
		fmt.Println("Error creating Telegram client")
		return Client{}, err
	}
	return Client{bot, settings}, nil
}

/*
StartBot adds necessary handlers to the client and then connects,
returning any errors that occur
*/
func (tg *Client) StartBot(errChan chan<- error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tg.api.GetUpdatesChan(u)
	if err != nil {
		errChan <- err
	}

	// go tg.messageHandler(updates)
	// return nil

	for update := range updates {
		if update.Message == nil {
			continue
		}
		// Repeat sent message back to user
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		tg.api.Send(msg)
	}
	errChan <- nil
}
