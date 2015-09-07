package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func stdin2yaml() (s string) {
	if json, err := ioutil.ReadAll(os.Stdin); err != nil {
		s = err.Error()
	} else {
		if y, err := yaml.JSONToYAML(json); err != nil {
			s = err.Error()
		} else {
			s = string(y)
		}
	}
	return s
}

// read json from stdin
func main() {
	if fi, err := os.Stdin.Stat(); err != nil {
		panic(err)
	} else {
		if fi.Size() > 0 {
			fmt.Println(stdin2yaml())
		} else {
			fmt.Println("stdin is empty")
		}
	}
}
