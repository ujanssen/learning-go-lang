package pngimage

import (
	"image"
	"image/png"
	"os"
)

type Pngimage struct {
	Name string
	Img  *image.RGBA
}

func NewPngimage(maxX, maxY int, name string) *Pngimage {
	img := image.NewRGBA(image.Rect(0, 0, maxX, maxY))

	pi := new(Pngimage)
	pi.Img = img
	pi.Name = name

	return pi
}

func (pi *Pngimage) Save() {
	f, err := os.Create(pi.Name)
	if err != nil {
		panic(err)
	}
	if err := png.Encode(f, pi.Img); err != nil {
		panic(err)
	}
}

type PixelTask struct {
	X, Y int
}

func (pi *Pngimage) PixelTasks() chan PixelTask {

	b := pi.Img.Bounds()
	input := make(chan PixelTask, b.Max.Y*b.Max.X)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			input <- PixelTask{X: x, Y: y}
		}
	}
	close(input)
	return input
}
