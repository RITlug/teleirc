package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kyokomi/emoji"
	"github.com/ritlug/teleirc/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartFull_On(t *testing.T) {
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

func TestPartFull_Off(t *testing.T) {
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

func TestJoinFull_On(t *testing.T) {
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

func TestJoinFull_Off(t *testing.T) {
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
