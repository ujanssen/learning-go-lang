package main

import "fmt"
import "time"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}
func main() {
	tasks := 9
	jobs := make(chan int, tasks)
	results := make(chan int, tasks)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= tasks; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= tasks; a++ {
		<-results
	}
}
