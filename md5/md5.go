package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func HashString(input string) string {
	h := md5.New()
	io.WriteString(h, input)
	md5Sum := h.Sum(nil)
	return fmt.Sprintf("%x", md5Sum)
}

func HashFile(file string) (hash string, err error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return hash, err
	}
	_, err = f.Stat()
	if err != nil {
		return hash, err
	}
	h := md5.New()
	io.Copy(h, f)
	b := h.Sum(nil)
	hash = fmt.Sprintf("%x", b)
	return hash, err
}
