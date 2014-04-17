package main

import (
	"hue"
	"image/color"
	"pngimage"
)

func main() {
	maxX := 3 * 360
	pi := pngimage.NewPngimage(maxX-1, 240, "hue.png")
	m := pi.Img
	b := m.Bounds()
	var c color.Color
	var err error
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if c, err = hue.Color(float64(x) / float64(maxX)); err != nil {
				panic(err)
			}
			m.Set(x, y, c)
		}
	}
	pi.Save()
}
