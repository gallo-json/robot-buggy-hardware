package main

import "remote-buggy/utils/hardware"

func main() {
	hardware.Setup()
	for {
		hardware.Backward()
	}
}
