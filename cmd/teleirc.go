// Package main contains all logic relating to running TeleIRC
package main

import (
	"flag"
	"fmt"

	"github.com/ritlug/teleirc/internal"
	"github.com/ritlug/teleirc/internal/handlers/irc"
	tg "github.com/ritlug/teleirc/internal/handlers/telegram"
)

var (
	version = "v2.0"

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
		fmt.Println(err)
		return
	}

	tgClient, _ := tg.NewClient(settings.Telegram)
	tgChan := make(chan error)
	go tgClient.StartBot(tgChan)

	ircClient := irc.NewClient(settings.IRC)
	ircChan := make(chan error)
	go ircClient.StartBot(ircChan)

	select {
	case ircErr := <-ircChan:
		fmt.Println(ircErr)
	case tgErr := <-tgChan:
		fmt.Println(tgErr)
	}
}
