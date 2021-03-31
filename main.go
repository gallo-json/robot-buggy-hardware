package main

import (
	"os"
	"os/signal"
	"remote-buggy/utils/hardware"
	"syscall"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		rpio.Close()
		os.Exit(1)
	}()
	hardware.Setup()

	for {
		hardware.Forward()
	}
}
