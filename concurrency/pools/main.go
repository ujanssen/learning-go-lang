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
	output := make(chan task)

	// create tasks
	for id := 1; id <= *numTasks; id++ {
		input <- task{id: id}
	}

	// create Worker
	for w := 1; w <= *numWorker; w++ {
		go doWork(w, input, output)
	}

	// input queue len logger
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("input len:", len(input))
		}
	}()

	// read results
	for d := range output {
		*numTasks--
		fmt.Println("result: task", d.id, "result", d.result)
		if *numTasks == 0 {
			close(input)
			break
		}
	}
}

func doWork(id int, input, output chan task) {
	for t := range input {
		fmt.Println("doWork: worker", id, "processing task", t.id, " state:", t.state)
		time.Sleep(100 * time.Millisecond)
		t.state++
		if t.state < 10 {
			input <- t
		} else {
			t.result = t.id * 2
			output <- t
		}
	}
}

type task struct {
	id     int
	result int
	state  int
}
