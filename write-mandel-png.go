package main

import (
	"image"
	"image/color"
	"image/png"
	"mandel"
	"os"
)

func getColor(c int) color.RGBA {
	rgb := uint8(c % 255)
	return color.RGBA{rgb, rgb, rgb, 255}
}

func main() {
	const maxX = 1024
	const maxY = 800

	var d = mandel.NewData(640, 400, -0.7435669, 0.1314023, 0.0022878)

	f, err := os.Create("mandel.png")
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, d.MaxX, d.MaxY))
	for y := 0; y < d.MaxY; y++ {
		for x := 0; x < d.MaxX; x++ {
			dx := float64(x) * d.StepX
			dy := float64(y) * d.StepY
			iter := mandel.Iterate(d.UpperLeftRe+dx, d.UpperLeftIm+dy, 1024)
			c := getColor(iter)
			m.Set(x, d.MaxY-y, c)
		}
	}

	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
