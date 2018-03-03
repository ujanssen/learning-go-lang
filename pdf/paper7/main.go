package main

import (
	"math/rand"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Helvetica", "", 12)

	unit := 5.0
	x := unit + 5.0
	y := unit + 5.0

	rand.Seed(782195697562498)

	for r := 0.0; r < 54; r++ {
		for c := 0.0; c < 38; c++ {

			cx := x + c*unit
			cy := y + r*unit

			sign := 1.0
			if rand.Float64() > 0.5 {
				sign = -1.0
			}
			g := int(rand.Float64() * 255)
			pdf.SetFillColor(g, g, g)
			if r > 2 {
				rx := rand.Float64() * r / 24.0
				ry := rand.Float64() * r / 25.0

				pdf.TransformBegin()
				pdf.TransformRotate(sign*rand.Float64()*r, cx, cy)
				pdf.Rect(cx+rx, cy+ry, unit, unit, "F")
				pdf.TransformEnd()
			} else {
				pdf.Rect(cx, cy, unit, unit, "F")
			}
		}
	}

	pdf.Text(x, 280+y, "paper7")
	pdf.OutputFileAndClose("paper7.pdf")

}
