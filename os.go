package main

import (
	"fmt"
	"os"
	"os/signal"
)

func reportSignals() {
	c := make(chan os.Signal, 10)

	signal.Notify(c)

	for s := range c {
		fmt.Println(s)
	}
}
