// Package main contains all logic relating to running TeleIRC
package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ritlug/teleirc/internal"
	"github.com/ritlug/teleirc/internal/handlers/irc"
	tg "github.com/ritlug/teleirc/internal/handlers/telegram"
)

func startTelegram() {
	tgbot, err := tg.StartBot()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tgbot)
}

func main() {
	// optional path flag if user does not want .env file in root dir
	var path string
	flag.StringVar(&path, "p", ".env", "Path to .env")

	flag.Parse()

	settings, err := internal.LoadConfig(path)

	ircChan := make(chan error)
	waitDur, _ := time.ParseDuration("1s")

	if err != nil {
		fmt.Println(err)
	} else {
		startTelegram()
		client := irc.NewClient(settings.IRC)
		go client.StartBot(ircChan)
		for {
			select {
			case err := <-ircChan:
				fmt.Println(err)
			case <-time.After(waitDur):
			}
		}
	}
}
