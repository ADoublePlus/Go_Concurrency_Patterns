// Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

// Golang multiplexing (fan-in) function to allow multiple channels go through one channel

package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := fanIn(generator("Hello"), generator("Bye"))

	for i := 0; i < 10; i++ {
		fmt.Println(<- ch)
	}
}

// fanIn is itself a generator
func fanIn(ch1, ch2 <- chan string) <- chan string { // Receives two read-only channels
	new_ch := make(chan string)

	// Launch two goroutine while loops to continuously pipe to new channel
	go func() {
		for {
			new_ch <- <- ch1
		}
	}()

	go func() {
		for {
			new_ch <- <- ch2
		}
	}()

	return new_ch
}

func generator(msg string) <- chan string { // Returns receive-only channel
	ch := make(chan string)

	go func() { // Anonymous goroutine
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()

	return ch
}