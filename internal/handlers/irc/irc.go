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
	if err := c.ConnectDialer(&net.Dialer{Timeout: 10 * time.Second}); err != nil {
		errChan <- err
		c.logger.LogError(err)
	} else {
		errChan <- nil
	}
}

/*
AddHandler registers the handler function for the given event.
*/
func (c Client) AddHandler(eventType string, cb func(*girc.Client, girc.Event)) {
	c.Handlers.Add(eventType, cb)
}

/*
ConnectDialer allows you to specify your own custom dialer which implements
the Dialer interface.
*/
func (c Client) ConnectDialer(dialer girc.Dialer) error {
	return c.DialerConnect(dialer)
}

/*
Message sends a PRIVMSG to target (either channel, service, or user).
*/
func (c Client) Message(channel string, msg string) {
	c.Cmd.Message(channel, msg)
}

/*
Join attempts to enter a list of IRC channels, at bulk if possible to
prevent sending extensive JOIN commands.
*/
func (c Client) Join(channels ...string) {
	c.Cmd.Join(channels...)
}

/*
JoinKey attempts to enter an IRC channel with a password.
*/
func (c Client) JoinKey(channel string, key string) {
	c.Cmd.JoinKey(channel, key)
}

/*
Logger returns the DebugLogger to be used for logging
*/
func (c Client) Logger() internal.DebugLogger {
	return c.logger
}

/*
SendToTg sends a message to Telegram
*/
func (c Client) SendToTg(msg string) {
	c.sendToTg(msg)
}

/*
IRCSettings returns the IRCSettings struct associated with this client
*/
func (c Client) IRCSettings() *internal.IRCSettings {
	return c.Settings
}

/*
TgSettings returns the TgSettings struct associated with this client
*/
func (c Client) TgSettings() *internal.TelegramSettings {
	return c.TelegramSettings
}

/*
SendMessage sends a message to the IRC channel specified in the
settings
*/
func (c Client) SendMessage(msg string) {
	c.Message(c.Settings.Channel, msg)
}

/*
addHandlers adds handlers for the client struct based on the settings
that were passed in to NewClient
*/
func (c Client) addHandlers() {
	for eventType, handler := range getHandlerMapping() {
		c.logger.LogDebug("Adding IRC event handler:", eventType)
		c.AddHandler(eventType, handler(c))
	}
}
