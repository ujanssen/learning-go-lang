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
	const x = 15.0
	const y = 15.0

	width, height := pdf.GetPageSize()

	halfw := width / 2.0
	halfh := height / 2.0

	for p := 1; p <= 100; p++ {

		pdf.AddPage()

		for i := 0; i < p; i++ {
			x0 := halfw + rand.Float64()*halfw
			y0 := halfh + rand.Float64()*halfh

			cx0 := rand.Float64() * halfw
			cy0 := halfh + rand.Float64()*halfh

			cx1 := rand.Float64() * halfw
			cy1 := rand.Float64() * height

			x1 := halfw + rand.Float64()*halfw
			y1 := rand.Float64() * halfh

			pdf.CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1, "D")
		}
		pdf.Text(x, height-y, fmt.Sprintf("bezier3-%d", p))
	}

	pdf.OutputFileAndClose("bezier3.pdf")
}
