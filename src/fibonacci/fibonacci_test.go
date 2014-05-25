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
}

func TestFib15(t *testing.T) {
	n, e := 15, 610
	r := fibonacci.Fib(n)
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
}
