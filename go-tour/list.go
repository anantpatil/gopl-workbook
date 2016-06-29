// implement simple list in golang

package main

type Node struct {
	value interface{}
	next Node
}

func (n *Node) append(val int) Node {
	next = Node{value: val}
	n.next = next
	return next
}

func main() {
	head := Node{0}
	tail := head
	for i := 1; i <= 5; i++ {
		tail = tail.append(i)
	}
}
