package main

import (
	"fmt"

	"github.com/ritlug/teleirc/internal"
	"github.com/ritlug/teleirc/internal/handlers/irc"
)

func main() {
	settings, err := internal.LoadConfig("../.env")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(settings)
	}
	client := irc.NewClient(settings)
	if err := client.StartBot(); err != nil {
		fmt.Println(err)
	}
}
