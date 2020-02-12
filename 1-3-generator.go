// Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

// Golang generator pattern: functions that return channels

package main

import (
	"fmt"
	"time"
)

// Goroutine is launched inside the called function (more idiomatic)
// Multiple instances of the generator may be called

func main()  {
	ch := generator("Hello")

	for i := 0; i < 5; i++ {
		fmt.Println(<- ch)
	}
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