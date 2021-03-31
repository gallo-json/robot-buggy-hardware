package hardware

import (
	"github.com/stianeikeland/go-rpio/v4"

	"log"
)

var clockwiseA rpio.Pin
var clockwiseB rpio.Pin
var counterClockwiseA rpio.Pin
var counterClockwiseB rpio.Pin

// Setup Configures the analog ports on the RaspberryPI/motor controllers
func Setup() {
	err := rpio.Open()

	clockwiseA = rpio.Pin(19)
	clockwiseB = rpio.Pin(6)
	counterClockwiseA = rpio.Pin(13)
	counterClockwiseB = rpio.Pin(5)

	clockwiseA.Output()
	clockwiseB.Output()
	counterClockwiseA.Output()
	counterClockwiseB.Output()

	if err != nil {
		log.Fatal(err)
		rpio.Close()
	}
}

// Forward moves the car forward
func Forward() {
	clockwiseA.High()
	clockwiseB.High()
}

// Backward moves the car backward
func Backward() {
	counterClockwiseA.High()
	counterClockwiseB.High()
}

// Left moves the car left
func Left() {
	clockwiseA.High()
	counterClockwiseB.High()
}

// Right moves the car right
func Right() {
	clockwiseB.High()
	counterClockwiseA.High()
}
