package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
TestDocument_plain checks the behavior of the document handlers if only required
fields are available.
*/
func TestDocument_plain(t *testing.T) {
	correct := "test shared a file"
	updateObj := &tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
			},
			Document: &tgbotapi.Document{
				FileID: "https://teleirc.com/file.txt",
			},
		},
	}
	clientObj := &Client{
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	documentHandler(clientObj, updateObj.Message)
}

/*
TestDocument_basic checks the behavior of the document handlers when
the update just has required informations in addition to the caption.
*/
func TestDocument_basic(t *testing.T) {
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

/*
TestDocument_with_mime checks the behavior of the document handlers when
the document contains the mimetype information.
*/
func TestDocument_with_mime(t *testing.T) {
	correct := "test shared a file (test/txt) on Telegram with caption: 'Random Caption'."
	updateObj := &tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
			},
			Document: &tgbotapi.Document{
				FileID:   "https://teleirc.com/file.txt",
				MimeType: "test/txt",
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

/*
TestDocument_with_bothNames checks the behavior of the document handlers when
both firstname and username exist. It also incorporates the availability of a mimetype.
*/
func TestDocument_with_bothNames(t *testing.T) {
	correct := "user shared a file (test/txt) on Telegram with caption: 'Random Caption'."
	updateObj := &tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
				UserName:  "user",
			},
			Document: &tgbotapi.Document{
				FileID:   "https://teleirc.com/file.txt",
				MimeType: "test/txt",
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

/*
TestDocument_wout_caption checks the behavior of the document handlers when neither
a caption nor a username is attached to the document. It also test a case where
both filename and mimetype exist.
*/
func TestDocument_wout_caption(t *testing.T) {
	correct := "test shared a file (test/txt) on Telegram with title: 'test.txt'."
	updateObj := &tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
			},
			Document: &tgbotapi.Document{
				FileID:   "https://teleirc.com/file.txt",
				MimeType: "test/txt",
				FileName: "test.txt",
			},
			Caption: "",
		},
	}
	clientObj := &Client{
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	documentHandler(clientObj, updateObj.Message)
}

/*
TestDocument_with_CapAndFile checks the behavior of the document handlers when
both caption and filename exist. It also incorporates the availability of both
firstname and username
*/
func TestDocument_with_CapAndFile(t *testing.T) {
	correct := "user shared a file (test/txt) on Telegram with caption: 'Random Caption'."
	updateObj := &tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
				UserName:  "user",
			},
			Document: &tgbotapi.Document{
				FileID:   "https://teleirc.com/file.txt",
				MimeType: "test/txt",
				FileName: "test.txt",
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
