package main

import (
	"fmt"
	"hue"
	"image/color"
	"mandel"
	"math"
	"pngimage"
)

func getColor(iter float64) color.Color {
	var c color.Color
	var err error
	h := iter / 360.0
	h = h - math.Floor(h)
	if c, err = hue.Color(h); err != nil {
		panic(err)
	}
	return c
}

var maxIter = 10 * 360

func iterateSubPixelTasks(d *mandel.Data, pi *pngimage.Pngimage, subPixel int) {
	spX := d.StepX / float64(subPixel)
	spY := d.StepY / float64(subPixel)
	b := pi.Img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			sumIter := 0
			for sx := 1; sx <= subPixel; sx++ {
				for sy := 1; sy <= subPixel; sy++ {
					cRe := d.UpperLeftRe + float64(x)*d.StepX + float64(sx)*spX
					cIm := d.UpperLeftIm + float64(y)*d.StepY + float64(sy)*spY
					iter := mandel.Iterate(cRe, cIm, maxIter)
					sumIter += iter
				}
			}
			fIter := float64(sumIter) / float64(subPixel*subPixel)
			pi.Img.Set(x, d.MaxY-y, getColor(fIter))
		}
	}
}

func main() {
	var d = mandel.NewData(800, 600, -0.7435669, 0.1314023, 0.0022878)

	for subPixel := 1; subPixel <= 16; subPixel++ {
		n := fmt.Sprintf("mandel-subpixel-%03d.png", subPixel*subPixel)
		fmt.Println("Render:", n)
		pi := pngimage.NewPngimage(d.MaxX, d.MaxY, n)
		iterateSubPixelTasks(d, pi, subPixel)
		pi.Save()
	}
}
