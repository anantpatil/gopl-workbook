package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered channel which can hold two values
	ch <- 1 // doesn't block, buff not full
	ch <- 2 // doesn't block, buff not full
	// ch <- 3 // blocks and error since nothing is there to recieve the values
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
