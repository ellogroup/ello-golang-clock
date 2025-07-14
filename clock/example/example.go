package main

import (
	"fmt"
	"github.com/ellogroup/ello-golang-clock/clock"
	"time"
)

func main() {
	// Returns a system clock
	system := clock.NewSystem()
	fmt.Printf("System time: %s\n\n", system.Now().String())

	// Returns a fixed clock
	fixed := clock.NewFixed(time.Date(2030, 1, 2, 12, 34, 56, 0, time.UTC))
	fmt.Printf("Fixed time: %s\n\n", fixed.Now().String())
}
