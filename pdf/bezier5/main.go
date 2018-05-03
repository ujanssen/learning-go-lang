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

	for p := 1; p <= 100; p++ {
		pdf.AddPage()

		for h := 0; h < 6; h++ {

			x0 := rand.Float64() * width
			y0 := rand.Float64() * height

			cx0 := rand.Float64() * width
			cy0 := rand.Float64() * height

			cx1 := rand.Float64() * width
			cy1 := rand.Float64() * height

			x1 := rand.Float64() * width
			y1 := rand.Float64() * height

			for i := 0; i < p; i++ {
				x0 = x0 + 1.0
				y0 = y0 + 1.0

				x1 = x1 + 1.0
				y1 = y1 + 1.0

				cx0 = cx0 + 1.0
				cy0 = cy0 + 1.0

				cx1 = cx1 + 1.0
				cy1 = cy1 + 1.0

				pdf.CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1, "D")
			}
		}
		const x = 15.0
		const y = 15.0
		pdf.Text(x, height-y, fmt.Sprintf("bezier5-%d", p))
	}

	pdf.OutputFileAndClose("bezier5.pdf")
}
