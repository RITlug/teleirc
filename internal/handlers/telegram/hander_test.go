package telegram_test

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kyokomi/emoji"
	"github.com/ritlug/teleirc/internal"
	bridge "github.com/ritlug/teleirc/internal/handlers/telegram"
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

	clientObj := &bridge.Client{}

	clientObj.Settings = internal.TelegramSettings{
		Prefix: "<",
		Suffix: ">",
	}
	clientObj.SendToIrc = func(s string) {
		assert.Equal(t, correct, s)
	}

	// clientObj.handler.stickerHandler(clientObj, updateObj)

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
	clientObj := &bridge.Client{}

	clientObj.Settings = internal.TelegramSettings{
		Prefix: "<",
		Suffix: ">",
	}
	clientObj.SendToIrc = func(s string) {
		assert.Equal(t, correct, s)
	}

	// clientObj.handler.messageHandler(clientObj, updateObj)
}
