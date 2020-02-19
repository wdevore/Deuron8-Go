package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/wdevore/Deuron8-Go/config"
	logg "github.com/wdevore/Deuron8-Go/log"
	"github.com/wdevore/Deuron8-Go/simulation"
)

const configFile = "config/config.json"

func init() {
	config.API = config.New(configFile)
	logg.API = logg.New(config.API)
}

func main() {
	fmt.Println("Welcome to Deuron8 Go edition")

	defer logg.API.Close()
	defer config.API.Save() // Last-in means First-out, this will run before logg-close.

	reader := bufio.NewReader(os.Stdin)
	printHelp()

	end := false

	ch := make(chan string)
	// Start simulation thread. It will idle by default.
	go simulation.Boot(ch)

	for !end {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "q":
			ch <- "exit"
			end = true
		case "r":
			ch <- "run"
		case "p":
			ch <- "pause"
		case "u":
			ch <- "resume"
		case "s":
			ch <- "reset"
		case "t":
			ch <- "stop"
		case "a":
			ch <- "status"
		default:
			fmt.Println("*********************")
			fmt.Println("** Unknown command **")
			fmt.Println("*********************")
		}

		if !end {
			printHelp()
		}
	}

	fmt.Println("Goodbye.")
}

func printHelp() {
	fmt.Println("-----------------------------")
	fmt.Println("Commands:")
	fmt.Println("  q: quit")
	fmt.Println("  r: run simulation")
	fmt.Println("  p: pause simulation")
	fmt.Println("  u: resume simulation")
	fmt.Println("  s: reset simulation. Sim is stopped after reset.")
	fmt.Println("  t: stop simulation")
	fmt.Println("  a: status of simulation")
	fmt.Println("  h: this help menu")
	fmt.Println("-----------------------------")
	fmt.Print("> ")
}
