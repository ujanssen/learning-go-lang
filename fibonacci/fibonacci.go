package fibonacci

var f [61]int

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
BenchmarkFib10	 5000000	       330 ns/op
BenchmarkFib15	  500000	      3684 ns/op
BenchmarkFib20	   50000	     40786 ns/op
BenchmarkFib25	    5000	    450551 ns/op
BenchmarkFib30	     500	   5043135 ns/op
BenchmarkFib35	      50	  55598520 ns/op
BenchmarkFib40	       2	 618482580 ns/op
BenchmarkMfib10	500000000	         5.68 ns/op
BenchmarkMfib15	500000000	         5.68 ns/op
BenchmarkMfib20	500000000	         5.66 ns/op
BenchmarkMfib25	500000000	         5.72 ns/op
BenchmarkMfib30	500000000	         5.63 ns/op
BenchmarkMfib35	500000000	         5.66 ns/op
BenchmarkMfib40	500000000	         5.65 ns/op
ok  	command-line-arguments	40.218s

*/
