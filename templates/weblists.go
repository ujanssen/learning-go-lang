package main

import (
	"html/template"
	"log"
	"net/http"
)

type Materials struct {
	Material []string
}

func main() {
	http.Handle("/", http.HandlerFunc(HandleMaterials))
	log.Fatal("ListenAndServe:", http.ListenAndServe(":8088", nil))
}

func HandleMaterials(w http.ResponseWriter, req *http.Request) {
	materials := Materials{Material: []string{"gold", "uran", "wood"}}
	tmpl, err := template.New("materials").Parse(templateStr)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, materials)
	if err != nil {
		panic(err)
	}
}

const templateStr = `
<html>
<head><title>Materials</title></head>
<body>
The materials are:<ul>{{range .Material}}
  <li>{{.}}</li>
{{end}}</ul>
</body>
</html>
`
