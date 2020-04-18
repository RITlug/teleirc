package telegram_test

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kyokomi/emoji"
	"github.com/ritlug/teleirc/internal"
	brige "github.com/ritlug/teleirc/internal/handlers/telegram"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStickerSmile(t *testing.T) {
	correct := "<test>😄"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName: "test",
			},
			Sticker: &tgbotapi.Sticker{
				Emoji: emoji.Sprint(":smile:"),
			},
		},
	}
	clientObj := &brige.Client{}
	clientObj.Settings = internal.TelegramSettings{
		Prefix: "<",
		Suffix: ">",
	}
	clientObj

Settings: internal.TelegramSettings{
	Prefix: "<",
	Suffix: ">",
},
	sendToIrc: func(s string) {
		assert.Equal(t, correct, s)
	},

	stickerHandler(clientObj, updateObj)
}

func TestMessageRandom(t *testing.T) {
	correct := "<test>Random Text"
	updateObj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName: "test",
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
