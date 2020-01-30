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

func init() {
	config.API = config.New()
	logg.API = logg.New(config.API)
}

func main() {
	fmt.Println("Welcome to Deuron8 Go edition")

	defer config.API.Save()
	defer logg.API.Close()

	reader := bufio.NewReader(os.Stdin)
	printHelp()

	end := false

	ch := make(chan string)
	// Launch simulation thread. It will idle by default.
	go simulation.Entry(ch)

	for !end {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "q":
			// send(ch, "exit")
			ch <- "exit"
			end = true
		case "r":
			// go send(ch, "run")
			ch <- "run"
		case "p":
			// go send(ch, "pause")
			ch <- "pause"
		case "u":
			// go send(ch, "resume")
			ch <- "resume"
		case "s":
			// go send(ch, "reset")
			ch <- "reset"
		case "t":
			// go send(ch, "stop")
			ch <- "stop"
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

func send(ch chan string, msg string) {
	ch <- msg
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
