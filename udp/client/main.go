package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	server, err := net.ResolveUDPAddr("udp", "127.0.0.1:6000")
	checkError(err)

	conn, err := net.DialUDP("udp", nil, server)
	checkError(err)

	defer conn.Close()
	msg := ""
	for {
		buf := []byte(msg)
		_, err := conn.Write(buf)
		checkError(err)

		fmt.Printf("Send %d bytes to %s\n", len(msg), server)

		msg += "X"
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
