package main

import (
	"fmt"

	"github.com/ujanssen/learning-go-lang/mandel"
)

func main() {
	const maxIter = 1024

	var d = mandel.NewData(320, 200, -0.7435669, 0.1314023, 0.0022878)

	m := make(map[int]int)

	var maxV = 0
	for y := 0; y < d.MaxY; y++ {
		for x := 0; x < d.MaxX; x++ {
			dx := float64(x) * d.StepX
			dy := float64(y) * d.StepY
			iter := mandel.Iterate(d.UpperLeftRe+dx, d.UpperLeftIm+dy, maxIter)
			m[iter] = m[iter] + 1
			if m[iter] > maxV {
				maxV = m[iter]
			}
			fmt.Println(x, d.MaxY-y, iter)
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
