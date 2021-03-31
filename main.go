package main

import (
	"remote-buggy/utils/hardware"
)

func main() {
	for {
		hardware.Forward()
	}
}
