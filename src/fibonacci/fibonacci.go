package fibonacci

import (
	"fmt"
)

var f [51]int

func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-2) + Fib(n-1)
}

func Mfib(n int) int {
	if n == 1 || n == 2 {
		f[n] = 1
		return 1
	}
	if f[n] == 0 {
		f[n] = Mfib(n-1) + Mfib(n-2)
	}
	return f[n-1] + f[n-2]
}

/*

$ go test -bench=. src/fibonacci/fibonacci_test.go
PASS
BenchmarkFib10	 5000000	       348 ns/op
BenchmarkFib15	  500000	      3896 ns/op
BenchmarkFib20	   50000	     43294 ns/op
BenchmarkFib25	    5000	    480761 ns/op
BenchmarkMfib10	500000000	         5.66 ns/op
BenchmarkMfib15	500000000	         5.31 ns/op
BenchmarkMfib20	500000000	         5.67 ns/op
BenchmarkMfib25	500000000	         5.29 ns/op
ok  	command-line-arguments	22.353s

*/
