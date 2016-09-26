package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
)

type Tree struct {
	val         int
	left, right *Tree
}

func (root *Tree) Walk(ch chan int) {
	// Walk in InOrder for sorted
	walk(root, root, ch)
}

func walk(t *Tree, root *Tree, ch chan int) {
	// print t in order
	if t == nil {
		return
	}

	walk(t.left, root, ch)

	if ch != nil {
		ch <- t.val
	} else {
		fmt.Println(t.val)
	}

	walk(t.right, root, ch)

	if t == root && ch != nil {
		close(ch)
	}
}

func (t *Tree) Serialize(writer io.Writer) {
	ch := make(chan int)
	go t.WalkPreOrder(ch)
	for x := range ch {
		writer.Write([]byte(strconv.Itoa(x)))
	}
}

func (root *Tree) WalkPreOrder(ch chan int) {
	_walkPreOrder(root, root, ch)
}

func _walkPreOrder(t *Tree, root *Tree, ch chan int) {
	if t == nil {
		return
	}

	if ch != nil {
		ch <- t.val
	} else {
		fmt.Println(t.val)
	}

	_walkPreOrder(t.left, root, ch)

	_walkPreOrder(t.right, root, ch)

	if t == root && ch != nil {
		close(ch)
	}

}

func Add(t *Tree, val int) *Tree {
	if t == nil {
		t = new(Tree)
		t.val = val
		return t
	}
	if val < t.val {
		t.left = Add(t.left, val)
	} else {
		t.right = Add(t.right, val)
	}
	return t
}

func main() {
	l := rand.Perm(10)
	fmt.Println(l)
	var root *Tree
	for _, v := range l {
		root = Add(root, v)
	}
	ch := make(chan int)
	go root.WalkPreOrder(ch)
	for v := range ch {
		fmt.Println(v)
	}
	writer := bufio.NewWriter(os.Stdout)
	root.Serialize(writer)
}
