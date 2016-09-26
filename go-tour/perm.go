package main

import "fmt"
import "math/rand"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(rand.Perm(i))
	}
}
