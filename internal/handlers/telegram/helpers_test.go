package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"testing"
)

func TestGetFullUsername(t *testing.T) {
	username := &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"}
	correct := username.FirstName + " (@" + username.UserName + ")"
	if name := GetUsername(username); name != correct {
		t.Errorf("Username was incorrect, got: %s, want: %s.", name, correct)
	}
}

func TestGetUserNoUsername(t *testing.T) {
	username := &tgbotapi.User{ID: 1, FirstName: "John"}
	correct := username.FirstName
	if name := GetUsername(username); name != correct {
		t.Errorf("FirstName was incorrect, got: %s, want: %s.", name, correct)
	}
}
