package mandel

import (
	"math/cmplx"
)

func Iterate(cx, cy float64, maxIter int) int {
	var i int
	var abs float64
	var z, c complex128
	c = complex(cx, cy)
	for {
		z = z*z + c
		if i >= maxIter {
			break
		}
		abs = cmplx.Abs(z)
		if abs > 2.0 {
			break
		}
		i = i + 1
	}
	return i
}
