package main

import (
	"fmt"
	"strconv"

	"github.com/ujanssen/learning-go-lang/fibonacci"
)

const max = 240

func main() {
	var v string
	var i uint64
	for {
		if i > 0 {
			v += ", "
		}
		f := fibonacci.Mfib(i)
		v += strconv.FormatUint(f, 10)
		i++
		if i > max {
			break
		}
	}

	fmt.Printf("Fib -> %d: %s\n", max, v)
}
