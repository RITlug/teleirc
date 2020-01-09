package irc

import (
	"github.com/lrstanley/girc"
)

/*
Handler specifies a function that handles an IRC event
In this case, we take an IRC client and return a function that
handles an IRC event
*/
type Handler = func(c Client) func(*girc.Client, girc.Event)

/*
connectHandler return a function to use as the connect handler for girc,
so that the specified channel is joined after the server connection is established
*/
func connectHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Cmd.Join(c.Settings.Channel)
	}
}

func privMsgHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		if pretty, ok := e.Pretty(); ok {
			c.SendMessage(pretty)
		}
	}
}

/*
GetHandlerMapping returns a mapping of girc event types to handlers
*/
func GetHandlerMapping() map[string]Handler {
	return map[string]Handler{
		girc.CONNECTED: connectHandler,
		girc.PRIVMSG:   privMsgHandler,
	}
}
