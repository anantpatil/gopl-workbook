package main

import (
	"fmt"
)

func main() {
	// the following doesn't work
	// m := map[int]string
	
	var m map[int]string
	// m is a nil map; no entries can be added, but it's size is zero
	v, ok := m[1]
	fmt.Println(v, ok) // nil and False
	fmt.Println(len(m))
	// no entries can be added to nil map
	// m[1] = "one" // error
	
	// make a map and then start using it
	m = make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	v, ok = m[1]
	fmt.Println(v, ok)
	fmt.Println(len(m))
}
