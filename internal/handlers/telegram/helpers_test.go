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
	tests := []struct {
		name     string
		user     *tgbotapi.User
		expected string
	}{
		{
			name:     "ascii",
			user:     &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"},
			expected: "John (@j\u200bsmith)",
		},
		{
			name:     "cyrillic",
			user:     &tgbotapi.User{ID: 1, FirstName: "Иван", UserName: "иван"},
			expected: "Иван (@и\u200bван)",
		},
		{
			name:     "japanese",
			user:     &tgbotapi.User{ID: 1, FirstName: "まこと", UserName: "まこと"},
			expected: "まこと (@ま\u200bこと)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GetFullUsername(true, test.user)
			assert.Equal(t, test.expected, actual)
		})
	}
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
	tests := []struct {
		name     string
		user     *tgbotapi.User
		expected string
	}{
		{
			name:     "ascii",
			user:     &tgbotapi.User{ID: 1, FirstName: "John", UserName: "jsmith"},
			expected: "j\u200bsmith",
		},
		{
			name:     "cyrillic",
			user:     &tgbotapi.User{ID: 1, FirstName: "Иван", UserName: "иван"},
			expected: "и\u200bван",
		},
		{
			name:     "japanese",
			user:     &tgbotapi.User{ID: 1, FirstName: "まこと", UserName: "まこと"},
			expected: "ま\u200bこと",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GetUsername(true, test.user)
			assert.Equal(t, test.expected, actual)
		})
	}
}
