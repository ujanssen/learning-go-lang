package main

//curl  http://fritz/login_sid.lua

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"unicode/utf16"
)

// SessionInfo has SID
type SessionInfo struct {
	SID       string `xml:"SID"`
	Challenge string `xml:"Challenge"`
	BlockTime string `xml:"BlockTime"`
}

func main() {
	var password = flag.String("password", "", "fritzbox screen password")
	var username = flag.String("username", "", "fritzbox screen username")
	flag.Parse()

	var s SessionInfo

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

		err = xml.Unmarshal(contents, &s)

		fmt.Printf("SessionInfo: %s\n", s)
	}

	text := s.Challenge + "-" + *password

	hash := md5Hash(UTF16LE(text))
	sid := s.Challenge + "-" + hash
	fmt.Printf("response -> %s\n", sid)

	values := url.Values{}
	values.Set("username", *username)
	values.Set("response", sid)
	response, err = http.PostForm("http://fritz/login_sid.lua", values)
	fmt.Printf("values -> %v\n", values)

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

		err = xml.Unmarshal(contents, &s)

		fmt.Printf("SessionInfo: %s\n", s)
	}
}

func UTF16LE(in string) []uint16 {
	runes := []rune(in)
	return utf16.Encode(runes)
}

func md5Hash(data []uint16) (hash string) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	hash = fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
	return hash
}
