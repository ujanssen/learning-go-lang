package main

import (
	"hue"
	"image"
	"image/color"
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
	var c color.Color
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			h := float64(x) / 60.0
			h = h - math.Floor(h)
			if c, err = hue.Color(h); err != nil {
				panic(err)
			}
			m.Set(x, y, c)
		}
	}
	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
