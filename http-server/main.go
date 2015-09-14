package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		defer r.Body.Close()
		contents, _ := ioutil.ReadAll(r.Body)
		log.Printf("%s %s %s\n", r.Method, r.URL.Path[1:], contents)

	default:
		log.Printf("%s %s\n", r.Method, r.URL.Path[1:])
	}

}

func main() {
	port := flag.String("port", "80", "port to listen to")
	flag.Parse()
	http.HandleFunc("/", handler)
	fmt.Println("listen on port ", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
