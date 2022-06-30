package main

import "fmt"

// empty struct doesn't take any space!
type Set map[string]struct{}

func main6() {
	s := make(Set)
	//add items
	s["Item1"] = struct{}{}
	s["Item2"] = struct{}{}
	//get and print items
	fmt.Println(getSetValues(s))
}

func getSetValues(s Set) []string {
	var retVal []string
	for k, _ := range s {
		retVal = append(retVal, k)
	}
	return retVal
}
