// implement simple list in golang

package main

import "fmt"

type List struct {
	val  interface{}
	next *List
}

func (head *List) hasCycle() bool {
	// see if there is a cycle in the list
	// run two pointers stepping 1 and 2 steps at a time and they should
	// meet or the the faster should exit
	first, second := head, head
	for second.next != nil && second.next.next != nil {
		first, second = first.next, second.next.next
		if second == first {
			return true
		}

	}
	return false
}

func (head *List) Walk() {
	//print values in list
	for ; head != nil; head = head.next {
		fmt.Printf("%v, ", head.val)
	}
	fmt.Println()
}

func (head *List) Add(val interface{}) {
	t := head
	for t.next != nil {
		t = t.next
	}
	t.next = &List{val: val}
}

func cycleEntryPoint(head *List) *List {
	// find the node that is entrypoint for cycle if there is a cycle,
	// else return nil
	slow, fast := head, head
	for fast.next != nil && fast.next.next != nil {
		slow, fast = slow.next, fast.next.next
		if slow == fast {
			// there is a cycle, and we need to find the entrypoint
			first := head
			for first != fast {
				first, fast = first.next, fast.next
			}

			return fast
		}
	}

	return nil
}

func Reverse(head *List) *List {
	c, n := head, head.next
	for n != nil {
		t := n.next
		n.next = c
		c, n = n, t
	}
	head.next = nil
	return c
}

func main() {
	list := List{val: 0}
	for i := 1; i <= 5; i++ {
		list.Add(i)
	}
	list.Walk()
	fmt.Println(list.hasCycle())
	reverse := Reverse(&list)
	fmt.Println("Reverse: ")
	reverse.Walk()

	// make a cycle
	h := &list
	for ; h.next != nil; h = h.next {

	}
	h.next = &list
	// fmt.Println(list.hasCycle())
	if list.hasCycle() {
		entry := cycleEntryPoint(&list)
		fmt.Println(entry.val)
	}
}
