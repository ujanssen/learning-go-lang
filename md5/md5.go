package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func HashString(input string) (hash string) {
	h := md5.New()
	io.WriteString(h, input)
	hash = fmt.Sprintf("%x", h.Sum(nil))
	return hash
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
	hash = fmt.Sprintf("%x", h.Sum(nil))
	return hash, err
}
