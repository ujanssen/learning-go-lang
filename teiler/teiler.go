package main

import (
	"fmt"
)

func main() {
	for z := 1; z < 1000000; z++ {
		teiler := make([]int, 0)
		for t := 1; t <= z; t++ {
			if z%t == 0 {
				teiler = append(teiler, t)
			}
		}
		fmt.Printf("%3d %6d %v\n", len(teiler), z, teiler)
	}
}
