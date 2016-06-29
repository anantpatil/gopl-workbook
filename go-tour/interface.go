package main

import (
	"fmt"
	"math"
)

// An type with Abs() implemented
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (v MyFloat) Abs() float64 {
	if v < -1 {
		// since v is MyFloat, type cast to float64
		return float64(-v)
	}
	// since v is MyFloat, type cast to float64
	return float64(v)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var f MyFloat = math.Sqrt2
	var i Abser

	i = f
	fmt.Printf("%T(%v)\n", i, i)
	fmt.Printf("Abs of %v is %v\n", i, i.Abs())

	var v Vertex = Vertex{3.0, 4.0}
	// the following will say Vertex doesn't implement Abser
	//i = v // compile error, since the reciever is pointer for Abs()
	i = &v // this will go through
	fmt.Printf("%T(%v)\n", i, i)
	fmt.Printf("Abs of %v is %v\n", i, i.Abs())
}
