package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var (
	x       uint8 = 255
	red           = color.RGBA{x, 0, 0, 255}
	yellow        = color.RGBA{x, x, 0, 255}
	green         = color.RGBA{0, x, 0, 255}
	cyan          = color.RGBA{0, x, x, 255}
	blue          = color.RGBA{0, 0, x, 255}
	magenta       = color.RGBA{x, 0, x, 255}
)

func getValue(hue float64) uint8 {
	return uint8(255.0 / 60.0 * hue)
}

func getColor(hue float64) color.RGBA {
	if hue < 0.0 || hue >= 360.0 {
		fmt.Println("Illegal hue value:", hue)
	}
	var r, g, b uint8

	// red->yellow
	if hue < 60.0 {
		r = x
		g = getValue(hue)
		b = 0
		return color.RGBA{r, g, b, 255}
	}
	// yellow->green
	if hue < 120.0 {
		r = 255 - getValue(hue-60.0)
		g = x
		b = 0
		return color.RGBA{r, g, b, 255}
	}
	// green->cyan
	if hue < 180.0 {
		r = 0
		g = x
		b = getValue(hue - 120.0)
		return color.RGBA{r, g, b, 255}
	}
	// cyan->blue
	if hue < 240.0 {
		r = 0
		g = 255 - getValue(hue-180.0)
		b = x
		return color.RGBA{r, g, b, 255}
	}
	// blue->magenta
	if hue < 300.0 {
		r = getValue(hue - 240.0)
		g = 0
		b = x
		return color.RGBA{r, g, b, 255}
	}
	// magenta->red
	if hue < 360.0 {
		r = x
		g = 0
		b = 255 - getValue(hue-300.0)
		return color.RGBA{r, g, b, 255}
	}
	return color.RGBA{0, 0, 0, 255}
}

func main() {
	f, err := os.Create("hue.png")
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, 3*360-1, 240))
	b := m.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			m.Set(x, y, getColor(float64(x/3.0)))
		}
	}
	if err = png.Encode(f, m); err != nil {
		panic(err)
	}
}
