package main

import (
	"fmt"

	"github.com/anantpatil/gopl-workbook/mastering-go/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	s := "mastering-go"
	fmt.Println(s)
	rs := morestrings.ReverseString(s)
	fmt.Println(rs)
	fmt.Println(cmp.Diff(s, rs))
}
