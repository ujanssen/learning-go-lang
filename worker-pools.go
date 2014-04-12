package main

import "fmt"
import "time"

type task struct {
	id     int
	result int
}

func worker(id int, input <-chan task, output chan<- task) {
	for t := range input {
		fmt.Println("worker", id, "processing task", t.id)
		t.result = t.id * 2
		time.Sleep(time.Second)
		output <- t
	}
}

func main() {
	tasks := 9
	input := make(chan task, tasks)
	output := make(chan task, tasks)

	for w := 1; w <= 3; w++ {
		go worker(w, input, output)
	}
	for j := 1; j <= tasks; j++ {
		input <- task{id: j}
	}
	close(input)

	for a := 1; a <= tasks; a++ {
		t := <-output
		fmt.Println("task", t.id, "result", t.result)
	}
	close(output)
}
