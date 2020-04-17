package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDocument(t *testing.T) {
	correct := "Someone shared a file (Random File) on Telegram with caption: 'Random Caption'."
	updateObj := &tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "Someone",
			},
			Document: &tgbotapi.Document{
				FileID: "Random File",
			},
			Caption: "Random Caption",
		},
	}
	// Call the String() function which checks if a username(Optional) exists
	updateObj.Message.From.UserName = updateObj.Message.From.String()

	// Assign the MimeType(Optional) to FileID(Required) if it does not exists
	if updateObj.Message.Document.MimeType == ""{
		updateObj.Message.Document.MimeType = updateObj.Message.Document.FileID
	}

	// Change to FileName(Optional) is the Caption(Optional) is not given
	if updateObj.Message.Caption == ""{
		updateObj.Message.Document.FileName = "Random Text"
	}
	clientObj := &Client{
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	documentHandler(clientObj, updateObj.Message)
}
