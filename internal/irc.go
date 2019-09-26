package internal

import "github.com/lrstanley/girc"

func NewClient(settings *Settings) *girc.Client {
	client := girc.New(girc.Config{
		Server: settings.IRC.Server,
		Port:   settings.IRC.Port,
		Nick:   settings.IRC.BotName,
		User:   settings.IRC.BotName,
	})
	return client
}
