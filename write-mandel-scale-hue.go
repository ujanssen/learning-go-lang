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

func main() {
	var err error
	var c color.Color
	var maxIter = 10 * 360
	var d = mandel.NewData(640, 400, -0.7435669, 0.1314023, 0.0022878)

	for colScale := 30; colScale < 120; colScale += 30 {

		n := fmt.Sprintf("mandel-hue-%04d.png", colScale)
		fmt.Println("Render:", n)
		pi := pngimage.NewPngimage(d.MaxX, d.MaxY, n)
		it := task.IterateTasks(d, pi)

		for t := range it {
			iter := mandel.Iterate(t.CRe, t.CIm, maxIter)
			h := float64(iter) / float64(colScale)
			h = h - math.Floor(h)
			if c, err = hue.Color(h); err != nil {
				panic(err)
			}
			pi.Img.Set(t.X, d.MaxY-t.Y, c)
		}
		pi.Save()
	}
}
