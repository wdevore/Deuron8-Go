package simulation

import (
	"fmt"
	"time"

	"github.com/wdevore/Deuron8-Go/config"
	logg "github.com/wdevore/Deuron8-Go/log"
)

var paused = false
var completed = false

// Boot is the simulation bootstrap. The simulation isn't
// running until told to do so.
func Boot(c chan string) {
	debug = 0
	loop := true
	running := false

	for loop {
		select {
		case cmd := <-c:
			switch cmd {
			case "exit":
				if running {
					config.API.SetExitState("Terminated")
					logg.API.LogInfo("Terminating simulation...")
				} else {
					if !completed {
						config.API.SetExitState("Exited")
					}
				}
				loop = false
			case "run":
				if running {
					// Starting a run while the sim is already running
					// isn't allowed.
					continue
				}
				logg.API.LogInfo("Simulation starting...")
				construct()

				logg.API.LogInfo("Simulation running...")
				running = true
				completed = false
			case "pause":
				if paused {
					continue
				}
				config.API.SetExitState("Paused")
				logg.API.LogInfo("Simulation paused")
				paused = true
			case "resume":
				if !paused {
					continue
				}
				logg.API.LogInfo("Simulation resumed")
				paused = false
			case "reset":
				logg.API.LogInfo("Simulation resetting...")
				logg.API.LogInfo("Simulation reset and paused")
				paused = true
				running = false
				completed = false
			case "stop":
				config.API.SetExitState("Stopped")
				logg.API.LogInfo("Stopping simulation...")
				running = false
				completed = false
				logg.API.LogInfo("Simulation stopped before completion")
			case "status":
				fmt.Printf("Status: %d\n", debug)
			}
		default:
			if !running {
				time.Sleep(10 * time.Millisecond)
				continue
			}

			if paused {
				logg.API.LogInfo("-- paused --")
			} else {
				running = simulate()
				if !running {
					config.API.SetExitState("Completed")
					completed = true
					logg.API.LogInfo("Simulation has completed and stopped")
					debug = 0
				}
			}

			time.Sleep(1000 * time.Millisecond)
		}
	}

	logg.API.LogInfo("Simulation goroutine is exiting...")
}
