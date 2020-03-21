package irc

import (
	"net"
	"time"

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
	verbose  internal.DebugLogger
	sendToTg func(string)
}

/*
NewClient returns a new IRCClient based on the provided settings
*/
func NewClient(settings internal.IRCSettings, debug internal.DebugLogger) Client {
	debug.LogInfo("Creating new IRC bot client...")
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
	return Client{client, settings, debug, nil}
}

/*
StartBot adds necessary handlers to the client and then connects,
returns any errors that occur
*/
func (c Client) StartBot(errChan chan<- error, sendMessage func(string)) {
	c.verbose.LogInfo("Starting up IRC bot...")
	c.sendToTg = sendMessage
	c.addHandlers()
	// TODO: Currently just set to 5 seconds,
	// if we want this to be configurable we can add a config option
	if err := c.DialerConnect(&net.Dialer{Timeout: 5 * time.Second}); err != nil {
		errChan <- err
		c.verbose.LogError(err)
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
	c.verbose.LogInfo("Adding IRC event handlers...")
	for eventType, handler := range getHandlerMapping() {
		c.Handlers.Add(eventType, handler(c))
	}
}
