package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

/*
GetLargestPhoto takes a slice of PhotoSize objects and returns the instance with the
largest file size. This is used to upload the highest-resolution photo to a photo-
hosting service like Imgur, but could be used for other purposes.
*/
// func GetLargestPhoto(u tgbotapi.Message) *tgbotapi.PhotoSize {
// 	var largestPhoto = new(tgbotapi.PhotoSize)
// 	var photos := *u.Message.Photo
// 	for _, img := range photos {
// 		if img.FileSize > largestPhoto.FileSize {
// 			largestPhoto = img
// 			continue
// 		}
// 	}
// 	return largestPhoto
// }

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
	return u.FirstName + " (@" + u.UserName[:1] + "​" + u.UserName[1:] + ")"
}

/*
ZwspUsername adds a zero-width space after the first character of a Telegram user's
username.
*/
func ZwspUsername(u *tgbotapi.User) string {
	// Add ZWSP to prevent pinging across platforms
	// See https://github.com/42wim/matterbridge/issues/175
	return u.UserName[:1] + "​" + u.UserName[1:]
}
