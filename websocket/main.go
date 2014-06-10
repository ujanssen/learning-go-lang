package main

import (
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

func listen() {
	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				log.Println("Error:", err)
			}
		}()
		log.Println("Created onConnected handler")
		log.Println("onConnected: ", ws)

	}
	http.Handle("/websocket", websocket.Handler(onConnected))

}
func main() {

	// websocket server
	go listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
