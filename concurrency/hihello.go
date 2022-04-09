package main

import (
	"fmt"
	"time"
)

func hi(ch chan int) {
	// this func will only produce value of 1 and add to the channel and
	// wait for it to be consumed
	for {
		fmt.Println("hi")
		time.Sleep(time.Second * 2)
		ch <- 1
		ch <- 1
	}
}

func hello(ch chan int) {
	// only consume value
	for {
		<-ch
		fmt.Println("hello ")
		time.Sleep(time.Second * 2)
		<-ch
	}
}

// hi1 and hello1 works when time.Sleep is removed, but doesn't work as expected
// with time.Sleep
func hi1(ch chan int) {
	// only produce values
	for {
		fmt.Println("Sending 1...")
		ch <- 1
		fmt.Println("hi1")
		time.Sleep(time.Second * 2)
		fmt.Println("Sending2 1...")
		ch <- 1
	}
}

func hello1(ch chan int) {
	// only consume value
	for {
		<-ch
		fmt.Println("Recving 1...")
		fmt.Println("hello1 ")
		time.Sleep(time.Second * 2)
		<-ch
		fmt.Println("Recving2 1...")
	}
}

func main() {
	// same as make(chan int) - unbuffered channel
	ch := make(chan int, 0)
	go hi1(ch)
	go hello1(ch)
	quit := make(chan int, 0)
	<-quit
}
