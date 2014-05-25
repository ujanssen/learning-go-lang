package fibonacci

import (
	"fibonacci"
	"testing"
)

/*

go test src/fibonacci/fibonacci_test.go
ok  	command-line-arguments	0.019s

*/

func testFib(n, e int, t *testing.T) {
	r := fibonacci.Fib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
	r = fibonacci.Mfib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
}

func TestFib10(t *testing.T) {
	testFib(10, 55, t)
}

func TestFib15(t *testing.T) {
	testFib(15, 610, t)
}

func TestFib20(t *testing.T) {
	testFib(20, 6765, t)
}

func TestFib25(t *testing.T) {
	testFib(25, 75025, t)
}

/*

go test -bench=. src/fibonacci/fibonacci_test.go

*/

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Fib(i)
	}
}

func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib15(b *testing.B) { benchmarkFib(15, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib25(b *testing.B) { benchmarkFib(25, b) }
func BenchmarkFib30(b *testing.B) { benchmarkFib(30, b) }
func BenchmarkFib35(b *testing.B) { benchmarkFib(35, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func benchmarkMfib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Mfib(i)
	}
}

func BenchmarkMfib10(b *testing.B) { benchmarkMfib(10, b) }
func BenchmarkMfib15(b *testing.B) { benchmarkMfib(15, b) }
func BenchmarkMfib20(b *testing.B) { benchmarkMfib(20, b) }
func BenchmarkMfib25(b *testing.B) { benchmarkMfib(25, b) }
func BenchmarkMfib30(b *testing.B) { benchmarkMfib(30, b) }
func BenchmarkMfib35(b *testing.B) { benchmarkMfib(35, b) }
func BenchmarkMfib40(b *testing.B) { benchmarkMfib(40, b) }
