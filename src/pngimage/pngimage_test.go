package pngimage

import (
	"pngimage"
	"testing"
)

func TestPixelTasks(t *testing.T) {
	pi := pngimage.NewPngimage(16, 9, "image.png")

	c := pi.PixelTasks()

	if len(c) != 16*9 {
		t.Errorf("len(%v) = %v, want %v", c, len(c), 16*9)
	}

}
