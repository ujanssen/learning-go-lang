package main

import (
	"io"
	"log"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

type Message struct {
	Head string `json:"head"`
	Body string `json:"body"`
}

func (self *Message) String() string {
	return self.Head + " with " + self.Body
}

func receive(ws *websocket.Conn) {
	log.Println("Listening receiving from client")
	var msg Message
	err := websocket.JSON.Receive(ws, &msg)
	if err == io.EOF {
		log.Println("Receive io.EOF:", err)
	} else if err != nil {
		log.Println("Receive Error:", err, msg)
	} else {
		log.Println("Received: ", msg)
	}
}

func listen() {
	log.Println("Listening ws server on /websocket ...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				log.Println("Close Error:", err)
			}
		}()
		log.Println("Created onConnected handler")
		log.Println("onConnected: ", ws)
		go receive(ws)

		websocket.JSON.Send(ws, Message{"nead", "body"})

	}
	http.Handle("/websocket", websocket.Handler(onConnected))

}

func main() {

	// websocket server
	go listen()

	// static files
	http.Handle("/", http.FileServer(http.Dir("static")))

	log.Println("Listening web server on :8044 ...")
	log.Fatal(http.ListenAndServe(":8044", nil))
}
