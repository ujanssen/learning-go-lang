package main

import (
	"testing"

	"github.com/ujanssen/learning-go-lang/fibonacci"
)

/*

go test src/fibonacci/fibonacci_test.go
ok  	command-line-arguments	0.019s

*/

func testFib(n, e int, t *testing.T) {
	var r int

	if n <= 20 {
		r = fibonacci.Fib(n)
		if e != r {
			t.Errorf("Fib(%v) = %v, want %v", n, r, e)
		}
	}
	r = fibonacci.Mfib(n)
	if e != r {
		t.Errorf("Fib(%v) = %v, want %v", n, r, e)
	}
}

func TestFib10(t *testing.T) {
	testFib(10, 55, t)
}

func TestFib20(t *testing.T) {
	testFib(20, 6765, t)
}

func TestFib30(t *testing.T) {
	testFib(30, 832040, t)
}

func TestFib40(t *testing.T) {
	testFib(40, 102334155, t)
}

func TestFib50(t *testing.T) {
	testFib(50, 12586269025, t)
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Fib(i)
	}
}

func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib30(b *testing.B) { benchmarkFib(30, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }

func benchmarkMfib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Mfib(i)
	}
}

func BenchmarkMfib10(b *testing.B) { benchmarkMfib(10, b) }
func BenchmarkMfib20(b *testing.B) { benchmarkMfib(20, b) }
func BenchmarkMfib30(b *testing.B) { benchmarkMfib(30, b) }
func BenchmarkMfib40(b *testing.B) { benchmarkMfib(40, b) }

func benchmarkLfib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		fibonacci.Mfib(i)
	}
}

func BenchmarkLfib10(b *testing.B) { benchmarkLfib(10, b) }
func BenchmarkLfib20(b *testing.B) { benchmarkLfib(20, b) }
func BenchmarkLfib30(b *testing.B) { benchmarkLfib(30, b) }
func BenchmarkLfib40(b *testing.B) { benchmarkLfib(40, b) }
