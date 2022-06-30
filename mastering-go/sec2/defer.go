package main

import "fmt"

func main3() {
	fmt.Println("Counting")
	// defer will invoke in LIFO order
	for i := 1; i <= 4; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
