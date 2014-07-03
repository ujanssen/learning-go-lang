Use the switch functions of the [FRITZ!DECT 200](http://avm.de/produkte/fritzdect/fritzdect-200/).


	package main

	import (
		"flag"
		"fmt"
		"github.com/ujanssen/learning-go-lang/fritzbox"
	)

	func main() {
		var (
			host     = flag.String("host", "", "fritzbox host or ip address")
			password = flag.String("password", "", "fritzbox screen password")
			username = flag.String("username", "", "fritzbox screen username")
			command  = flag.String("command", "list", "list, on, off, state")
			ain      = flag.String("ain", "", "ain")
		)
		flag.Parse()
		
		box := fritzbox.NewFritzbox(*host, *username, *password)
		fmt.Printf("box -> %v\n", box)

		switch *command {
		case "list":
			box.Switchlist()
		case "on":
			box.SwitchOn(*ain)
		case "off":
			box.SwitchOff(*ain)
		case "state":
			box.SwitchState(*ain)
		}
	}