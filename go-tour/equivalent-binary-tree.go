package main

import (
	"fmt"
	"golang.org/x/tour/tree" // provides a tree impl
)

// Tree struct for reference
// type Tree struct {
// 	Left *Tree
//	Value int
// 	Right *Tree
// }

func Walk(t *tree.Tree, ch chan int) {
	// send elements of tree to chan
	// send in sorted order so that equivalence can be tested
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	// compare two trees and return true or false
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		x, y := <-ch1, <-ch2
		if x != y {
			return false
		}
	}
	return true
}

func main() {
	// tree.New(k) returns a random tree of size 10 with elements
	// 1k, 2k, 3k...10k

	// first test walk by making a tree and printing it,
	// the output should be sorted 1k, 2k..10k
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false
}
