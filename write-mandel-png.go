package main

import (
	"image/color"
	"mandel"
	"pngimage"
)

func getColor(c int) color.RGBA {
	rgb := uint8(c % 255)
	return color.RGBA{rgb, rgb, rgb, 255}
}

func main() {
	d := mandel.NewData(1024, 800, -0.7435669, 0.1314023, 0.0022878)
	pi := pngimage.NewPngimage(d.MaxX, d.MaxY, "mandel.png")

	for y := 0; y < d.MaxY; y++ {
		for x := 0; x < d.MaxX; x++ {
			dx := float64(x) * d.StepX
			dy := float64(y) * d.StepY
			iter := mandel.Iterate(d.UpperLeftRe+dx, d.UpperLeftIm+dy, 1024)
			c := getColor(iter)
			pi.Img.Set(x, d.MaxY-y, c)
		}
	}
	pi.Save()
}
