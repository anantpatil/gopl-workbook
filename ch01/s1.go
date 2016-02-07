// This is go doc starts before package statement

// package main in special, it is used to build executables
package main

// package is followed by import statements

import "fmt"
import "os"

// then functions, variables, constants and types follow; introduced by
// func, var, const and type respectively. No order is imposed for
// these.

// function main is special, it is entry point for executable
func main() {
	var s, sep string
	// with i from 0 prints program name and its arguments
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
