package main

import (
	"fmt"
	"hue"
	"image/color"
	"mandel"
	"pngimage"
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

	var err error
	var c color.Color
	for colorShift := 0; colorShift < 36; colorShift++ {
		n := fmt.Sprintf("mandel-hue-colorshift-%02d.png", colorShift)
		pi := pngimage.NewPngimage(d.MaxX, d.MaxY, n)

		for y := 0; y < d.MaxY; y++ {
			for x := 0; x < d.MaxX; x++ {
				iter := iteration[x][y]
				h := float64((iter+colorShift*10)%360) / 360.0
				if c, err = hue.Color(h); err != nil {
					panic(err)
				}

				pi.Img.Set(x, d.MaxY-y, c)
			}
		}
		fmt.Println("write file", n)
		pi.Save()
	}
}
