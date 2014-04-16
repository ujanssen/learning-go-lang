package main

import (
	"hue"
	"image"
	"image/png"
	"math"
	"os"
)

func main() {
	f, err := os.Create("use-hue4.png")
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, 360-1, 240))
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := float64(x) / 60.0
			c = c - math.Floor(c)
			m.Set(x, y, hue.Color(c))
		}
	}
	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
