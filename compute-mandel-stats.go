package main

import (
	"fmt"
	"math/cmplx"
)

const maxIter = 1024

func computeMandel(cx, cy float64) int {
	var i int
	var z, c complex128
	c = complex(cx, cy)
	for {
		i = i + 1
		z = z*z + c
		if i == maxIter {
			break
		}
		if cmplx.Abs(z) > 2.0 {
			break
		}
	}
	return i
}

func main() {
	const maxX = 320
	const maxY = 200

	const centerRe, centerIm float64 = -0.7435669, 0.1314023
	const hd float64 = 0.0022878
	const hv float64 = float64(maxY) / float64(maxX) * hd

	const hdStepX = hd / maxX
	const hdStepY = hv / maxY

	const upperLeftRe = centerRe - (float64(maxX/2) * hdStepX)
	const upperLeftIm = centerIm - (float64(maxY/2) * hdStepY)

	m := make(map[int]int)

	var maxV = 0
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			dx := float64(x) * hdStepX
			dy := float64(y) * hdStepY
			iter := computeMandel(upperLeftRe+dx, upperLeftIm+dy)
			m[iter] = m[iter] + 1
			if m[iter] > maxV {
				maxV = m[iter]
			}
			fmt.Println(x, maxY-y, iter)
		}
	}
	for x := 1; x <= maxV; x++ {
		for k, v := range m {
			if x == v {
				fmt.Println("iter", k, " -> count", v)
			}
		}
	}
}
