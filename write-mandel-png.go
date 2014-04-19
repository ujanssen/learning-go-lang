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

	pt := pi.PixelTasks()

	for t := range pt {
		dx := float64(t.X) * d.StepX
		dy := float64(t.Y) * d.StepY
		iter := mandel.Iterate(d.UpperLeftRe+dx, d.UpperLeftIm+dy, 1024)
		c := getColor(iter)
		pi.Img.Set(t.X, d.MaxY-t.Y, c)
	}

	pi.Save()
}
