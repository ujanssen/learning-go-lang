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
	Type        XapiTypeName
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

func (name XapiName) String() string {
	s := string(name)
	return strings.Title(s)
}

func (name XapiTypeName) String() string {
	s := string(name)
	if s == "string" || s == "bool" {
		return s
	}
	s = strings.Replace(s, " ", "_", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	s = strings.Replace(s, "->", "", -1)

	if s == "string__string_map" {
		return "map[string]string"
	}

	return "string // " + s
}

type XapiName string
type XapiTypeName string

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
