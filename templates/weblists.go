package main

import (
	"html/template"
	"log"
	"net/http"
)

type Materials struct {
	Material []string
}
type Templ struct {
	Text string
}

func main() {
	http.Handle("/", http.HandlerFunc(HandleMaterials))
	log.Fatal("ListenAndServe:", http.ListenAndServe(":8088", nil))
}

func HandleMaterials(w http.ResponseWriter, req *http.Request) {

	t, err := template.New("page").ParseFiles("weblists.html", "materials.html")
	if err != nil {
		panic(err)
	}

	materials := Materials{Material: []string{"gold", "uran", "wood"}}

	err = t.ExecuteTemplate(w, "weblist", materials)
	if err != nil {
		panic(err)
	}
}
