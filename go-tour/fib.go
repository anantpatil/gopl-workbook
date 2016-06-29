package main

import "fmt"

func fibonacci() func() uint {
	var x, y uint = 0, 1
	f := func() uint {
		fib := x
		x, y = y, x+y
		return fib
	}
	return f
}

func main() {
	fib := fibonacci()
	for i := 0; i <= 25; i++ {
		fmt.Println(fib())
	}
}
