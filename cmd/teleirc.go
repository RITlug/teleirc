// Package main contains all logic relating to running TeleIRC
package main

import (
	"flag"
	"fmt"
	"github.com/ritlug/teleirc/internal"
	"github.com/ritlug/teleirc/internal/handlers/irc"
	tg "github.com/ritlug/teleirc/internal/handlers/telegram"
)

func startIrc() {
	ircbot, err := irc.StartBot()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ircbot)
}

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
	flag.StringVar(&path, "p", "../.env", "Path to .env")

	flag.Parse()

	// TODO: change _ to settings once used
	_, err := internal.LoadConfig(path)
	if err != nil {
		fmt.Println(err)
	} else {
		startIrc()
		startTelegram()
	}
}
