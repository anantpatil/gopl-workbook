package main

import "fmt"

func reverseWords(line string) string {
	// reverse the string
	// then iterate through them and reverse the individual words
	r := []rune(reverse(line))
	space := rune(' ')
	i, j := 0, 0
	for j < len(r) {
		if r[j] == space {
			i++
			j++
			continue
		}
		for j < len(r) && r[j] != space {
			j++
		}
		// encountered space; reverse r[i:j]
		reverseRunes(r, i, j)
		i = j
	}

	return string(r)
}

func reverseRunes(r []rune, start, end int) {
	// reverse the sequence of runes
	for i, j := start, end-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
}

func reverse(s string) string {
	// r := []byte(s) // this also works
	r := []rune(s) // but this is unicode aware
	reverseRunes(r, 0, len(r))
	return string(r)
}

func main() {
	var s string = "something"
	// for _, x := range s {
	//fmt.Printf("%T(%c)\n", x, x)
	//}
	r := reverse(s)
	if r != "gnihtemos" {
		fmt.Println("Error", s)
	}

	s = ""
	r = reverse("")
	if r != "" {
		fmt.Println("Error", s)
	}

	line := "A quick brown fox jumped over the lazy dog"
	fmt.Println(reverseWords(line))
}
