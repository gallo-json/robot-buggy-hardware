package main

import (
	"os"
	"os/signal"
	"remote-buggy/utils/hardware"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	hardware.Setup()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		rpio.Close()
		os.Exit(1)
	}()
	for {
		hardware.Forward()
	}
}
