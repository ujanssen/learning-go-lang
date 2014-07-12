package main

import (
	"fmt"
	"github.com/ujanssen/learning-go-lang/gen_rand"
	"runtime"
	"time"
)

var mio int = 100000000

func main() {
	start := time.Now()
	d := gen_rand.GenRand(mio)
	fmt.Println("c ", d, " Dur ", time.Since(start))

	maxProcs := runtime.GOMAXPROCS(0)
	fmt.Println("maxProcs ", maxProcs)
	numCPU := runtime.NumCPU()
	fmt.Println("numCPU ", numCPU)

	maxProcs = runtime.GOMAXPROCS(numCPU)
	maxProcs = runtime.GOMAXPROCS(numCPU)
	fmt.Println("maxProcs ", maxProcs)

	runtime.GOMAXPROCS(numCPU)

	start = time.Now()
	gen_rand.GenRandGo(mio, numCPU)
	fmt.Println("c ", d, " Dur ", time.Since(start))
}
