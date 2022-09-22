package main

import (
	"fmt"
	"os"
	"wire/msg"
)

func main() {
	e, err := msg.InitializeEvent("hi there!")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
