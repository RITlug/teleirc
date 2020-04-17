package irc

import (
	"fmt"
	"strings"

	"github.com/lrstanley/girc"
)

const (
	joinFmt = "* %s joins"
	partFmt = "* %s parts"
	quitFmt = "* %s quit (%s)"
)

/*
Handler specifies a function that handles an IRC event
In this case, we take an IRC client and return a function that
handles an IRC event
*/
type Handler = func(c ClientInterface) func(*girc.Client, girc.Event)

/*
checkBlacklist checks the IRC blacklist for a name, and returns whether
or not the name is in the blacklist
*/
func checkBlacklist(c ClientInterface, toCheck string) bool {
	for _, name := range c.IRCSettings().IRCBlacklist {
		if strings.EqualFold(toCheck, name) {
			return true
		}
	}
	return false
}

/*
connectHandler returns a function to use as the connect handler for girc,
so that the specified channel is joined after the server connection is established
*/
func connectHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("connectHandler triggered")
		if c.IRCSettings().ChannelKey != "" {
			c.JoinKey(c.IRCSettings().Channel, c.IRCSettings().ChannelKey)
		} else {
			c.Join(c.IRCSettings().Channel)
		}
	}
}

/*
messageHandler handles the PRIVMSG IRC event, which entails both private
and channel messages. However, it only cares about channel messages
*/
func messageHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("messageHandler triggered")
		// Only send if user is not in blacklist
		if !(checkBlacklist(c, e.Source.Name)) {
			formatted := c.IRCSettings().Prefix + e.Source.Name + c.IRCSettings().Suffix + " " + e.Params[1]
			if e.IsFromChannel() {
				c.SendToTg(formatted)
			}
		}
	}
}

func joinHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("joinHandler triggered")
		if c.TgSettings().ShowJoinMessage {
			c.SendToTg(fmt.Sprintf(joinFmt, e.Source.Name))
		}
	}
}

func partHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("partHandler triggered")
		if c.TgSettings().ShowLeaveMessage {
			c.SendToTg(fmt.Sprintf(partFmt, e.Source.Name))
		}
	}
}

func quitHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("quitHandler triggered")
		if c.TgSettings().ShowLeaveMessage {
			c.SendToTg(fmt.Sprintf(quitFmt, e.Source.Name, e.Params[0]))
		}
	}
}

/*
getHandlerMapping returns a mapping of girc event types to handlers
*/
func getHandlerMapping() map[string]Handler {
	return map[string]Handler{
		girc.CONNECTED: connectHandler,
		girc.PRIVMSG:   messageHandler,
		girc.PART:      partHandler,
		girc.QUIT:      quitHandler,
		girc.JOIN:      joinHandler,
	}
}
