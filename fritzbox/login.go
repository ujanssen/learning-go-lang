package main

//curl  http://192.168.12.1/login_sid.lua

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// SessionInfo has SID
type SessionInfo struct {
	SID       string `xml:"SID"`
	Challenge string `xml:"Challenge"`
	BlockTime string `xml:"BlockTime"`
}

func main() {
	response, err := http.Get("http://fritz/login_sid.lua")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))

		var s SessionInfo
		err = xml.Unmarshal(contents, &s)

		fmt.Printf("SessionInfo: %s\n", s)
	}
}
