// Package hue can be used to convert a hue float value to a color.
package hue

import (
	"errors"
	"fmt"
	"image/color"
	"math"
)

var segmentSize float64 = 1.0 / 6.0

// Get color of a hue value.
// The hue value must be >= 0.0 and < 1.0.
func Color(hue float64) (color.Color, error) {
	if hue < 0.0 || hue >= 1.0 {
		errMsg := fmt.Sprintf("mandel.Color: arg hue float64 must be >= 0.0 and < 1.0, was %v", hue)
		return color.Black, errors.New(errMsg)
	}

	var segment = uint8(math.Floor(hue / segmentSize))
	var v = uint8(math.Floor(0.5 + 255.0/segmentSize*(hue-float64(segment)*segmentSize)))

	var r, g, b uint8
	switch segment {
	case 0: // red->yellow 0->60
		r = 255
		g = v
		b = 0
	case 1: // yellow->green 60->120
		r = 255 - v
		g = 255
		b = 0
	case 2: //green->cyan 120->180
		r = 0
		g = 255
		b = v
	case 3: // cyan->blue 180->240
		r = 0
		g = 255 - v
		b = 255
	case 4: // blue->magenta 240->300
		r = v
		g = 0
		b = 255
	case 5: //magenta->red 300->360
		r = 255
		g = 0
		b = 255 - v
	}
	return color.RGBA{r, g, b, 255}, nil
}
