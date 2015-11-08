package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	server, err := net.ResolveUDPAddr("udp", "127.0.0.1:6000")
	checkError(err)

	conn, err := net.DialUDP("udp", nil, server)
	checkError(err)

	defer conn.Close()
	i := 0
	for {
		msg := fmt.Sprintf("A%d", i)
		i++
		buf := []byte(msg)
		_, err := conn.Write(buf)
		fmt.Printf("Send %s %v to %s\n", msg, buf, server)

		checkError(err)
		time.Sleep(time.Second * 1)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
