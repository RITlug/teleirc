package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUsername(t *testing.T) {
	username := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := username.FirstName + " (@" + username.UserName + ")"
	name := GetUsername(username)

	assert.Equal(t, correct, name)
}

func TestGetNoUsername(t *testing.T) {
	username := &tgbotapi.User{ID: 1, FirstName: "John"}
	correct := username.FirstName
	name := GetUsername(username)

	assert.Equal(t, correct, name)
}
