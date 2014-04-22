package main

import (
	"image/color"
	"mandel"
	"pngimage"
	"task"
)

func getColor(c int) color.RGBA {
	rgb := uint8(c % 255)
	return color.RGBA{rgb, rgb, rgb, 255}
}

func main() {
	d := mandel.NewData(1024, 800, -0.7435669, 0.1314023, 0.0022878)
	pi := pngimage.NewPngimage(d.MaxX, d.MaxY, "mandel.png")
	it := task.IterateTasks(d, pi)

	for t := range it {
		iter := mandel.Iterate(t.CRe, t.CIm, 1024)
		c := getColor(iter)
		pi.Img.Set(t.X, d.MaxY-t.Y, c)
	}
	pi.Save()
}
