package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kyokomi/emoji"
	"github.com/ritlug/teleirc/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
		Settings: internal.TelegramSettings{
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
		Settings: internal.TelegramSettings{
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
		Settings: internal.TelegramSettings{
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
		Settings: internal.TelegramSettings{
			Prefix: "<",
			Suffix: ">",
		},
		sendToIrc: func(s string) {
			assert.Equal(t, correct, s)
		},
	}

	messageHandler(clientObj, updateObj)
}
