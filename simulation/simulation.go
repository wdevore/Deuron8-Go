package simulation

import (
	"time"

	"github.com/wdevore/Deuron8-Go/config"
	logg "github.com/wdevore/Deuron8-Go/log"
)

var paused = false
var completed = false

var debug = 0

// Entry is the main simulation entry point
func Entry(c chan string) {
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
					logg.API.LogInfo("Terminating Exiting simulation...")
				} else {
					if !completed {
						config.API.SetExitState("Exited")
					}
				}
				loop = false
			case "run":
				if running {
					continue
				}
				logg.API.LogInfo("Simulation starting...")
				logg.API.LogInfo("Simulation running...")
				running = true
				completed = false
			case "pause":
				config.API.SetExitState("Paused")
				logg.API.LogInfo("Simulation paused")
				paused = true
			case "resume":
				logg.API.LogInfo("Simulation resumed")
				paused = false
			case "reset":
				logg.API.LogInfo("Simulation resetting...")
				logg.API.LogInfo("Simulation reset and paused")
				paused = true
			case "stop":
				config.API.SetExitState("Stopped")
				logg.API.LogInfo("Stopping simulation...")
				running = false
				completed = false
				logg.API.LogInfo("Simulation stopped before completion")
			}
		default:
			if !running {
				time.Sleep(10 * time.Millisecond)
				continue
			}

			// A simulation runs in chunks. This gives the "application" a chance
			// to check on user input. A chunk size depends on how long a group
			// of simulation steps take.
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

func simulate() bool {
	debug++
	if debug > 5 {
		return false
	}
	logg.API.LogInfo("Simulating...")

	return true
}
