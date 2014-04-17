package main

import (
	"hue"
	"image"
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
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			m.Set(x, y, hue.Color(float64(x)/float64(maxX)))
		}
	}
	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
