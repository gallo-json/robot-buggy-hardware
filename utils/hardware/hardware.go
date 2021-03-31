package hardware

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

var clockwiseA rpio.Pin
var clockwiseB rpio.Pin
var counterClockwiseA rpio.Pin
var counterClockwiseB rpio.Pin

// Setup configures the analog ports on the RaspberryPI/motor controllers
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

// Backward moves the car backward
func Backward() {
	clockwiseA.High()
	counterClockwiseA.Low()
	clockwiseB.High()
	counterClockwiseB.Low()
}

// Forward moves the car forward
func Forward() {
	clockwiseA.Low()
	counterClockwiseA.High()
	clockwiseB.Low()
	counterClockwiseB.High()
}

// Left moves the car left
func Left() {
	clockwiseA.High()
	counterClockwiseA.Low()
	clockwiseB.Low()
	counterClockwiseB.High()
}

// Right moves the car right
func Right() {

	clockwiseA.Low()
	counterClockwiseA.High()
	clockwiseB.High()
	counterClockwiseB.Low()
}
