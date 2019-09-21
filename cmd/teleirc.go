package main

import (
	"fmt"
	"github.com/ritlug/teleirc/internal"
)

func main() {
	settings, err := internal.LoadConfig("../env.example")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(settings)
	}
}
