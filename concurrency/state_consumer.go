package main

import (
	"code.google.com/p/go.net/websocket"
	"github.com/ujanssen/learning-go-lang/concurrency/state"
	"log"
	"time"
)

var origin = "http://localhost/"
var url = "ws://localhost" + state.Port + state.Path

func main() {
	for {
		ws := dial()
		consume(ws)
	}
}

func dial() *websocket.Conn {
	for {
		ws, err := websocket.Dial(url, "", origin)
		if err == nil {
			return ws
		} else {
			log.Println(err)
			time.Sleep(1 * time.Second)
		}
	}
}

func consume(ws *websocket.Conn) (err error) {
	var s state.Message
	for {
		err = websocket.JSON.Receive(ws, &s)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Receive: %s\n", s)
	}
}
