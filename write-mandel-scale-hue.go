package main

import (
	"fmt"
	"hue"
	"image/color"
	"mandel"
	"math"
	"pngimage"
)

func main() {
	var err error
	var c color.Color
	var maxIter = 10 * 360
	var d = mandel.NewData(640, 400, -0.7435669, 0.1314023, 0.0022878)

	for colScale := 30; colScale < 120; colScale += 30 {

		n := fmt.Sprintf("mandel-hue-%04d.png", colScale)
		fmt.Println("Render n:", n)
		pi := pngimage.NewPngimage(d.MaxX, d.MaxY, n)

		for y := 0; y < d.MaxY; y++ {
			for x := 0; x < d.MaxX; x++ {
				cRe := d.UpperLeftRe + float64(x)*d.StepX
				cIm := d.UpperLeftIm + float64(y)*d.StepY
				iter := mandel.Iterate(cRe, cIm, maxIter)
				h := float64(iter) / float64(colScale)
				h = h - math.Floor(h)
				if c, err = hue.Color(h); err != nil {
					panic(err)
				}
				pi.Img.Set(x, d.MaxY-y, c)
			}
		}
		pi.Save()
	}
}
