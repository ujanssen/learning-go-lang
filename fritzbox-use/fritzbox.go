package main

import (
	"flag"
	"fmt"
	"github.com/ujanssen/learning-go-lang/fritzbox"
)

func main() {
	var password = flag.String("password", "", "fritzbox screen password")
	var username = flag.String("username", "", "fritzbox screen username")
	flag.Parse()

	box := fritzbox.NewFritzbox(*username, *password)

	fmt.Printf("box -> %v\n", box)
}
