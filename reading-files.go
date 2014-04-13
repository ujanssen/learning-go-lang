package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	dat, err := ioutil.ReadFile("reading-files.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println("dat", dat)
}
