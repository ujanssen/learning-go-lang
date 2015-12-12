package main

import "fmt"
import "time"
import "flag"
import "runtime"

//go run main.go -worker=2 -tasks=20
func main() {
	numWorker := flag.Int("worker", runtime.NumCPU(), "num worker")
	numTasks := flag.Int("tasks", 20, "num tasks")
	flag.Parse()

	input := make(chan task, *numTasks)
	output := make(chan task, *numTasks)

	// create tasks
	for id := 1; id <= *numTasks; id++ {
		input <- task{id: id, output: output}
	}
	close(input)

	// create Worker
	for w := 1; w <= *numWorker; w++ {
		go doWork(w, input)
	}

	// read results
	for d := range output {
		*numTasks--
		fmt.Println("result: task", d.id, "result", d.result)
		if *numTasks == 0 {
			break
		}
	}
}

func doWork(id int, input <-chan task) {
	for t := range input {
		fmt.Println("doWork: worker", id, "processing task", t.id)
		time.Sleep(time.Second)
		t.result = t.id * 2
		t.output <- t
	}
}

type task struct {
	id     int
	result int
	output chan task
}
