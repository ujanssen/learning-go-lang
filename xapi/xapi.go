package main

import (
	"errors"
	"fmt"
	"github.com/nilshell/xmlrpc"
)

type XenAPIClient struct {
	Session  interface{}
	Host     string
	Url      string
	Username string
	Password string
	RPC      *xmlrpc.Client
}

type APIResult struct {
	Status           string
	Value            interface{}
	ErrorDescription string
}

type XenAPIObject struct {
	Ref    string
	Client *XenAPIClient
}

func (c *XenAPIClient) RPCCall(result interface{}, method string, params []interface{}) (err error) {
	p := new(xmlrpc.Params)
	p.Params = params
	err = c.RPC.Call(method, *p, result)
	return err
}

func (client *XenAPIClient) APICall(result *APIResult, method string, params ...interface{}) (err error) {
	if client.Session == nil {
		return fmt.Errorf("No session. Unable to make call")
	}

	//Make a params slice which will include the session
	p := make([]interface{}, len(params)+1)
	p[0] = client.Session

	if params != nil {
		for idx, element := range params {
			p[idx+1] = element
		}
	}

	res := xmlrpc.Struct{}

	err = client.RPCCall(&res, method, p)

	if err != nil {
		return err
	}

	result.Status = res["Status"].(string)

	if result.Status != "Success" {
		return fmt.Errorf("API Error: %s", res["ErrorDescription"])
	} else {
		result.Value = res["Value"]
	}
	return
}

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

func (Session *Session) login_with_password(uname string, pwd string, version string, originator string) {
}

func (client *XenAPIClient) Login() (err error) {
	//Do loging call
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)
	params[0] = client.Username
	params[1] = client.Password

	err = client.RPCCall(&result, "session.login_with_password", params)
	if err == nil {
		// err might not be set properly, so check the reference
		if result["Value"] == nil {
			return errors.New("Invalid credentials supplied")
		}
	}
	client.Session = result["Value"]
	return err
}

func main() {
	fmt.Println("it compiles")
}
