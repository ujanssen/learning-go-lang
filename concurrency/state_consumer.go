package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/ujanssen/learning-go-lang/concurrency/state"
	"log"
)

var origin = "http://localhost/"
var url = "ws://localhost" + state.Port + state.Path

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	for {
		_, err = ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive: %s\n", msg)
	}
}
