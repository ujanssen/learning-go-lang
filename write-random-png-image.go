package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func getRandomColor() color.RGBA {
	return color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
}

func main() {
	f, err := os.OpenFile("imgage.png", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			m.Set(x, y, getRandomColor())
		}
	}
	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
