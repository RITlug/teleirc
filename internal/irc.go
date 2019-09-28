package internal

import (
	"fmt"

	"github.com/lrstanley/girc"
)

type Client struct {
	*girc.Client
	Settings IRCSettings
}

func NewClient(settings *Settings) Client {
	client := girc.New(girc.Config{
		Server: settings.IRC.Server,
		Port:   settings.IRC.Port,
		Nick:   settings.IRC.BotName,
		User:   settings.IRC.BotName,
		SASL: &girc.SASLPlain{
			User: settings.IRC.BotName,
			Pass: settings.IRC.NickServPassword,
		},
	})
	return Client{client, settings.IRC}
}

func (c Client) AddHandlers() {
	c.Handlers.Add(girc.ALL_EVENTS, func(c *girc.Client, e girc.Event) {
		fmt.Println(e.String())
	})
	c.Handlers.Add(girc.CONNECTED, connectHandlerTest(c))
}

func connectHandlerTest(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Cmd.Join(c.Settings.Channel)
	}
}
