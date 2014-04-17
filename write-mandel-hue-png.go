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

	const maxX = 600
	const maxY = 480

	var iteration [maxX][maxY]int

	const centerRe, centerIm float64 = -0.7435669, 0.1314023
	const hd float64 = 0.0022878
	const hv float64 = float64(maxY) / float64(maxX) * hd

	const hdStepX = hd / maxX
	const hdStepY = hv / maxY

	const upperLeftRe = centerRe - (float64(maxX/2) * hdStepX)
	const upperLeftIm = centerIm - (float64(maxY/2) * hdStepY)

	fmt.Println("Iterate")
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			dx := float64(x) * hdStepX
			dy := float64(y) * hdStepY
			iteration[x][y] = mandel.Iterate(upperLeftRe+dx, upperLeftIm+dy, maxIter)
		}
	}

	var f *os.File
	var err error
	var c color.Color
	m := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	b := m.Bounds()
	for colorShift := 0; colorShift < 36; colorShift++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				iter := iteration[x][y]
				h := float64((iter+colorShift*10)%360) / 360.0
				if c, err = hue.Color(h); err != nil {
					panic(err)
				}

				m.Set(x, maxY-y, c)
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
