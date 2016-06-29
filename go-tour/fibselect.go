package main

import "fmt"

func fib(ch chan int, quit chan int) {
	// send over ch and when you recieve anything over quit, return
	x, y := 0, 1
	// you should run in for-ever loop, as select is executed exactly once
	for {
		select {
		// selects one of the cases when they become ready
		case ch <- x:
			// x was sent over ch
			x, y = y, x+y
		case <-quit:
			// recieved from quit, time to return back
			fmt.Println("Quitting...")
			return
		}
	}
}

func main() {
	ch := make(chan int)   // blocking
	quit := make(chan int) // another blocking
	// write one func to recieve from ch and send on quit when done
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	fib(ch, quit)
}
