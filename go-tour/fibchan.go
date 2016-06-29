package main

import "fmt"

func fib(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	// always sender closes the channel to tell reciever (range operator)
	// that the sender will no longer send any values. This is not mandatory
	// like file close
	close(ch)
}

func main() {
	ch := make(chan int, 10)
	// cap give the buff capacity of channel
	go fib(cap(ch), ch)
	// range recieves values from channel until it is closed
	// range is reciever side of channel
	for x := range ch {
		fmt.Println(x)
	}
}
