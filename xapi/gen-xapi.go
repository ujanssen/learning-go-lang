package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type XapiTypes struct {
	Xapi []XapiType
}

type XapiType struct {
	Name        XapiName
	Description string
	Fields      []XapiField
	Messages    []XapiMessage
}

type XapiField struct {
	Name        XapiName
	Description string
	Type        string
	Qualifier   string
}

type XapiMessage struct {
	Name        string
	Description string
	Params      []XapiParam
	Result      []string
}

type XapiParam struct {
	Name string
	Type string
	Doc  string
}

func (name XapiName) Title() string {
	return strings.Title(string(name))
}

type XapiName string

func main() {
	file, e := ioutil.ReadFile("./xapi.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var types XapiTypes
	json.Unmarshal(file, &types)
	t, err := template.New("xapi.template").ParseFiles("xapi.template")
	if err != nil {
		fmt.Printf("execution failed: %s\n", err)
	}

	err = t.ExecuteTemplate(os.Stdout, "xapi.template", types)
	if err != nil {
		fmt.Printf("execution failed: %s\n", err)
	}
}
