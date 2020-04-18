package irc

import (
	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
)

/*
ClientInterface represents an IRC client
*/
type ClientInterface interface {
	SendMessage(string)
	StartBot(chan<- error, func(string))
	Logger() internal.DebugLogger
	addHandlers()
	SendToTg(string)
	IRCSettings() *internal.IRCSettings
	TgSettings() *internal.TelegramSettings

	AddHandler(string, func(*girc.Client, girc.Event))
	ConnectDialer(girc.Dialer) error
	Message(string, string)
	JoinKey(string, string)
	Join(...string)
}
