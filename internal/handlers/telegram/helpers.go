package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
GetUsername takes showZWSP condition and user then returns username with or without ​.
*/
func GetUsername(showZWSP bool, u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	if showZWSP {
		return ZwspUsername(u)
	}
	return u.UserName
}

/*
GetFullUsername takes showZWSP condition and user then returns full username with or without ​.
*/
func GetFullUsername(showZWSP bool, u *tgbotapi.User) string {
	if u.UserName == "" {
		return u.FirstName
	}
	if showZWSP {
		return GetFullUserZwsp(u)
	}
	return u.FirstName + " (@" + u.UserName + ")"
}

/*
GetFullUserZwsp returns both the Telegram user's first name and username, if available.
Adds ZWSP to username to prevent username pinging across platform.
*/
func GetFullUserZwsp(u *tgbotapi.User) string {
	// Add ZWSP to prevent pinging across platforms
	// See https://github.com/42wim/matterbridge/issues/175
	userNameAsRunes := []rune(u.UserName)
	return u.FirstName + " (@" + string(userNameAsRunes[:1]) + "\u200b" + string(userNameAsRunes[1:]) + ")"
}

/*
ZwspUsername adds a zero-width space after the first character of a Telegram user's
username.
*/
func ZwspUsername(u *tgbotapi.User) string {
	// Add ZWSP to prevent pinging across platforms
	// See https://github.com/42wim/matterbridge/issues/175
	userNameAsRunes := []rune(u.UserName)
	return string(userNameAsRunes[:1]) + "\u200b" + string(userNameAsRunes[1:])
}

/*
uploadImage uploads a Photo object from Telegram to the Imgur API and
returns a string with the Imgur URL.
*/
func uploadImage(tg *Client, u tgbotapi.Update) string {
	photo := (*u.Message.Photo)[len(*u.Message.Photo)-1]

	// Get Telegram file URL
	tgLink, err := tg.api.GetFileDirectURL(photo.FileID)
	if err != nil {
		tg.logger.LogError("Could not get Telegram Photo URL:", err)
	}

	return getImgurLink(tg, tgLink)
}
