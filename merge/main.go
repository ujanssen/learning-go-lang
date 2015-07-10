package main

import (
	"fmt"
	"sync"
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

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	in := gen(1000 * 1000 * 1000 * 1000 * 1000 * 1000)
	c1 := sq(in)
	c2 := sq(in)
	c3 := sq(in)
	c4 := sq(in)

	// Consume the output.
	for n := range merge(c1, c2, c3, c4) {
		fmt.Println(n) // 1,4,9...
	}
}
