package internal

import (
	"fmt"

	"github.com/lrstanley/girc"
)

/*
IRCClient contains information for our IRC bridge, including the girc Client
and the IRCSettings that were passed into NewClient
*/
type IRCClient struct {
	*girc.Client
	Settings IRCSettings
}

/*
NewClient returns a new IRCClient based on the provided settings
*/
func NewClient(settings *Settings) IRCClient {
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
	return IRCClient{client, settings.IRC}
}

/*
AddHandlers adds handlers for the client struct based on the settings
that were passed in to NewClient
*/
func (c IRCClient) AddHandlers() {
	c.Handlers.Add(girc.ALL_EVENTS, func(c *girc.Client, e girc.Event) {
		fmt.Println(e.String())
	})
	c.Handlers.Add(girc.CONNECTED, connectHandlerTest(c))
}

func connectHandlerTest(c IRCClient) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Cmd.Join(c.Settings.Channel)
	}
}
