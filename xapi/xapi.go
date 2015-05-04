package main

import (
	"fmt"
)

// A session
type Session struct {
	Uuid               string
	This_host          string // host_ref
	This_user          string // user_ref
	Last_active        string // datetime
	Pool               bool
	Other_config       map[string]string
	Is_local_superuser bool
	Subject            string // subject_ref
	Validation_time    string // datetime
	Auth_user_sid      string
	Auth_user_name     string
	Rbac_permissions   string // string_set
	Tasks              string // task_ref_set
	Parent             string // session_ref
	Originator         string
}

func main() {
	fmt.Println("it compiles")
}
