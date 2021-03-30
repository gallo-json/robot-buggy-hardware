package hardware

import (
	"github.com/stianeikeland/go-rpio/v4"

	"log"
)

var speedA rpio.Pin
var speedB rpio.Pin
var clockwiseA rpio.Pin
var clockwiseB rpio.Pin
var counterClockwiseA rpio.Pin
var counterClockwiseB rpio.Pin

// Setup Configures the analog ports on the RaspberryPI/motor controllers
func Setup() {
	err := rpio.Open()

	speedA = rpio.Pin(26)
	speedB = rpio.Pin(11)
	clockwiseA = rpio.Pin(19)
	clockwiseB = rpio.Pin(6)
	counterClockwiseA = rpio.Pin(13)
	counterClockwiseB = rpio.Pin(5)
	if err != nil {
		log.Fatal(err)
	}
}

// Forward moves the car forward
func Forward() {
	for i := 0; i < 5; i++ {
		clockwiseA.Low()
		clockwiseB.Low()
	}
}
