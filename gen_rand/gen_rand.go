package gen_rand

import (
	"fmt"
	"math/rand"
	"time"
)

func GenRand(n int) int {

	fmt.Println("GenRand", n)
	sum := 0

	start := time.Now()

	for i := 0; i < n; i++ {
		sum += rand.Intn(1000)
	}
	fmt.Println("n ", n, " Dur ", time.Since(start))

	if n > 0 {
		return sum / n
	}
	return 0
}

func randGo(id int, ready, rands, result chan int, msg chan string) {
	fmt.Println("randGo start", id)
	sum, n := 0, 0
	start := time.Now()

	<-ready

	fmt.Println("randGo ", id, "ready", time.Since(start))

	start = time.Now()
	for r := range rands {
		sum += r
		n++
	}
	msg <- fmt.Sprintf("%d-n %d Dur %s", id, n, time.Since(start))

	if n > 0 {
		result <- sum / n
	}
	result <- 0
}

func GenRandGo(n, g int) int {
	fmt.Println("GenRandGo", n, g)
	sum := 0
	results := make(chan int)
	rands := make(chan int, n)
	ready := make(chan int)
	msg := make(chan string)

	for i := 0; i < g; i++ {
		go randGo(i, ready, rands, results, msg)
	}

	for i := 0; i < n; i++ {
		rands <- rand.Intn(1000)
	}
	close(rands)
	close(ready)
	start := time.Now()

	for i := 0; i < g; i++ {
		fmt.Println(<-msg)
		sum += <-results
	}
	fmt.Println("n ", n, " Dur ", time.Since(start))
	return sum / n
}
