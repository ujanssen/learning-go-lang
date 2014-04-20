package main

import "fmt"

// go run write-mandel-scale-shift-html.go > mandel-scale-shift/mandel-scale-shift.html
func main() {
	fmt.Println("<table>")
	for colScale := 30; colScale < 360; colScale += 30 {
		fmt.Println("<tr>")
		for colShift := 0; colShift < colScale; colShift += 6 {
			fmt.Println("<td>")
			fmt.Printf("<img src=\"mandel-scale-%04d-shift-%03d.png\">\n", colScale, colShift)
			fmt.Println("</td>")
		}
		fmt.Println("</tr>")
	}
	fmt.Println("</table>")
}
