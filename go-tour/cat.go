package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func cat(fd *os.File) {
	reader := bufio.NewReader(fd)
	io.Copy(os.Stdout, reader)
}

func main() {
	// open the file and read and write to os.Stdout
	if len(os.Args) == 1 {
		fmt.Println("Please give file names")
		return
	}
	for _, file := range os.Args[1:] {
		// read the file and print it on STDOUT
		fd, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file: ", file, err)
			continue
		}
		defer fd.Close()
		cat(fd)
	}
}
