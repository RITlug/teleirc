// Package main contains all logic relating to running TeleIRC
package main

import (
	"flag"
	"fmt"
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
	flagDebug   = flag.Bool("debug", false, "enable debugging")
	flagVersion = flag.Bool("version", false, "displays current version of TeleIRC")
)

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Printf("Current TeleIRC version: %s\n", version)
		return
	}

	// TODO: Build out debugging capabilities for more verbose output
	if *flagDebug {
		fmt.Printf("Debug mode currently set to: %t\n", *flagDebug)
	}

	settings, err := internal.LoadConfig(*flagPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var tgapi *tgbotapi.BotAPI
	tgClient := tg.NewClient(settings.Telegram, tgapi)
	tgChan := make(chan error)

	ircClient := irc.NewClient(settings.IRC)
	ircChan := make(chan error)
	go ircClient.StartBot(ircChan)
	go tgClient.StartBot(tgChan, ircClient.SendMessage)

	select {
	case ircErr := <-ircChan:
		fmt.Println(ircErr)
	case tgErr := <-tgChan:
		fmt.Println(tgErr)
	}
}
