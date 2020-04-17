package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocument(t *testing.T) {
	correct := "test shared a file on Telegram with caption: 'Random Caption'."
	updateObj := &tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
			},
			Document: &tgbotapi.Document{
				FileID: "https://teleirc.com/file.txt",
			},
			Caption: "Random Caption",
		},
	}
	clientObj := &Client{
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	documentHandler(clientObj, updateObj.Message)
}
