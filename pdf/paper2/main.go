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

			rx, ry := 0.0, 0.0

			if r > 1 {
				rx = rand.Float64() * r / 3.0
				ry = rand.Float64() * r / 3.0
			}

			pdf.Rect(rx+x+c*unit, ry+y+r*unit, unit, unit, "D")
		}
	}

	pdf.Text(x, 260+y, "paper2")
	pdf.OutputFileAndClose("paper2.pdf")

}



package main

import (
	"math/rand"

	"github.com/jung-kurt/gofpdf"
)
