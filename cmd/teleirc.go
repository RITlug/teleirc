// Package main contains all logic relating to running TeleIRC
package main

import (
	"flag"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ritlug/teleirc/internal"
	"github.com/ritlug/teleirc/internal/handlers/irc"
	tg "github.com/ritlug/teleirc/internal/handlers/telegram"
)

const (
	version = "v2.0"
)

var (
	flagPath    = flag.String("conf", ".env", "config file")
	flagDebug   = flag.Bool("debug", false, "disable debugging")
	flagVersion = flag.Bool("version", false, "displays current version of TeleIRC")
)

func main() {
	flag.Parse()
	debug := internal.Debug{DebugLevel: *flagDebug}

	if *flagVersion {
		debug.PrintVersion("Current TeleIRC version:", version)
		return
	}

	debug.LogInfo("Current TeleIRC version:", version)
	// Notify that debug is enabled
	debug.LogDebug("Debug mode enabled!")

	settings, err := internal.LoadConfig(*flagPath)
	if err != nil {
		debug.LogError(err)
		os.Exit(1)
	}

	var tgapi *tgbotapi.BotAPI
	tgClient := tg.NewClient(settings.Telegram, tgapi, debug)
	tgChan := make(chan error)

	ircClient := irc.NewClient(settings.IRC, debug)
	ircChan := make(chan error)
	go ircClient.StartBot(ircChan, tgClient.SendMessage)
	go tgClient.StartBot(tgChan, ircClient.SendMessage)

	select {
	case ircErr := <-ircChan:
		debug.LogError(ircErr)
	case tgErr := <-tgChan:
		debug.LogError(tgErr)
	}
}
