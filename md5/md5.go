package md5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Hash(input string) string {
	h := md5.New()
	io.WriteString(h, input)
	md5Sum := h.Sum(nil)
	return fmt.Sprintf("%x", md5Sum)
}
