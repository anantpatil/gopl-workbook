package main

// some type which has a Read impmplemented
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	// return number of bytes and error if any
	//r := rune('A')
	// ba := []byte(string(r)) // ba = [65]
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	myR := MyReader{}
	// implement a reader which makes sure that the MyReader reads infinitely
	// https://github.com/golang/tour/blob/master/reader/validate.go
}




