package irc

import (
	"fmt"

	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
)

const (
	plainMessageFormat = "%s%%s%s %%s"
)

/*
Client contains information for our IRC bridge, including the girc Client
and the IRCSettings that were passed into NewClient
*/
type Client struct {
	*girc.Client
	Settings      internal.IRCSettings
	MessageFormat string
}

/*
NewClient returns a new IRCClient based on the provided settings
*/
func NewClient(settings internal.IRCSettings) Client {
	fmt.Println("Creating new IRC bot client...")
	client := girc.New(girc.Config{
		Server: settings.Server,
		Port:   settings.Port,
		Nick:   settings.BotName,
		User:   settings.BotName,
	})
	if settings.NickServPassword != "" {
		client.Config.SASL = &girc.SASLPlain{
			User: settings.BotName,
			Pass: settings.NickServPassword,
		}
	}
	msgFmt := fmt.Sprintf(plainMessageFormat, settings.Prefix, settings.Suffix)
	return Client{client, settings, msgFmt}
}

/*
StartBot adds necessary handlers to the client and then connects,
returns any errors that occur
*/
func (c Client) StartBot(errChan chan<- error) {
	fmt.Println("Starting up IRC bot...")
	c.addHandlers()
	if err := c.Connect(); err != nil {
		errChan <- err
	} else {
		errChan <- nil
	}
}

/*
SendMessage sends a message to the IRC channel specified in the
settings
*/
func (c Client) SendMessage(msg string) {
	c.Cmd.Message(c.Settings.Channel, msg)
}

/*
addHandlers adds handlers for the client struct based on the settings
that were passed in to NewClient
*/
func (c Client) addHandlers() {
	fmt.Println("Adding IRC event handlers...")
	for eventType, handler := range GetHandlerMapping() {
		c.Handlers.Add(eventType, handler(c))
	}
}
