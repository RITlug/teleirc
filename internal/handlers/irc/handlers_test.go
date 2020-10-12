package irc

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
)

func TestJoinHandler_On(t *testing.T) {
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

func TestJoinHandler_Off(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowJoinMessage: false,
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
		SendToTg(gomock.Any()).
		MaxTimes(0)

	myHandler := joinHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestPartHandler_On(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowLeaveMessage: true,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("partHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Eq("* TEST_NAME parts"))

	myHandler := partHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "TEST_NAME",
		},
	})
}

func TestPartHandler_Off(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowLeaveMessage: false,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("partHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Any()).
		MaxTimes(0)

	myHandler := partHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestQuitHandler_On(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowLeaveMessage: true,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("quitHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Eq("* TEST_NAME quit (TEST_REASON)"))

	myHandler := quitHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "TEST_NAME",
		},
		Params: []string{
			"TEST_REASON",
		},
	})
}

func TestQuitHandler_Off(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowLeaveMessage: false,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("quitHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Any()).
		MaxTimes(0)

	myHandler := quitHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestKickHandler_On(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowKickMessage: true,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("kickHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Eq("* TEST_NAME kicked TEST_KICKEDNAME from TEST_GROUP: TEST_REASON"))

	myHandler := kickHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "TEST_NAME",
		},
		Params: []string{
			"TEST_GROUP",
			"TEST_KICKEDNAME",
			"TEST_REASON",
		},
	})
}

func TestKickHandler_Off(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowKickMessage: false,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("kickHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Any()).
		MaxTimes(0)

	myHandler := kickHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestKickHandlerNoReason(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowKickMessage: true,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("kickHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings)
	mockClient.
		EXPECT().
		// SendToTg(gomock.Eq("* TEST_NAME kicked TEST_KICKEDNAME from TEST_GROUP: TEST_KICKEDNAME"))
		SendToTg(gomock.Eq("* TEST_NAME kicked TEST_KICKEDNAME from TEST_GROUP: Reason Undefined"))

	myHandler := kickHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "TEST_NAME",
		},
		Params: []string{
			"TEST_GROUP",
			"TEST_KICKEDNAME",
		},
	})
}

func TestConnectHandlerKey(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ircSettings := internal.IRCSettings{
		Channel:    "SomeChannel",
		ChannelKey: "SomeKey",
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("connectHandler triggered"))
	mockClient.
		EXPECT().
		IRCSettings().
		Return(&ircSettings).
		AnyTimes()
	mockClient.
		EXPECT().
		JoinKey(gomock.Eq("SomeChannel"), gomock.Eq("SomeKey"))

	myHandler := connectHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestConnectHandlerNoKey(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ircSettings := internal.IRCSettings{
		Channel: "SomeChannel",
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("connectHandler triggered"))
	mockClient.
		EXPECT().
		IRCSettings().
		Return(&ircSettings).
		AnyTimes()
	mockClient.
		EXPECT().
		Join(gomock.Eq("SomeChannel"))

	myHandler := connectHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestDisconnectHandlerWhenDisabled(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tgSettings := internal.TelegramSettings{
		ShowDisconnectMesssage: false,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("disconnectHandler triggered"))
	// We are disabled, should never be called.
	mockClient.
		EXPECT().
		SendToTg(gomock.Any()).
		MaxTimes(0)
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings).
		AnyTimes()

	myHandler := disconnectHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestDisconnectHandlerWhenEnabled(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ircSettings := internal.IRCSettings{
		Server:  "irc.someserver.org",
		Channel: "#somechannel",
	}

	tgSettings := internal.TelegramSettings{
		ShowDisconnectMesssage: true,
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("disconnectHandler triggered"))
	mockClient.
		EXPECT().
		TgSettings().
		Return(&tgSettings).
		AnyTimes()
	mockClient.
		EXPECT().
		IRCSettings().
		Return(&ircSettings).
		AnyTimes()
	mockClient.
		EXPECT().
		SendToTg("Lost connection to '" + ircSettings.Channel + "' on '" + ircSettings.Server + "'").
		MaxTimes(1)

	myHandler := disconnectHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{})
}

func TestMessageHandlerInBlacklist(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ircSettings := internal.IRCSettings{
		IRCBlacklist: []string{
			"SomeUser",
		},
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("messageHandler triggered"))
	mockClient.
		EXPECT().
		IRCSettings().
		Return(&ircSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Any()).
		MaxTimes(0)

	myHandler := messageHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "SomeUser",
		},
	})
}

func TestMessageHandlerNotChannel(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ircSettings := internal.IRCSettings{
		IRCBlacklist: []string{},
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("messageHandler triggered"))
	mockClient.
		EXPECT().
		IRCSettings().
		Return(&ircSettings)
	mockClient.
		EXPECT().
		SendToTg(gomock.Any()).
		MaxTimes(0)

	myHandler := messageHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "SomeUser",
		},
		// Need to not be PRIVMSG or NOTICE
		Command: girc.KICK,
	})
}

func TestMessageHandlerNoForward(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ircSettings := internal.IRCSettings{
		IRCBlacklist: []string{},
		Prefix:       "<<",
		Suffix:       ">>",
		NoForwardPrefix: "[off]",
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("messageHandler triggered"))
	mockClient.
		EXPECT().
		IRCSettings().
		Return(&ircSettings).
		AnyTimes()
	mockClient.
		EXPECT().
		SendToTg(gomock.Any()).
		MaxTimes(0)

	myHandler := messageHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "SomeUser",
		},
		// Need to be PRIVMSG
		Command: girc.PRIVMSG,
		Params: []string{
			"#testchannel",
			"[off] a message",
		},
	})

}

func TestMessageHandlerFull(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	ircSettings := internal.IRCSettings{
		IRCBlacklist: []string{},
		Prefix:       "<<",
		Suffix:       ">>",
	}

	mockClient := NewMockClientInterface(ctrl)
	mockLogger := internal.NewMockDebugLogger(ctrl)
	mockClient.
		EXPECT().
		Logger().
		Return(mockLogger)
	mockLogger.
		EXPECT().
		LogDebug(gomock.Eq("messageHandler triggered"))
	mockClient.
		EXPECT().
		IRCSettings().
		Return(&ircSettings).
		AnyTimes()
	mockClient.
		EXPECT().
		SendToTg(gomock.Eq("<<SomeUser>> a message"))

	myHandler := messageHandler(mockClient)
	myHandler(&girc.Client{}, girc.Event{
		Source: &girc.Source{
			Name: "SomeUser",
		},
		// Need to be PRIVMSG
		Command: girc.PRIVMSG,
		Params: []string{
			"#testchannel",
			"a message",
		},
	})
}
