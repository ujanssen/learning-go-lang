package task

import (
	"github.com/ujanssen/learning-go-lang/mandel"
	"github.com/ujanssen/learning-go-lang/pngimage"
)

// IterateTask with x,z,re and img
type IterateTask struct {
	X, Y     int
	CRe, CIm float64
}

// IterateTasks generator
func IterateTasks(d *mandel.Data, pi *pngimage.Pngimage) chan IterateTask {
	b := pi.Img.Bounds()
	input := make(chan IterateTask, b.Max.Y*b.Max.X)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			cRe := d.UpperLeftRe + float64(x)*d.StepX
			cIm := d.UpperLeftIm + float64(y)*d.StepY
			input <- IterateTask{X: x, Y: y, CRe: cRe, CIm: cIm}
		}
	}
	close(input)
	return input
}
