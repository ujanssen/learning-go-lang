package main

import (
	"hue"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	f, err := os.Create("hue.png")
	if err != nil {
		panic(err)
	}
	maxX := 3 * 360
	m := image.NewRGBA(image.Rect(0, 0, maxX-1, 240))
	b := m.Bounds()
	var c color.Color

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if c, err = hue.Color(float64(x) / float64(maxX)); err != nil {
				panic(err)
			}
			m.Set(x, y, c)
		}
	}
	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
