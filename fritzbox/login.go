package main

//curl  http://fritz/login_sid.lua

import (
	// "encoding/xml"
	"crypto/md5"
	"fmt"

	// "io/ioutil"
	// "net/http"
	// "os"
	"bytes"
	"encoding/binary"
	"unicode/utf16"
)

// SessionInfo has SID
type SessionInfo struct {
	SID       string `xml:"SID"`
	Challenge string `xml:"Challenge"`
	BlockTime string `xml:"BlockTime"`
}

func main() {
	// response, err := http.Get("http://fritz/login_sid.lua")
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// 	os.Exit(1)
	// } else {
	// 	defer response.Body.Close()
	// 	contents, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		fmt.Printf("%s", err)
	// 		os.Exit(1)
	// 	}
	// 	fmt.Printf("%s\n", string(contents))

	// 	var s SessionInfo
	// 	err = xml.Unmarshal(contents, &s)

	// 	fmt.Printf("SessionInfo: %s\n", s)
	// }

	challenge := "1234567z"
	password := "äbc"
	response := "1234567z-9e224a41eeefa284df7bb0f26c2913e2"
	text := challenge + "-" + password

	hash := md5Hash(UTF16LE(text))
	fmt.Printf("response -> %s\n", challenge+"-"+hash)
	fmt.Printf("response -> %s\n", response)
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
