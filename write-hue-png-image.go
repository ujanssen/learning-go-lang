package main

import (
	"hue"
	"image/color"
	"pngimage"
)

func main() {
	var err error
	var c color.Color
	maxX := 3 * 360
	pi := pngimage.NewPngimage(maxX-1, 240, "hue.png")
	b := pi.Img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			if c, err = hue.Color(float64(x) / float64(maxX)); err != nil {
				panic(err)
			}
			pi.Img.Set(x, y, c)
		}
	}
	pi.Save()
}
