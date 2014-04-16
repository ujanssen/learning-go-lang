package main

import (
	"fmt"
	"hue"
	"image"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func computeMandel(cx, cy float64) int {
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
				iter := computeMandel(upperLeftRe+dx, upperLeftIm+dy)
				c := float64(iter) / float64(colScale)
				c = c - math.Floor(c)
				m.Set(x, maxY-y, hue.Color(c))
			}
		}

		if err = png.Encode(f, m); err != nil {
			panic(err)
		}
	}
}
