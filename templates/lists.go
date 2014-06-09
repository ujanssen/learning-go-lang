package main

import (
	"os"
	"text/template"
)

type Materials struct {
	Material []string
}

func main() {
	materials := Materials{Material: []string{"gold", "uran", "wood"}}
	tmpl, err := template.New("materials").Parse("The materials are:\n{{range .Material}}- {{.}}\n{{end}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, materials)
	if err != nil {
		panic(err)
	}
}
