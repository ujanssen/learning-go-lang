package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

var errEmpty = errors.New("stdin is empty")

// read json from stdin
// print as yaml
func main() {
	if err := checkStdin(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(readStdinConvertToYaml())
}

func checkStdin() (noError error) {
	fi, err := os.Stdin.Stat()

	if err != nil {
		return err
	}
	if fi.Size() == 0 {
		return errEmpty
	}

	return noError
}

func readStdinConvertToYaml() (s string) {
	if json, err := ioutil.ReadAll(os.Stdin); err != nil {
		s = err.Error()
	} else {
		s = json2Yaml(json)
	}
	return
}

func json2Yaml(json []byte) (s string) {
	if y, err := yaml.JSONToYAML(json); err != nil {
		s = err.Error()
	} else {
		s = string(y)
	}
	return
}
