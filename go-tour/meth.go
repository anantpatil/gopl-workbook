package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	// small case also works
	X, y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.y*v.y)
}

// same as Abs but Abs is a function and this becomes method
// the v Vertex here is reciever and the func can refer to it as v.<field>
func (v Vertex) Something() float64 {
	return (v.X + v.y) / 2
}

// pass the pointer to object as reciever so that the object state can be
// changed
func (v *Vertex) Transform() {
	// change the vertex in place
	v.X *= 2
	v.y *= 2
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
	fmt.Println(v.Something())
	v.Transform()
	fmt.Println(Abs(v))
	fmt.Println(v.Something())
}
