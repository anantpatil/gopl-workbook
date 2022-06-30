package main

import "fmt"

func main7() {
	n := 1
	// A for loop
	for i := 10; i > 0; i-- {
		n *= i
	}
	fmt.Println("Result: ", n)

	n = 1
	// A Go While
	for n <= 50 {
		n += n
	}
	fmt.Println("Result: ", n)

	// A Do-While loop?
	n = 1
	for {
		n += n
		if n >= 50 {
			break
		}
	}
	fmt.Println("Result: ", n)

	// Another way
	n = 1
	for ok := true; ok; ok = n < 50 {
		n += n
	}

	fmt.Println("Result: ", n)
}
