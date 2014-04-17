package main

import (
	"fmt"
	"hue"
	"image"
	"image/color"
	"image/png"
	"mandel"
	"os"
)

func main() {
	const maxIter = 10 * 360
	var d = mandel.NewData(640, 400, -0.7435669, 0.1314023, 0.0022878)

	var iteration [640][400]int

	fmt.Println("Iterate")
	for y := 0; y < d.MaxY; y++ {
		for x := 0; x < d.MaxX; x++ {
			dx := float64(x) * d.StepX
			dy := float64(y) * d.StepY
			iteration[x][y] = mandel.Iterate(d.UpperLeftRe+dx, d.UpperLeftIm+dy, maxIter)
		}
	}

	var f *os.File
	var err error
	var c color.Color
	m := image.NewRGBA(image.Rect(0, 0, d.MaxX, d.MaxY))
	for colorShift := 0; colorShift < 36; colorShift++ {
		for y := 0; y < d.MaxY; y++ {
			for x := 0; x < d.MaxX; x++ {
				iter := iteration[x][y]
				h := float64((iter+colorShift*10)%360) / 360.0
				if c, err = hue.Color(h); err != nil {
					panic(err)
				}

				m.Set(x, d.MaxY-y, c)
			}
		}
		n := fmt.Sprintf("mandel-hue-colorshift-%02d.png", colorShift)
		fmt.Println("write file", n)
		f, err = os.Create(n)
		if err != nil {
			panic(err)
		}
		if err = png.Encode(f, m); err != nil {
			panic(err)
		}
	}
}
