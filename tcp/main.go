package main

import (
	"fmt"
	"net"
)

// starting a TCP echo server using a random port
//
// test with:
// echo  "test out the server" | nc localhost 56670

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		handleError(err)
	}
	port := l.Addr().(*net.TCPAddr).Port
	fmt.Printf("Listen on port %d\n", port)

	for {
		conn, err := l.Accept()
		if err != nil {
			handleError(err)
		}
		go func(c net.Conn) {
			buf := make([]byte, 1024)
			len, err := c.Read(buf)
			if err != nil {
				handleError(err)
			}
			c.Write(buf[0:len])
			c.Close()
		}(conn)
	}
}

func handleError(err error) {
	panic(err.Error())
}

