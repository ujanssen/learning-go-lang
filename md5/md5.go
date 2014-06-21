package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	empty := "d41d8cd98f00b204e9800998ecf8427e"
	h := md5.New()
	io.WriteString(h, "")
	md5Sum := h.Sum(nil)
	fmt.Printf("empty  %s\n", empty)
	fmt.Printf("md5Sum %x\n", md5Sum)
}
