package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/ujanssen/learning-go-lang/mandel"
	"github.com/ujanssen/learning-go-lang/pngimage"

	"github.com/ujanssen/learning-go-lang/task"

	"github.com/ujanssen/learning-go-lang/hue"
)

func getColor(iter, colScale, colShift int) color.Color {
	var c color.Color
	var err error
	h := float64((iter+colShift)%colScale) / float64(colScale)

	h = h - math.Floor(h)
	if c, err = hue.Color(h); err != nil {
		panic(err)
	}
	return c
}

func main() {
	var maxIter = 10 * 360
	var d = mandel.NewData(640, 400, -0.7435669, 0.1314023, 0.0022878)

	for colScale := 30; colScale < 360; colScale += 30 {
		for colShift := 0; colShift < colScale; colShift += 6 {

			n := fmt.Sprintf("mandel-scale-%04d-shift-%03d.png", colScale, colShift)
			fmt.Println("Render:", n)
			pi := pngimage.NewPngimage(d.MaxX, d.MaxY, n)
			it := task.IterateTasks(d, pi)

			for t := range it {
				iter := mandel.Iterate(t.CRe, t.CIm, maxIter)
				pi.Img.Set(t.X, d.MaxY-t.Y, getColor(iter, colScale, colShift))
			}
			pi.Save()
		}
	}
}
