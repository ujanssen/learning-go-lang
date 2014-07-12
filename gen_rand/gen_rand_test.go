package gen_rand_test

import (
	"github.com/ujanssen/learning-go-lang/gen_rand"
	"testing"
)

var mio int = 1000000

func TestGenRand(t *testing.T) {
	if gen_rand.GenRand(0) < 0 {
		t.Errorf("fail")
	}
}

func GenRandGo(t *testing.T) {
	if gen_rand.GenRandGo(0, 0) < 0 {
		t.Errorf("fail")
	}
}

// go test -bench=. gen_rand_test.go

func genRand(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		gen_rand.GenRand(i)
	}
}

func BenchmarkGenRand(b *testing.B) { genRand(mio, b) }

func genRandGo(i, g int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		gen_rand.GenRandGo(i, g)
	}
}

func BenchmarkGenRand1(b *testing.B) { genRandGo(mio, 1, b) }
func BenchmarkGenRand2(b *testing.B) { genRandGo(mio, 2, b) }
func BenchmarkGenRand3(b *testing.B) { genRandGo(mio, 3, b) }
func BenchmarkGenRand4(b *testing.B) { genRandGo(mio, 4, b) }
func BenchmarkGenRand5(b *testing.B) { genRandGo(mio, 5, b) }
func BenchmarkGenRand6(b *testing.B) { genRandGo(mio, 6, b) }
func BenchmarkGenRand7(b *testing.B) { genRandGo(mio, 7, b) }
func BenchmarkGenRand8(b *testing.B) { genRandGo(mio, 8, b) }
