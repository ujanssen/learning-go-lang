package main

import (
	"fmt"
	"mandel"
)

func main() {
	const maxIter = 1024

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
			iter := mandel.Iterate(upperLeftRe+dx, upperLeftIm+dy, maxIter)
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
