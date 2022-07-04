package main

import "fmt"

func printHi(ch1 chan bool, ch2 chan bool) {
	for {
		<-ch1
		fmt.Println("hi")
		ch2 <- true
	}
}

func printHello(ch1 chan bool, ch2 chan bool) {
	for {
		<-ch2
		fmt.Println("hello")
		ch1 <- true
	}
}

func main2() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go printHi(ch1, ch2)
	go printHello(ch1, ch2)
	ch1 <- true
	quit := make(chan struct{})
	<-quit
}
