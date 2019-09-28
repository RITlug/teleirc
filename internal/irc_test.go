package internal

import (
	"testing"
	"time"

	"github.com/lrstanley/girc"
	"github.com/stretchr/testify/assert"
)

func TestNewClientBasic(t *testing.T) {
	ircSettings := IRCSettings{
		Server:  "test_server",
		Port:    1234,
		BotName: "test_name",
	}
	client := NewClient(&Settings{IRC: ircSettings})

	expectedPing, _ := time.ParseDuration("20s")
	expectedConfig := girc.Config{
		Server:    "test_server",
		Port:      1234,
		Nick:      "test_name",
		User:      "test_name",
		PingDelay: expectedPing,
	}
	assert.Equal(t, client.Settings, ircSettings, "Client settings should be properly set")
	assert.Equal(t, client.Config, expectedConfig, "girc config should be properly set")
}

func TestNewClientFull(t *testing.T) {
	ircSettings := IRCSettings{
		Server:           "test_server",
		Port:             1234,
		BotName:          "test_name",
		NickServPassword: "test_pass",
	}
	client := NewClient(&Settings{IRC: ircSettings})
	expectedPing, _ := time.ParseDuration("20s")
	expectedConfig := girc.Config{
		Server:    "test_server",
		Port:      1234,
		Nick:      "test_name",
		User:      "test_name",
		PingDelay: expectedPing,
		SASL: &girc.SASLPlain{
			User: "test_name",
			Pass: "test_pass",
		},
	}
	assert.Equal(t, client.Settings, ircSettings, "Client settings should be properly set")
	assert.Equal(t, client.Config, expectedConfig, "girc config should be properly set")

}
