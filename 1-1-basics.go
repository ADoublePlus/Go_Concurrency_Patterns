// Patterns from Pike's Google I/O talk, "Go Concurrency Patterns"

// Exposition of Golang's concurrency primitives

package main

import (
	"fmt"
	"time"
)

func main()  {
	go regular_print("Hello")
	fmt.Println("Second print statement!")
	time.Sleep(3 * time.Second)
	fmt.Println("Third print statement!") // When main returns, the goroutines also end
}

func regular_print(msg string)  {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}