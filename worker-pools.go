package main

import "fmt"
import "time"

type task struct {
	id int
}

type result struct {
	task   task
	result int
}

func worker(id int, input <-chan task, output chan<- result) {
	for t := range input {
		fmt.Println("worker", id, "processing task", t)
		time.Sleep(time.Second)
		output <- result{task: t, result: t.id * 2}
	}
}

func main() {
	tasks := 9
	input := make(chan task, tasks)
	output := make(chan result, tasks)

	for w := 1; w <= 3; w++ {
		go worker(w, input, output)
	}
	for id := 1; id <= tasks; id++ {
		input <- task{id: id}
	}
	close(input)

	for a := 1; a <= tasks; a++ {
		d := <-output
		fmt.Println("task", d.task, "result", d.result)
	}
	close(output)
}
