package telegram

import (
	"errors"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/golang/mock/gomock"
	"github.com/ritlug/teleirc/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewClientBasic(t *testing.T) {
	tgRequiredSettings := &internal.TelegramSettings{
		Token:  "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatID: -0000000000000,
	}
	tgExpectedSettings := &internal.TelegramSettings{
		Token:               "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatID:              -0000000000000,
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
		ChatID:              -0000000000000,
		ShowJoinMessage:     true,
		ShowActionMessage:   true,
		ShowLeaveMessage:    true,
		ShowKickMessage:     true,
		MaxMessagePerMinute: 10,
	}
	tgDefaultSettings := &internal.TelegramSettings{
		Token:               "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
		ChatID:              -0000000000000,
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

/*
TestSendMessageNotConnected tests that SendMessage drops the message instead of
panicking when the Telegram bot has not connected yet
*/
func TestSendMessageNotConnected(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockLogger.EXPECT().LogWarning("Not connected to Telegram, dropping message:", "test message")

	client := &Client{Settings: &internal.TelegramSettings{}, logger: mockLogger}
	assert.NotPanics(t, func() { client.SendMessage("test message") })
}

/*
TestStartBotAuthFailure tests that StartBot reports a failed Telegram login and
returns, leaving the client safe to call SendMessage on
*/
func TestStartBotAuthFailure(t *testing.T) {
	authErr := errors.New("Unauthorized")
	defer func(orig func(string) (*tgbotapi.BotAPI, error)) { newBotAPI = orig }(newBotAPI)
	newBotAPI = func(string) (*tgbotapi.BotAPI, error) { return nil, authErr }

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockLogger.EXPECT().LogInfo("Starting up Telegram bot...")
	mockLogger.EXPECT().LogError("Failed to authenticate to Telegram:", authErr)
	mockLogger.EXPECT().LogWarning("Not connected to Telegram, dropping message:", "test message")

	client := &Client{Settings: &internal.TelegramSettings{}, logger: mockLogger}
	errChan := make(chan error, 1)

	assert.NotPanics(t, func() { client.StartBot(errChan, func(string) {}) })
	assert.Equal(t, authErr, <-errChan)
	assert.NotPanics(t, func() { client.SendMessage("test message") })
}
