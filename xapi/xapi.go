package main

import (
	"fmt"
	"github.com/nilshell/xmlrpc"
	"os"
)

type XenAPIClient struct {
	Session  interface{}
	Host     string
	Url      string
	Username string
	Password string
	RPC      *xmlrpc.Client
}

func NewXenAPIClient(host string) (client XenAPIClient) {
	client.Host = host
	client.Url = "http://" + host
	client.RPC, _ = xmlrpc.NewClient(client.Url, nil)
	return
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
	if err != nil {
		fmt.Errorf("RPC Call Error: %s", err)
	}
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
type session struct {
	uuid             string
	thisHost         interface{}
	thisUser         interface{}
	lastActive       interface{}
	pool             bool
	otherConfig      map[string]string
	isLocalSuperuser bool
	subject          interface{}
	validationTime   interface{}
	authUserSid      string
	authUserName     string
	rbacPermissions  interface{}
	tasks            interface{}
	parent           interface{}
	originator       string
}

func (client *XenAPIClient) session_login_with_password(uname string, pwd string, version string, originator string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = uname
	params[1] = pwd
	params[2] = version
	params[3] = originator

	err = client.RPCCall(&result, "session.login_with_password", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VM.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

const (
	XS_HOST   = "XS_HOST"
	XS_USER   = "XS_USER"
	XS_PASSWD = "XS_PASSWD"
)

func main() {
	fmt.Println("it compiles")
	host, user, passwd := os.Getenv(XS_HOST), os.Getenv(XS_USER), os.Getenv(XS_PASSWD)
	client := NewXenAPIClient(host)

	fmt.Println("NewXenAPIClient for: " + client.Host)
	fmt.Println("login_with_password: " + user + "/" + passwd)

	session, err := client.session_login_with_password(user, passwd, "", "")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("session: %+v\n", session)

	vms, err := client.VM_get_all(session)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("vms: %+v\n", vms)

	//    client.Username=user
	//    client.Password=passwd
	// err = client.Login()
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("session: %+v\n", client.Session)
}
