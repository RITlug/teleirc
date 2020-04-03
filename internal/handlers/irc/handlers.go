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
type Handler = func(c Client) func(*girc.Client, girc.Event)

/*
checkBlacklist checks the IRC blacklist for a name, and returns whether
or not the name is in the blacklist
*/
func checkBlacklist(c Client, toCheck string) bool {
	for _, name := range c.Settings.IRCBlacklist {
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
func connectHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.logger.LogDebug("connectHandler triggered")
		if c.Settings.ChannelKey != "" {
			c.Cmd.JoinKey(c.Settings.Channel, c.Settings.ChannelKey)
		} else {
			c.Cmd.Join(c.Settings.Channel)
		}
	}
}

/*
messageHandler handles the PRIVMSG IRC event, which entails both private
and channel messages. However, it only cares about channel messages
*/
func messageHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.logger.LogDebug("messageHandler triggered")
		// Only send if user is not in blacklist
		if !(checkBlacklist(c, e.Source.Name)) {
			formatted := c.Settings.Prefix + e.Source.Name + c.Settings.Suffix + " " + e.Params[1]
			if e.IsFromChannel() {
				c.sendToTg(formatted)
			}
		}
	}
}

func joinHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.logger.LogDebug("joinHandler triggered")
		if c.TelegramSettings != nil && c.TelegramSettings.ShowJoinMessage {
			c.sendToTg(fmt.Sprintf(joinFmt, e.Source.Name))
		}
	}
}

func partHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.logger.LogDebug("partHandler triggered")
		if c.TelegramSettings != nil && c.TelegramSettings.ShowLeaveMessage {
			c.sendToTg(fmt.Sprintf(partFmt, e.Source.Name))
		}
	}
}

func quitHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.logger.LogDebug("quitHandler triggered")
		if c.TelegramSettings != nil && c.TelegramSettings.ShowLeaveMessage {
			c.sendToTg(fmt.Sprintf(quitFmt, e.Source.Name, e.Params[0]))
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
