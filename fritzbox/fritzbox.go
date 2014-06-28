package fritzbox

//curl  http://fritz/login_sid.lua

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
	sid                      string
	ainList                  string
}

func NewFritzbox(host, username, password string) *Fritzbox {
	box := Fritzbox{
		Host:     host,
		Username: username,
		Password: password}

	(&box).getsid()
	return &box
}

func (box *Fritzbox) getsid() {
	challenge := box.challenge()
	box.login(challenge)
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
	// get ain
	values := url.Values{}
	values.Set("switchcmd", switchcmd)
	values.Set("sid", box.sid)
	if len(ain) > 0 {
		values.Set("ain", ain)
	}
	response, err := http.Get("http://" + box.Host + "/webservices/homeautoswitch.lua?" + values.Encode())
	fmt.Printf("values -> %v\n", values)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		contents := readBody(response.Body)
		resp = string(contents)
		fmt.Printf("response: %s\n", resp)
	}
	return resp
}

func utf16Encode(in string) []uint16 {
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

func (box *Fritzbox) login(challenge string) {
	text := challenge + "-" + box.Password

	hash := md5Hash(utf16Encode(text))
	resp := challenge + "-" + hash
	fmt.Printf("response -> %s\n", resp)

	values := url.Values{}
	values.Set("username", box.Username)
	values.Set("response", resp)
	response, err := http.PostForm("http://"+box.Host+"/login_sid.lua", values)
	fmt.Printf("values -> %v\n", values)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		contents := readBody(response.Body)
		var s SessionInfo
		fmt.Printf("%s\n", string(contents))
		err = xml.Unmarshal(contents, &s)
		fmt.Printf("SessionInfo: %s\n", s)
		box.sid = s.SID
	}
}
func (box *Fritzbox) challenge() string {
	var s SessionInfo
	response, err := http.Get("http://" + box.Host + "/login_sid.lua")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		contents := readBody(response.Body)
		fmt.Printf("%s\n", string(contents))
		err = xml.Unmarshal(contents, &s)
		fmt.Printf("SessionInfo: %s\n", s)
	}
	return s.Challenge
}

func readBody(body io.ReadCloser) []byte {
	defer body.Close()
	contents, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return contents
}
