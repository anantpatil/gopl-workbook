package main

import "testing"

// Go doesn't allow following - it doesn't allow arrays. maps, slices,
// to be constants Why? Because array is pointer to a list, and the
// contents of the list or map can be independently modified.
// const a = []int{1, 2, 3}

var a = []int{0, 1, 5, 13, 26, 47, 79, 81, 99, 106, 110}
var cases = []struct {
	input, expected int
}{
	{13, 3},
	{0, 0},
	{110, 10},
	{112, -1},
	{-1, -1},
}

func TestMyBinarySearch(t *testing.T) {
	for _, c := range cases {
		actual := MyBinarySearch(a, c.input)
		if actual != c.expected {
			t.Errorf("Failed search for %d, expected %d, got %d", c.input, c.expected, actual)
		}
	}
}

func TestMyBinarySearchIterative(t *testing.T) {
	for _, c := range cases {
		actual := MyBinarySearchIterative(a, c.input)
		if actual != c.expected {
			t.Errorf("Failed iterative search for %d, expected %d, got %d", c.input, c.expected, actual)
		}
	}
}
