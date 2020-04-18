package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
GetUsername returns the name and username of a user. Since usernames are optional
on Telegram, we first need to check to see if they have one set.
*/
func GetUsername(u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	return u.FirstName + " (@" + u.UserName + ")"
}
