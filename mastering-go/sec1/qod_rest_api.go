// Get Quote of the Day from a Public API endpoint
// Public API is at https://theysaidso.com/api/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// multiple main func in same package shows following in err
// go install will create a binay by name mastering-go (parent folder
// name) and copy to $GOPATH/bin
func main() {
	resp, err := http.Get("http://quotes.rest/qod")
	if err != nil {
		// Falal = Orint + sys,exit(1)
		log.Fatal(err)
	}
	// Body of resp is os type io.ReadCloser
	// It is streamed on demand from the server as the Body field is
	// read. It needs to be closed manually.
	defer resp.Body.Close()

	// ioutil.ReadAll will read all the streamed response from the
	// server into an array of bytes
	qod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", qod)
}
