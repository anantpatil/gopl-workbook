package main

import (
	"fmt"
)

func main2() {
	a := []int{0, 1, 5, 13, 26, 47, 79, 81, 99, 106, 110}
	target := 110
	v := MyBinarySearchIterative(a, target)
	fmt.Printf("%d is at %d\n", target, v)

}

func MyBinarySearch(in []int, target int) int {
	return _bs(in, target, 0, len(in)-1)
}

func _bs(in []int, target int, left int, right int) int {
	// base condition for recursion
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	if in[mid] == target {
		return mid
	} else if target < in[mid] {
		return _bs(in, target, left, mid-1)
	} else {
		return _bs(in, target, mid+1, right)
	}
}

func MyBinarySearchIterative(in []int, target int) int {
	l, r := 0, len(in)
	for l < r {
		mid := (l + r) / 2
		if in[mid] == target {
			return mid
		} else if target < in[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
