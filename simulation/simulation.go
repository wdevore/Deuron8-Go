package simulation

import (
	"fmt"
	"time"

	"github.com/wdevore/Deuron8-Go/configuration"
	"github.com/wdevore/Deuron8-Go/interfaces"
	"github.com/wdevore/Deuron8-Go/log"
)

// Logger is the main logger for the simulation
var Logger interfaces.ILogger

// Config is the runtime configuration
var Config interfaces.IConfig

func init() {
	Config = configuration.New()
	Logger = log.New(Config)
}

// Entry is the main simulation entry point
func Entry(c chan string) {

	Logger.LogInfo("Beginning simulation")

	fmt.Println("Err: " + Config.ErrLogFileName())

	Config.SetExitState("Terminated")
	Config.Save()
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
