package telegram

import (
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
)

func TestGetFullUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := user.FirstName + " (@" + user.UserName + ")"
	name := GetFullUsername(user)

	assert.Equal(t, correct, name)
}

func TestGetNoUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John"}
	correct := user.FirstName
	name := GetFullUsername(user)

	assert.Equal(t, correct, name)
}

func TestGetUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := user.UserName
	name := GetUsername(user)

	assert.Equal(t, correct, name)
}

func TestZwspUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := "j" + "​" + "smith"
	name := ZwspUsername(user)

	assert.Equal(t, correct, name)
}

func TestZwspNoUsername(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John"}
	correct := user.FirstName
	name := ZwspUsername(user)

	assert.Equal(t, correct, name)
}

func TestGetFullUserZwsp(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := user.FirstName + " (@" + user.UserName[:1] +
		"​" + user.UserName[1:] + ")"
	name := GetFullUserZwsp(user)

	assert.Equal(t, correct, name)
}

func TestGetFullNoUserZwsp(t *testing.T) {
	user := &tgbotapi.User{ID: 1, FirstName: "John"}
	correct := user.FirstName
	name := GetFullUserZwsp(user)

	assert.Equal(t, correct, name)
}
