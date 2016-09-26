package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	go say("world")   // runs in another go routine
	go say("namaste") // runs in another go routine

	// following runs in current routine; main has its own go routine
	// and as long as it is running, the run time willnot exit
	say("hello")
}
