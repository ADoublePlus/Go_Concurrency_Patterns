// Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

// Daisy chaining goroutines... all routines at once, so if one is fulfilled
// everything after should also have their blocking commitment (input) fulfilled

package main

import (
	"fmt"
)

// Takes two int channels. store right val (+1) into left
func f(left, right chan int)  {
	left <- 1 + <- right // After 1st right read, locks until left read
}

func main()  {
	const n = 10000

	// Construct an array of n+1 int channels
	var channels [n + 1]chan int

	for i := range channels {
		channels[i] = make(chan int)
	}

	// Wire n goroutines in a chain
	for i := 0; i < n; i++ {
		go f(channels[i], channels[i + 1])
	}

	// Insert a value into right-hand end
	go func(c chan <- int) {
		c <- 1
	}(channels[n])

	// Get value from the left-hand end
	fmt.Println(<- channels[0])
}