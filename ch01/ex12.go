package main

import (
	"fmt"
	"os"
)

func main() {
	// print index and values of arguments
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
