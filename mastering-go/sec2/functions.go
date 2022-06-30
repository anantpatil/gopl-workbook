package main

import "fmt"

// functions as user defined type
type action func(a, b int) int

func main4() {
	// function literal
	add2 := func(a, b int) int {
		return a + b
	}

	fmt.Println("add2: ", add2(3, 5))
	fmt.Println("Action: ", takeAction(5, add2))
}

func takeAction(x int, f action) int {
	x = f(x, x)
	return x
}
