package irc

import (
	"testing"

	"github.com/lrstanley/girc"

	"github.com/ritlug/teleirc/internal"

	gomock "github.com/golang/mock/gomock"
)

func TestJoinHandler(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowJoinMessage: true,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("joinHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Eq("* TEST_NAME joins"))

	myHandler := joinHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "TEST_NAME",
		},
	})
}
