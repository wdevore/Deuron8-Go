package main

import (
	"bufio"
	"fmt"
	"github.com/wdevore/Deuron8-Go/simulation"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Deuron8 Go edition")

	reader := bufio.NewReader(os.Stdin)
	printHelp()

	end := false

	ch := make(chan string)

	for !end {
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "q":
			ch <- "exit"
			end = true
		case "r":
			go simulation.Entry(ch)
		default:
			fmt.Println("*********************")
			fmt.Println("** Unknown command **")
			fmt.Println("*********************")
		}
		// if strings.Compare("hi", text) == 0 {
		// 	fmt.Println("hello, Yourself")
		// }

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
	fmt.Println("  s: reset simulation")
	fmt.Println("  a: status of simulation")
	fmt.Println("  h: this help menu")
	fmt.Println("-----------------------------")
	fmt.Print("> ")
}
