package main

import (
	"flag"
	"fmt"
	"github.com/ujanssen/learning-go-lang/fritzbox"
)

func main() {
	var host = flag.String("host", "", "fritzbox host or ip address")
	var password = flag.String("password", "", "fritzbox screen password")
	var username = flag.String("username", "", "fritzbox screen username")
	flag.Parse()

	box := fritzbox.NewFritzbox(*host, *username, *password)
	box.Switchlist()

	fmt.Printf("box -> %v\n", box)
}
