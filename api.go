package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := json.MarshalIndent(r, "", "\t")
		fmt.Fprintf(w, "%v", string(resp))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	/*
	   	// request:  http://localhost:8080/request
	   	// response:

	   {
	   	"Method": "GET",
	   	"URL": {
	   		"Scheme": "",
	   		"Opaque": "",
	   		"User": null,
	   		"Host": "",
	   		"Path": "/request",
	   		"RawQuery": "",
	   		"Fragment": ""
	   	},
	   	"Proto": "HTTP/1.1",
	   	"ProtoMajor": 1,
	   	"ProtoMinor": 1,
	   	"Header": {
	   		"Accept": [
	   			"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,* /*;q=0.8"
	   		],
	   		"Accept-Encoding": [
	   			"gzip,deflate,sdch"
	   		],
	   		"Accept-Language": [
	   			"de-DE,de;q=0.8,en-US;q=0.6,en;q=0.4"
	   		],
	   		"Cache-Control": [
	   			"max-age=0"
	   		],
	   		"Connection": [
	   			"keep-alive"
	   		],
	   		"Dnt": [
	   			"1"
	   		],
	   		"User-Agent": [
	   			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.137 Safari/537.36"
	   		]
	   	},
	   	"Body": {
	   		"Closer": {
	   			"Reader": null
	   		}
	   	},
	   	"ContentLength": 0,
	   	"TransferEncoding": null,
	   	"Close": false,
	   	"Host": "localhost:8080",
	   	"Form": null,
	   	"PostForm": null,
	   	"MultipartForm": null,
	   	"Trailer": null,
	   	"RemoteAddr": "[::1]:61199",
	   	"RequestURI": "/request",
	   	"TLS": null
	   }
	*/

}
