package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// Using strings library
	fmt.Println(strings.Join(os.Args, " "))

	// use fmt's built-in mechanism to print an array
	fmt.Println(os.Args)
}
