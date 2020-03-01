// Package main contains all logic relating to running TeleIRC
package main

import (
	"flag"
	"os"
	"log"
	"io/ioutil"

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

	logFlags    = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
	Info        = log.New(os.Stdout, "INFO: ", logFlags)
	Warning     = log.New(os.Stdout, "WARNING: ", logFlags)
	Error       = log.New(os.Stderr, "ERROR: ", logFlags)
	Ignored     = log.New(ioutil.Discard, "ERROR: ", logFlags)
)

func main() {
	flag.Parse()

	if *flagVersion {
		Info.Printf("Current TeleIRC version: %s\n", version)
		return
	}

	// TODO: Build out debugging capabilities for more verbose output
	if *flagDebug {
		Info.Printf("Debug mode currently set to: %t\n", *flagDebug)
	}

	settings, err := internal.LoadConfig(*flagPath)
	if err != nil {
		Error.Println(err)
		os.Exit(1)
	}

	var tgapi *tgbotapi.BotAPI
	tgClient := tg.NewClient(settings.Telegram, tgapi)
	tgChan := make(chan error)

	ircClient := irc.NewClient(settings.IRC)
	ircChan := make(chan error)
	go ircClient.StartBot(ircChan, tgClient.SendMessage)
	go tgClient.StartBot(tgChan, ircClient.SendMessage)

	select {
	case ircErr := <-ircChan:
		Error.Println(ircErr)
	case tgErr := <-tgChan:
		Error.Println(tgErr)
	}
}
