package main

import (
	"fmt"
)

type Request struct {
	args []int
	f    func([]int) int

	resultChan chan int
}

func main() {
	clientRequests := make(chan *Request)
	go handle(clientRequests)

	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	clientRequests <- request
	fmt.Printf("answer: %d\n", <-request.resultChan)

	request = &Request{[]int{10, 4, 5}, sum, make(chan int)}
	clientRequests <- request
	fmt.Printf("answer: %d\n", <-request.resultChan)
}

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}
