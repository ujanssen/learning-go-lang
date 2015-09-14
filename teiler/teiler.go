package main

import (
	"fmt"
)

func mainNew(max int) {
	for z := 1; z < max; z++ {
		var teiler []int
		teiler = append(teiler, 1)
		for t := 2; t <= z/2; t++ {
			if z%t == 0 {
				teiler = append(teiler, t)
			}
		}
		if z > 1 {
			teiler = append(teiler, z)
		}
		fmt.Printf("%3d %6d %v\n", len(teiler), z, teiler)
	}
}
func main() {
	max := 100000
	mainNew(max)
}
