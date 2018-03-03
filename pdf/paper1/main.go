package main

import "github.com/jung-kurt/gofpdf"

func main() {

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Helvetica", "", 12)

	const x = 15.0
	const y = 15.0

	for r := 0.0; r < 26; r++ {
		for c := 0.0; c < 18; c++ {
			pdf.Rect(x+c*10, y+r*10, 10.0, 10.0, "D")
		}
	}
	pdf.Text(x, 270+y, "paper1")
	pdf.OutputFileAndClose("paper1.pdf")

}
