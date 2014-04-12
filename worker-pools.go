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

func createChannels(num int) (chan task, chan result) {
	input := make(chan task, num)
	output := make(chan result, num)
	return input, output
}

func worker(id int, input <-chan task, output chan<- result) {
	for t := range input {
		fmt.Println("worker", id, "processing task", t)
		time.Sleep(time.Second)
		output <- result{task: t, result: t.id * 2}
	}
}

func createWorker(num int, input <-chan task, output chan<- result) {
	for w := 1; w <= num; w++ {
		go worker(w, input, output)
	}
}

func createTasks(num int, input chan<- task) {
	for id := 1; id <= num; id++ {
		input <- task{id: id}
	}
	close(input)
}

func readResults(num int, output <-chan result) {
	for a := 1; a <= num; a++ {
		d := <-output
		fmt.Println("task", d.task, "result", d.result)
	}
}

func main() {
	tasks := 9
	input, output := createChannels(tasks)

	createTasks(tasks, input)
	createWorker(3, input, output)

	readResults(tasks, output)
}
