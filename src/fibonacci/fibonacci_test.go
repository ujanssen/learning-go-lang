package fibonacci

import (
	"fibonacci"
	"testing"
)

/*

go test src/fibonacci/fibonacci_test.go
ok  	command-line-arguments	0.019s

*/

func TestFib10(t *testing.T) {
	n, e := 10, 55
	r := fibonacci.Fib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
	r = fibonacci.Mfib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
}

func TestFib15(t *testing.T) {
	n, e := 15, 610
	r := fibonacci.Fib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
	r = fibonacci.Mfib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
}

func TestFib20(t *testing.T) {
	n, e := 20, 6765
	r := fibonacci.Fib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
	r = fibonacci.Mfib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
}

func TestFib25(t *testing.T) {
	n, e := 25, 75025
	r := fibonacci.Fib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
	r = fibonacci.Mfib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
}

/*

go test -bench=. src/fibonacci/fibonacci_test.go

*/
func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Fib(10)
	}
}
func BenchmarkFib15(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Fib(15)
	}
}
func BenchmarkFib20(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Fib(20)
	}
}
func BenchmarkFib25(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Fib(25)
	}
}

func BenchmarkMfib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Mfib(10)
	}
}
func BenchmarkMfib15(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Mfib(15)
	}
}
func BenchmarkMfib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Mfib(20)
	}
}
func BenchmarkMfib25(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Mfib(25)
	}
}
