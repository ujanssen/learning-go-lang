package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	server, err := net.ResolveUDPAddr("udp", ":6000")
	checkError(err)

	conn, err := net.ListenUDP("udp", server)
	checkError(err)
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		checkError(err)
		fmt.Printf("Received %s %v from %s\n", string(buf[0:n]), buf[0:n], addr)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
