package internal

import (
	"fmt"

	"github.com/lrstanley/girc"
)

type Client struct {
	*girc.Client
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
	return Client{client}
}

func (c Client) AddHandlers(settings *Settings) {
	c.Handlers.Add(girc.ALL_EVENTS, GenericHandler)
}

func GenericHandler(c *girc.Client, e girc.Event) {
	fmt.Println(e.String())
}
