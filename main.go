package main

import (
	"remote-buggy/utils/hardware"
	"time"
)

func main() {
	hardware.Setup()
	
	hardware.Forward()
	time.Sleep(5)
	hardware.Right()
	time.Sleep(3)
	hardware.Backward()
	time.Sleep(5)
	hardware.Left()
	time.Sleep(3)
	hardware.Stop()
}
