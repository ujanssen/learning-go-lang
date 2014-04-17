package main

import (
	"fmt"
	"hue"
	"image"
	"image/color"
	"image/png"
	"mandel"
	"math"
	"os"
)

func main() {
	var c color.Color
	var maxIter = 10 * 360
	var d = mandel.NewData(640, 400, -0.7435669, 0.1314023, 0.0022878)

	for colScale := 30; colScale < 120; colScale += 30 {
		n := fmt.Sprintf("mandel-hue-%04d.png", colScale)
		fmt.Println("Render n:", n)
		f, err := os.Create(n)
		if err != nil {
			panic(err)
		}
		m := image.NewRGBA(image.Rect(0, 0, d.MaxX, d.MaxY))
		for y := 0; y < d.MaxY; y++ {
			for x := 0; x < d.MaxX; x++ {
				dx := float64(x) * d.StepX
				dy := float64(y) * d.StepY
				iter := mandel.Iterate(d.UpperLeftRe+dx, d.UpperLeftIm+dy, maxIter)
				h := float64(iter) / float64(colScale)
				h = h - math.Floor(h)
				if c, err = hue.Color(h); err != nil {
					panic(err)
				}
				m.Set(x, d.MaxY-y, c)
			}
		}

		if err = png.Encode(f, m); err != nil {
			panic(err)
		}
	}
}
