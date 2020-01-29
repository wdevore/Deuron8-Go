package simulation

import (
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

var paused = false
var debug = 0

// Entry is the main simulation entry point
func Entry(c chan string) {
	debug = 0

	// fmt.Println("Err: " + Config.ErrLogFileName())

	loop := true
	run := false

	for loop {
		select {
		case cmd := <-c:
			switch cmd {
			case "exit":
				Config.SetExitState("Terminated")
				Logger.LogInfo("Exiting simulation...")
				loop = false
			case "run":
				Logger.LogInfo("Simulation starting...")
				Logger.LogInfo("Simulation running...")
				run = true
			case "pause":
				Config.SetExitState("Paused")
				Logger.LogInfo("Simulation paused")
				paused = true
			case "resume":
				Logger.LogInfo("Simulation resumed")
				paused = false
			case "reset":
				Logger.LogInfo("Simulation resetting...")
				Logger.LogInfo("Simulation reset and paused")
				paused = true
			case "stop":
				Config.SetExitState("Stopped")
				Logger.LogInfo("Stopping simulation...")
				run = false
				loop = false
			}
		default:
			if !run {
				time.Sleep(10 * time.Millisecond)
				continue
			}

			// A simulation runs in chunks. This gives the "application" a chance
			// to check on user input. A chunk size depends on how long a group
			// of simulation steps take.
			if paused {
				Logger.LogInfo("-- paused --")
			} else {
				run = simulate()
				if !run {
					Logger.LogInfo("Simulation has completed and stopped")
					debug = 0
				}
			}

			time.Sleep(1000 * time.Millisecond)
		}
	}

	Logger.LogInfo("Simulation goroutine is exiting...")

	Config.Save()
}

func simulate() bool {
	debug++
	Logger.LogInfo("Simulating...")
	if debug > 5 {
		Config.SetExitState("Completed")
		return false
	}

	return true
}
