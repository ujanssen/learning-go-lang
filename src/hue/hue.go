package hue

import (
	"fmt"
	"image/color"
	"math"
)

var (
	max  float64 = 1.0
	step float64 = max / 6
)

func Color(hue float64) color.RGBA {
	if hue < 0.0 || hue >= max {
		fmt.Println("Illegal hue value:", hue)
	}
	var r, g, b uint8

	var part uint8 = uint8(math.Floor(hue / step))
	var v = uint8(255.0 / step * (hue - float64(part)*step))

	switch part {
	case 0: // red->yellow 0->60
		r = 255
		g = v
		b = 0
		return color.RGBA{r, g, b, 255}
	case 1: // yellow->green 60->120
		r = 255 - v
		g = 255
		b = 0
		return color.RGBA{r, g, b, 255}
	case 2: //green->cyan 120->180
		r = 0
		g = 255
		b = v
		return color.RGBA{r, g, b, 255}
	case 3: // cyan->blue 180->240
		r = 0
		g = 255 - v
		b = 255
		return color.RGBA{r, g, b, 255}
	case 4: // blue->magenta 240->300
		r = v
		g = 0
		b = 255
		return color.RGBA{r, g, b, 255}
	case 5: //magenta->red 300->360
		r = 255
		g = 0
		b = 255 - v
		return color.RGBA{r, g, b, 255}
	}
	return color.RGBA{0, 0, 0, 255}
}
