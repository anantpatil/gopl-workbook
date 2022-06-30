package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {

	var v1 int //declare a variable
	v1 = 1     // assign the variable
	var v2 = 2 //declare and assign a variable (variable with intializers)
	v3 := 3    //declare and assign variable via type inference
	fmt.Println(v1 + v2 + v3)

	//A pointer holds the memory address of a variable

	var p1 *int //declare a pointer
	i := 4      //create an integer using type inference
	p1 = &i     //assign the address of the integer to the pointer
	fmt.Println(p1)

	p2 := &i         //create a pointer via type inference
	fmt.Println(*p2) //use the * in order to 'dereference' the pointer and get the original value

	n := 32
	f := float64(n)
	const pi = 3.14
	fmt.Println(pi * f)

	//arrays
	var a [2]int
	a[0], a[1] = 4, 5
	fmt.Println("Array a", a, len(a))

	b := [3]int{9, 8, 10}
	fmt.Println("Array b", b, len(b))

	// a slice: note it doesn't specify any size
	c := []int{}
	fmt.Printf("a: %T, b: %T, c: %T\n", a, b, c)
	c = append(c, 12, 13)
	c = append(c, 82)
	fmt.Println("Slice c", c, len(c))

	//functions pass by value, so we need the argument to be a pointer if we want to change the value of the passed variable
	X := 5
	ChangeX(&X)
	fmt.Println("Value of X: ", X)

	//function value
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(computeMultiplyVal(2, add))

	//function closer
	inc := incrementer()
	fmt.Println(inc(), inc(), inc(), inc())

	inc2 := incrementer()
	fmt.Println(inc2(), inc2(), inc2(), inc2())

	samples := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	rg := randomGenerator(samples)
	fmt.Println(rg(), rg(), rg(), rg(), rg())

	//if statement
	if i := inc(); i < 0 {
		fmt.Println("i is a negative number ")
	} else if i == 0 {
		fmt.Println("i is zero")
	} else {
		fmt.Println("i is a positive number")
	}

	//switch statements
	switch i := inc(); {
	case i < 0:
		fmt.Println("i is a negative number")
	case i == 0:
		fmt.Println("i is zero")
	default:
		fmt.Println("i is a positive number")
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		defer fmt.Println("Exiting function.... ")
		fmt.Println("Entering function")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)

		defer fmt.Println("Exiting function.... ")
		fmt.Println("Entering function")

	}

}

const (
	big   = 1 << 100
	small = big >> 99
)

func computeMultiplyVal(val int, fn func(a, b int) int) int {
	return val * fn(val, val)
}

func ChangeX(X *int) {
	*X = 10
}

// function closure - returns a function that is enclosed and has access
// to local variable of parent func
func incrementer() func() int {
	//intializes i
	i := 0
	return func() int {
		//the new value of it will be retained
		i++
		return i
	}
}

// the above could also be a generator
func randomGenerator(in []int) func() int {
	return func() int {
		r := rand.Intn(len(in)-0) + 0
		return in[r]
	}
}
