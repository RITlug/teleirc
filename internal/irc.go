package internal

import (
	"fmt"

	"github.com/lrstanley/girc"
)

type Client struct {
	*girc.Client
}

var SETTINGS *Settings

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
	SETTINGS = settings
	return Client{client}
}

func (c Client) AddHandlers() {
	c.Handlers.Add(girc.ALL_EVENTS, GenericHandler)
	c.Handlers.Add(girc.CONNECTED, ConnectHandler)
}

func GenericHandler(c *girc.Client, e girc.Event) {
	fmt.Println(e.String())
}

func ConnectHandler(c *girc.Client, e girc.Event) {
	c.Cmd.Join(SETTINGS.IRC.Channel)
}
