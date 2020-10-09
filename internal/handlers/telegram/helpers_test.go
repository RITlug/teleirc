package telegram

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
)

func TestGetFullUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := user.FirstName + " (@" + user.UserName + ")"
	name := GetFullUsername(false, user)

	assert.Equal(t, correct, name)
}

func TestGetFullUserZwsp(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := user.FirstName + " (@" + user.UserName[:1] +
		"​" + user.UserName[1:] + ")"
	name := GetFullUsername(true, user)

	assert.Equal(t, correct, name)
}

func TestGetFullNoUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John"}
	correct := user.FirstName
	name := GetFullUsername(false, user)

	assert.Equal(t, correct, name)
}

func TestGetNoUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John"}
	correct := user.FirstName
	name := GetFullUsername(false, user)

	assert.Equal(t, correct, name)
}

func TestGetUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := user.UserName
	name := GetUsername(false, user)

	assert.Equal(t, correct, name)
}

func TestZwspUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := "j" + "​" + "smith"
	name := GetUsername(true, user)

	assert.Equal(t, correct, name)
}
