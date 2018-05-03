package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	clientInput, err := net.ResolveUDPAddr("udp", ":6000")
	checkError(err)

	serverOutput, err := net.ResolveUDPAddr("udp", ":8000")
	checkError(err)

	server, err := net.DialUDP("udp", nil, serverOutput)
	checkError(err)

	client, err := net.ListenUDP("udp", clientInput)
	checkError(err)

	sclient, err := net.ListenUDP("udp", serverOutput)
	checkError(err)

	go func() {
		for {
			buf := make([]byte, 9*1024)
			n, addr, err := client.ReadFromUDP(buf)
			checkError(err)
			fmt.Printf("Received %d bytes from %s\n", n, addr)
			fmt.Printf("%v\n", buf[:n])
			server.Write(buf[:n])
		}
	}()
	go func() {
		for {
			buf := make([]byte, 9*1024)
			n, addr, err := sclient.ReadFromUDP(buf)
			checkError(err)
			fmt.Printf("Received %d bytes from %s\n", n, addr)
			fmt.Printf("%v\n", buf[:n])
		}
	}()

	select {}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
