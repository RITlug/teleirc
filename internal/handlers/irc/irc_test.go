package irc

import (
	"testing"
	"time"

	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewClientBasic(t *testing.T) {
	ircSettings := &internal.IRCSettings{
		Server:   "irc.batcave.intl",
		Port:     1337,
		BotIdent: "alfred",
		BotName:  "Alfred Pennyworth",
		BotNick:  "alfred-p",
	}
	logger := internal.Debug{
		DebugLevel: false,
	}
	client := NewClient(ircSettings, nil, logger)

	expectedPing, _ := time.ParseDuration("20s")
	expectedConfig := girc.Config{
		Server:    "irc.batcave.intl",
		Port:      1337,
		Nick:      "alfred-p",
		Name:      "Alfred Pennyworth",
		User:      "alfred",
		PingDelay: expectedPing,
	}
	assert.Equal(t, client.Settings, ircSettings, "Client settings should be properly set")
	assert.Equal(t, client.Config, expectedConfig, "girc config should be properly set")
}

func TestNewClientFull(t *testing.T) {
	ircSettings := &internal.IRCSettings{
		BindAddress:      "129.21.13.37",
		Server:           "irc.batcave.intl",
		ServerPass:       "BatmanNeverDies!",
		Port:             1337,
		BotIdent:         "alfred",
		BotName:          "Alfred Pennyworth",
		BotNick:          "alfred-p",
		NickServUser:     "irc_moderators",
		NickServPassword: "ProtectGotham",
	}
	logger := internal.Debug{
		DebugLevel: false,
	}
	client := NewClient(ircSettings, nil, logger)
	expectedPing, _ := time.ParseDuration("20s")
	expectedConfig := girc.Config{
		Bind:       "129.21.13.37",
		Server:     "irc.batcave.intl",
		ServerPass: "BatmanNeverDies!",
		Port:       1337,
		Nick:       "alfred-p",
		Name:       "Alfred Pennyworth",
		User:       "alfred",
		PingDelay:  expectedPing,
		SASL: &girc.SASLPlain{
			User: "irc_moderators",
			Pass: "ProtectGotham",
		},
	}
	assert.Equal(t, client.Settings, ircSettings, "Client settings should be properly set")
	assert.Equal(t, client.Config, expectedConfig, "girc config should be properly set")
}
