package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	boom := time.Tick(5 * time.Second)
	for {
		select {
		case <-tick:
			// a tick
			fmt.Println("tick")
		case <-boom:
			// boom, done
			fmt.Println("BOOM")
			return
		default:
			// anything else comes here
			fmt.Println("\t.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
