package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kyokomi/emoji"
	"github.com/ritlug/teleirc/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
TestPartFullOn tests the ability of the partHandler to send messages
when ShowLeaveMessage is set to true
*/
func TestPartFullOn(t *testing.T) {
	testUser := &tgbotapi.User{
		ID:        1,
		FirstName: "test",
		UserName:  "testUser",
	}
	correct := testUser.FirstName + " (@" + testUser.UserName + ") has left the Telegram Group!"
	clientObj := &Client{
		IRCSettings: &internal.IRCSettings{
			Prefix:           "<",
			Suffix:           ">",
			ShowLeaveMessage: true,
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	partHandler(clientObj, testUser)
}

/*
TestPartFullOff tests the ability of the partHandler to not send messages
when ShowLeaveMessage is set to false
*/
func TestPartFullOff(t *testing.T) {
	testUser := &tgbotapi.User{
		ID:        1,
		FirstName: "test",
		UserName:  "testUser",
	}
	correct := ""
	clientObj := &Client{
		IRCSettings: &internal.IRCSettings{
			Prefix:           "<",
			Suffix:           ">",
			ShowLeaveMessage: false,
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	partHandler(clientObj, testUser)
}

/*
TestPartNoUsername tests the ability of the partHandler to send correctly
formatted messages when a TG user has no username
*/
func TestPartNoUsername(t *testing.T) {
	testUser := &tgbotapi.User{
		ID:        1,
		FirstName: "test",
	}
	correct := testUser.FirstName + " has left the Telegram Group!"
	clientObj := &Client{
		IRCSettings: &internal.IRCSettings{
			Prefix:           "<",
			Suffix:           ">",
			ShowLeaveMessage: true,
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	partHandler(clientObj, testUser)
}

/*
TestJoinFullOn tests the ability of the joinHandler to send messages
when ShowJoinMessage is set to true
*/
func TestJoinFullOn(t *testing.T) {
	testListUser := &[]tgbotapi.User{
		tgbotapi.User{
			ID:        1,
			FirstName: "test",
			UserName:  "testUser",
		},
	}
	correct := "test (@testUser) has joined the Telegram Group!"
	clientObj := &Client{
		IRCSettings: &internal.IRCSettings{
			Prefix:          "<",
			Suffix:          ">",
			ShowJoinMessage: true,
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	joinHandler(clientObj, testListUser)
}

/*
TestJoinFullOff tests the ability of the joinHandler to not send messages
when ShowJoinMessage is set to false
*/
func TestJoinFullOff(t *testing.T) {
	testListUser := &[]tgbotapi.User{
		tgbotapi.User{
			ID:        1,
			FirstName: "test",
			UserName:  "testUser",
		},
	}
	correct := ""
	clientObj := &Client{
		IRCSettings: &internal.IRCSettings{
			Prefix:          "<",
			Suffix:          ">",
			ShowJoinMessage: false,
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	joinHandler(clientObj, testListUser)
}

/*
TestJoinNoUsername tests the ability of the joinHandler to send correctly
formatted messages when a TG user has no username
*/
func TestJoinNoUsername(t *testing.T) {
	testListUser := &[]tgbotapi.User{
		tgbotapi.User{
			ID:        1,
			FirstName: "test",
		},
	}
	correct := "test has joined the Telegram Group!"
	clientObj := &Client{
		IRCSettings: &internal.IRCSettings{
			Prefix:          "<",
			Suffix:          ">",
			ShowJoinMessage: true,
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	joinHandler(clientObj, testListUser)
}

/*
TestDocumentPlain checks the behavior of the document handlers if only required
fields are available.
*/
func TestDocumentPlain(t *testing.T) {
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
TestDocumentBasic checks the behavior of the document handlers when
the update just has required informations in addition to the caption.
*/
func TestDocumentBasic(t *testing.T) {
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
TestDocumentMime checks the behavior of the document handlers when
the document contains the mimetype information.
*/
func TestDocumentMime(t *testing.T) {
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
TestDocumentUsername checks the behavior of the document handlers when
both firstname and username exist. It also incorporates the availability of a mimetype.
*/
func TestDocumentUsername(t *testing.T) {
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
TestDocumentNoCaption checks the behavior of the document handlers when neither
a caption nor a username is attached to the document. It also test a case where
both filename and mimetype exist.
*/
func TestDocumentNoCaption(t *testing.T) {
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
TestDocumentFull checks the behavior of the document handlers when
both caption and filename exist. It also incorporates the availability of both
firstname and username
*/
func TestDocumentFull(t *testing.T) {
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

/*
TestPhotoFull tests a complete Photo object
*/
func TestPhotoFull(t *testing.T) {
	correct := "user shared a photo on Telegram with caption: 'Random Caption'"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
				UserName:  "user",
			},
			Photo: &[]tgbotapi.PhotoSize{
				tgbotapi.PhotoSize{
					FileID:   "https://teleirc.com/file.png",
					Width:    1,
					Height:   1,
					FileSize: 1,
				},
			},
			Caption: "Random Caption",
		},
	}
	clientObj := &Client{
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	photoHandler(clientObj, updateObj)
}

/*
TestPhotoNoUsername tests a Photo object with no username present. Should default
to user's FirstName
*/
func TestPhotoNoUsername(t *testing.T) {
	correct := "test shared a photo on Telegram with caption: 'Random Caption'"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
			},
			Photo: &[]tgbotapi.PhotoSize{
				tgbotapi.PhotoSize{
					FileID:   "https://teleirc.com/file.png",
					Width:    1,
					Height:   1,
					FileSize: 1,
				},
			},
			Caption: "Random Caption",
		},
	}
	clientObj := &Client{
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	photoHandler(clientObj, updateObj)
}

/*
TestPhotoNoCaption tests messages are correctly formatted when a photo
is uploaded without a caption
*/
func TestPhotoNoCaption(t *testing.T) {
	correct := "user shared a photo on Telegram with caption: ''"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				FirstName: "test",
				UserName:  "user",
			},
			Photo: &[]tgbotapi.PhotoSize{
				tgbotapi.PhotoSize{
					FileID:   "https://teleirc.com/file.png",
					Width:    1,
					Height:   1,
					FileSize: 1,
				},
			},
			Caption: "",
		},
	}
	clientObj := &Client{
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}
	photoHandler(clientObj, updateObj)
}

func TestStickerSmileWithUsername(t *testing.T) {
	testUser := &tgbotapi.User{
		ID:        1,
		UserName:  "test",
		FirstName: "testing",
		LastName:  "123",
	}
	correct := fmt.Sprintf("<%s> 😄", testUser.String())
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: testUser,
			Sticker: &tgbotapi.Sticker{
				Emoji: emoji.Sprint(":smile:"),
			},
		},
	}

	clientObj := &Client{
		Settings: &internal.TelegramSettings{
			Prefix: "<",
			Suffix: ">",
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}

	stickerHandler(clientObj, updateObj)

}

func TestStickerSmileWithoutUsername(t *testing.T) {
	testUser := &tgbotapi.User{
		ID:        1,
		UserName:  "",
		FirstName: "testing",
		LastName:  "123",
	}
	correct := fmt.Sprintf("<%s> 😄", testUser.String())
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: testUser,
			Sticker: &tgbotapi.Sticker{
				Emoji: emoji.Sprint(":smile:"),
			},
		},
	}

	clientObj := &Client{
		Settings: &internal.TelegramSettings{
			Prefix: "<",
			Suffix: ">",
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}

	stickerHandler(clientObj, updateObj)
}

func TestMessageRandomWithUsername(t *testing.T) {
	testUser := &tgbotapi.User{
		ID:        1,
		UserName:  "test",
		FirstName: "testing",
		LastName:  "123",
	}
	correct := fmt.Sprintf("<%s> Random Text", testUser.String())

	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: testUser,
			Text: "Random Text",
		},
	}

	clientObj := &Client{
		Settings: &internal.TelegramSettings{
			Prefix: "<",
			Suffix: ">",
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}

	messageHandler(clientObj, updateObj)
}

func TestMessageRandomWithoutUsername(t *testing.T) {
	testUser := &tgbotapi.User{
		ID:        1,
		UserName:  "",
		FirstName: "testing",
		LastName:  "123",
	}
	correct := fmt.Sprintf("<%s> Random Text", testUser.String())

	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: testUser,
			Text: "Random Text",
		},
	}
	clientObj := &Client{
		Settings: &internal.TelegramSettings{
			Prefix: "<",
			Suffix: ">",
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}

	messageHandler(clientObj, updateObj)
}
