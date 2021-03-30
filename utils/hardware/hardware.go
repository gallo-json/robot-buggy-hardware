package hardware

import (
	"github.com/stianeikeland/go-rpio/v4"

	"log"
)

func Setup() {
	err := rpio.Open()
	if err != nil {
		log.Fatal(err)
	}
}


