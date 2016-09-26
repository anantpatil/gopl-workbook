package main

import (
	"fmt"
	"os"
)

func rot13(s string) string {
	runes := []rune(s)
	for i, x := range runes {
		if (x >= 'A' && x <= 'M') || (x >= 'a' && x <= 'm') {
			x += 13
		} else if (x >= 'N' && x <= 'Z') || (x >= 'n' && x <= 'z') {
			x -= 13
		}
		runes[i] = x
	}
	return string(runes)
}

func rot13_test() {
	s := ""
	s1 := rot13(s)
	if s1 != "" {
		fmt.Println("Error ", s, s1)
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {
	rot13_test()
}
