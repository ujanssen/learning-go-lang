package main

import (
	"math/rand"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Helvetica", "", 12)

	unit := 20.0
	x := unit + 5.0
	y := unit + 5.0

	rand.Seed(782195697562498)

	for r := 0.0; r < 12; r++ {
		for c := 0.0; c < 8; c++ {

			cx := x + c*unit
			cy := y + r*unit

			sign := 1.0
			if rand.Float64() > 0.5 {
				sign = -1.0
			}

			if r > 2 {
				rx := rand.Float64() * r / 3.0
				ry := rand.Float64() * r / 3.0

				pdf.TransformBegin()
				pdf.TransformRotate(sign*rand.Float64()*r, cx, cy)
				pdf.Rect(cx+rx, cy+ry, unit, unit, "D")
				pdf.TransformEnd()
			} else {
				pdf.Rect(cx, cy, unit, unit, "D")
			}
		}
	}

	pdf.Text(x, 260+y, "paper3")
	pdf.OutputFileAndClose("paper3.pdf")

}
