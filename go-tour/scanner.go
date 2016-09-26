package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// scan from input and print
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Anant") {
			fmt.Println("Line has name\n")
		}
	}
}
