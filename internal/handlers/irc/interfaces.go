package irc

import (
	"github.com/lrstanley/girc"
	"github.com/ritlug/teleirc/internal"
)

type ClientInterface interface {
	NewClient(internal.IRCSettings, internal.DebugLogger)
	SendMessage(string)
	StartBot(chan<- error, func(string))
	Logger() internal.DebugLogger
	addHandlers()

	AddHandler(string, func(*girc.Client, girc.Event))
	DialerConnect(girc.Dialer)
	Message()
}
