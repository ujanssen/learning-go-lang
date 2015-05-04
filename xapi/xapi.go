package main

import (
	"fmt"
)

// A session
type Session struct {
	Uuid             string
	ThisHost         string // host_ref
	ThisUser         string // user_ref
	LastActive       string // datetime
	Pool             bool
	OtherConfig      map[string]string
	IsLocalSuperuser bool
	Subject          string // subject_ref
	ValidationTime   string // datetime
	AuthUserSid      string
	AuthUserName     string
	RbacPermissions  string // string_set
	Tasks            string // task_ref_set
	Parent           string // session_ref
	Originator       string
}

func (Session *Session) login_with_password() {
}

func main() {
	fmt.Println("it compiles")
}
