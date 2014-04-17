package main

import (
	"image/color"
	"math/rand"
	"pngimage"
)

func getRandomColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}

func main() {
	pi := pngimage.NewPngimage(640, 480, "image.png")
	m := pi.Img
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			m.Set(x, y, getRandomColor())
		}
	}
	pi.Save()
}
