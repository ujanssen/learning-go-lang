package main

// emmit a state to all connected clients

import (
	"github.com/ujanssen/learning-go-lang/concurrency/state"
	"log"
	"net/http"
	"time"
)

func main() {

	// websocket server
	server := state.NewServer(state.Path)
	go server.Listen()
	go emitState(server)

	log.Println("Listening web server on " + state.Path + state.Port + " ...")
	log.Fatal(http.ListenAndServe(state.Port, nil))

}

func emitState(server *state.Server) {
	for {
		time.Sleep(1 * time.Second)
		if server.HasClients() {
			log.Println("emit state")
			server.SendAll(&state.Message{"hello", "world"})
		} else {
			log.Println("wait for clients")
		}
	}
}
