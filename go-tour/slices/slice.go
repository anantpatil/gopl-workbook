package main

import (
	"fmt"
)

// write a program to cerate a multidiemnsional slice

func main() {
	// 2D slice of unit
	var v [][]uint
	v = make([][]uint, 5) // 5 rows
	for i := 0; i < 5; i++ {
		v[i] = make([]uint, 6) // 6 columns
		for j := 0; j < 6; j++ {
			v[i][j] = (uint)((i + j) / 2)
		}
		fmt.Println(v[i])
	}
}
