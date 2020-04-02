package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kyokomi/emoji"
	"github.com/ritlug/teleirc/internal"
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
	clientObj := &Client{
		Settings: internal.TelegramSettings{
			Prefix: "<",
			Suffix: ">",
		},
		sendToIrc: func(s string) {
			if s != correct {
				t.Errorf(fmt.Sprintf("stickerHandler(\":smile:\") = \"%s\"; want \"%s\"", s, correct))
			}
		},
	}

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
			if s != correct {
				t.Errorf(fmt.Sprintf("messageHandler(\"Random Text\") = \"%s\"; want \"%s\"", s, correct))
			}
		},
	}

	messageHandler(clientObj, updateObj)
}
