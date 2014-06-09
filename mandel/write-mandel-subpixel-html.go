package main

import "fmt"

// produce a html page to show all mandelbrot set images
// go run write-mandel-subpixel-html.go > write-mandel-subpixel/write-mandel-subpixel.html
func main() {
	fmt.Println("<table>")

	for subPixel := 1; subPixel <= 16; subPixel++ {
		fmt.Println("<tr><td>")
		fmt.Println("<td>")
		fmt.Printf("<img src=\"mandel-subpixel-%03d.png\">\n", subPixel*subPixel)
		fmt.Println("</td>")
		fmt.Println("</td></tr>")
	}
	fmt.Println("</table>")
}
