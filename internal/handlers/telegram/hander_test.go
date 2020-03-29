package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kyokomi/emoji"
	"github.com/ritlug/teleirc/internal"
)

func ExampleStickerHandler() {

	update_obj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName:     "test",
			},
			Sticker: &tgbotapi.Sticker{
				Emoji: emoji.Sprint(":smile:"),
			},
		},
	}
	client_obj := &Client{
		Settings: internal.TelegramSettings{
			Prefix:              "<",
			Suffix:              ">",
		},
		sendToIrc: func(s string) {
			fmt.Println(s)
		},
	}

	stickerHandler(client_obj, update_obj)
	// Output:
	// <test>😄
}

func ExampleMessageHandler() {

	update_obj := tgbotapi.Update{
		Message: &tgbotapi.Message{
			From: &tgbotapi.User{
				UserName:     "test",
			},
			Text: "Random Text",
		},
	}
	client_obj := &Client{
		Settings: internal.TelegramSettings{
			Prefix:              "<",
			Suffix:              ">",
		},
		sendToIrc: func(s string) {
			fmt.Println(s)
		},
	}

	messageHandler(client_obj, update_obj)
	// Output:
	// <test>Random Text
}