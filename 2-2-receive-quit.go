// Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

// Deterministically quit goroutine with quit channel option in select

package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	quit := make(chan string)
	ch := generator("Hi!", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<- ch, i)
	}

	quit <- "Bye!"
	fmt.Printf("Generator says %s", <- quit)
}

func generator(msg string, quit chan string) <- chan string { // Returns receive-only channel
	ch := make(chan string)

	go func() { // Anonymous goroutine
		for {
			select {
				case ch <- fmt.Sprintf("%s", msg):
					// Nothing

				case <- quit:
					quit <- "See you!"
					return
			}
		}
	}()

	return ch
}