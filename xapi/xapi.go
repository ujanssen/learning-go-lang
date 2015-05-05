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

// A session

func (client *XenAPIClient) session_logout_subject_identifier(session_id interface{}, subject_identifier string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = subject_identifier

	err = client.RPCCall(&result, "session.logout_subject_identifier", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_all_subject_identifiers(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "session.get_all_subject_identifiers", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_local_logout(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "session.local_logout", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_slave_local_login_with_password(uname string, pwd string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = uname
	params[1] = pwd

	err = client.RPCCall(&result, "session.slave_local_login_with_password", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_change_password(session_id interface{}, old_pwd string, new_pwd string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = old_pwd
	params[2] = new_pwd

	err = client.RPCCall(&result, "session.change_password", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_logout(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "session.logout", params)
	resultValue = result["Value"]

	return resultValue, err
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

func (client *XenAPIClient) session_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "session.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "session.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "session.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_originator(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_originator", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_parent(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_parent", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_tasks(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_tasks", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_rbac_permissions(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_rbac_permissions", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_auth_user_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_auth_user_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_auth_user_sid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_auth_user_sid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_validation_time(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_validation_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_subject(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_subject", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_is_local_superuser(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_is_local_superuser", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_pool(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_pool", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_last_active(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_last_active", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_this_user(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_this_user", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_this_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_this_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "session.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) session_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "session.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// Management of remote authentication services

func (client *XenAPIClient) auth_get_group_membership(session_id interface{}, subject_identifier string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = subject_identifier

	err = client.RPCCall(&result, "auth.get_group_membership", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) auth_get_subject_information_from_identifier(session_id interface{}, subject_identifier string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = subject_identifier

	err = client.RPCCall(&result, "auth.get_subject_information_from_identifier", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) auth_get_subject_identifier(session_id interface{}, subject_name string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = subject_name

	err = client.RPCCall(&result, "auth.get_subject_identifier", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A user or group that can log in xapi

func (client *XenAPIClient) subject_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "subject.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "subject.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_permissions_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "subject.get_permissions_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_remove_from_roles(session_id interface{}, self interface{}, role interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = role

	err = client.RPCCall(&result, "subject.remove_from_roles", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_add_to_roles(session_id interface{}, self interface{}, role interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = role

	err = client.RPCCall(&result, "subject.add_to_roles", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_roles(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "subject.get_roles", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "subject.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_subject_identifier(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "subject.get_subject_identifier", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "subject.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "subject.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "subject.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "subject.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) subject_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "subject.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A set of permissions associated with a subject

func (client *XenAPIClient) role_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "role.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "role.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_by_permission_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "role.get_by_permission_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_by_permission(session_id interface{}, permission interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = permission

	err = client.RPCCall(&result, "role.get_by_permission", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_permissions_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "role.get_permissions_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_permissions(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "role.get_permissions", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_subroles(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "role.get_subroles", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "role.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "role.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "role.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "role.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "role.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) role_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "role.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A long-running asynchronous task

func (client *XenAPIClient) task_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "task.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "task.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_cancel(session_id interface{}, task interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = task

	err = client.RPCCall(&result, "task.cancel", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_create(session_id interface{}, label string, description string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = label
	params[2] = description

	err = client.RPCCall(&result, "task.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "task.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "task.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "task.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_backtrace(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_backtrace", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_subtasks(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_subtasks", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_subtask_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_subtask_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_error_info(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_error_info", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_result(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_result", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_progress(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_progress", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_resident_on(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_resident_on", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_status(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_status", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_finished(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_finished", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_created(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_created", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "task.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "task.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) task_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "task.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// Asynchronous event registration and handling

func (client *XenAPIClient) event_inject(session_id interface{}, class string, ref string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = class
	params[2] = ref

	err = client.RPCCall(&result, "event.inject", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) event_get_current_id(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "event.get_current_id", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) event_from(session_id interface{}, classes interface{}, token string, timeout interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = classes
	params[2] = token
	params[3] = timeout

	err = client.RPCCall(&result, "event.from", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) event_next(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "event.next", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) event_unregister(session_id interface{}, classes interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = classes

	err = client.RPCCall(&result, "event.unregister", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) event_register(session_id interface{}, classes interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = classes

	err = client.RPCCall(&result, "event.register", params)
	resultValue = result["Value"]

	return resultValue, err
}

// Pool-wide information

func (client *XenAPIClient) pool_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_apply_edition(session_id interface{}, self interface{}, edition string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = edition

	err = client.RPCCall(&result, "pool.apply_edition", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_license_state(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_license_state", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_disable_local_storage_caching(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.disable_local_storage_caching", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_enable_local_storage_caching(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.enable_local_storage_caching", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_test_archive_target(session_id interface{}, self interface{}, config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = config

	err = client.RPCCall(&result, "pool.test_archive_target", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_vswitch_controller(session_id interface{}, address string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = address

	err = client.RPCCall(&result, "pool.set_vswitch_controller", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_disable_redo_log(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.disable_redo_log", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_enable_redo_log(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "pool.enable_redo_log", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_certificate_sync(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.certificate_sync", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_crl_list(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.crl_list", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_crl_uninstall(session_id interface{}, name string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = name

	err = client.RPCCall(&result, "pool.crl_uninstall", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_crl_install(session_id interface{}, name string, cert string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = name
	params[2] = cert

	err = client.RPCCall(&result, "pool.crl_install", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_certificate_list(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.certificate_list", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_certificate_uninstall(session_id interface{}, name string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = name

	err = client.RPCCall(&result, "pool.certificate_uninstall", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_certificate_install(session_id interface{}, name string, cert string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = name
	params[2] = cert

	err = client.RPCCall(&result, "pool.certificate_install", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_send_test_post(session_id interface{}, host string, port interface{}, body string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = host
	params[2] = port
	params[3] = body

	err = client.RPCCall(&result, "pool.send_test_post", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_retrieve_wlb_recommendations(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.retrieve_wlb_recommendations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_retrieve_wlb_configuration(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.retrieve_wlb_configuration", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_send_wlb_configuration(session_id interface{}, config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = config

	err = client.RPCCall(&result, "pool.send_wlb_configuration", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_deconfigure_wlb(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.deconfigure_wlb", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_initialize_wlb(session_id interface{}, wlb_url string, wlb_username string, wlb_password string, xenserver_username string, xenserver_password string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 6)

	params[0] = session_id
	params[1] = wlb_url
	params[2] = wlb_username
	params[3] = wlb_password
	params[4] = xenserver_username
	params[5] = xenserver_password

	err = client.RPCCall(&result, "pool.initialize_wlb", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_detect_nonhomogeneous_external_auth(session_id interface{}, pool interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = pool

	err = client.RPCCall(&result, "pool.detect_nonhomogeneous_external_auth", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_disable_external_auth(session_id interface{}, pool interface{}, config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = pool
	params[2] = config

	err = client.RPCCall(&result, "pool.disable_external_auth", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_enable_external_auth(session_id interface{}, pool interface{}, config map[string]string, service_name string, auth_type string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = pool
	params[2] = config
	params[3] = service_name
	params[4] = auth_type

	err = client.RPCCall(&result, "pool.enable_external_auth", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_create_new_blob(session_id interface{}, pool interface{}, name string, mime_type string, public bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = pool
	params[2] = name
	params[3] = mime_type
	params[4] = public

	err = client.RPCCall(&result, "pool.create_new_blob", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_ha_host_failures_to_tolerate(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_ha_host_failures_to_tolerate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_ha_compute_vm_failover_plan(session_id interface{}, failed_hosts interface{}, failed_vms interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = failed_hosts
	params[2] = failed_vms

	err = client.RPCCall(&result, "pool.ha_compute_vm_failover_plan", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_ha_compute_hypothetical_max_host_failures_to_tolerate(session_id interface{}, configuration interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = configuration

	err = client.RPCCall(&result, "pool.ha_compute_hypothetical_max_host_failures_to_tolerate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_ha_compute_max_host_failures_to_tolerate(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.ha_compute_max_host_failures_to_tolerate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_ha_failover_plan_exists(session_id interface{}, n interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = n

	err = client.RPCCall(&result, "pool.ha_failover_plan_exists", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_ha_prevent_restarts_for(session_id interface{}, seconds interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = seconds

	err = client.RPCCall(&result, "pool.ha_prevent_restarts_for", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_designate_new_master(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "pool.designate_new_master", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_sync_database(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.sync_database", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_disable_ha(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.disable_ha", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_enable_ha(session_id interface{}, heartbeat_srs interface{}, configuration map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = heartbeat_srs
	params[2] = configuration

	err = client.RPCCall(&result, "pool.enable_ha", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_create_VLAN_from_PIF(session_id interface{}, pif interface{}, network interface{}, VLAN interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = pif
	params[2] = network
	params[3] = VLAN

	err = client.RPCCall(&result, "pool.create_VLAN_from_PIF", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_create_VLAN(session_id interface{}, device string, network interface{}, VLAN interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = device
	params[2] = network
	params[3] = VLAN

	err = client.RPCCall(&result, "pool.create_VLAN", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_recover_slaves(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.recover_slaves", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_emergency_reset_master(session_id interface{}, master_address string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = master_address

	err = client.RPCCall(&result, "pool.emergency_reset_master", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_emergency_transition_to_master(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "pool.emergency_transition_to_master", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_eject(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "pool.eject", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_join_force(session_id interface{}, master_address string, master_username string, master_password string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = master_address
	params[2] = master_username
	params[3] = master_password

	err = client.RPCCall(&result, "pool.join_force", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_join(session_id interface{}, master_address string, master_username string, master_password string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = master_address
	params[2] = master_username
	params[3] = master_password

	err = client.RPCCall(&result, "pool.join", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_wlb_verify_cert(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_wlb_verify_cert", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_wlb_enabled(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_wlb_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_remove_from_gui_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "pool.remove_from_gui_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_add_to_gui_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "pool.add_to_gui_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_gui_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_gui_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_remove_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.remove_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_add_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.add_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_tags(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_ha_allow_overcommit(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_ha_allow_overcommit", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "pool.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "pool.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_crash_dump_SR(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_crash_dump_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_suspend_image_SR(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_suspend_image_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_default_SR(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_default_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "pool.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_metadata_VDIs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_metadata_VDIs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_restrictions(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_restrictions", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_vswitch_controller(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_vswitch_controller", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_redo_log_vdi(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_redo_log_vdi", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_redo_log_enabled(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_redo_log_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_wlb_verify_cert(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_wlb_verify_cert", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_wlb_enabled(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_wlb_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_wlb_username(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_wlb_username", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_wlb_url(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_wlb_url", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_gui_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_gui_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_tags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_blobs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_blobs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_ha_overcommitted(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_ha_overcommitted", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_ha_allow_overcommit(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_ha_allow_overcommit", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_ha_plan_exists_for(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_ha_plan_exists_for", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_ha_host_failures_to_tolerate(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_ha_host_failures_to_tolerate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_ha_statefiles(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_ha_statefiles", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_ha_configuration(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_ha_configuration", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_ha_enabled(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_ha_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_crash_dump_SR(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_crash_dump_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_suspend_image_SR(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_suspend_image_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_default_SR(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_default_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_master(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_master", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "pool.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) pool_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "pool.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// Pool-wide patches

func (client *XenAPIClient) poolPatch_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "poolPatch.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "poolPatch.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_clean_on_host(session_id interface{}, self interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = host

	err = client.RPCCall(&result, "poolPatch.clean_on_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_pool_clean(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.pool_clean", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_clean(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.clean", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_precheck(session_id interface{}, self interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = host

	err = client.RPCCall(&result, "poolPatch.precheck", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_pool_apply(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.pool_apply", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_apply(session_id interface{}, self interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = host

	err = client.RPCCall(&result, "poolPatch.apply", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "poolPatch.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "poolPatch.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "poolPatch.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_after_apply_guidance(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_after_apply_guidance", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_host_patches(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_host_patches", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_pool_applied(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_pool_applied", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_size(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "poolPatch.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "poolPatch.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) poolPatch_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "poolPatch.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A virtual machine (or 'guest').

func (client *XenAPIClient) VM_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VM.get_all_records", params)
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

func (client *XenAPIClient) VM_call_plugin(session_id interface{}, vm interface{}, plugin string, fn string, args map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = vm
	params[2] = plugin
	params[3] = fn
	params[4] = args

	err = client.RPCCall(&result, "VM.call_plugin", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_query_services(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.query_services", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_appliance(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_appliance", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_import_convert(session_id interface{}, a_type string, username string, password string, sr interface{}, remote_config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 6)

	params[0] = session_id
	params[1] = a_type
	params[2] = username
	params[3] = password
	params[4] = sr
	params[5] = remote_config

	err = client.RPCCall(&result, "VM.import_convert", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_recover(session_id interface{}, self interface{}, session_to interface{}, force bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = session_to
	params[3] = force

	err = client.RPCCall(&result, "VM.recover", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_SRs_required_for_recovery(session_id interface{}, self interface{}, session_to interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = session_to

	err = client.RPCCall(&result, "VM.get_SRs_required_for_recovery", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_assert_can_be_recovered(session_id interface{}, self interface{}, session_to interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = session_to

	err = client.RPCCall(&result, "VM.assert_can_be_recovered", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_suspend_VDI(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_suspend_VDI", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_order(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_order", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_shutdown_delay(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_shutdown_delay", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_start_delay(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_start_delay", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_protection_policy(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_protection_policy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_copy_bios_strings(session_id interface{}, vm interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vm
	params[2] = host

	err = client.RPCCall(&result, "VM.copy_bios_strings", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_retrieve_wlb_recommendations(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.retrieve_wlb_recommendations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_assert_agile(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.assert_agile", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_create_new_blob(session_id interface{}, vm interface{}, name string, mime_type string, public bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = vm
	params[2] = name
	params[3] = mime_type
	params[4] = public

	err = client.RPCCall(&result, "VM.create_new_blob", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_assert_can_boot_here(session_id interface{}, self interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = host

	err = client.RPCCall(&result, "VM.assert_can_boot_here", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_possible_hosts(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.get_possible_hosts", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_allowed_VIF_devices(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.get_allowed_VIF_devices", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_allowed_VBD_devices(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.get_allowed_VBD_devices", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_update_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.update_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_assert_operation_valid(session_id interface{}, self interface{}, op interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = op

	err = client.RPCCall(&result, "VM.assert_operation_valid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_forget_data_source_archives(session_id interface{}, self interface{}, data_source string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = data_source

	err = client.RPCCall(&result, "VM.forget_data_source_archives", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_query_data_source(session_id interface{}, self interface{}, data_source string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = data_source

	err = client.RPCCall(&result, "VM.query_data_source", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_record_data_source(session_id interface{}, self interface{}, data_source string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = data_source

	err = client.RPCCall(&result, "VM.record_data_source", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_data_sources(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_data_sources", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_boot_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_boot_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_assert_can_migrate(session_id interface{}, vm interface{}, dest map[string]string, live bool, vdi_map interface{}, vif_map interface{}, options map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 7)

	params[0] = session_id
	params[1] = vm
	params[2] = dest
	params[3] = live
	params[4] = vdi_map
	params[5] = vif_map
	params[6] = options

	err = client.RPCCall(&result, "VM.assert_can_migrate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_migrate_send(session_id interface{}, vm interface{}, dest map[string]string, live bool, vdi_map interface{}, vif_map interface{}, options map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 7)

	params[0] = session_id
	params[1] = vm
	params[2] = dest
	params[3] = live
	params[4] = vdi_map
	params[5] = vif_map
	params[6] = options

	err = client.RPCCall(&result, "VM.migrate_send", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_maximise_memory(session_id interface{}, self interface{}, total interface{}, approximate bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = total
	params[3] = approximate

	err = client.RPCCall(&result, "VM.maximise_memory", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_send_trigger(session_id interface{}, vm interface{}, trigger string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vm
	params[2] = trigger

	err = client.RPCCall(&result, "VM.send_trigger", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_send_sysrq(session_id interface{}, vm interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vm
	params[2] = key

	err = client.RPCCall(&result, "VM.send_sysrq", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_VCPUs_at_startup(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_VCPUs_at_startup", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_VCPUs_max(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_VCPUs_max", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_shadow_multiplier_live(session_id interface{}, self interface{}, multiplier interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = multiplier

	err = client.RPCCall(&result, "VM.set_shadow_multiplier_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_HVM_shadow_multiplier(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_HVM_shadow_multiplier", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_cooperative(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_cooperative", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_wait_memory_target_live(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.wait_memory_target_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_target_live(session_id interface{}, self interface{}, target interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = target

	err = client.RPCCall(&result, "VM.set_memory_target_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_limits(session_id interface{}, self interface{}, static_min interface{}, static_max interface{}, dynamic_min interface{}, dynamic_max interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 6)

	params[0] = session_id
	params[1] = self
	params[2] = static_min
	params[3] = static_max
	params[4] = dynamic_min
	params[5] = dynamic_max

	err = client.RPCCall(&result, "VM.set_memory_limits", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_static_range(session_id interface{}, self interface{}, min interface{}, max interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = min
	params[3] = max

	err = client.RPCCall(&result, "VM.set_memory_static_range", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_static_min(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_memory_static_min", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_static_max(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_memory_static_max", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_dynamic_range(session_id interface{}, self interface{}, min interface{}, max interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = min
	params[3] = max

	err = client.RPCCall(&result, "VM.set_memory_dynamic_range", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_dynamic_min(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_memory_dynamic_min", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_memory_dynamic_max(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_memory_dynamic_max", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_compute_memory_overhead(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.compute_memory_overhead", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_ha_always_run(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_ha_always_run", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_ha_restart_priority(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_ha_restart_priority", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_to_VCPUs_params_live(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM.add_to_VCPUs_params_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_VCPUs_number_live(session_id interface{}, self interface{}, nvcpu interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = nvcpu

	err = client.RPCCall(&result, "VM.set_VCPUs_number_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_pool_migrate(session_id interface{}, vm interface{}, host interface{}, options map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = vm
	params[2] = host
	params[3] = options

	err = client.RPCCall(&result, "VM.pool_migrate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_resume_on(session_id interface{}, vm interface{}, host interface{}, start_paused bool, force bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = vm
	params[2] = host
	params[3] = start_paused
	params[4] = force

	err = client.RPCCall(&result, "VM.resume_on", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_resume(session_id interface{}, vm interface{}, start_paused bool, force bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = vm
	params[2] = start_paused
	params[3] = force

	err = client.RPCCall(&result, "VM.resume", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_suspend(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.suspend", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_hard_reboot(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.hard_reboot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_power_state_reset(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.power_state_reset", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_hard_shutdown(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.hard_shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_clean_reboot(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.clean_reboot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_shutdown(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_clean_shutdown(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.clean_shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_unpause(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.unpause", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_pause(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.pause", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_start_on(session_id interface{}, vm interface{}, host interface{}, start_paused bool, force bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = vm
	params[2] = host
	params[3] = start_paused
	params[4] = force

	err = client.RPCCall(&result, "VM.start_on", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_start(session_id interface{}, vm interface{}, start_paused bool, force bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = vm
	params[2] = start_paused
	params[3] = force

	err = client.RPCCall(&result, "VM.start", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_provision(session_id interface{}, vm interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vm

	err = client.RPCCall(&result, "VM.provision", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_checkpoint(session_id interface{}, vm interface{}, new_name string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vm
	params[2] = new_name

	err = client.RPCCall(&result, "VM.checkpoint", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_revert(session_id interface{}, snapshot interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = snapshot

	err = client.RPCCall(&result, "VM.revert", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_copy(session_id interface{}, vm interface{}, new_name string, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = vm
	params[2] = new_name
	params[3] = sr

	err = client.RPCCall(&result, "VM.copy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_clone(session_id interface{}, vm interface{}, new_name string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vm
	params[2] = new_name

	err = client.RPCCall(&result, "VM.clone", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_snapshot_with_quiesce(session_id interface{}, vm interface{}, new_name string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vm
	params[2] = new_name

	err = client.RPCCall(&result, "VM.snapshot_with_quiesce", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_snapshot(session_id interface{}, vm interface{}, new_name string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vm
	params[2] = new_name

	err = client.RPCCall(&result, "VM.snapshot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_suspend_SR(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_suspend_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_remove_from_blocked_operations(session_id interface{}, self interface{}, key interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM.remove_from_blocked_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_to_blocked_operations(session_id interface{}, self interface{}, key interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM.add_to_blocked_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_blocked_operations(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_blocked_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_remove_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.remove_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.add_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_tags(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_remove_from_xenstore_data(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM.remove_from_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_to_xenstore_data(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM.add_to_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_xenstore_data(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_recommendations(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_recommendations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_PCI_bus(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_PCI_bus", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_remove_from_platform(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM.remove_from_platform", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_to_platform(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM.add_to_platform", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_platform(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_platform", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_remove_from_HVM_boot_params(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM.remove_from_HVM_boot_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_to_HVM_boot_params(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM.add_to_HVM_boot_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_HVM_boot_params(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_HVM_boot_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_HVM_boot_policy(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_HVM_boot_policy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_PV_legacy_args(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_PV_legacy_args", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_PV_bootloader_args(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_PV_bootloader_args", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_PV_args(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_PV_args", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_PV_ramdisk(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_PV_ramdisk", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_PV_kernel(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_PV_kernel", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_PV_bootloader(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_PV_bootloader", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_actions_after_crash(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_actions_after_crash", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_actions_after_reboot(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_actions_after_reboot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_actions_after_shutdown(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_actions_after_shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_remove_from_VCPUs_params(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM.remove_from_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_add_to_VCPUs_params(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM.add_to_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_VCPUs_params(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_affinity(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_affinity", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_is_a_template(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_is_a_template", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_user_version(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_user_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_generation_id(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_generation_id", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_suspend_SR(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_suspend_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_attached_PCIs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_attached_PCIs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_VGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_VGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_order(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_order", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_shutdown_delay(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_shutdown_delay", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_start_delay(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_start_delay", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_appliance(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_appliance", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_is_snapshot_from_vmpp(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_is_snapshot_from_vmpp", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_protection_policy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_protection_policy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_bios_strings(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_bios_strings", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_children(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_children", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_parent(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_parent", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_snapshot_metadata(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_snapshot_metadata", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_snapshot_info(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_snapshot_info", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_blocked_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_blocked_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_tags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_blobs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_blobs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_transportable_snapshot_id(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_transportable_snapshot_id", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_snapshot_time(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_snapshot_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_snapshots(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_snapshots", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_snapshot_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_snapshot_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_is_a_snapshot(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_is_a_snapshot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_ha_restart_priority(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_ha_restart_priority", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_ha_always_run(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_ha_always_run", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_xenstore_data(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_recommendations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_recommendations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_last_booted_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_last_booted_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_guest_metrics(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_guest_metrics", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_metrics(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_metrics", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_is_control_domain(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_is_control_domain", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_last_boot_CPU_flags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_last_boot_CPU_flags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_domarch(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_domarch", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_domid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_domid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_PCI_bus(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_PCI_bus", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_platform(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_platform", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_HVM_shadow_multiplier(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_HVM_shadow_multiplier", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_HVM_boot_params(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_HVM_boot_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_HVM_boot_policy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_HVM_boot_policy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_PV_legacy_args(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_PV_legacy_args", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_PV_bootloader_args(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_PV_bootloader_args", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_PV_args(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_PV_args", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_PV_ramdisk(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_PV_ramdisk", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_PV_kernel(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_PV_kernel", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_PV_bootloader(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_PV_bootloader", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_VTPMs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_VTPMs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_crash_dumps(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_crash_dumps", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_VBDs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_VBDs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_VIFs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_VIFs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_consoles(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_consoles", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_actions_after_crash(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_actions_after_crash", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_actions_after_reboot(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_actions_after_reboot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_actions_after_shutdown(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_actions_after_shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_VCPUs_at_startup(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_VCPUs_at_startup", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_VCPUs_max(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_VCPUs_max", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_VCPUs_params(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_memory_static_min(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_memory_static_min", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_memory_dynamic_min(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_memory_dynamic_min", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_memory_dynamic_max(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_memory_dynamic_max", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_memory_static_max(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_memory_static_max", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_memory_target(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_memory_target", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_memory_overhead(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_memory_overhead", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_affinity(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_affinity", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_resident_on(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_resident_on", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_suspend_VDI(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_suspend_VDI", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_is_a_template(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_is_a_template", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_user_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_user_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_power_state(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_power_state", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "VM.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "VM.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VM.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// The metrics associated with a VM

func (client *XenAPIClient) VM_metrics_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VM_metrics.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VM_metrics.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM_metrics.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM_metrics.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM_metrics.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_last_updated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_last_updated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_install_time(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_install_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_start_time(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_start_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_state(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_state", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_flags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_VCPUs_flags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_params(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_CPU(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_VCPUs_CPU", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_utilisation(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_VCPUs_utilisation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_number(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_VCPUs_number", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_memory_actual(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_memory_actual", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VM_metrics.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_metrics_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_metrics.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// The metrics reported by the guest (as opposed to inferred from outside)

func (client *XenAPIClient) VM_guest_metrics_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VM_guest_metrics.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VM_guest_metrics.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VM_guest_metrics.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VM_guest_metrics.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VM_guest_metrics.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_live(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_last_updated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_last_updated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_other(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_other", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_networks(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_networks", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_disks(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_disks", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_memory(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_memory", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_PV_drivers_up_to_date(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_PV_drivers_up_to_date", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_PV_drivers_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_PV_drivers_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_os_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_os_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VM_guest_metrics.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VM_guest_metrics_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VM_guest_metrics.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// VM Protection Policy

func (client *XenAPIClient) VMPP_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VMPP.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VMPP.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_archive_last_run_time(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_archive_last_run_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_backup_last_run_time(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_backup_last_run_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_remove_from_alarm_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VMPP.remove_from_alarm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_remove_from_archive_schedule(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VMPP.remove_from_archive_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_remove_from_archive_target_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VMPP.remove_from_archive_target_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_remove_from_backup_schedule(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VMPP.remove_from_backup_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_add_to_alarm_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VMPP.add_to_alarm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_add_to_archive_schedule(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VMPP.add_to_archive_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_add_to_archive_target_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VMPP.add_to_archive_target_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_add_to_backup_schedule(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VMPP.add_to_backup_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_alarm_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_alarm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_is_alarm_enabled(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_is_alarm_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_archive_target_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_archive_target_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_archive_target_type(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_archive_target_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_archive_schedule(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_archive_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_archive_frequency(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_archive_frequency", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_backup_schedule(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_backup_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_backup_frequency(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_backup_frequency", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_backup_retention_value(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_backup_retention_value", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_alerts(session_id interface{}, vmpp interface{}, hours_from_now interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vmpp
	params[2] = hours_from_now

	err = client.RPCCall(&result, "VMPP.get_alerts", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_archive_now(session_id interface{}, snapshot interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = snapshot

	err = client.RPCCall(&result, "VMPP.archive_now", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_protect_now(session_id interface{}, vmpp interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vmpp

	err = client.RPCCall(&result, "VMPP.protect_now", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_backup_type(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_backup_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_is_policy_enabled(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_is_policy_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMPP.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_recent_alerts(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_recent_alerts", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_alarm_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_alarm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_is_alarm_enabled(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_is_alarm_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_VMs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_VMs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_archive_last_run_time(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_archive_last_run_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_is_archive_running(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_is_archive_running", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_archive_schedule(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_archive_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_archive_frequency(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_archive_frequency", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_archive_target_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_archive_target_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_archive_target_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_archive_target_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_backup_last_run_time(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_backup_last_run_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_is_backup_running(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_is_backup_running", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_backup_schedule(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_backup_schedule", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_backup_frequency(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_backup_frequency", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_backup_retention_value(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_backup_retention_value", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_backup_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_backup_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_is_policy_enabled(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_is_policy_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "VMPP.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "VMPP.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VMPP.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMPP_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMPP.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// VM appliance

func (client *XenAPIClient) VMAppliance_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VMAppliance.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VMAppliance.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_recover(session_id interface{}, self interface{}, session_to interface{}, force bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = session_to
	params[3] = force

	err = client.RPCCall(&result, "VMAppliance.recover", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_SRs_required_for_recovery(session_id interface{}, self interface{}, session_to interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = session_to

	err = client.RPCCall(&result, "VMAppliance.get_SRs_required_for_recovery", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_assert_can_be_recovered(session_id interface{}, self interface{}, session_to interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = session_to

	err = client.RPCCall(&result, "VMAppliance.assert_can_be_recovered", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_shutdown(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_hard_shutdown(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.hard_shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_clean_shutdown(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.clean_shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_start(session_id interface{}, self interface{}, paused bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = paused

	err = client.RPCCall(&result, "VMAppliance.start", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMAppliance.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VMAppliance.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_VMs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.get_VMs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "VMAppliance.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "VMAppliance.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VMAppliance.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VMAppliance_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VMAppliance.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// DR task

func (client *XenAPIClient) DRTask_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "DRTask.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) DRTask_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "DRTask.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) DRTask_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "DRTask.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) DRTask_create(session_id interface{}, a_type string, device_config map[string]string, whitelist interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = a_type
	params[2] = device_config
	params[3] = whitelist

	err = client.RPCCall(&result, "DRTask.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) DRTask_get_introduced_SRs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "DRTask.get_introduced_SRs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) DRTask_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "DRTask.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) DRTask_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "DRTask.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) DRTask_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "DRTask.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A physical host

func (client *XenAPIClient) host_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_disable_display(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.disable_display", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_enable_display(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.enable_display", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_declare_dead(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.declare_dead", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_migrate_receive(session_id interface{}, host interface{}, network interface{}, options map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = host
	params[2] = network
	params[3] = options

	err = client.RPCCall(&result, "host.migrate_receive", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_disable_local_storage_caching(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.disable_local_storage_caching", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_enable_local_storage_caching(session_id interface{}, host interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = sr

	err = client.RPCCall(&result, "host.enable_local_storage_caching", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_reset_cpu_features(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.reset_cpu_features", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_cpu_features(session_id interface{}, host interface{}, features string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = features

	err = client.RPCCall(&result, "host.set_cpu_features", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_power_on_mode(session_id interface{}, self interface{}, power_on_mode string, power_on_config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = power_on_mode
	params[3] = power_on_config

	err = client.RPCCall(&result, "host.set_power_on_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_refresh_pack_info(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.refresh_pack_info", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_apply_edition(session_id interface{}, host interface{}, edition string, force bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = host
	params[2] = edition
	params[3] = force

	err = client.RPCCall(&result, "host.apply_edition", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_server_certificate(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.get_server_certificate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_retrieve_wlb_evacuate_recommendations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.retrieve_wlb_evacuate_recommendations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_disable_external_auth(session_id interface{}, host interface{}, config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = config

	err = client.RPCCall(&result, "host.disable_external_auth", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_enable_external_auth(session_id interface{}, host interface{}, config map[string]string, service_name string, auth_type string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = host
	params[2] = config
	params[3] = service_name
	params[4] = auth_type

	err = client.RPCCall(&result, "host.enable_external_auth", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_server_localtime(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.get_server_localtime", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_servertime(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.get_servertime", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_call_plugin(session_id interface{}, host interface{}, plugin string, fn string, args map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = host
	params[2] = plugin
	params[3] = fn
	params[4] = args

	err = client.RPCCall(&result, "host.call_plugin", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_create_new_blob(session_id interface{}, host interface{}, name string, mime_type string, public bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = host
	params[2] = name
	params[3] = mime_type
	params[4] = public

	err = client.RPCCall(&result, "host.create_new_blob", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_backup_rrds(session_id interface{}, host interface{}, delay interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = delay

	err = client.RPCCall(&result, "host.backup_rrds", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_sync_data(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.sync_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_compute_memory_overhead(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.compute_memory_overhead", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_compute_free_memory(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.compute_free_memory", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_hostname_live(session_id interface{}, host interface{}, hostname string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = hostname

	err = client.RPCCall(&result, "host.set_hostname_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_shutdown_agent(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host.shutdown_agent", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_restart_agent(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.restart_agent", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_system_status_capabilities(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.get_system_status_capabilities", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_management_interface(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.get_management_interface", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_management_disable(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host.management_disable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_local_management_reconfigure(session_id interface{}, a_interface string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = a_interface

	err = client.RPCCall(&result, "host.local_management_reconfigure", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_management_reconfigure(session_id interface{}, pif interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = pif

	err = client.RPCCall(&result, "host.management_reconfigure", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_syslog_reconfigure(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.syslog_reconfigure", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_evacuate(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.evacuate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_uncooperative_resident_VMs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_uncooperative_resident_VMs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_vms_which_prevent_evacuation(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_vms_which_prevent_evacuation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_assert_can_evacuate(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.assert_can_evacuate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_forget_data_source_archives(session_id interface{}, host interface{}, data_source string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = data_source

	err = client.RPCCall(&result, "host.forget_data_source_archives", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_query_data_source(session_id interface{}, host interface{}, data_source string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = data_source

	err = client.RPCCall(&result, "host.query_data_source", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_record_data_source(session_id interface{}, host interface{}, data_source string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = data_source

	err = client.RPCCall(&result, "host.record_data_source", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_data_sources(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.get_data_sources", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_emergency_ha_disable(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host.emergency_ha_disable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_power_on(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.power_on", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_license_apply(session_id interface{}, host interface{}, contents string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = contents

	err = client.RPCCall(&result, "host.license_apply", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_list_methods(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host.list_methods", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_bugreport_upload(session_id interface{}, host interface{}, url string, options map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = host
	params[2] = url
	params[3] = options

	err = client.RPCCall(&result, "host.bugreport_upload", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_send_debug_keys(session_id interface{}, host interface{}, keys string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = host
	params[2] = keys

	err = client.RPCCall(&result, "host.send_debug_keys", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_log(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.get_log", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_dmesg_clear(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.dmesg_clear", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_dmesg(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.dmesg", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_reboot(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.reboot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_shutdown(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.shutdown", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_enable(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.enable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_disable(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "host.disable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_display(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_display", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_remove_from_guest_VCPUs_params(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "host.remove_from_guest_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_add_to_guest_VCPUs_params(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "host.add_to_guest_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_guest_VCPUs_params(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_guest_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_remove_from_license_server(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "host.remove_from_license_server", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_add_to_license_server(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "host.add_to_license_server", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_license_server(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_license_server", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_remove_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.remove_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_add_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.add_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_tags(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_address(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_address", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_hostname(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_hostname", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_crash_dump_sr(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_crash_dump_sr", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_suspend_image_sr(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_suspend_image_sr", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_remove_from_logging(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "host.remove_from_logging", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_add_to_logging(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "host.add_to_logging", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_logging(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_logging", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "host.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "host.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_display(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_display", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_guest_VCPUs_params(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_guest_VCPUs_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_PGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_PGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_PCIs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_PCIs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_chipset_info(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_chipset_info", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_local_cache_sr(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_local_cache_sr", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_power_on_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_power_on_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_power_on_mode(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_power_on_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_bios_strings(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_bios_strings", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_license_server(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_license_server", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_edition(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_edition", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_external_auth_configuration(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_external_auth_configuration", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_external_auth_service_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_external_auth_service_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_external_auth_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_external_auth_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_tags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_blobs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_blobs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_ha_network_peers(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_ha_network_peers", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_ha_statefiles(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_ha_statefiles", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_license_params(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_license_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_metrics(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_metrics", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_address(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_address", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_hostname(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_hostname", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_cpu_info(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_cpu_info", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_host_CPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_host_CPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_PBDs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_PBDs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_patches(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_patches", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_crashdumps(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_crashdumps", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_crash_dump_sr(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_crash_dump_sr", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_suspend_image_sr(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_suspend_image_sr", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_PIFs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_PIFs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_logging(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_logging", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_resident_VMs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_resident_VMs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_supported_bootloaders(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_supported_bootloaders", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_sched_policy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_sched_policy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_cpu_configuration(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_cpu_configuration", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_capabilities(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_capabilities", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_software_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_software_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_enabled(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_API_version_vendor_implementation(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_API_version_vendor_implementation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_API_version_vendor(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_API_version_vendor", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_API_version_minor(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_API_version_minor", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_API_version_major(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_API_version_major", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_memory_overhead(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_memory_overhead", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "host.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "host.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// Represents a host crash dump

func (client *XenAPIClient) hostCrashdump_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "hostCrashdump.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "hostCrashdump.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_upload(session_id interface{}, self interface{}, url string, options map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = url
	params[3] = options

	err = client.RPCCall(&result, "hostCrashdump.upload", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCrashdump.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "hostCrashdump.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "hostCrashdump.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "hostCrashdump.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCrashdump.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_size(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCrashdump.get_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_timestamp(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCrashdump.get_timestamp", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCrashdump.get_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCrashdump.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "hostCrashdump.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCrashdump_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCrashdump.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// Represents a patch stored on a server

func (client *XenAPIClient) hostPatch_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "hostPatch.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "hostPatch.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_apply(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.apply", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "hostPatch.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "hostPatch.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "hostPatch.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_pool_patch(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_pool_patch", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_size(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_timestamp_applied(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_timestamp_applied", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_applied(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_applied", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "hostPatch.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "hostPatch.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostPatch_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostPatch.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// The metrics associated with a host

func (client *XenAPIClient) host_metrics_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host_metrics.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "host_metrics.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "host_metrics.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "host_metrics.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "host_metrics.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host_metrics.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_last_updated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host_metrics.get_last_updated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_live(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host_metrics.get_live", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_memory_free(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host_metrics.get_memory_free", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_memory_total(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host_metrics.get_memory_total", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host_metrics.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "host_metrics.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) host_metrics_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "host_metrics.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A physical CPU

func (client *XenAPIClient) hostCpu_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "hostCpu.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "hostCpu.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "hostCpu.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "hostCpu.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "hostCpu.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_utilisation(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_utilisation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_features(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_features", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_flags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_flags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_stepping(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_stepping", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_model(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_model", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_family(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_family", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_modelname(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_modelname", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_speed(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_speed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_vendor(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_vendor", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_number(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_number", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "hostCpu.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) hostCpu_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "hostCpu.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A virtual network

func (client *XenAPIClient) network_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "network.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "network.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_set_default_locking_mode(session_id interface{}, network interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = network
	params[2] = value

	err = client.RPCCall(&result, "network.set_default_locking_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_create_new_blob(session_id interface{}, network interface{}, name string, mime_type string, public bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = network
	params[2] = name
	params[3] = mime_type
	params[4] = public

	err = client.RPCCall(&result, "network.create_new_blob", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_remove_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "network.remove_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_add_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "network.add_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_set_tags(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "network.set_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "network.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "network.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "network.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_set_MTU(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "network.set_MTU", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "network.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "network.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_assigned_ips(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_assigned_ips", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_default_locking_mode(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_default_locking_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_tags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_blobs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_blobs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_bridge(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_bridge", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_MTU(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_MTU", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_PIFs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_PIFs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_VIFs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_VIFs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "network.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "network.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "network.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) network_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "network.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A virtual network interface

func (client *XenAPIClient) VIF_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VIF.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VIF.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_remove_ipv6_allowed(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.remove_ipv6_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_add_ipv6_allowed(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.add_ipv6_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_set_ipv6_allowed(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.set_ipv6_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_remove_ipv4_allowed(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.remove_ipv4_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_add_ipv4_allowed(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.add_ipv4_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_set_ipv4_allowed(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.set_ipv4_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_set_locking_mode(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.set_locking_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_unplug_force(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.unplug_force", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_unplug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.unplug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_plug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.plug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_remove_from_qos_algorithm_params(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VIF.remove_from_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_add_to_qos_algorithm_params(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VIF.add_to_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_set_qos_algorithm_params(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.set_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_set_qos_algorithm_type(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.set_qos_algorithm_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VIF.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VIF.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_ipv6_allowed(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_ipv6_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_ipv4_allowed(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_ipv4_allowed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_locking_mode(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_locking_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_MAC_autogenerated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_MAC_autogenerated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_metrics(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_metrics", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_qos_supported_algorithms(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_qos_supported_algorithms", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_qos_algorithm_params(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_qos_algorithm_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_qos_algorithm_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_runtime_properties(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_runtime_properties", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_status_detail(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_status_detail", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_status_code(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_status_code", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_currently_attached(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_currently_attached", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_MTU(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_MTU", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_MAC(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_MAC", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_VM(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_VM", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_network(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_network", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_device(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_device", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "VIF.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VIF.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// The metrics associated with a virtual network device

func (client *XenAPIClient) VIF_metrics_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VIF_metrics.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VIF_metrics.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VIF_metrics.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VIF_metrics.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VIF_metrics.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF_metrics.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_last_updated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF_metrics.get_last_updated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF_metrics.get_io_write_kbs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF_metrics.get_io_read_kbs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF_metrics.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VIF_metrics.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VIF_metrics_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VIF_metrics.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A physical network interface (note separate VLANs are represented as several PIFs)

func (client *XenAPIClient) PIF_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PIF.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PIF.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_set_property(session_id interface{}, self interface{}, name string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = name
	params[3] = value

	err = client.RPCCall(&result, "PIF.set_property", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_db_forget(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.db_forget", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_db_introduce(session_id interface{}, device string, network interface{}, host interface{}, MAC string, MTU interface{}, VLAN interface{}, physical bool, ip_configuration_mode interface{}, IP string, netmask string, gateway string, DNS string, bond_slave_of interface{}, VLAN_master_of interface{}, management bool, other_config map[string]string, disallow_unplug bool, ipv6_configuration_mode interface{}, IPv6 interface{}, ipv6_gateway string, primary_address_type interface{}, managed bool, properties map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 24)

	params[0] = session_id
	params[1] = device
	params[2] = network
	params[3] = host
	params[4] = MAC
	params[5] = MTU
	params[6] = VLAN
	params[7] = physical
	params[8] = ip_configuration_mode
	params[9] = IP
	params[10] = netmask
	params[11] = gateway
	params[12] = DNS
	params[13] = bond_slave_of
	params[14] = VLAN_master_of
	params[15] = management
	params[16] = other_config
	params[17] = disallow_unplug
	params[18] = ipv6_configuration_mode
	params[19] = IPv6
	params[20] = ipv6_gateway
	params[21] = primary_address_type
	params[22] = managed
	params[23] = properties

	err = client.RPCCall(&result, "PIF.db_introduce", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_plug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.plug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_unplug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.unplug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_forget(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.forget", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_introduce(session_id interface{}, host interface{}, MAC string, device string, managed bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = host
	params[2] = MAC
	params[3] = device
	params[4] = managed

	err = client.RPCCall(&result, "PIF.introduce", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_scan(session_id interface{}, host interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = host

	err = client.RPCCall(&result, "PIF.scan", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_set_primary_address_type(session_id interface{}, self interface{}, primary_address_type interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = primary_address_type

	err = client.RPCCall(&result, "PIF.set_primary_address_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_reconfigure_ipv6(session_id interface{}, self interface{}, mode interface{}, IPv6 string, gateway string, DNS string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 6)

	params[0] = session_id
	params[1] = self
	params[2] = mode
	params[3] = IPv6
	params[4] = gateway
	params[5] = DNS

	err = client.RPCCall(&result, "PIF.reconfigure_ipv6", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_reconfigure_ip(session_id interface{}, self interface{}, mode interface{}, IP string, netmask string, gateway string, DNS string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 7)

	params[0] = session_id
	params[1] = self
	params[2] = mode
	params[3] = IP
	params[4] = netmask
	params[5] = gateway
	params[6] = DNS

	err = client.RPCCall(&result, "PIF.reconfigure_ip", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_create_VLAN(session_id interface{}, device string, network interface{}, host interface{}, VLAN interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = device
	params[2] = network
	params[3] = host
	params[4] = VLAN

	err = client.RPCCall(&result, "PIF.create_VLAN", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_set_disallow_unplug(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PIF.set_disallow_unplug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "PIF.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "PIF.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PIF.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_properties(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_properties", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_managed(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_managed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_primary_address_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_primary_address_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_ipv6_gateway(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_ipv6_gateway", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_IPv6(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_IPv6", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_ipv6_configuration_mode(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_ipv6_configuration_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_tunnel_transport_PIF_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_tunnel_transport_PIF_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_tunnel_access_PIF_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_tunnel_access_PIF_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_disallow_unplug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_disallow_unplug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_management(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_management", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_VLAN_slave_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_VLAN_slave_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_VLAN_master_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_VLAN_master_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_bond_master_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_bond_master_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_bond_slave_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_bond_slave_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_DNS(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_DNS", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_gateway(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_gateway", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_netmask(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_netmask", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_IP(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_IP", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_ip_configuration_mode(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_ip_configuration_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_currently_attached(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_currently_attached", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_physical(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_physical", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_metrics(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_metrics", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_VLAN(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_VLAN", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_MTU(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_MTU", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_MAC(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_MAC", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_network(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_network", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_device(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_device", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "PIF.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// The metrics associated with a physical network interface

func (client *XenAPIClient) PIF_metrics_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PIF_metrics.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PIF_metrics.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "PIF_metrics.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "PIF_metrics.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PIF_metrics.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_last_updated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_last_updated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_pci_bus_path(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_pci_bus_path", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_duplex(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_duplex", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_speed(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_speed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_device_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_device_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_device_id(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_device_id", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_vendor_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_vendor_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_vendor_id(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_vendor_id", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_carrier(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_carrier", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_io_write_kbs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_io_read_kbs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "PIF_metrics.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PIF_metrics_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PIF_metrics.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

//

func (client *XenAPIClient) Bond_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "Bond.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "Bond.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_set_property(session_id interface{}, self interface{}, name string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = name
	params[3] = value

	err = client.RPCCall(&result, "Bond.set_property", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_set_mode(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "Bond.set_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_create(session_id interface{}, network interface{}, members interface{}, MAC string, mode interface{}, properties map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 6)

	params[0] = session_id
	params[1] = network
	params[2] = members
	params[3] = MAC
	params[4] = mode
	params[5] = properties

	err = client.RPCCall(&result, "Bond.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "Bond.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "Bond.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "Bond.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_links_up(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_links_up", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_properties(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_properties", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_mode(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_primary_slave(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_primary_slave", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_slaves(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_slaves", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_master(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_master", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "Bond.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) Bond_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "Bond.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A VLAN mux/demux

func (client *XenAPIClient) VLAN_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VLAN.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VLAN.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VLAN.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_create(session_id interface{}, tagged_PIF interface{}, tag interface{}, network interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = tagged_PIF
	params[2] = tag
	params[3] = network

	err = client.RPCCall(&result, "VLAN.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VLAN.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VLAN.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VLAN.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VLAN.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_tag(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VLAN.get_tag", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_untagged_PIF(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VLAN.get_untagged_PIF", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_tagged_PIF(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VLAN.get_tagged_PIF", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VLAN.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VLAN.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VLAN_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VLAN.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A storage manager plugin

func (client *XenAPIClient) SM_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "SM.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "SM.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "SM.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "SM.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SM.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_driver_filename(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_driver_filename", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_features(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_features", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_capabilities(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_capabilities", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_configuration(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_configuration", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_required_api_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_required_api_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_version(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_version", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_copyright(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_copyright", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_vendor(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_vendor", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "SM.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "SM.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SM_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SM.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A storage repository

func (client *XenAPIClient) SR_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "SR.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "SR.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_disable_database_replication(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.disable_database_replication", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_enable_database_replication(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.enable_database_replication", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_assert_supports_database_replication(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.assert_supports_database_replication", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_assert_can_host_ha_statefile(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.assert_can_host_ha_statefile", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_physical_utilisation(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.set_physical_utilisation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_virtual_allocation(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.set_virtual_allocation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_physical_size(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.set_physical_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_create_new_blob(session_id interface{}, sr interface{}, name string, mime_type string, public bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = sr
	params[2] = name
	params[3] = mime_type
	params[4] = public

	err = client.RPCCall(&result, "SR.create_new_blob", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_name_description(session_id interface{}, sr interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = sr
	params[2] = value

	err = client.RPCCall(&result, "SR.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_name_label(session_id interface{}, sr interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = sr
	params[2] = value

	err = client.RPCCall(&result, "SR.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_shared(session_id interface{}, sr interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = sr
	params[2] = value

	err = client.RPCCall(&result, "SR.set_shared", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_probe(session_id interface{}, host interface{}, device_config map[string]string, a_type string, sm_config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = host
	params[2] = device_config
	params[3] = a_type
	params[4] = sm_config

	err = client.RPCCall(&result, "SR.probe", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_scan(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.scan", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_supported_types(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "SR.get_supported_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_update(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.update", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_forget(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.forget", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_destroy(session_id interface{}, sr interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = sr

	err = client.RPCCall(&result, "SR.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_make(session_id interface{}, host interface{}, device_config map[string]string, physical_size interface{}, name_label string, name_description string, a_type string, content_type string, sm_config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 9)

	params[0] = session_id
	params[1] = host
	params[2] = device_config
	params[3] = physical_size
	params[4] = name_label
	params[5] = name_description
	params[6] = a_type
	params[7] = content_type
	params[8] = sm_config

	err = client.RPCCall(&result, "SR.make", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_introduce(session_id interface{}, uuid string, name_label string, name_description string, a_type string, content_type string, shared bool, sm_config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 8)

	params[0] = session_id
	params[1] = uuid
	params[2] = name_label
	params[3] = name_description
	params[4] = a_type
	params[5] = content_type
	params[6] = shared
	params[7] = sm_config

	err = client.RPCCall(&result, "SR.introduce", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_create(session_id interface{}, host interface{}, device_config map[string]string, physical_size interface{}, name_label string, name_description string, a_type string, content_type string, shared bool, sm_config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 10)

	params[0] = session_id
	params[1] = host
	params[2] = device_config
	params[3] = physical_size
	params[4] = name_label
	params[5] = name_description
	params[6] = a_type
	params[7] = content_type
	params[8] = shared
	params[9] = sm_config

	err = client.RPCCall(&result, "SR.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_remove_from_sm_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "SR.remove_from_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_add_to_sm_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "SR.add_to_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_sm_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.set_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_remove_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.remove_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_add_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.add_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_tags(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.set_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "SR.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "SR.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "SR.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_introduced_by(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_introduced_by", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_local_cache_enabled(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_local_cache_enabled", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_blobs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_blobs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_sm_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_tags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_shared(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_shared", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_content_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_content_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_physical_size(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_physical_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_physical_utilisation(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_physical_utilisation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_virtual_allocation(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_virtual_allocation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_PBDs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_PBDs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_VDIs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_VDIs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "SR.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "SR.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) SR_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "SR.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A virtual disk image

func (client *XenAPIClient) VDI_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VDI.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VDI.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_pool_migrate(session_id interface{}, vdi interface{}, sr interface{}, options map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = vdi
	params[2] = sr
	params[3] = options

	err = client.RPCCall(&result, "VDI.pool_migrate", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_read_database_pool_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.read_database_pool_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_open_database(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.open_database", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_allow_caching(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_allow_caching", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_on_boot(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_on_boot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_metadata_of_pool(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_metadata_of_pool", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_snapshot_time(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_snapshot_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_snapshot_of(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_snapshot_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_is_a_snapshot(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_is_a_snapshot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_physical_utilisation(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_physical_utilisation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_virtual_size(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_virtual_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_missing(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_missing", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_read_only(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_read_only", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_sharable(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_sharable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_forget(session_id interface{}, vdi interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vdi

	err = client.RPCCall(&result, "VDI.forget", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_managed(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_managed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_copy(session_id interface{}, vdi interface{}, sr interface{}, base_vdi interface{}, into_vdi interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 5)

	params[0] = session_id
	params[1] = vdi
	params[2] = sr
	params[3] = base_vdi
	params[4] = into_vdi

	err = client.RPCCall(&result, "VDI.copy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_update(session_id interface{}, vdi interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vdi

	err = client.RPCCall(&result, "VDI.update", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_db_forget(session_id interface{}, vdi interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vdi

	err = client.RPCCall(&result, "VDI.db_forget", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_db_introduce(session_id interface{}, uuid string, name_label string, name_description string, SR interface{}, a_type interface{}, sharable bool, read_only bool, other_config map[string]string, location string, xenstore_data map[string]string, sm_config map[string]string, managed bool, virtual_size interface{}, physical_utilisation interface{}, metadata_of_pool interface{}, is_a_snapshot bool, snapshot_time interface{}, snapshot_of interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 19)

	params[0] = session_id
	params[1] = uuid
	params[2] = name_label
	params[3] = name_description
	params[4] = SR
	params[5] = a_type
	params[6] = sharable
	params[7] = read_only
	params[8] = other_config
	params[9] = location
	params[10] = xenstore_data
	params[11] = sm_config
	params[12] = managed
	params[13] = virtual_size
	params[14] = physical_utilisation
	params[15] = metadata_of_pool
	params[16] = is_a_snapshot
	params[17] = snapshot_time
	params[18] = snapshot_of

	err = client.RPCCall(&result, "VDI.db_introduce", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_introduce(session_id interface{}, uuid string, name_label string, name_description string, SR interface{}, a_type interface{}, sharable bool, read_only bool, other_config map[string]string, location string, xenstore_data map[string]string, sm_config map[string]string, managed bool, virtual_size interface{}, physical_utilisation interface{}, metadata_of_pool interface{}, is_a_snapshot bool, snapshot_time interface{}, snapshot_of interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 19)

	params[0] = session_id
	params[1] = uuid
	params[2] = name_label
	params[3] = name_description
	params[4] = SR
	params[5] = a_type
	params[6] = sharable
	params[7] = read_only
	params[8] = other_config
	params[9] = location
	params[10] = xenstore_data
	params[11] = sm_config
	params[12] = managed
	params[13] = virtual_size
	params[14] = physical_utilisation
	params[15] = metadata_of_pool
	params[16] = is_a_snapshot
	params[17] = snapshot_time
	params[18] = snapshot_of

	err = client.RPCCall(&result, "VDI.introduce", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_resize_online(session_id interface{}, vdi interface{}, size interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vdi
	params[2] = size

	err = client.RPCCall(&result, "VDI.resize_online", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_resize(session_id interface{}, vdi interface{}, size interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vdi
	params[2] = size

	err = client.RPCCall(&result, "VDI.resize", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_clone(session_id interface{}, vdi interface{}, driver_params map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vdi
	params[2] = driver_params

	err = client.RPCCall(&result, "VDI.clone", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_snapshot(session_id interface{}, vdi interface{}, driver_params map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vdi
	params[2] = driver_params

	err = client.RPCCall(&result, "VDI.snapshot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_remove_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.remove_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_add_tags(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.add_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_tags(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_remove_from_sm_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VDI.remove_from_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_add_to_sm_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VDI.add_to_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_sm_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_remove_from_xenstore_data(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VDI.remove_from_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_add_to_xenstore_data(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VDI.add_to_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_xenstore_data(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VDI.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VDI.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VDI.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_metadata_latest(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_metadata_latest", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_metadata_of_pool(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_metadata_of_pool", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_on_boot(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_on_boot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_allow_caching(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_allow_caching", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_tags(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_tags", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_snapshot_time(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_snapshot_time", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_snapshots(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_snapshots", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_snapshot_of(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_snapshot_of", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_is_a_snapshot(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_is_a_snapshot", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_sm_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_sm_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_xenstore_data(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_xenstore_data", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_parent(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_parent", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_missing(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_missing", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_managed(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_managed", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_location(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_location", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_storage_lock(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_storage_lock", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_read_only(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_read_only", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_sharable(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_sharable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_physical_utilisation(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_physical_utilisation", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_virtual_size(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_virtual_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_crash_dumps(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_crash_dumps", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_VBDs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_VBDs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_SR(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "VDI.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "VDI.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VDI.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VDI_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VDI.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A virtual block device

func (client *XenAPIClient) VBD_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VBD.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VBD.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_assert_attachable(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.assert_attachable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_unplug_force(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.unplug_force", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_unplug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.unplug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_plug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.plug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_insert(session_id interface{}, vbd interface{}, vdi interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = vbd
	params[2] = vdi

	err = client.RPCCall(&result, "VBD.insert", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_eject(session_id interface{}, vbd interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = vbd

	err = client.RPCCall(&result, "VBD.eject", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_remove_from_qos_algorithm_params(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VBD.remove_from_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_add_to_qos_algorithm_params(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VBD.add_to_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_qos_algorithm_params(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_qos_algorithm_type(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_qos_algorithm_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VBD.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VBD.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_unpluggable(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_unpluggable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_type(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_mode(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_bootable(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_bootable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_set_userdevice(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD.set_userdevice", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_metrics(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_metrics", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_qos_supported_algorithms(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_qos_supported_algorithms", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_qos_algorithm_params(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_qos_algorithm_params", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_qos_algorithm_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_qos_algorithm_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_runtime_properties(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_runtime_properties", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_status_detail(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_status_detail", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_status_code(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_status_code", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_currently_attached(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_currently_attached", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_empty(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_empty", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_storage_lock(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_storage_lock", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_unpluggable(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_unpluggable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_mode(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_mode", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_bootable(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_bootable", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_userdevice(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_userdevice", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_device(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_device", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_VDI(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_VDI", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_VM(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_VM", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_current_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_current_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_allowed_operations(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_allowed_operations", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "VBD.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VBD.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// The metrics associated with a virtual block device

func (client *XenAPIClient) VBD_metrics_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VBD_metrics.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VBD_metrics.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VBD_metrics.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VBD_metrics.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VBD_metrics.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD_metrics.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_last_updated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD_metrics.get_last_updated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD_metrics.get_io_write_kbs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD_metrics.get_io_read_kbs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD_metrics.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VBD_metrics.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VBD_metrics_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VBD_metrics.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// The physical block devices through which hosts access SRs

func (client *XenAPIClient) PBD_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PBD.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PBD.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_set_device_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PBD.set_device_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_unplug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.unplug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_plug(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.plug", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "PBD.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "PBD.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PBD.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_currently_attached(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.get_currently_attached", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_device_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.get_device_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_SR(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.get_SR", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.get_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "PBD.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "PBD.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PBD_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PBD.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A VM crashdump

func (client *XenAPIClient) crashdump_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "crashdump.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "crashdump.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "crashdump.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "crashdump.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "crashdump.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "crashdump.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "crashdump.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_get_VDI(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "crashdump.get_VDI", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_get_VM(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "crashdump.get_VM", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "crashdump.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "crashdump.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) crashdump_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "crashdump.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A virtual TPM device

func (client *XenAPIClient) VTPM_get_backend(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VTPM.get_backend", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VTPM_get_VM(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VTPM.get_VM", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VTPM_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VTPM.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VTPM_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VTPM.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VTPM_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "VTPM.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VTPM_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VTPM.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VTPM_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VTPM.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A console

func (client *XenAPIClient) console_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "console.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "console.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "console.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "console.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "console.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "console.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_VM(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "console.get_VM", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_location(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "console.get_location", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_protocol(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "console.get_protocol", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "console.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "console.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "console.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "console.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) console_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "console.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A user of the system

func (client *XenAPIClient) user_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "user.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "user.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "user.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_set_fullname(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "user.set_fullname", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "user.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_get_fullname(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "user.get_fullname", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_get_short_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "user.get_short_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "user.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "user.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "user.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "user.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) user_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "user.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// Data sources for logging in RRDs

// A placeholder for a binary blob

func (client *XenAPIClient) blob_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "blob.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "blob.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_create(session_id interface{}, mime_type string, public bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = mime_type
	params[2] = public

	err = client.RPCCall(&result, "blob.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_set_public(session_id interface{}, self interface{}, value bool) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "blob.set_public", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "blob.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "blob.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_mime_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_mime_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_last_updated(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_last_updated", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_public(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_public", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_size(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "blob.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "blob.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) blob_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "blob.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// An message for the attention of the administrator

func (client *XenAPIClient) message_get_all_records_where(session_id interface{}, expr string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = expr

	err = client.RPCCall(&result, "message.get_all_records_where", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "message.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "message.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "message.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_get_since(session_id interface{}, since interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = since

	err = client.RPCCall(&result, "message.get_since", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "message.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_get(session_id interface{}, cls interface{}, obj_uuid string, since interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = cls
	params[2] = obj_uuid
	params[3] = since

	err = client.RPCCall(&result, "message.get", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "message.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) message_create(session_id interface{}, name string, priority interface{}, cls interface{}, obj_uuid string, body string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 6)

	params[0] = session_id
	params[1] = name
	params[2] = priority
	params[3] = cls
	params[4] = obj_uuid
	params[5] = body

	err = client.RPCCall(&result, "message.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A secret

func (client *XenAPIClient) secret_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "secret.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "secret.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "secret.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "secret.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "secret.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_set_value(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "secret.set_value", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "secret.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_get_value(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "secret.get_value", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "secret.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "secret.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_create(session_id interface{}, args interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = args

	err = client.RPCCall(&result, "secret.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "secret.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) secret_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "secret.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A tunnel for network traffic

func (client *XenAPIClient) tunnel_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "tunnel.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "tunnel.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "tunnel.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_create(session_id interface{}, transport_PIF interface{}, network interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = transport_PIF
	params[2] = network

	err = client.RPCCall(&result, "tunnel.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "tunnel.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "tunnel.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "tunnel.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_remove_from_status(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "tunnel.remove_from_status", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_add_to_status(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "tunnel.add_to_status", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_set_status(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "tunnel.set_status", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "tunnel.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_status(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "tunnel.get_status", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_transport_PIF(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "tunnel.get_transport_PIF", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_access_PIF(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "tunnel.get_access_PIF", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "tunnel.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "tunnel.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) tunnel_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "tunnel.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A PCI device

func (client *XenAPIClient) PCI_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PCI.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PCI.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "PCI.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "PCI.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PCI.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_subsystem_device_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_subsystem_device_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_subsystem_vendor_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_subsystem_vendor_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_dependencies(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_dependencies", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_pci_id(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_pci_id", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_device_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_device_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_vendor_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_vendor_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_class_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_class_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "PCI.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PCI_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PCI.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A physical GPU (pGPU)

func (client *XenAPIClient) PGPU_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PGPU.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "PGPU.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_disable_dom0_access(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.disable_dom0_access", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_enable_dom0_access(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.enable_dom0_access", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_remaining_capacity(session_id interface{}, self interface{}, vgpu_type interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = vgpu_type

	err = client.RPCCall(&result, "PGPU.get_remaining_capacity", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_set_GPU_group(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PGPU.set_GPU_group", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_set_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PGPU.set_enabled_VGPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_remove_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PGPU.remove_enabled_VGPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_add_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PGPU.add_enabled_VGPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "PGPU.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "PGPU.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "PGPU.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_is_system_display_device(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_is_system_display_device", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_dom0_access(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_dom0_access", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_supported_VGPU_max_capacities(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_supported_VGPU_max_capacities", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_resident_VGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_resident_VGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_enabled_VGPU_types(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_enabled_VGPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_supported_VGPU_types(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_supported_VGPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_host(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_host", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_GPU_group(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_GPU_group", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_PCI(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_PCI", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "PGPU.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) PGPU_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "PGPU.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A group of compatible GPUs across the resource pool

func (client *XenAPIClient) GPU_group_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "GPU_group.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "GPU_group.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_remaining_capacity(session_id interface{}, self interface{}, vgpu_type interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = vgpu_type

	err = client.RPCCall(&result, "GPU_group.get_remaining_capacity", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_create(session_id interface{}, name_label string, name_description string, other_config map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = name_label
	params[2] = name_description
	params[3] = other_config

	err = client.RPCCall(&result, "GPU_group.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_set_allocation_algorithm(session_id interface{}, self interface{}, value interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "GPU_group.set_allocation_algorithm", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "GPU_group.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "GPU_group.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "GPU_group.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_set_name_description(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "GPU_group.set_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_set_name_label(session_id interface{}, self interface{}, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "GPU_group.set_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_enabled_VGPU_types(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_enabled_VGPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_supported_VGPU_types(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_supported_VGPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_allocation_algorithm(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_allocation_algorithm", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_GPU_types(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_GPU_types", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_VGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_VGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_PGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_PGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_name_description(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_name_description", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_name_label(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_by_name_label(session_id interface{}, label string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = label

	err = client.RPCCall(&result, "GPU_group.get_by_name_label", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "GPU_group.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) GPU_group_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "GPU_group.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A virtual GPU (vGPU)

func (client *XenAPIClient) VGPU_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VGPU.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VGPU.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_destroy(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.destroy", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_create(session_id interface{}, VM interface{}, GPU_group interface{}, device string, other_config map[string]string, a_type interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 6)

	params[0] = session_id
	params[1] = VM
	params[2] = GPU_group
	params[3] = device
	params[4] = other_config
	params[5] = a_type

	err = client.RPCCall(&result, "VGPU.create", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_remove_from_other_config(session_id interface{}, self interface{}, key string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = key

	err = client.RPCCall(&result, "VGPU.remove_from_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 4)

	params[0] = session_id
	params[1] = self
	params[2] = key
	params[3] = value

	err = client.RPCCall(&result, "VGPU.add_to_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_set_other_config(session_id interface{}, self interface{}, value map[string]string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 3)

	params[0] = session_id
	params[1] = self
	params[2] = value

	err = client.RPCCall(&result, "VGPU.set_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_resident_on(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_resident_on", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_type(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_type", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_other_config(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_other_config", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_currently_attached(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_currently_attached", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_device(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_device", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_GPU_group(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_GPU_group", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_VM(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_VM", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VGPU.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPU_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPU.get_record", params)
	resultValue = result["Value"]

	return resultValue, err
}

// A type of virtual GPU

func (client *XenAPIClient) VGPUType_get_all_records(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VGPUType.get_all_records", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_all(session_id interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 1)

	params[0] = session_id

	err = client.RPCCall(&result, "VGPUType.get_all", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_enabled_on_GPU_groups(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_enabled_on_GPU_groups", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_supported_on_GPU_groups(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_supported_on_GPU_groups", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_VGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_VGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_enabled_on_PGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_enabled_on_PGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_supported_on_PGPUs(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_supported_on_PGPUs", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_max_resolution_y(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_max_resolution_y", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_max_resolution_x(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_max_resolution_x", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_max_heads(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_max_heads", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_framebuffer_size(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_framebuffer_size", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_model_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_model_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_vendor_name(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_vendor_name", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_uuid(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_by_uuid(session_id interface{}, uuid string) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = uuid

	err = client.RPCCall(&result, "VGPUType.get_by_uuid", params)
	resultValue = result["Value"]

	return resultValue, err
}

func (client *XenAPIClient) VGPUType_get_record(session_id interface{}, self interface{}) (resultValue interface{}, err error) {
	result := xmlrpc.Struct{}

	params := make([]interface{}, 2)

	params[0] = session_id
	params[1] = self

	err = client.RPCCall(&result, "VGPUType.get_record", params)
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

	res, err := client.session_local_logout(session)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("res: %+v\n", res)

	//    client.Username=user
	//    client.Password=passwd
	// err = client.Login()
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("session: %+v\n", client.Session)
}
