package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
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
	black         = color.RGBA{0, 0, 0, 255}
	white         = color.RGBA{x, x, x, 255}
)

func getValue(hue float64) uint8 {
	return uint8(255.0 / 60.0 * hue)
}

func getHueColor(hue float64) color.RGBA {
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

func getColor(c int) color.RGBA {
	return getHueColor(float64((c + palShift*10) % 360))
}

func computeMandel(cx, cy float64) color.RGBA {
	var i int
	var abs float64
	var z, c complex128
	c = complex(cx, cy)
	for {
		z = z*z + c
		if i >= maxIter {
			break
		}
		abs = cmplx.Abs(z)
		if abs > 2.0 {
			return getColor(i)
		}
		i = i + 1
	}
	return black
}

var maxIter = 10 * 360
var palShift int

func main() {
	const maxX = 1024
	const maxY = 800

	const centerRe, centerIm float64 = -0.7435669, 0.1314023
	const hd float64 = 0.0022878
	const hv float64 = float64(maxY) / float64(maxX) * hd

	const hdStepX = hd / maxX
	const hdStepY = hv / maxY

	const upperLeftRe = centerRe - (float64(maxX/2) * hdStepX)
	const upperLeftIm = centerIm - (float64(maxY/2) * hdStepY)

	for palShift = 0; palShift < 36; palShift++ {
		n := fmt.Sprintf("mandel-hue-%02d.png", palShift)
		fmt.Println("Render n:", n)
		f, err := os.Create(n)
		if err != nil {
			panic(err)
		}
		m := image.NewRGBA(image.Rect(0, 0, maxX, maxY))
		b := m.Bounds()
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				dx := float64(x) * hdStepX
				dy := float64(y) * hdStepY
				c := computeMandel(upperLeftRe+dx, upperLeftIm+dy)
				m.Set(x, maxY-y, c)
			}
		}

		if err = png.Encode(f, m); err != nil {
			panic(err)
		}
	}
}
