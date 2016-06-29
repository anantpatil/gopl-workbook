package main

import (
	"fmt"
)

// note that chan has a type int: <name> chan <type>
func sum(a []int, ch chan int) {
	total := 0
	for _, x := range a {
		total += x
	}
	ch <- total // send result to channel
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	// in go f(x, y, z) the evaluation of f, x, y, z happens in current
	// goroutine and the execution in another. The following index
	// computation happens in main goroutine and execution of sum happens
	// in another
	go sum(a[:len(a)/2], ch) // do half of work
	go sum(a[len(a)/2:], ch) // do the remaining half of work
	// the following is blocked till the two goroutines are done
	x, y := <-ch, <-ch
	fmt.Println(x, y, x+y)
}
