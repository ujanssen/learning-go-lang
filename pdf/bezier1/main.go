package main

import (
	"math/rand"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Helvetica", "", 12)
	rand.Seed(782195697562498)

	const x = 15.0
	const y = 15.0

	width, height := pdf.GetPageSize()
	bottom := height - 2*y
	half := width / 2.0

	for i := 0; i < 100; i++ {
		x0 := half
		y0 := bottom

		cx0 := half
		cy0 := rand.Float64() * height

		cx1 := rand.Float64() * width
		cy1 := rand.Float64() * height

		x1 := rand.Float64() * width
		y1 := rand.Float64() * height

		pdf.CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1, "D")
	}
	pdf.Text(x, height-y, "bezier1")
	pdf.OutputFileAndClose("bezier1.pdf")

}
