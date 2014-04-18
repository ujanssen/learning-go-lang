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

func Each(f callback(x,y int))  {
	for y := 0; y < d.MaxY; y++ {
		for x := 0; x < d.MaxX; x++ {
			callback(x,y)
		}
	}
}

