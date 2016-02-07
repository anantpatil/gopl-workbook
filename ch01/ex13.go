package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// print time taken to print arguments

	var s, sep string

	start := time.Now()

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	nsecs1 := time.Since(start).Nanoseconds()
	fmt.Println("Time elapsed(ns)", nsecs1)

	// using strings join
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	nsecs2 := time.Since(start).Nanoseconds()
	fmt.Println("Time elapsed(ns)", nsecs2)

	factor := nsecs1 / nsecs2
	fmt.Printf("Join is %d times quicker\n", factor)
}
