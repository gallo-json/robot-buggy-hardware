package main

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
)

// func main() {
// 	hardware.Setup()
// 	for {
// 		hardware.Forward()
// 	}
// }

var speedA rpio.Pin
var speedB rpio.Pin
var clockwiseA rpio.Pin
var clockwiseB rpio.Pin
var counterClockwiseA rpio.Pin
var counterClockwiseB rpio.Pin

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	speedA = rpio.Pin(26)
	speedB = rpio.Pin(11)
	clockwiseA = rpio.Pin(19)
	clockwiseB = rpio.Pin(6)
	counterClockwiseA = rpio.Pin(13)
	counterClockwiseB = rpio.Pin(5)

	speedA.Output()
	speedB.Output()
	clockwiseA.Output()
	clockwiseB.Output()
	counterClockwiseA.Output()
	counterClockwiseB.Output()

	for x := 0; x < 20; x++ {
		Forward()
	}
}

func Forward() {
	for i := 0; i < 5; i++ {
		clockwiseA.Low()
		clockwiseB.Low()
	}
}
