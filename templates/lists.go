package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material []string
}

func main() {
	sweaters := Inventory{Material: []string{"gold", "uran", "wood"}}
	tmpl, err := template.New("test").Parse("The items are made of:\n{{range .Material}}- {{.}}\n{{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
