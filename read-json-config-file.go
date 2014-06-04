package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Worker int    `json:"worker"`
	Tasks  int    `json:"tasks"`
	Type   string `json:"type"`
}

func main() {
	var conf Config
	dat, err := ioutil.ReadFile("read-json-config-file.conf")
	if err != nil {
		panic(err)
	}
	fmt.Println("dat", dat)

	if err = json.Unmarshal(dat, &conf); err != nil {
		panic(err)
	}
	fmt.Println("conf", conf)
}
