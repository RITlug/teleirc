package irc

import (
	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
)

/*
ClientInterface represents an IRC client
*/
type ClientInterface interface {
	NewClient(internal.IRCSettings, internal.DebugLogger)
	SendMessage(string)
	StartBot(chan<- error, func(string))
	Logger() internal.DebugLogger
	addHandlers()

	AddHandler(string, func(*girc.Client, girc.Event))
	ConnectDialer(girc.Dialer) error
	Message(string, string)
}
