package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
GetUsername returns a Telegram user's username if one is set, or first name otherwise.
*/
func GetUsername(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	return u.UserName
}

/*
GetFullUsername returns both the Telegram user's first name and username, if available.
*/
func GetFullUsername(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	return u.FirstName + " (@" + u.UserName + ")"
}

/*
GetFullUserZwsp returns both the Telegram user's first name and username, if available.
Adds ZWSP to username to prevent username pinging across platform.
*/
func GetFullUserZwsp(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	// Add ZWSP to prevent pinging across platforms
	// See https://github.com/42wim/matterbridge/issues/175
	return u.FirstName + " (@" + u.UserName[:1] + "​" + u.UserName[1:] + ")"
}

/*
ZwspUsername adds a zero-width space after the first character of a Telegram user's
username.
*/
func ZwspUsername(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	// Add ZWSP to prevent pinging across platforms
	// See https://github.com/42wim/matterbridge/issues/175
	return u.UserName[:1] + "​" + u.UserName[1:]
}
