package main

import (
	"fmt"
)

func gen(num int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	c := gen(10000)
	out := sq(c)

	// Consume the output.
	for n := range out {
		fmt.Println(n) // 1,4,9...
	}
}
