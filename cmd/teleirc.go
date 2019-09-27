package main

import (
	"fmt"

	"github.com/ritlug/teleirc/internal"
)

func main() {
	settings, err := internal.LoadConfig("../.env")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(settings)
	}
	client := internal.NewClient(settings)
	client.AddHandlers()
	if err := client.Connect(); err != nil {
		fmt.Println(err)
	}
}
