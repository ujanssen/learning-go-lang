package main

import (
	"fmt"
	"math/rand"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	rand.Seed(782195697562498)

	width, height := pdf.GetPageSize()

	const x = 15.0
	const y = 15.0

	x0 := width / 2
	y0 := y

	x1 := width / 2
	y1 := height - y

	for p := 1; p <= 100; p++ {
		pdf.AddPage()

		cx0 := width/4 + rand.Float64()*width/2
		cy0 := height/4 + rand.Float64()*height/2

		cx1 := cx0 + rand.Float64()*width/8
		cy1 := cy0 + rand.Float64()*height/8

		for i := 0; i < p; i++ {

			cx0 = cx0 + 2.0
			cy0 = cy0 + 2.0

			cx1 = cx1 + 2.0
			cy1 = cy1 + 2.0

			pdf.CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1, "D")
		}
		pdf.Text(x, height-y, fmt.Sprintf("bezier8-%d", p))
	}

	pdf.OutputFileAndClose("bezier8.pdf")
}
