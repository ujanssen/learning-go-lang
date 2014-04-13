package main

import "fmt"
import "time"
import "flag"
import "runtime"

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

func doWork(id int, input <-chan task, output chan<- result) {
	for t := range input {
		fmt.Println("worker", id, "processing task", t)
		time.Sleep(time.Second)
		output <- result{task: t, result: t.id * 2}
	}
}

func createWorker(num int, input <-chan task, output chan<- result) {
	for w := 1; w <= num; w++ {
		go doWork(w, input, output)
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

//go run worker-pools.go -worker=2 -tasks=20
func getFromCommandLine(t int, w int) (int, int) {
	numWorker := flag.Int("worker", w, "num worker")
	numWTasks := flag.Int("tasks", t, "num tasks")
	flag.Parse()
	return *numWTasks, *numWorker
}

func main() {
	tasks, worker := getFromCommandLine(9, runtime.NumCPU())
	input, output := createChannels(tasks)

	createTasks(tasks, input)
	createWorker(worker, input, output)

	readResults(tasks, output)
}
