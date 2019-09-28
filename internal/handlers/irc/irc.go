package irc

import (
	"fmt"

	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
)

/*
Client contains information for our IRC bridge, including the girc Client
and the IRCSettings that were passed into NewClient
*/
type Client struct {
	*girc.Client
	Settings internal.IRCSettings
}

/*
NewClient returns a new IRCClient based on the provided settings
*/
func NewClient(settings *internal.Settings) Client {
	client := girc.New(girc.Config{
		Server: settings.IRC.Server,
		Port:   settings.IRC.Port,
		Nick:   settings.IRC.BotName,
		User:   settings.IRC.BotName,
	})
	if settings.IRC.NickServPassword != "" {
		client.Config.SASL = &girc.SASLPlain{
			User: settings.IRC.BotName,
			Pass: settings.IRC.NickServPassword,
		}
	}
	return Client{client, settings.IRC}
}

func (c Client) StartBot() error {
	c.addHandlers()
	if err := c.Connect(); err != nil {
		return err
	}
	return nil
}

/*
AddHandlers adds handlers for the client struct based on the settings
that were passed in to NewClient
*/
func (c Client) addHandlers() {
	c.Handlers.Add(girc.ALL_EVENTS, func(c *girc.Client, e girc.Event) {
		fmt.Println(e.String())
	})
	c.Handlers.Add(girc.CONNECTED, connectHandler(c))
}

func connectHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Cmd.Join(c.Settings.Channel)
	}
}
