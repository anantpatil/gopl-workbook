package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// a struct that has an reading end of stream and implments a reader below
type rot13Reader struct {
	r io.Reader
}

// rot13Reader is-a io.Reader as it implmenets a Read method
// rot13Reader has a reader and is-a reader. From its reader, it reads data and converts it to
// by rotating 13 positions of english alphabet and returns back
func (r *rot13Reader) Read(b []byte) (int, error) {
	// rotate by 13 places A -> N, B -> O...
	// r has a read end of some stream of data,
	// convert it by applying rot13 and write back to b
	n, err := r.r.Read(b)
	if err != nil {
		if err != io.EOF {
			return n, err
		}
	}
	for i, x := range b {
		c := x
		if (c >= 'a' && c <= 'm') || (c >= 'A' && c <= 'M') {
			c += 13
		} else if (c >= 'n' && c <= 'z') || (c >= 'N' && c <= 'Z') {
			c -= 13
		} else {
			// pass
		}
		b[i] = c
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println() // just add new line in end to push promt to next line
}
