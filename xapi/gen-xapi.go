package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type XapiTypes struct {
	Xapi []XapiType
}

type XapiType struct {
	Name        string
	Description string
	Fields      []XapiField
	Messages    []XapiMessage
}

type XapiField struct {
	Name        string
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
