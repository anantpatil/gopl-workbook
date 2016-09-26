package main

import "fmt"

const N = 4

type board [N][N]int

func isSafe(b *board, row int, col int) bool {
	// we are filling one column at a time
	// check anything in same row; since we start from left, just
	// check on the left side only
	for i := 0; i < col; i++ {
		if b[row][i] == 1 {
			return false
		}
	}

	// check is there is a queen in upper left
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b[i][j] == 1 {
			return false
		}
	}

	// check if there is a queen in lower left
	for i, j := row, col; i < N && j >= 0; i, j = i+1, j-1 {
		if b[i][j] == 1 {
			return false
		}
	}

	return true
}

func solveNQueens(b *board, col int) bool {
	// go through columns one by one and recursively try for rows
	if col == N {
		return true
	}

	for row := 0; row < N; row++ {
		if isSafe(b, row, col) {
			b[row][col] = 1
			if solveNQueens(b, col+1) {
				return true
			}
			//backtrack
			b[row][col] = 0
		}
	}

	return false
}

func printBoard(b *board) {
	for _, v := range b {
		fmt.Println(v)
	}
	fmt.Println("\n")
}

func main() {
	myboard := board{}
	if solveNQueens(&myboard, 0) {
		printBoard(&myboard)
		return
	}
	fmt.Println("Not solvable")
}
