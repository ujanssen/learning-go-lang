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
	var command = flag.String("command", "list", "list, on, off, state")
	var ain = flag.String("ain", "", "ain")

	flag.Parse()

	box := fritzbox.NewFritzbox(*host, *username, *password)
	fmt.Printf("box -> %v\n", box)

	switch *command {
	case "list":
		box.Switchlist()
	case "on":
		box.SwitchOn(*ain)
	}

}
