package main

import (
	"fmt"
)

// n: 50
// fib: 12586269025
//
// real	1m21.730s
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib(n-2) + fib(n-1)
}

// n: 50
// fmib: 12586269025
//
// real	0m0.279s
func mfib(n int) int {
	if n == 1 || n == 2 {
		f[n] = 1
		return 1
	}
	if f[n] == 0 {
		f[n] = mfib(n-1) + mfib(n-2)
	}
	return f[n-1] + f[n-2]
}

var f [51]int

func main() {
	n := 50
	fmt.Println("n:", n)
	//	fmt.Println("fib:", fib(n))

	fmt.Println("fmib:", mfib(n))
}
