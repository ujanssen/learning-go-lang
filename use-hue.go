package main

import (
	"hue"
	"image/color"
	"math"
	"pngimage"
)

func main() {
	var c color.Color
	var err error
	pi := pngimage.NewPngimage(360-1, 240, "use-hue.png")
	b := pi.Img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			h := float64(x) / 60.0
			h = h - math.Floor(h)
			if c, err = hue.Color(h); err != nil {
				panic(err)
			}
			pi.Img.Set(x, y, c)
		}
	}
	pi.Save()
}
