package main

import (
	"fmt"
	"hue"
	"image/color"
	"mandel"
	"math"
	"pngimage"
	"task"
)

func getColor(iter, colScale int) color.Color {
	var c color.Color
	var err error
	h := float64(iter) / float64(colScale)
	h = h - math.Floor(h)
	if c, err = hue.Color(h); err != nil {
		panic(err)
	}
	return c
}

func main() {
	var maxIter = 10 * 360
	var d = mandel.NewData(640, 400, -0.7435669, 0.1314023, 0.0022878)

	for colScale := 30; colScale < 120; colScale += 30 {

		n := fmt.Sprintf("mandel-hue-%04d.png", colScale)
		fmt.Println("Render:", n)
		pi := pngimage.NewPngimage(d.MaxX, d.MaxY, n)
		it := task.IterateTasks(d, pi)

		for t := range it {
			iter := mandel.Iterate(t.CRe, t.CIm, maxIter)
			pi.Img.Set(t.X, d.MaxY-t.Y, getColor(iter, colScale))
		}
		pi.Save()
	}
}
