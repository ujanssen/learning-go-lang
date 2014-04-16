package hue

import (
	"hue"
	"image/color"
	"testing"
)

var (
	step          = 1.0 / 6.0
	x       uint8 = 255
	red           = color.RGBA{x, 0, 0, 255}
	yellow        = color.RGBA{x, x, 0, 255}
	green         = color.RGBA{0, x, 0, 255}
	cyan          = color.RGBA{0, x, x, 255}
	blue          = color.RGBA{0, 0, x, 255}
	magenta       = color.RGBA{x, 0, x, 255}
)

func TestRed(t *testing.T) {
	color := hue.Color(0 * step)
	if red != color {
		t.Errorf("hue.Color(%v) = %v, want %v", 0*step, red, color)
	}
}
func TestYellow(t *testing.T) {
	color := hue.Color(1 * step)
	if yellow != color {
		t.Errorf("hue.Color(%v) = %v, want %v", 1*step, yellow, color)
	}
}

func TestGreen(t *testing.T) {
	color := hue.Color(2 * step)
	if green != color {
		t.Errorf("hue.Color(%v) = %v, want %v", 2*step, green, color)
	}
}
func TestCyan(t *testing.T) {
	color := hue.Color(3 * step)
	if cyan != color {
		t.Errorf("hue.Color(%v) = %v, want %v", 3*step, cyan, color)
	}
}

func TestBlue(t *testing.T) {
	color := hue.Color(4 * step)
	if blue != color {
		t.Errorf("hue.Color(%v) = %v, want %v", 4*step, blue, color)
	}
}
func TestMagenta(t *testing.T) {
	color := hue.Color(5 * step)
	if magenta != color {
		t.Errorf("hue.Color(%v) = %v, want %v", 5*step, magenta, color)
	}
}
