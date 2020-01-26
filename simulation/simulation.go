package simulation

import "time"

import "fmt"

var loop = true

// Entry is the main simulation entry point
func Entry(c chan string) {
	fmt.Println(("Beginning simulation"))
	for loop {
		select {
		case cmd := <-c:
			if cmd == "exit" {
				fmt.Println("Exiting sim")
				loop = false
			}
			fmt.Println(("cmd: " + cmd))
		default:
			fmt.Println(("Simulating..."))
			time.Sleep(1000 * time.Millisecond)
		}
	}

	fmt.Println(("Entry done."))
}
