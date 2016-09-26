package main

import (
	"fmt"
	"time"
)

func main() {
	// everything in go is about duration, convert any time from
	// your unit to durations and then from duration you can drive
	// any other unit
	seconds := 64
	fmt.Println(time.Duration(seconds)) // Duration is in nanoseconds
	// to convert to seconds, multiply by time.Seconds
	duration := time.Duration(seconds) * time.Second
	fmt.Println("Duration is ", duration)

	minutes := 196
	duration = time.Duration(minutes) * time.Minute
	fmt.Println("Duration is ", duration)
	fmt.Println("Hours elapsed", duration.Hours())
	fmt.Println("Seconds elapsed", duration.Seconds())
}
