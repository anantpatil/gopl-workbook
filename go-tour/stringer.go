package main

import (
	"fmt"
	"strconv"
)

type IPv4Addr [4]byte

func (ip IPv4Addr)String() string {
	// return a string representation of IP address like "127.0.0.1"
	var s string = ""
	for i := 0; i < len(ip) - 1; i++ {
		s += strconv.Itoa(int(ip[i]))
		s += "."
	}
	s += strconv.Itoa(int(ip[len(ip) - 1]))
	return s
}

func main() {
	hosts := map[string]IPv4Addr{
		"loopback": {127, 1, 1, 1},
		"google.com": {8, 8, 8, 8},
	}
	for name, addr := range hosts {
		fmt.Printf("%v: %v\n", name, addr)
	}
}
