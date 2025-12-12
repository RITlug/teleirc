package telegram

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ritlug/teleirc/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewClientBasic(t *testing.T) {
	tgRequiredSettings := &internal.TelegramSettings{
		Token:   "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatIDs: []int64{-0000000000000},
	}
	tgExpectedSettings := &internal.TelegramSettings{
		Token:               "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatIDs:             []int64{-0000000000000},
		ShowJoinMessage:     false,
		ShowActionMessage:   false,
		ShowLeaveMessage:    false,
		ShowKickMessage:     false,
		MaxMessagePerMinute: 0,
	}
	imgurSettings := &internal.ImgurSettings{
		ImgurClientID: "7d6b00b87043f58",
	}
	logger := internal.Debug{
		DebugLevel: false,
	}
	var tgapi *tgbotapi.BotAPI
	client := NewClient(tgRequiredSettings, nil, imgurSettings, tgapi, logger)
	assert.Equal(t, client.Settings, tgExpectedSettings, "Basic client settings should be properly set")
}

func TestNewClientFull(t *testing.T) {
	tgSettings := &internal.TelegramSettings{
		Token:               "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatIDs:             []int64{-0000000000000},
		ShowJoinMessage:     true,
		ShowActionMessage:   true,
		ShowLeaveMessage:    true,
		ShowKickMessage:     true,
		MaxMessagePerMinute: 10,
	}
	tgDefaultSettings := &internal.TelegramSettings{
		Token:               "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatIDs:             []int64{-0000000000000},
		ShowJoinMessage:     false,
		ShowActionMessage:   false,
		ShowLeaveMessage:    false,
		ShowKickMessage:     false,
		MaxMessagePerMinute: 0,
	}
	imgurSettings := &internal.ImgurSettings{
		ImgurClientID: "7d6b00b87043f58",
	}
	logger := internal.Debug{
		DebugLevel: false,
	}
	var tgapi *tgbotapi.BotAPI
	client := NewClient(tgSettings, nil, imgurSettings, tgapi, logger)
	assert.Equal(t, client.Settings, tgSettings, "All client settings should be properly set")
	assert.NotEqual(t, client.Settings, tgDefaultSettings, "tgSettings should override defaults")
}

func TestNewClientMultipleChats(t *testing.T) {
	tgSettings := &internal.TelegramSettings{
		Token:   "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatIDs: []int64{-1000000000000, -2000000000000},
	}
	imgurSettings := &internal.ImgurSettings{
		ImgurClientID: "7d6b00b87043f58",
	}
	logger := internal.Debug{
		DebugLevel: false,
	}
	var tgapi *tgbotapi.BotAPI
	client := NewClient(tgSettings, nil, imgurSettings, tgapi, logger)
	assert.Equal(t, client.Settings.ChatIDs, []int64{-1000000000000, -2000000000000}, "Client should support multiple ChatIDs")
	assert.Len(t, client.Settings.ChatIDs, 2, "Should have two chat IDs")
