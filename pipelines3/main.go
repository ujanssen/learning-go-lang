package main

import (
	"fmt"
	"time"
)

func main() {
	steps := 100 * 1000
	line := make([]*pipeline, steps)

	for i := 0; i < steps; i++ {
		line[i] = &pipeline{
			Out:     make(chan int),
			Compute: func(in int) int { return in + 1 }}
	}

	line[0].In = make(chan int)
	// connect the input from step n with output
	// from step n-1
	for n := 1; n < steps; n++ {
		line[n].In = line[n-1].Out
	}

	// start computation
	for i := 0; i < steps; i++ {
		go func() {
			line[i].Out <- line[i].Compute(<-line[i].In)
		}()
	}

	start := time.Now()
	// start pipeline with value 0
	line[0].In <- 0
	// wait for result
	println(<-line[steps-1].Out)
	fmt.Printf("Duration: %v\n", time.Since(start))
}

type pipeline struct {
	In, Out chan int
	Compute func(int) int
}
