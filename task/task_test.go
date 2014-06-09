package task

import (
	"mandel"
	"pngimage"
	"task"
	"testing"
)

func TestPixelTasks(t *testing.T) {
	pi := pngimage.NewPngimage(16, 9, "image.png")
	d := mandel.NewData(1024, 800, -0.7435669, 0.1314023, 0.0022878)
	it := task.IterateTasks(d, pi)

	if len(it) != 16*9 {
		t.Errorf("len(%v) = %v, want %v", it, len(it), 16*9)
	}

}
