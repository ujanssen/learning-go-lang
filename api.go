package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// request:  http://localhost:8080/bar
	// response: Thanks for the GET!
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thanks for the %s!", r.Method)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
