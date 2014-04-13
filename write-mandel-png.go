package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func getColor(c int) color.RGBA {
	rgb := uint8(c % 255)
	return color.RGBA{rgb, rgb, rgb, 255}
}

func computeMandel(cx, cy float64) color.RGBA {
	var i int
	var absq float64
	var z, c complex128
	c = complex(cx, cy)
	for {
		z = z*z + c
		absq = cmplx.Abs(z)
		if i >= 1024 {
			return getColor(255)
		}
		if absq > 4.0 {
			return getColor(i)
		}
		i = i + 1
	}
	return getColor(0)
}

func main() {
	const maxX = 1024
	const maxY = 800

	const centerRe, centerIm float64 = -0.7435669, 0.1314023
	const hd float64 = 0.0022878
	const hv float64 = float64(maxY) / float64(maxX) * hd

	const hdStepX = hd / maxX
	const hdStepY = hv / maxY

	const upperLeftRe = centerRe - (float64(maxX/2) * hdStepX)
	const upperLeftIm = centerIm - (float64(maxY/2) * hdStepY)

	f, err := os.Create("mandel.png")
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			dx := float64(x) * hdStepX
			dy := float64(y) * hdStepY
			c := computeMandel(upperLeftRe+dx, upperLeftIm+dy)
			m.Set(x, maxY-y, c)
		}
	}

	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
