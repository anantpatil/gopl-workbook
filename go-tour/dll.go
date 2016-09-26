package main

import (
	"fmt"
	"math/rand"
)

// Node in a doubly-linked list
type Node struct {
	val        interface{}
	prev, next *Node
}

func (head *Node) Append(val interface{}) *Node {
	// appand val to liest
	n := &Node{val: val}
	if head == nil {
		return n
	}
	t := head
	for t.next != nil {
		t = t.next
	}
	t.next, n.prev = n, t
	return head
}

func Walk(head *Node) {
	for t := head; t != nil; t = t.next {
		fmt.Printf("%v ", t.val)
	}
	fmt.Println()
}

func Reverse(head *Node) *Node {
	// return head of reverse list
	return nil
}

func HasCycle(head *Node) bool {
	// return False if there the list doesn't have a cycle
	return false
}

func CycleEntryPoint(head *Node) *Node {
	// return pointer to entry point of cycle in list
	return nil
}

func main() {
	// test
	var head *Node
	l := rand.Perm(10)
	fmt.Println(l)
	for _, v := range l {
		head = head.Append(v)
	}
	// just print
	Walk(head)
}
