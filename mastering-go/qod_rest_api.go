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
func main() {
	resp, err := http.Get("http://quotes.rest/qod")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	qod, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", qod)
}
