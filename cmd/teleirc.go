package main

import (
	"fmt"
	"time"

	"github.com/ritlug/teleirc/internal"
)

func main() {
	settings, err := internal.LoadConfig("../env.example")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(settings)
	}
	client := internal.NewClient(settings)
	client.AddHandlers(settings)
	go client.Connect()
	dur, _ := time.ParseDuration("10s")
	time.Sleep(dur)
	for {
		if client.IsConnected() {
			client.Cmd.Join(settings.IRC.Channel)
			break
		}
	}
	for {
		time.Sleep(dur)
	}
}
