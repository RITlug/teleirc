package irc

import (
	"fmt"
	"regexp"
	"strings"
	"regexp"

	"github.com/lrstanley/girc"
)

const (
	joinFmt         = "* %s joins"
	partFmt         = "* %s parts"
	quitFmt         = "* %s quit (%s)"
	kickFmt         = "* %s kicked %s from %s: %s"
	topicChangeFmt  = "* %s changed topic to: %s"
	topicClearedFmt = "* %s removed topic"
	nickFmt         = "* %s is now known as: %s"
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

func shouldSendJoin(c ClientInterface, toCheck string) bool {
	var settings = c.TgSettings()
	if settings.ShowJoinMessage {
		return true
	} else if settings.JoinMessageAllowList == nil {
		return false
	}

	for _, name := range settings.JoinMessageAllowList {
		if strings.EqualFold(toCheck, name) {
			return true
		}
	}
	return false
}

func shouldSendLeave(c ClientInterface, toCheck string) bool {
	var settings = c.TgSettings()
	if settings.ShowLeaveMessage {
		return true
	} else if settings.LeaveMessageAllowList == nil {
		return false
	}

	for _, name := range settings.LeaveMessageAllowList {
		if strings.EqualFold(toCheck, name) {
			return true
		}
	}
	return false
}

func hasNoForwardPrefix(c ClientInterface, toCheck string) bool {
	noForwardPrefix := c.IRCSettings().NoForwardPrefix

	if noForwardPrefix == "" {
		return false
	}

	if strings.HasPrefix(toCheck, c.IRCSettings().NoForwardPrefix) {
		return true
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

func disconnectHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("disconnectHandler triggered")
		if c.TgSettings().ShowDisconnectMessage {
			c.SendToTg("Lost connection to '" + c.IRCSettings().Channel + "' on '" + c.IRCSettings().Server + "'")
		}
	}
}

/*
messageHandler handles the PRIVMSG IRC event, which entails both private
and channel messages. However, it only cares about channel messages
*/
func messageHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	var colorStripper = regexp.MustCompile(`[\x02\x1F\x0F\x16]|\x03(\d\d?(,\d\d?)?)?`)

	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("messageHandler triggered")
		// Only send if user is not in blacklist
		if !(checkBlacklist(c, e.Source.Name)) {

			if e.IsFromChannel() {
				formatted := ""
				if e.IsAction() {
					msg := e.Last()
					// Strips out ACTION word from text
					formatted = "* " + e.Source.Name + " " + msg[8:len(msg)-1]
				} else {
					formatted = c.IRCSettings().Prefix + e.Source.Name + c.IRCSettings().Suffix + " " + e.Params[1]
				}

				if hasNoForwardPrefix(c, e.Params[1]) {
					return // sender didn't want this forwarded
				}

				// Strip of mIRC formatting
				formatted = colorStripper.ReplaceAllString(formatted, "")

				c.SendToTg(formatted)
			}
		}
	}
}

func joinHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("joinHandler triggered")
		if (e.Source != nil) && shouldSendJoin(c, e.Source.Name) {
			c.SendToTg(fmt.Sprintf(joinFmt, e.Source.Name))
		}
	}
}

func partHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("partHandler triggered")
		if (e.Source != nil) && shouldSendLeave(c, e.Source.Name) {
			c.SendToTg(fmt.Sprintf(partFmt, e.Source.Name))
		}
	}
}

func topicHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("topicHandler triggered")
		if c.TgSettings().ShowTopicMessage {
			// e.Source.Name is the user who changed the topic.
			// e.Params[0] is the channel where the topic changed.
			// e.Params[1] is the new topic.  We should assume that
			// this may or may not appear as its possible to clear a topic.
			if len(e.Params) <= 1 {
				c.SendToTg(fmt.Sprintf(topicClearedFmt, e.Source.Name))
			} else {
				c.SendToTg(fmt.Sprintf(topicChangeFmt, e.Source.Name, e.Params[1]))
			}
		}
	}
}

func quitHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("quitHandler triggered")
		if (e.Source != nil) && shouldSendLeave(c, e.Source.Name) {
			c.SendToTg(fmt.Sprintf(quitFmt, e.Source.Name, e.Params[0]))
		}
	}
}

/*
kickHandler handles the event when a user is kicked from the IRC channel.
*/
func kickHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("kickHandler triggered")
		if c.TgSettings().ShowKickMessage {
			var reason string
			// Params are obtained from the kick command: /kick #channel nickname [reason]
			if len(e.Params) == 2 {
				reason = "Reason Undefined"
			} else {
				reason = e.Last()
			}
			c.SendToTg(fmt.Sprintf(kickFmt, e.Source.Name, e.Params[1], e.Params[0], reason))
		}
	}
}

func nickHandler(c ClientInterface) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Logger().LogDebug("nickHandler triggered")
		if c.TgSettings().ShowNickMessage {
			// e.Source.Name is the original name.
			// e.Params[0] is the new nick name.
			// However, let's assume it is possible (though unlikely)
			// e.Params can be empty.
			var newName string
			if len(e.Params) == 0 {
				newName = "Unspecified Name"
			} else {
				newName = e.Params[0]
			}
			c.SendToTg(fmt.Sprintf(nickFmt, e.Source.Name, newName))
		}
	}
}

/*
getHandlerMapping returns a mapping of girc event types to handlers
*/
func getHandlerMapping() map[string]Handler {
	return map[string]Handler{
		girc.CONNECTED:    connectHandler,
		girc.DISCONNECTED: disconnectHandler,
		girc.JOIN:         joinHandler,
		girc.KICK:         kickHandler,
		girc.NICK:         nickHandler,
		girc.PRIVMSG:      messageHandler,
		girc.PART:         partHandler,
		girc.TOPIC:        topicHandler,
		girc.QUIT:         quitHandler,
	}
}
