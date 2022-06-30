package main

import "fmt"

type SLLNode struct {
	next  *SLLNode
	value int
}

func NewSLLNode(val int) *SLLNode {
	n := new(SLLNode)
	n.value = val
	return n
}

func Walk(head *SLLNode) <-chan int {
	send := make(chan int)
	go func() {
		defer close(send)
		for curr := head; curr != nil; curr = curr.next {
			send <- curr.value
		}
	}()
	return send
}

func main() {
	head := NewSLLNode(0)
	curr := head
	for i := 1; i < 10; i++ {
		n := NewSLLNode(i)
		curr.next = n
		curr = n
	}

	list := Walk(head)
	for i := range list {
		fmt.Println(i)
	}
}
