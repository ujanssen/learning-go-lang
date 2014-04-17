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

var maxIter = 10 * 360

func main() {
	const maxX = 640
	const maxY = 400

	const centerRe, centerIm float64 = -0.7435669, 0.1314023
	const hd float64 = 0.0022878
	const hv float64 = float64(maxY) / float64(maxX) * hd

	const hdStepX = hd / maxX
	const hdStepY = hv / maxY

	const upperLeftRe = centerRe - (float64(maxX/2) * hdStepX)
	const upperLeftIm = centerIm - (float64(maxY/2) * hdStepY)

	var c color.Color

	for colScale := 30; colScale < 360*2; colScale += 30 {
		n := fmt.Sprintf("mandel-hue-%04d.png", colScale)
		fmt.Println("Render n:", n)
		f, err := os.Create(n)
		if err != nil {
			panic(err)
		}
		m := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
		b := m.Bounds()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				dx := float64(x) * hdStepX
				dy := float64(y) * hdStepY
				iter := mandel.Iterate(upperLeftRe+dx, upperLeftIm+dy, maxIter)
				h := float64(iter) / float64(colScale)
				h = h - math.Floor(h)
				if c, err = hue.Color(h); err != nil {
					panic(err)
				}
				m.Set(x, maxY-y, c)
			}
		}

		if err = png.Encode(f, m); err != nil {
			panic(err)
		}
	}
}
