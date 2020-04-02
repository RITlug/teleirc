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
	Settings         *internal.IRCSettings
	TelegramSettings *internal.TelegramSettings
	logger           internal.DebugLogger
	sendToTg         func(string)
}

/*
NewClient returns a new IRCClient based on the provided settings
*/
func NewClient(settings *internal.IRCSettings, telegramSettings *internal.TelegramSettings, logger internal.DebugLogger) Client {
	logger.LogInfo("Creating new IRC bot client...")
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
	return Client{client, settings, telegramSettings, logger, nil}
}

/*
StartBot adds necessary handlers to the client and then connects,
returns any errors that occur
*/
func (c Client) StartBot(errChan chan<- error, sendMessage func(string)) {
	c.logger.LogInfo("Starting up IRC bot...")
	c.sendToTg = sendMessage
	c.addHandlers()
	// 10 second timeout for connection
	if err := c.DialerConnect(&net.Dialer{Timeout: 10 * time.Second}); err != nil {
		errChan <- err
		c.logger.LogError(err)
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
	for eventType, handler := range getHandlerMapping() {
		c.logger.LogDebug("Adding IRC event handler:", eventType)
		c.Handlers.Add(eventType, handler(c))
	}
}
