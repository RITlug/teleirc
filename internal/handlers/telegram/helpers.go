package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
ResolveUserName does basic cleanup if a user does not have a username on Telegram.
*/
func ResolveUserName(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	// Add ZWSP to prevent pinging across platforms
	// See https://github.com/42wim/matterbridge/issues/175
	username := u.UserName[:1] + "" + u.UserName[1:]
	return username
}

/*
GetFullUsername returns the name and username of a user. Since usernames are optional
on Telegram, we first need to check to see if they have one set.
*/
func GetFullUsername(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	return u.FirstName + " (@" + u.UserName[:1] + "" + u.UserName[1:] + ")"
}
