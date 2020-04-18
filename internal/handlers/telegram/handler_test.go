package telegram

import (
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

func TestStickerSmile_withusername(t *testing.T) {
	correct := "<test> 😄"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName:  "test",
				FirstName: "testing",
				LastName:  "123",
			},
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

func TestStickerSmile_withoutusername(t *testing.T) {
	correct := "<testing 123> 😄"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName:  "",
				FirstName: "testing",
				LastName:  "123",
			},
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

func TestMessageRandom_withusername(t *testing.T) {
	correct := "<test> Random Text"

	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName:  "test",
				FirstName: "testing",
				LastName:  "123",
			},
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

func TestMessageRandom_withoutusername(t *testing.T) {
	correct := "<testing 123> Random Text"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName:  "",
				FirstName: "testing",
				LastName:  "123",
			},
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
