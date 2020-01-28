package simulation

import (
	"fmt"
	"time"

	"github.com/wdevore/Deuron8-Go/configuration"
	"github.com/wdevore/Deuron8-Go/log"
)

// Entry is the main simulation entry point
func Entry(c chan string) {
	logger := log.New()

	logger.LogInfo("Beginning simulation")

	conf := configuration.New()

	fmt.Println("Err: " + conf.ErrLogFileName())

	loop := true

	for loop {
		select {
		case cmd := <-c:
			if cmd == "exit" {
				fmt.Println("Exiting sim")
				loop = false
			}
		default:
			// Perform a chunk of steps.

			fmt.Println("Simulating...")
			time.Sleep(1000 * time.Millisecond)
		}
	}
}
