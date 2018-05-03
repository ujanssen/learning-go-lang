package main

import (
	"fmt"
	"math/rand"

	"github.com/jung-kurt/gofpdf"
)

// Rand
const (
	X = 15.0
	Y = 15.0
)

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Helvetica", "", 12)
	rand.Seed(782195697562498)

	for p := 1; p <= 100; p++ {

		pdf.AddPage()

		width, height := pdf.GetPageSize()

		for i := 0; i < p; i++ {
			x0, y0 := getRandomPoint(width, height)
			cx0, cy0 := getRandomPoint(width, height)
			cx1, cy1 := getRandomPoint(width, height)
			x1, y1 := getRandomPoint(width, height)

			pdf.CurveBezierCubic(x0, y0, cx0, cy0, cx1, cy1, x1, y1, "D")
		}
		pdf.Text(X, height-Y, fmt.Sprintf("bezier2-%d", p))
	}

	pdf.OutputFileAndClose("bezier2.pdf")

}

func getRandomPoint(width, height float64) (x, y float64) {
	x = X + rand.Float64()*(width-2*X)
	y = Y + rand.Float64()*(height-2*Y)
	return
}
