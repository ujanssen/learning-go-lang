package fritzbox

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io"
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

type Fritzbox struct {
	Host, Username, Password string

	LoginURL, CommandURL string

	sid     string
	ainList string
}

func NewFritzbox(host, username, password string) *Fritzbox {
	box := Fritzbox{
		Host:       host,
		Username:   username,
		Password:   password,
		LoginURL:   "http://" + host + "/login_sid.lua",
		CommandURL: "http://" + host + "/webservices/homeautoswitch.lua"}

	(&box).login()
	return &box
}

func (box *Fritzbox) SwitchOn(ain string) {
	box.switchCommand("setswitchon", ain)
}
func (box *Fritzbox) SwitchOff(ain string) {
	box.switchCommand("setswitchoff", ain)
}
func (box *Fritzbox) SwitchState(ain string) {
	box.switchCommand("getswitchstate", ain)
}
func (box *Fritzbox) Switchlist() {
	ain := box.switchCommand("getswitchlist", "")
	ain = ain[0 : len(ain)-1]
	fmt.Printf("ain: %s\n", ain)
	box.ainList = ain
}

func (box *Fritzbox) switchCommand(switchcmd, ain string) (resp string) {
	values := url.Values{}
	values.Set("switchcmd", switchcmd)
	values.Set("sid", box.sid)
	if len(ain) > 0 {
		values.Set("ain", ain)
	}
	response, err := http.Get(box.CommandURL + "?" + values.Encode())
	fmt.Printf("values -> %v\n", values)

	checkError(err)
	contents := readBody(response.Body)
	resp = string(contents)
	fmt.Printf("response: %s\n", resp)
	return resp
}

func (box *Fritzbox) login() {
	challenge := box.challenge()

	text := challenge + "-" + box.Password

	hash := md5Hash(utf16Encode(text))
	resp := challenge + "-" + hash
	fmt.Printf("response -> %s\n", resp)

	values := url.Values{}
	values.Set("username", box.Username)
	values.Set("response", resp)
	response, err := http.PostForm(box.LoginURL, values)
	fmt.Printf("values -> %v\n", values)

	checkError(err)
	contents := readBody(response.Body)
	var s SessionInfo
	fmt.Printf("%s\n", string(contents))
	err = xml.Unmarshal(contents, &s)
	fmt.Printf("SessionInfo: %s\n", s)
	box.sid = s.SID
}
func (box *Fritzbox) challenge() string {
	var s SessionInfo
	response, err := http.Get(box.LoginURL)
	checkError(err)
	contents := readBody(response.Body)
	fmt.Printf("%s\n", string(contents))
	err = xml.Unmarshal(contents, &s)
	fmt.Printf("SessionInfo: %s\n", s)
	return s.Challenge
}

func utf16Encode(test string) []uint16 {
	runes := []rune(test)
	return utf16.Encode(runes)
}

func md5Hash(data []uint16) (hash string) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, data)
	checkError(err)
	hash = fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
	return hash
}

func readBody(body io.ReadCloser) []byte {
	defer body.Close()
	contents, err := ioutil.ReadAll(body)
	checkError(err)
	return contents
}
func checkError(err error) {
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
}
