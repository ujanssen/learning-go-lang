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

func (c *XenAPIClient) RPCCall(method string, argi ...interface{}) (resultValue interface{}, err error) {

	params := make([]interface{}, len(argi))
	for i, element := range argi {
		params[i] = element
	}

	xmlrpcParams := new(xmlrpc.Params)
	xmlrpcParams.Params = params

	result := xmlrpc.Struct{}

	err = c.RPC.Call(method, *xmlrpcParams, &result)

	if err != nil {
		fmt.Errorf("RPC Call Error: %s", err)
	} else {
		resultValue = result["Value"]
	}
	return resultValue, err
}

// A session

func (client *XenAPIClient) session_logout_subject_identifier(session_id interface{}, subject_identifier string) (i interface{}, err error) {
	return client.RPCCall("session.logout_subject_identifier", session_id, subject_identifier)
}

func (client *XenAPIClient) session_get_all_subject_identifiers(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_all_subject_identifiers", session_id)
}

func (client *XenAPIClient) session_local_logout(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("session.local_logout", session_id)
}

func (client *XenAPIClient) session_slave_local_login_with_password(uname string, pwd string) (i interface{}, err error) {
	return client.RPCCall("session.slave_local_login_with_password", uname, pwd)
}

func (client *XenAPIClient) session_change_password(session_id interface{}, old_pwd string, new_pwd string) (i interface{}, err error) {
	return client.RPCCall("session.change_password", session_id, old_pwd, new_pwd)
}

func (client *XenAPIClient) session_logout(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("session.logout", session_id)
}

func (client *XenAPIClient) session_login_with_password(uname string, pwd string, version string, originator string) (i interface{}, err error) {
	return client.RPCCall("session.login_with_password", uname, pwd, version, originator)
}

func (client *XenAPIClient) session_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("session.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) session_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("session.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) session_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("session.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) session_get_originator(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_originator", session_id, self)
}

func (client *XenAPIClient) session_get_parent(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_parent", session_id, self)
}

func (client *XenAPIClient) session_get_tasks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_tasks", session_id, self)
}

func (client *XenAPIClient) session_get_rbac_permissions(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_rbac_permissions", session_id, self)
}

func (client *XenAPIClient) session_get_auth_user_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_auth_user_name", session_id, self)
}

func (client *XenAPIClient) session_get_auth_user_sid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_auth_user_sid", session_id, self)
}

func (client *XenAPIClient) session_get_validation_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_validation_time", session_id, self)
}

func (client *XenAPIClient) session_get_subject(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_subject", session_id, self)
}

func (client *XenAPIClient) session_get_is_local_superuser(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_is_local_superuser", session_id, self)
}

func (client *XenAPIClient) session_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_other_config", session_id, self)
}

func (client *XenAPIClient) session_get_pool(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_pool", session_id, self)
}

func (client *XenAPIClient) session_get_last_active(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_last_active", session_id, self)
}

func (client *XenAPIClient) session_get_this_user(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_this_user", session_id, self)
}

func (client *XenAPIClient) session_get_this_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_this_host", session_id, self)
}

func (client *XenAPIClient) session_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_uuid", session_id, self)
}

func (client *XenAPIClient) session_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("session.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) session_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_record", session_id, self)
}

// Management of remote authentication services

func (client *XenAPIClient) auth_get_group_membership(session_id interface{}, subject_identifier string) (i interface{}, err error) {
	return client.RPCCall("auth.get_group_membership", session_id, subject_identifier)
}

func (client *XenAPIClient) auth_get_subject_information_from_identifier(session_id interface{}, subject_identifier string) (i interface{}, err error) {
	return client.RPCCall("auth.get_subject_information_from_identifier", session_id, subject_identifier)
}

func (client *XenAPIClient) auth_get_subject_identifier(session_id interface{}, subject_name string) (i interface{}, err error) {
	return client.RPCCall("auth.get_subject_identifier", session_id, subject_name)
}

// A user or group that can log in xapi

func (client *XenAPIClient) subject_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_all_records", session_id)
}

func (client *XenAPIClient) subject_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_all", session_id)
}

func (client *XenAPIClient) subject_get_permissions_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_permissions_name_label", session_id, self)
}

func (client *XenAPIClient) subject_remove_from_roles(session_id interface{}, self interface{}, role interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.remove_from_roles", session_id, self, role)
}

func (client *XenAPIClient) subject_add_to_roles(session_id interface{}, self interface{}, role interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.add_to_roles", session_id, self, role)
}

func (client *XenAPIClient) subject_get_roles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_roles", session_id, self)
}

func (client *XenAPIClient) subject_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_other_config", session_id, self)
}

func (client *XenAPIClient) subject_get_subject_identifier(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_subject_identifier", session_id, self)
}

func (client *XenAPIClient) subject_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_uuid", session_id, self)
}

func (client *XenAPIClient) subject_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.destroy", session_id, self)
}

func (client *XenAPIClient) subject_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.create", session_id, args)
}

func (client *XenAPIClient) subject_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("subject.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) subject_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_record", session_id, self)
}

// A set of permissions associated with a subject

func (client *XenAPIClient) role_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_all_records", session_id)
}

func (client *XenAPIClient) role_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_all", session_id)
}

func (client *XenAPIClient) role_get_by_permission_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("role.get_by_permission_name_label", session_id, label)
}

func (client *XenAPIClient) role_get_by_permission(session_id interface{}, permission interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_by_permission", session_id, permission)
}

func (client *XenAPIClient) role_get_permissions_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_permissions_name_label", session_id, self)
}

func (client *XenAPIClient) role_get_permissions(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_permissions", session_id, self)
}

func (client *XenAPIClient) role_get_subroles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_subroles", session_id, self)
}

func (client *XenAPIClient) role_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_name_description", session_id, self)
}

func (client *XenAPIClient) role_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_name_label", session_id, self)
}

func (client *XenAPIClient) role_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_uuid", session_id, self)
}

func (client *XenAPIClient) role_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("role.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) role_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("role.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) role_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_record", session_id, self)
}

// A long-running asynchronous task

func (client *XenAPIClient) task_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_all_records", session_id)
}

func (client *XenAPIClient) task_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_all", session_id)
}

func (client *XenAPIClient) task_cancel(session_id interface{}, task interface{}) (i interface{}, err error) {
	return client.RPCCall("task.cancel", session_id, task)
}

func (client *XenAPIClient) task_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.destroy", session_id, self)
}

func (client *XenAPIClient) task_create(session_id interface{}, label string, description string) (i interface{}, err error) {
	return client.RPCCall("task.create", session_id, label, description)
}

func (client *XenAPIClient) task_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("task.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) task_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("task.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) task_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("task.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) task_get_backtrace(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_backtrace", session_id, self)
}

func (client *XenAPIClient) task_get_subtasks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_subtasks", session_id, self)
}

func (client *XenAPIClient) task_get_subtask_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_subtask_of", session_id, self)
}

func (client *XenAPIClient) task_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_other_config", session_id, self)
}

func (client *XenAPIClient) task_get_error_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_error_info", session_id, self)
}

func (client *XenAPIClient) task_get_result(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_result", session_id, self)
}

func (client *XenAPIClient) task_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_type", session_id, self)
}

func (client *XenAPIClient) task_get_progress(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_progress", session_id, self)
}

func (client *XenAPIClient) task_get_resident_on(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_resident_on", session_id, self)
}

func (client *XenAPIClient) task_get_status(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_status", session_id, self)
}

func (client *XenAPIClient) task_get_finished(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_finished", session_id, self)
}

func (client *XenAPIClient) task_get_created(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_created", session_id, self)
}

func (client *XenAPIClient) task_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_current_operations", session_id, self)
}

func (client *XenAPIClient) task_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) task_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_name_description", session_id, self)
}

func (client *XenAPIClient) task_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_name_label", session_id, self)
}

func (client *XenAPIClient) task_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_uuid", session_id, self)
}

func (client *XenAPIClient) task_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("task.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) task_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("task.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) task_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_record", session_id, self)
}

// Asynchronous event registration and handling

func (client *XenAPIClient) event_inject(session_id interface{}, class string, ref string) (i interface{}, err error) {
	return client.RPCCall("event.inject", session_id, class, ref)
}

func (client *XenAPIClient) event_get_current_id(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("event.get_current_id", session_id)
}

func (client *XenAPIClient) event_from(session_id interface{}, classes interface{}, token string, timeout interface{}) (i interface{}, err error) {
	return client.RPCCall("event.from", session_id, classes, token, timeout)
}

func (client *XenAPIClient) event_next(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("event.next", session_id)
}

func (client *XenAPIClient) event_unregister(session_id interface{}, classes interface{}) (i interface{}, err error) {
	return client.RPCCall("event.unregister", session_id, classes)
}

func (client *XenAPIClient) event_register(session_id interface{}, classes interface{}) (i interface{}, err error) {
	return client.RPCCall("event.register", session_id, classes)
}

// Pool-wide information

func (client *XenAPIClient) pool_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_all_records", session_id)
}

func (client *XenAPIClient) pool_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_all", session_id)
}

func (client *XenAPIClient) pool_apply_edition(session_id interface{}, self interface{}, edition string) (i interface{}, err error) {
	return client.RPCCall("pool.apply_edition", session_id, self, edition)
}

func (client *XenAPIClient) pool_get_license_state(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_license_state", session_id, self)
}

func (client *XenAPIClient) pool_disable_local_storage_caching(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.disable_local_storage_caching", session_id, self)
}

func (client *XenAPIClient) pool_enable_local_storage_caching(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.enable_local_storage_caching", session_id, self)
}

func (client *XenAPIClient) pool_test_archive_target(session_id interface{}, self interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.test_archive_target", session_id, self, config)
}

func (client *XenAPIClient) pool_set_vswitch_controller(session_id interface{}, address string) (i interface{}, err error) {
	return client.RPCCall("pool.set_vswitch_controller", session_id, address)
}

func (client *XenAPIClient) pool_disable_redo_log(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.disable_redo_log", session_id)
}

func (client *XenAPIClient) pool_enable_redo_log(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.enable_redo_log", session_id, sr)
}

func (client *XenAPIClient) pool_certificate_sync(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_sync", session_id)
}

func (client *XenAPIClient) pool_crl_list(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.crl_list", session_id)
}

func (client *XenAPIClient) pool_crl_uninstall(session_id interface{}, name string) (i interface{}, err error) {
	return client.RPCCall("pool.crl_uninstall", session_id, name)
}

func (client *XenAPIClient) pool_crl_install(session_id interface{}, name string, cert string) (i interface{}, err error) {
	return client.RPCCall("pool.crl_install", session_id, name, cert)
}

func (client *XenAPIClient) pool_certificate_list(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_list", session_id)
}

func (client *XenAPIClient) pool_certificate_uninstall(session_id interface{}, name string) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_uninstall", session_id, name)
}

func (client *XenAPIClient) pool_certificate_install(session_id interface{}, name string, cert string) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_install", session_id, name, cert)
}

func (client *XenAPIClient) pool_send_test_post(session_id interface{}, host string, port interface{}, body string) (i interface{}, err error) {
	return client.RPCCall("pool.send_test_post", session_id, host, port, body)
}

func (client *XenAPIClient) pool_retrieve_wlb_recommendations(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.retrieve_wlb_recommendations", session_id)
}

func (client *XenAPIClient) pool_retrieve_wlb_configuration(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.retrieve_wlb_configuration", session_id)
}

func (client *XenAPIClient) pool_send_wlb_configuration(session_id interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.send_wlb_configuration", session_id, config)
}

func (client *XenAPIClient) pool_deconfigure_wlb(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.deconfigure_wlb", session_id)
}

func (client *XenAPIClient) pool_initialize_wlb(session_id interface{}, wlb_url string, wlb_username string, wlb_password string, xenserver_username string, xenserver_password string) (i interface{}, err error) {
	return client.RPCCall("pool.initialize_wlb", session_id, wlb_url, wlb_username, wlb_password, xenserver_username, xenserver_password)
}

func (client *XenAPIClient) pool_detect_nonhomogeneous_external_auth(session_id interface{}, pool interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.detect_nonhomogeneous_external_auth", session_id, pool)
}

func (client *XenAPIClient) pool_disable_external_auth(session_id interface{}, pool interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.disable_external_auth", session_id, pool, config)
}

func (client *XenAPIClient) pool_enable_external_auth(session_id interface{}, pool interface{}, config map[string]string, service_name string, auth_type string) (i interface{}, err error) {
	return client.RPCCall("pool.enable_external_auth", session_id, pool, config, service_name, auth_type)
}

func (client *XenAPIClient) pool_create_new_blob(session_id interface{}, pool interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("pool.create_new_blob", session_id, pool, name, mime_type, public)
}

func (client *XenAPIClient) pool_set_ha_host_failures_to_tolerate(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_ha_host_failures_to_tolerate", session_id, self, value)
}

func (client *XenAPIClient) pool_ha_compute_vm_failover_plan(session_id interface{}, failed_hosts interface{}, failed_vms interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_compute_vm_failover_plan", session_id, failed_hosts, failed_vms)
}

func (client *XenAPIClient) pool_ha_compute_hypothetical_max_host_failures_to_tolerate(session_id interface{}, configuration interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_compute_hypothetical_max_host_failures_to_tolerate", session_id, configuration)
}

func (client *XenAPIClient) pool_ha_compute_max_host_failures_to_tolerate(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_compute_max_host_failures_to_tolerate", session_id)
}

func (client *XenAPIClient) pool_ha_failover_plan_exists(session_id interface{}, n interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_failover_plan_exists", session_id, n)
}

func (client *XenAPIClient) pool_ha_prevent_restarts_for(session_id interface{}, seconds interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_prevent_restarts_for", session_id, seconds)
}

func (client *XenAPIClient) pool_designate_new_master(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.designate_new_master", session_id, host)
}

func (client *XenAPIClient) pool_sync_database(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.sync_database", session_id)
}

func (client *XenAPIClient) pool_disable_ha(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.disable_ha", session_id)
}

func (client *XenAPIClient) pool_enable_ha(session_id interface{}, heartbeat_srs interface{}, configuration map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.enable_ha", session_id, heartbeat_srs, configuration)
}

func (client *XenAPIClient) pool_create_VLAN_from_PIF(session_id interface{}, pif interface{}, network interface{}, VLAN interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.create_VLAN_from_PIF", session_id, pif, network, VLAN)
}

func (client *XenAPIClient) pool_create_VLAN(session_id interface{}, device string, network interface{}, VLAN interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.create_VLAN", session_id, device, network, VLAN)
}

func (client *XenAPIClient) pool_recover_slaves(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.recover_slaves", session_id)
}

func (client *XenAPIClient) pool_emergency_reset_master(session_id interface{}, master_address string) (i interface{}, err error) {
	return client.RPCCall("pool.emergency_reset_master", session_id, master_address)
}

func (client *XenAPIClient) pool_emergency_transition_to_master(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.emergency_transition_to_master", session_id)
}

func (client *XenAPIClient) pool_eject(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.eject", session_id, host)
}

func (client *XenAPIClient) pool_join_force(session_id interface{}, master_address string, master_username string, master_password string) (i interface{}, err error) {
	return client.RPCCall("pool.join_force", session_id, master_address, master_username, master_password)
}

func (client *XenAPIClient) pool_join(session_id interface{}, master_address string, master_username string, master_password string) (i interface{}, err error) {
	return client.RPCCall("pool.join", session_id, master_address, master_username, master_password)
}

func (client *XenAPIClient) pool_set_wlb_verify_cert(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("pool.set_wlb_verify_cert", session_id, self, value)
}

func (client *XenAPIClient) pool_set_wlb_enabled(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("pool.set_wlb_enabled", session_id, self, value)
}

func (client *XenAPIClient) pool_remove_from_gui_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("pool.remove_from_gui_config", session_id, self, key)
}

func (client *XenAPIClient) pool_add_to_gui_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("pool.add_to_gui_config", session_id, self, key, value)
}

func (client *XenAPIClient) pool_set_gui_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.set_gui_config", session_id, self, value)
}

func (client *XenAPIClient) pool_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.remove_tags", session_id, self, value)
}

func (client *XenAPIClient) pool_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.add_tags", session_id, self, value)
}

func (client *XenAPIClient) pool_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_tags", session_id, self, value)
}

func (client *XenAPIClient) pool_set_ha_allow_overcommit(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("pool.set_ha_allow_overcommit", session_id, self, value)
}

func (client *XenAPIClient) pool_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("pool.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) pool_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("pool.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) pool_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) pool_set_crash_dump_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_crash_dump_SR", session_id, self, value)
}

func (client *XenAPIClient) pool_set_suspend_image_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_suspend_image_SR", session_id, self, value)
}

func (client *XenAPIClient) pool_set_default_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_default_SR", session_id, self, value)
}

func (client *XenAPIClient) pool_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) pool_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) pool_get_metadata_VDIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_metadata_VDIs", session_id, self)
}

func (client *XenAPIClient) pool_get_restrictions(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_restrictions", session_id, self)
}

func (client *XenAPIClient) pool_get_vswitch_controller(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_vswitch_controller", session_id, self)
}

func (client *XenAPIClient) pool_get_redo_log_vdi(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_redo_log_vdi", session_id, self)
}

func (client *XenAPIClient) pool_get_redo_log_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_redo_log_enabled", session_id, self)
}

func (client *XenAPIClient) pool_get_wlb_verify_cert(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_verify_cert", session_id, self)
}

func (client *XenAPIClient) pool_get_wlb_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_enabled", session_id, self)
}

func (client *XenAPIClient) pool_get_wlb_username(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_username", session_id, self)
}

func (client *XenAPIClient) pool_get_wlb_url(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_url", session_id, self)
}

func (client *XenAPIClient) pool_get_gui_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_gui_config", session_id, self)
}

func (client *XenAPIClient) pool_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_tags", session_id, self)
}

func (client *XenAPIClient) pool_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_blobs", session_id, self)
}

func (client *XenAPIClient) pool_get_ha_overcommitted(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_overcommitted", session_id, self)
}

func (client *XenAPIClient) pool_get_ha_allow_overcommit(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_allow_overcommit", session_id, self)
}

func (client *XenAPIClient) pool_get_ha_plan_exists_for(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_plan_exists_for", session_id, self)
}

func (client *XenAPIClient) pool_get_ha_host_failures_to_tolerate(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_host_failures_to_tolerate", session_id, self)
}

func (client *XenAPIClient) pool_get_ha_statefiles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_statefiles", session_id, self)
}

func (client *XenAPIClient) pool_get_ha_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_configuration", session_id, self)
}

func (client *XenAPIClient) pool_get_ha_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_enabled", session_id, self)
}

func (client *XenAPIClient) pool_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_other_config", session_id, self)
}

func (client *XenAPIClient) pool_get_crash_dump_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_crash_dump_SR", session_id, self)
}

func (client *XenAPIClient) pool_get_suspend_image_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_suspend_image_SR", session_id, self)
}

func (client *XenAPIClient) pool_get_default_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_default_SR", session_id, self)
}

func (client *XenAPIClient) pool_get_master(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_master", session_id, self)
}

func (client *XenAPIClient) pool_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_name_description", session_id, self)
}

func (client *XenAPIClient) pool_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_name_label", session_id, self)
}

func (client *XenAPIClient) pool_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_uuid", session_id, self)
}

func (client *XenAPIClient) pool_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("pool.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) pool_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_record", session_id, self)
}

// Pool-wide patches

func (client *XenAPIClient) poolPatch_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_all_records", session_id)
}

func (client *XenAPIClient) poolPatch_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_all", session_id)
}

func (client *XenAPIClient) poolPatch_clean_on_host(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.clean_on_host", session_id, self, host)
}

func (client *XenAPIClient) poolPatch_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.destroy", session_id, self)
}

func (client *XenAPIClient) poolPatch_pool_clean(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.pool_clean", session_id, self)
}

func (client *XenAPIClient) poolPatch_clean(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.clean", session_id, self)
}

func (client *XenAPIClient) poolPatch_precheck(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.precheck", session_id, self, host)
}

func (client *XenAPIClient) poolPatch_pool_apply(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.pool_apply", session_id, self)
}

func (client *XenAPIClient) poolPatch_apply(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.apply", session_id, self, host)
}

func (client *XenAPIClient) poolPatch_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) poolPatch_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) poolPatch_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) poolPatch_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_other_config", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_after_apply_guidance(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_after_apply_guidance", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_host_patches(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_host_patches", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_pool_applied(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_pool_applied", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_size", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_version", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_name_description", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_name_label", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_uuid", session_id, self)
}

func (client *XenAPIClient) poolPatch_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) poolPatch_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) poolPatch_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_record", session_id, self)
}

// A virtual machine (or 'guest').

func (client *XenAPIClient) VM_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_all_records", session_id)
}

func (client *XenAPIClient) VM_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_all", session_id)
}

func (client *XenAPIClient) VM_call_plugin(session_id interface{}, vm interface{}, plugin string, fn string, args map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.call_plugin", session_id, vm, plugin, fn, args)
}

func (client *XenAPIClient) VM_query_services(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.query_services", session_id, self)
}

func (client *XenAPIClient) VM_set_appliance(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_appliance", session_id, self, value)
}

func (client *XenAPIClient) VM_import_convert(session_id interface{}, a_type string, username string, password string, sr interface{}, remote_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.import_convert", session_id, a_type, username, password, sr, remote_config)
}

func (client *XenAPIClient) VM_recover(session_id interface{}, self interface{}, session_to interface{}, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.recover", session_id, self, session_to, force)
}

func (client *XenAPIClient) VM_get_SRs_required_for_recovery(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_SRs_required_for_recovery", session_id, self, session_to)
}

func (client *XenAPIClient) VM_assert_can_be_recovered(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_can_be_recovered", session_id, self, session_to)
}

func (client *XenAPIClient) VM_set_suspend_VDI(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_suspend_VDI", session_id, self, value)
}

func (client *XenAPIClient) VM_set_order(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_order", session_id, self, value)
}

func (client *XenAPIClient) VM_set_shutdown_delay(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_shutdown_delay", session_id, self, value)
}

func (client *XenAPIClient) VM_set_start_delay(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_start_delay", session_id, self, value)
}

func (client *XenAPIClient) VM_set_protection_policy(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_protection_policy", session_id, self, value)
}

func (client *XenAPIClient) VM_copy_bios_strings(session_id interface{}, vm interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.copy_bios_strings", session_id, vm, host)
}

func (client *XenAPIClient) VM_retrieve_wlb_recommendations(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.retrieve_wlb_recommendations", session_id, vm)
}

func (client *XenAPIClient) VM_assert_agile(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_agile", session_id, self)
}

func (client *XenAPIClient) VM_create_new_blob(session_id interface{}, vm interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("VM.create_new_blob", session_id, vm, name, mime_type, public)
}

func (client *XenAPIClient) VM_assert_can_boot_here(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_can_boot_here", session_id, self, host)
}

func (client *XenAPIClient) VM_get_possible_hosts(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_possible_hosts", session_id, vm)
}

func (client *XenAPIClient) VM_get_allowed_VIF_devices(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_allowed_VIF_devices", session_id, vm)
}

func (client *XenAPIClient) VM_get_allowed_VBD_devices(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_allowed_VBD_devices", session_id, vm)
}

func (client *XenAPIClient) VM_update_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.update_allowed_operations", session_id, self)
}

func (client *XenAPIClient) VM_assert_operation_valid(session_id interface{}, self interface{}, op interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_operation_valid", session_id, self, op)
}

func (client *XenAPIClient) VM_forget_data_source_archives(session_id interface{}, self interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("VM.forget_data_source_archives", session_id, self, data_source)
}

func (client *XenAPIClient) VM_query_data_source(session_id interface{}, self interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("VM.query_data_source", session_id, self, data_source)
}

func (client *XenAPIClient) VM_record_data_source(session_id interface{}, self interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("VM.record_data_source", session_id, self, data_source)
}

func (client *XenAPIClient) VM_get_data_sources(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_data_sources", session_id, self)
}

func (client *XenAPIClient) VM_get_boot_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_boot_record", session_id, self)
}

func (client *XenAPIClient) VM_assert_can_migrate(session_id interface{}, vm interface{}, dest map[string]string, live bool, vdi_map interface{}, vif_map interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.assert_can_migrate", session_id, vm, dest, live, vdi_map, vif_map, options)
}

func (client *XenAPIClient) VM_migrate_send(session_id interface{}, vm interface{}, dest map[string]string, live bool, vdi_map interface{}, vif_map interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.migrate_send", session_id, vm, dest, live, vdi_map, vif_map, options)
}

func (client *XenAPIClient) VM_maximise_memory(session_id interface{}, self interface{}, total interface{}, approximate bool) (i interface{}, err error) {
	return client.RPCCall("VM.maximise_memory", session_id, self, total, approximate)
}

func (client *XenAPIClient) VM_send_trigger(session_id interface{}, vm interface{}, trigger string) (i interface{}, err error) {
	return client.RPCCall("VM.send_trigger", session_id, vm, trigger)
}

func (client *XenAPIClient) VM_send_sysrq(session_id interface{}, vm interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.send_sysrq", session_id, vm, key)
}

func (client *XenAPIClient) VM_set_VCPUs_at_startup(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_at_startup", session_id, self, value)
}

func (client *XenAPIClient) VM_set_VCPUs_max(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_max", session_id, self, value)
}

func (client *XenAPIClient) VM_set_shadow_multiplier_live(session_id interface{}, self interface{}, multiplier interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_shadow_multiplier_live", session_id, self, multiplier)
}

func (client *XenAPIClient) VM_set_HVM_shadow_multiplier(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_HVM_shadow_multiplier", session_id, self, value)
}

func (client *XenAPIClient) VM_get_cooperative(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_cooperative", session_id, self)
}

func (client *XenAPIClient) VM_wait_memory_target_live(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.wait_memory_target_live", session_id, self)
}

func (client *XenAPIClient) VM_set_memory_target_live(session_id interface{}, self interface{}, target interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_target_live", session_id, self, target)
}

func (client *XenAPIClient) VM_set_memory_limits(session_id interface{}, self interface{}, static_min interface{}, static_max interface{}, dynamic_min interface{}, dynamic_max interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_limits", session_id, self, static_min, static_max, dynamic_min, dynamic_max)
}

func (client *XenAPIClient) VM_set_memory_static_range(session_id interface{}, self interface{}, min interface{}, max interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_static_range", session_id, self, min, max)
}

func (client *XenAPIClient) VM_set_memory_static_min(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_static_min", session_id, self, value)
}

func (client *XenAPIClient) VM_set_memory_static_max(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_static_max", session_id, self, value)
}

func (client *XenAPIClient) VM_set_memory_dynamic_range(session_id interface{}, self interface{}, min interface{}, max interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_dynamic_range", session_id, self, min, max)
}

func (client *XenAPIClient) VM_set_memory_dynamic_min(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_dynamic_min", session_id, self, value)
}

func (client *XenAPIClient) VM_set_memory_dynamic_max(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_dynamic_max", session_id, self, value)
}

func (client *XenAPIClient) VM_compute_memory_overhead(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.compute_memory_overhead", session_id, vm)
}

func (client *XenAPIClient) VM_set_ha_always_run(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VM.set_ha_always_run", session_id, self, value)
}

func (client *XenAPIClient) VM_set_ha_restart_priority(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_ha_restart_priority", session_id, self, value)
}

func (client *XenAPIClient) VM_add_to_VCPUs_params_live(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_VCPUs_params_live", session_id, self, key, value)
}

func (client *XenAPIClient) VM_set_VCPUs_number_live(session_id interface{}, self interface{}, nvcpu interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_number_live", session_id, self, nvcpu)
}

func (client *XenAPIClient) VM_pool_migrate(session_id interface{}, vm interface{}, host interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.pool_migrate", session_id, vm, host, options)
}

func (client *XenAPIClient) VM_resume_on(session_id interface{}, vm interface{}, host interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.resume_on", session_id, vm, host, start_paused, force)
}

func (client *XenAPIClient) VM_resume(session_id interface{}, vm interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.resume", session_id, vm, start_paused, force)
}

func (client *XenAPIClient) VM_suspend(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.suspend", session_id, vm)
}

func (client *XenAPIClient) VM_hard_reboot(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.hard_reboot", session_id, vm)
}

func (client *XenAPIClient) VM_power_state_reset(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.power_state_reset", session_id, vm)
}

func (client *XenAPIClient) VM_hard_shutdown(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.hard_shutdown", session_id, vm)
}

func (client *XenAPIClient) VM_clean_reboot(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.clean_reboot", session_id, vm)
}

func (client *XenAPIClient) VM_shutdown(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.shutdown", session_id, vm)
}

func (client *XenAPIClient) VM_clean_shutdown(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.clean_shutdown", session_id, vm)
}

func (client *XenAPIClient) VM_unpause(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.unpause", session_id, vm)
}

func (client *XenAPIClient) VM_pause(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.pause", session_id, vm)
}

func (client *XenAPIClient) VM_start_on(session_id interface{}, vm interface{}, host interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.start_on", session_id, vm, host, start_paused, force)
}

func (client *XenAPIClient) VM_start(session_id interface{}, vm interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.start", session_id, vm, start_paused, force)
}

func (client *XenAPIClient) VM_provision(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.provision", session_id, vm)
}

func (client *XenAPIClient) VM_checkpoint(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.checkpoint", session_id, vm, new_name)
}

func (client *XenAPIClient) VM_revert(session_id interface{}, snapshot interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.revert", session_id, snapshot)
}

func (client *XenAPIClient) VM_copy(session_id interface{}, vm interface{}, new_name string, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.copy", session_id, vm, new_name, sr)
}

func (client *XenAPIClient) VM_clone(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.clone", session_id, vm, new_name)
}

func (client *XenAPIClient) VM_snapshot_with_quiesce(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.snapshot_with_quiesce", session_id, vm, new_name)
}

func (client *XenAPIClient) VM_snapshot(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.snapshot", session_id, vm, new_name)
}

func (client *XenAPIClient) VM_set_suspend_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_suspend_SR", session_id, self, value)
}

func (client *XenAPIClient) VM_remove_from_blocked_operations(session_id interface{}, self interface{}, key interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_blocked_operations", session_id, self, key)
}

func (client *XenAPIClient) VM_add_to_blocked_operations(session_id interface{}, self interface{}, key interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_blocked_operations", session_id, self, key, value)
}

func (client *XenAPIClient) VM_set_blocked_operations(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_blocked_operations", session_id, self, value)
}

func (client *XenAPIClient) VM_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_tags", session_id, self, value)
}

func (client *XenAPIClient) VM_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_tags", session_id, self, value)
}

func (client *XenAPIClient) VM_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_tags", session_id, self, value)
}

func (client *XenAPIClient) VM_remove_from_xenstore_data(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_xenstore_data", session_id, self, key)
}

func (client *XenAPIClient) VM_add_to_xenstore_data(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_xenstore_data", session_id, self, key, value)
}

func (client *XenAPIClient) VM_set_xenstore_data(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_xenstore_data", session_id, self, value)
}

func (client *XenAPIClient) VM_set_recommendations(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_recommendations", session_id, self, value)
}

func (client *XenAPIClient) VM_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VM_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VM_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VM_set_PCI_bus(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PCI_bus", session_id, self, value)
}

func (client *XenAPIClient) VM_remove_from_platform(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_platform", session_id, self, key)
}

func (client *XenAPIClient) VM_add_to_platform(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_platform", session_id, self, key, value)
}

func (client *XenAPIClient) VM_set_platform(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_platform", session_id, self, value)
}

func (client *XenAPIClient) VM_remove_from_HVM_boot_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_HVM_boot_params", session_id, self, key)
}

func (client *XenAPIClient) VM_add_to_HVM_boot_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_HVM_boot_params", session_id, self, key, value)
}

func (client *XenAPIClient) VM_set_HVM_boot_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_HVM_boot_params", session_id, self, value)
}

func (client *XenAPIClient) VM_set_HVM_boot_policy(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_HVM_boot_policy", session_id, self, value)
}

func (client *XenAPIClient) VM_set_PV_legacy_args(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_legacy_args", session_id, self, value)
}

func (client *XenAPIClient) VM_set_PV_bootloader_args(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_bootloader_args", session_id, self, value)
}

func (client *XenAPIClient) VM_set_PV_args(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_args", session_id, self, value)
}

func (client *XenAPIClient) VM_set_PV_ramdisk(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_ramdisk", session_id, self, value)
}

func (client *XenAPIClient) VM_set_PV_kernel(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_kernel", session_id, self, value)
}

func (client *XenAPIClient) VM_set_PV_bootloader(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_bootloader", session_id, self, value)
}

func (client *XenAPIClient) VM_set_actions_after_crash(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_actions_after_crash", session_id, self, value)
}

func (client *XenAPIClient) VM_set_actions_after_reboot(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_actions_after_reboot", session_id, self, value)
}

func (client *XenAPIClient) VM_set_actions_after_shutdown(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_actions_after_shutdown", session_id, self, value)
}

func (client *XenAPIClient) VM_remove_from_VCPUs_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_VCPUs_params", session_id, self, key)
}

func (client *XenAPIClient) VM_add_to_VCPUs_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_VCPUs_params", session_id, self, key, value)
}

func (client *XenAPIClient) VM_set_VCPUs_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_params", session_id, self, value)
}

func (client *XenAPIClient) VM_set_affinity(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_affinity", session_id, self, value)
}

func (client *XenAPIClient) VM_set_is_a_template(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VM.set_is_a_template", session_id, self, value)
}

func (client *XenAPIClient) VM_set_user_version(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_user_version", session_id, self, value)
}

func (client *XenAPIClient) VM_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) VM_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) VM_get_generation_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_generation_id", session_id, self)
}

func (client *XenAPIClient) VM_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_version", session_id, self)
}

func (client *XenAPIClient) VM_get_suspend_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_suspend_SR", session_id, self)
}

func (client *XenAPIClient) VM_get_attached_PCIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_attached_PCIs", session_id, self)
}

func (client *XenAPIClient) VM_get_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VGPUs", session_id, self)
}

func (client *XenAPIClient) VM_get_order(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_order", session_id, self)
}

func (client *XenAPIClient) VM_get_shutdown_delay(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_shutdown_delay", session_id, self)
}

func (client *XenAPIClient) VM_get_start_delay(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_start_delay", session_id, self)
}

func (client *XenAPIClient) VM_get_appliance(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_appliance", session_id, self)
}

func (client *XenAPIClient) VM_get_is_snapshot_from_vmpp(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_snapshot_from_vmpp", session_id, self)
}

func (client *XenAPIClient) VM_get_protection_policy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_protection_policy", session_id, self)
}

func (client *XenAPIClient) VM_get_bios_strings(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_bios_strings", session_id, self)
}

func (client *XenAPIClient) VM_get_children(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_children", session_id, self)
}

func (client *XenAPIClient) VM_get_parent(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_parent", session_id, self)
}

func (client *XenAPIClient) VM_get_snapshot_metadata(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_metadata", session_id, self)
}

func (client *XenAPIClient) VM_get_snapshot_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_info", session_id, self)
}

func (client *XenAPIClient) VM_get_blocked_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_blocked_operations", session_id, self)
}

func (client *XenAPIClient) VM_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_tags", session_id, self)
}

func (client *XenAPIClient) VM_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_blobs", session_id, self)
}

func (client *XenAPIClient) VM_get_transportable_snapshot_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_transportable_snapshot_id", session_id, self)
}

func (client *XenAPIClient) VM_get_snapshot_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_time", session_id, self)
}

func (client *XenAPIClient) VM_get_snapshots(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshots", session_id, self)
}

func (client *XenAPIClient) VM_get_snapshot_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_of", session_id, self)
}

func (client *XenAPIClient) VM_get_is_a_snapshot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_a_snapshot", session_id, self)
}

func (client *XenAPIClient) VM_get_ha_restart_priority(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_ha_restart_priority", session_id, self)
}

func (client *XenAPIClient) VM_get_ha_always_run(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_ha_always_run", session_id, self)
}

func (client *XenAPIClient) VM_get_xenstore_data(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_xenstore_data", session_id, self)
}

func (client *XenAPIClient) VM_get_recommendations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_recommendations", session_id, self)
}

func (client *XenAPIClient) VM_get_last_booted_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_last_booted_record", session_id, self)
}

func (client *XenAPIClient) VM_get_guest_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_guest_metrics", session_id, self)
}

func (client *XenAPIClient) VM_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_metrics", session_id, self)
}

func (client *XenAPIClient) VM_get_is_control_domain(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_control_domain", session_id, self)
}

func (client *XenAPIClient) VM_get_last_boot_CPU_flags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_last_boot_CPU_flags", session_id, self)
}

func (client *XenAPIClient) VM_get_domarch(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_domarch", session_id, self)
}

func (client *XenAPIClient) VM_get_domid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_domid", session_id, self)
}

func (client *XenAPIClient) VM_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_other_config", session_id, self)
}

func (client *XenAPIClient) VM_get_PCI_bus(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PCI_bus", session_id, self)
}

func (client *XenAPIClient) VM_get_platform(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_platform", session_id, self)
}

func (client *XenAPIClient) VM_get_HVM_shadow_multiplier(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_HVM_shadow_multiplier", session_id, self)
}

func (client *XenAPIClient) VM_get_HVM_boot_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_HVM_boot_params", session_id, self)
}

func (client *XenAPIClient) VM_get_HVM_boot_policy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_HVM_boot_policy", session_id, self)
}

func (client *XenAPIClient) VM_get_PV_legacy_args(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_legacy_args", session_id, self)
}

func (client *XenAPIClient) VM_get_PV_bootloader_args(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_bootloader_args", session_id, self)
}

func (client *XenAPIClient) VM_get_PV_args(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_args", session_id, self)
}

func (client *XenAPIClient) VM_get_PV_ramdisk(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_ramdisk", session_id, self)
}

func (client *XenAPIClient) VM_get_PV_kernel(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_kernel", session_id, self)
}

func (client *XenAPIClient) VM_get_PV_bootloader(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_bootloader", session_id, self)
}

func (client *XenAPIClient) VM_get_VTPMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VTPMs", session_id, self)
}

func (client *XenAPIClient) VM_get_crash_dumps(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_crash_dumps", session_id, self)
}

func (client *XenAPIClient) VM_get_VBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VBDs", session_id, self)
}

func (client *XenAPIClient) VM_get_VIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VIFs", session_id, self)
}

func (client *XenAPIClient) VM_get_consoles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_consoles", session_id, self)
}

func (client *XenAPIClient) VM_get_actions_after_crash(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_actions_after_crash", session_id, self)
}

func (client *XenAPIClient) VM_get_actions_after_reboot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_actions_after_reboot", session_id, self)
}

func (client *XenAPIClient) VM_get_actions_after_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_actions_after_shutdown", session_id, self)
}

func (client *XenAPIClient) VM_get_VCPUs_at_startup(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VCPUs_at_startup", session_id, self)
}

func (client *XenAPIClient) VM_get_VCPUs_max(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VCPUs_max", session_id, self)
}

func (client *XenAPIClient) VM_get_VCPUs_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VCPUs_params", session_id, self)
}

func (client *XenAPIClient) VM_get_memory_static_min(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_static_min", session_id, self)
}

func (client *XenAPIClient) VM_get_memory_dynamic_min(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_dynamic_min", session_id, self)
}

func (client *XenAPIClient) VM_get_memory_dynamic_max(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_dynamic_max", session_id, self)
}

func (client *XenAPIClient) VM_get_memory_static_max(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_static_max", session_id, self)
}

func (client *XenAPIClient) VM_get_memory_target(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_target", session_id, self)
}

func (client *XenAPIClient) VM_get_memory_overhead(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_overhead", session_id, self)
}

func (client *XenAPIClient) VM_get_affinity(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_affinity", session_id, self)
}

func (client *XenAPIClient) VM_get_resident_on(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_resident_on", session_id, self)
}

func (client *XenAPIClient) VM_get_suspend_VDI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_suspend_VDI", session_id, self)
}

func (client *XenAPIClient) VM_get_is_a_template(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_a_template", session_id, self)
}

func (client *XenAPIClient) VM_get_user_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_user_version", session_id, self)
}

func (client *XenAPIClient) VM_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_name_description", session_id, self)
}

func (client *XenAPIClient) VM_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_name_label", session_id, self)
}

func (client *XenAPIClient) VM_get_power_state(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_power_state", session_id, self)
}

func (client *XenAPIClient) VM_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_current_operations", session_id, self)
}

func (client *XenAPIClient) VM_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) VM_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_uuid", session_id, self)
}

func (client *XenAPIClient) VM_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VM.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) VM_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.destroy", session_id, self)
}

func (client *XenAPIClient) VM_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.create", session_id, args)
}

func (client *XenAPIClient) VM_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VM.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VM_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_record", session_id, self)
}

// The metrics associated with a VM

func (client *XenAPIClient) VM_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_all_records", session_id)
}

func (client *XenAPIClient) VM_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_all", session_id)
}

func (client *XenAPIClient) VM_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VM_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VM_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VM_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_other_config", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_last_updated", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_install_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_install_time", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_start_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_start_time", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_state(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_state", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_flags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_flags", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_params", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_CPU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_CPU", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_utilisation", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_VCPUs_number(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_number", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_memory_actual(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_memory_actual", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_uuid", session_id, self)
}

func (client *XenAPIClient) VM_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VM_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_record", session_id, self)
}

// The metrics reported by the guest (as opposed to inferred from outside)

func (client *XenAPIClient) VM_guest_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_all_records", session_id)
}

func (client *XenAPIClient) VM_guest_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_all", session_id)
}

func (client *XenAPIClient) VM_guest_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VM_guest_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VM_guest_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VM_guest_metrics_get_live(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_live", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_other_config", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_last_updated", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_other(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_other", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_networks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_networks", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_disks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_disks", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_memory(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_memory", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_PV_drivers_up_to_date(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_PV_drivers_up_to_date", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_PV_drivers_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_PV_drivers_version", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_os_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_os_version", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_uuid", session_id, self)
}

func (client *XenAPIClient) VM_guest_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VM_guest_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_record", session_id, self)
}

// VM Protection Policy

func (client *XenAPIClient) VMPP_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_all_records", session_id)
}

func (client *XenAPIClient) VMPP_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_all", session_id)
}

func (client *XenAPIClient) VMPP_set_archive_last_run_time(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_last_run_time", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_backup_last_run_time(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_last_run_time", session_id, self, value)
}

func (client *XenAPIClient) VMPP_remove_from_alarm_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_alarm_config", session_id, self, key)
}

func (client *XenAPIClient) VMPP_remove_from_archive_schedule(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_archive_schedule", session_id, self, key)
}

func (client *XenAPIClient) VMPP_remove_from_archive_target_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_archive_target_config", session_id, self, key)
}

func (client *XenAPIClient) VMPP_remove_from_backup_schedule(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_backup_schedule", session_id, self, key)
}

func (client *XenAPIClient) VMPP_add_to_alarm_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_alarm_config", session_id, self, key, value)
}

func (client *XenAPIClient) VMPP_add_to_archive_schedule(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_archive_schedule", session_id, self, key, value)
}

func (client *XenAPIClient) VMPP_add_to_archive_target_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_archive_target_config", session_id, self, key, value)
}

func (client *XenAPIClient) VMPP_add_to_backup_schedule(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_backup_schedule", session_id, self, key, value)
}

func (client *XenAPIClient) VMPP_set_alarm_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_alarm_config", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_is_alarm_enabled(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_is_alarm_enabled", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_archive_target_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_target_config", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_archive_target_type(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_target_type", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_archive_schedule(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_schedule", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_archive_frequency(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_frequency", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_backup_schedule(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_schedule", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_backup_frequency(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_frequency", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_backup_retention_value(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_retention_value", session_id, self, value)
}

func (client *XenAPIClient) VMPP_get_alerts(session_id interface{}, vmpp interface{}, hours_from_now interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_alerts", session_id, vmpp, hours_from_now)
}

func (client *XenAPIClient) VMPP_archive_now(session_id interface{}, snapshot interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.archive_now", session_id, snapshot)
}

func (client *XenAPIClient) VMPP_protect_now(session_id interface{}, vmpp interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.protect_now", session_id, vmpp)
}

func (client *XenAPIClient) VMPP_set_backup_type(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_type", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_is_policy_enabled(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_is_policy_enabled", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) VMPP_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) VMPP_get_recent_alerts(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_recent_alerts", session_id, self)
}

func (client *XenAPIClient) VMPP_get_alarm_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_alarm_config", session_id, self)
}

func (client *XenAPIClient) VMPP_get_is_alarm_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_alarm_enabled", session_id, self)
}

func (client *XenAPIClient) VMPP_get_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_VMs", session_id, self)
}

func (client *XenAPIClient) VMPP_get_archive_last_run_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_last_run_time", session_id, self)
}

func (client *XenAPIClient) VMPP_get_is_archive_running(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_archive_running", session_id, self)
}

func (client *XenAPIClient) VMPP_get_archive_schedule(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_schedule", session_id, self)
}

func (client *XenAPIClient) VMPP_get_archive_frequency(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_frequency", session_id, self)
}

func (client *XenAPIClient) VMPP_get_archive_target_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_target_config", session_id, self)
}

func (client *XenAPIClient) VMPP_get_archive_target_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_target_type", session_id, self)
}

func (client *XenAPIClient) VMPP_get_backup_last_run_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_last_run_time", session_id, self)
}

func (client *XenAPIClient) VMPP_get_is_backup_running(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_backup_running", session_id, self)
}

func (client *XenAPIClient) VMPP_get_backup_schedule(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_schedule", session_id, self)
}

func (client *XenAPIClient) VMPP_get_backup_frequency(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_frequency", session_id, self)
}

func (client *XenAPIClient) VMPP_get_backup_retention_value(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_retention_value", session_id, self)
}

func (client *XenAPIClient) VMPP_get_backup_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_type", session_id, self)
}

func (client *XenAPIClient) VMPP_get_is_policy_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_policy_enabled", session_id, self)
}

func (client *XenAPIClient) VMPP_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_name_description", session_id, self)
}

func (client *XenAPIClient) VMPP_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_name_label", session_id, self)
}

func (client *XenAPIClient) VMPP_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_uuid", session_id, self)
}

func (client *XenAPIClient) VMPP_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) VMPP_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.destroy", session_id, self)
}

func (client *XenAPIClient) VMPP_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.create", session_id, args)
}

func (client *XenAPIClient) VMPP_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VMPP_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_record", session_id, self)
}

// VM appliance

func (client *XenAPIClient) VMAppliance_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_all_records", session_id)
}

func (client *XenAPIClient) VMAppliance_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_all", session_id)
}

func (client *XenAPIClient) VMAppliance_recover(session_id interface{}, self interface{}, session_to interface{}, force bool) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.recover", session_id, self, session_to, force)
}

func (client *XenAPIClient) VMAppliance_get_SRs_required_for_recovery(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_SRs_required_for_recovery", session_id, self, session_to)
}

func (client *XenAPIClient) VMAppliance_assert_can_be_recovered(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.assert_can_be_recovered", session_id, self, session_to)
}

func (client *XenAPIClient) VMAppliance_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.shutdown", session_id, self)
}

func (client *XenAPIClient) VMAppliance_hard_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.hard_shutdown", session_id, self)
}

func (client *XenAPIClient) VMAppliance_clean_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.clean_shutdown", session_id, self)
}

func (client *XenAPIClient) VMAppliance_start(session_id interface{}, self interface{}, paused bool) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.start", session_id, self, paused)
}

func (client *XenAPIClient) VMAppliance_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) VMAppliance_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) VMAppliance_get_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_VMs", session_id, self)
}

func (client *XenAPIClient) VMAppliance_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_current_operations", session_id, self)
}

func (client *XenAPIClient) VMAppliance_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) VMAppliance_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_name_description", session_id, self)
}

func (client *XenAPIClient) VMAppliance_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_name_label", session_id, self)
}

func (client *XenAPIClient) VMAppliance_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_uuid", session_id, self)
}

func (client *XenAPIClient) VMAppliance_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) VMAppliance_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.destroy", session_id, self)
}

func (client *XenAPIClient) VMAppliance_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.create", session_id, args)
}

func (client *XenAPIClient) VMAppliance_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VMAppliance_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_record", session_id, self)
}

// DR task

func (client *XenAPIClient) DRTask_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_all_records", session_id)
}

func (client *XenAPIClient) DRTask_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_all", session_id)
}

func (client *XenAPIClient) DRTask_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.destroy", session_id, self)
}

func (client *XenAPIClient) DRTask_create(session_id interface{}, a_type string, device_config map[string]string, whitelist interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.create", session_id, a_type, device_config, whitelist)
}

func (client *XenAPIClient) DRTask_get_introduced_SRs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_introduced_SRs", session_id, self)
}

func (client *XenAPIClient) DRTask_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_uuid", session_id, self)
}

func (client *XenAPIClient) DRTask_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) DRTask_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_record", session_id, self)
}

// A physical host

func (client *XenAPIClient) host_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_all_records", session_id)
}

func (client *XenAPIClient) host_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_all", session_id)
}

func (client *XenAPIClient) host_disable_display(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.disable_display", session_id, host)
}

func (client *XenAPIClient) host_enable_display(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.enable_display", session_id, host)
}

func (client *XenAPIClient) host_declare_dead(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.declare_dead", session_id, host)
}

func (client *XenAPIClient) host_migrate_receive(session_id interface{}, host interface{}, network interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.migrate_receive", session_id, host, network, options)
}

func (client *XenAPIClient) host_disable_local_storage_caching(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.disable_local_storage_caching", session_id, host)
}

func (client *XenAPIClient) host_enable_local_storage_caching(session_id interface{}, host interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("host.enable_local_storage_caching", session_id, host, sr)
}

func (client *XenAPIClient) host_reset_cpu_features(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.reset_cpu_features", session_id, host)
}

func (client *XenAPIClient) host_set_cpu_features(session_id interface{}, host interface{}, features string) (i interface{}, err error) {
	return client.RPCCall("host.set_cpu_features", session_id, host, features)
}

func (client *XenAPIClient) host_set_power_on_mode(session_id interface{}, self interface{}, power_on_mode string, power_on_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_power_on_mode", session_id, self, power_on_mode, power_on_config)
}

func (client *XenAPIClient) host_refresh_pack_info(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.refresh_pack_info", session_id, host)
}

func (client *XenAPIClient) host_apply_edition(session_id interface{}, host interface{}, edition string, force bool) (i interface{}, err error) {
	return client.RPCCall("host.apply_edition", session_id, host, edition, force)
}

func (client *XenAPIClient) host_get_server_certificate(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_server_certificate", session_id, host)
}

func (client *XenAPIClient) host_retrieve_wlb_evacuate_recommendations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.retrieve_wlb_evacuate_recommendations", session_id, self)
}

func (client *XenAPIClient) host_disable_external_auth(session_id interface{}, host interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.disable_external_auth", session_id, host, config)
}

func (client *XenAPIClient) host_enable_external_auth(session_id interface{}, host interface{}, config map[string]string, service_name string, auth_type string) (i interface{}, err error) {
	return client.RPCCall("host.enable_external_auth", session_id, host, config, service_name, auth_type)
}

func (client *XenAPIClient) host_get_server_localtime(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_server_localtime", session_id, host)
}

func (client *XenAPIClient) host_get_servertime(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_servertime", session_id, host)
}

func (client *XenAPIClient) host_call_plugin(session_id interface{}, host interface{}, plugin string, fn string, args map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.call_plugin", session_id, host, plugin, fn, args)
}

func (client *XenAPIClient) host_create_new_blob(session_id interface{}, host interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("host.create_new_blob", session_id, host, name, mime_type, public)
}

func (client *XenAPIClient) host_backup_rrds(session_id interface{}, host interface{}, delay interface{}) (i interface{}, err error) {
	return client.RPCCall("host.backup_rrds", session_id, host, delay)
}

func (client *XenAPIClient) host_sync_data(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.sync_data", session_id, host)
}

func (client *XenAPIClient) host_compute_memory_overhead(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.compute_memory_overhead", session_id, host)
}

func (client *XenAPIClient) host_compute_free_memory(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.compute_free_memory", session_id, host)
}

func (client *XenAPIClient) host_set_hostname_live(session_id interface{}, host interface{}, hostname string) (i interface{}, err error) {
	return client.RPCCall("host.set_hostname_live", session_id, host, hostname)
}

func (client *XenAPIClient) host_shutdown_agent(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.shutdown_agent", session_id)
}

func (client *XenAPIClient) host_restart_agent(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.restart_agent", session_id, host)
}

func (client *XenAPIClient) host_get_system_status_capabilities(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_system_status_capabilities", session_id, host)
}

func (client *XenAPIClient) host_get_management_interface(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_management_interface", session_id, host)
}

func (client *XenAPIClient) host_management_disable(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.management_disable", session_id)
}

func (client *XenAPIClient) host_local_management_reconfigure(session_id interface{}, a_interface string) (i interface{}, err error) {
	return client.RPCCall("host.local_management_reconfigure", session_id, a_interface)
}

func (client *XenAPIClient) host_management_reconfigure(session_id interface{}, pif interface{}) (i interface{}, err error) {
	return client.RPCCall("host.management_reconfigure", session_id, pif)
}

func (client *XenAPIClient) host_syslog_reconfigure(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.syslog_reconfigure", session_id, host)
}

func (client *XenAPIClient) host_evacuate(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.evacuate", session_id, host)
}

func (client *XenAPIClient) host_get_uncooperative_resident_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_uncooperative_resident_VMs", session_id, self)
}

func (client *XenAPIClient) host_get_vms_which_prevent_evacuation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_vms_which_prevent_evacuation", session_id, self)
}

func (client *XenAPIClient) host_assert_can_evacuate(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.assert_can_evacuate", session_id, host)
}

func (client *XenAPIClient) host_forget_data_source_archives(session_id interface{}, host interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("host.forget_data_source_archives", session_id, host, data_source)
}

func (client *XenAPIClient) host_query_data_source(session_id interface{}, host interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("host.query_data_source", session_id, host, data_source)
}

func (client *XenAPIClient) host_record_data_source(session_id interface{}, host interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("host.record_data_source", session_id, host, data_source)
}

func (client *XenAPIClient) host_get_data_sources(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_data_sources", session_id, host)
}

func (client *XenAPIClient) host_emergency_ha_disable(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.emergency_ha_disable", session_id)
}

func (client *XenAPIClient) host_power_on(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.power_on", session_id, host)
}

func (client *XenAPIClient) host_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.destroy", session_id, self)
}

func (client *XenAPIClient) host_license_apply(session_id interface{}, host interface{}, contents string) (i interface{}, err error) {
	return client.RPCCall("host.license_apply", session_id, host, contents)
}

func (client *XenAPIClient) host_list_methods(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.list_methods", session_id)
}

func (client *XenAPIClient) host_bugreport_upload(session_id interface{}, host interface{}, url string, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.bugreport_upload", session_id, host, url, options)
}

func (client *XenAPIClient) host_send_debug_keys(session_id interface{}, host interface{}, keys string) (i interface{}, err error) {
	return client.RPCCall("host.send_debug_keys", session_id, host, keys)
}

func (client *XenAPIClient) host_get_log(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_log", session_id, host)
}

func (client *XenAPIClient) host_dmesg_clear(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.dmesg_clear", session_id, host)
}

func (client *XenAPIClient) host_dmesg(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.dmesg", session_id, host)
}

func (client *XenAPIClient) host_reboot(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.reboot", session_id, host)
}

func (client *XenAPIClient) host_shutdown(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.shutdown", session_id, host)
}

func (client *XenAPIClient) host_enable(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.enable", session_id, host)
}

func (client *XenAPIClient) host_disable(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.disable", session_id, host)
}

func (client *XenAPIClient) host_set_display(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_display", session_id, self, value)
}

func (client *XenAPIClient) host_remove_from_guest_VCPUs_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_guest_VCPUs_params", session_id, self, key)
}

func (client *XenAPIClient) host_add_to_guest_VCPUs_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_guest_VCPUs_params", session_id, self, key, value)
}

func (client *XenAPIClient) host_set_guest_VCPUs_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_guest_VCPUs_params", session_id, self, value)
}

func (client *XenAPIClient) host_remove_from_license_server(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_license_server", session_id, self, key)
}

func (client *XenAPIClient) host_add_to_license_server(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_license_server", session_id, self, key, value)
}

func (client *XenAPIClient) host_set_license_server(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_license_server", session_id, self, value)
}

func (client *XenAPIClient) host_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.remove_tags", session_id, self, value)
}

func (client *XenAPIClient) host_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_tags", session_id, self, value)
}

func (client *XenAPIClient) host_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_tags", session_id, self, value)
}

func (client *XenAPIClient) host_set_address(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_address", session_id, self, value)
}

func (client *XenAPIClient) host_set_hostname(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_hostname", session_id, self, value)
}

func (client *XenAPIClient) host_set_crash_dump_sr(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_crash_dump_sr", session_id, self, value)
}

func (client *XenAPIClient) host_set_suspend_image_sr(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_suspend_image_sr", session_id, self, value)
}

func (client *XenAPIClient) host_remove_from_logging(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_logging", session_id, self, key)
}

func (client *XenAPIClient) host_add_to_logging(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_logging", session_id, self, key, value)
}

func (client *XenAPIClient) host_set_logging(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_logging", session_id, self, value)
}

func (client *XenAPIClient) host_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) host_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) host_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) host_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) host_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) host_get_display(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_display", session_id, self)
}

func (client *XenAPIClient) host_get_guest_VCPUs_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_guest_VCPUs_params", session_id, self)
}

func (client *XenAPIClient) host_get_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PGPUs", session_id, self)
}

func (client *XenAPIClient) host_get_PCIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PCIs", session_id, self)
}

func (client *XenAPIClient) host_get_chipset_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_chipset_info", session_id, self)
}

func (client *XenAPIClient) host_get_local_cache_sr(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_local_cache_sr", session_id, self)
}

func (client *XenAPIClient) host_get_power_on_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_power_on_config", session_id, self)
}

func (client *XenAPIClient) host_get_power_on_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_power_on_mode", session_id, self)
}

func (client *XenAPIClient) host_get_bios_strings(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_bios_strings", session_id, self)
}

func (client *XenAPIClient) host_get_license_server(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_license_server", session_id, self)
}

func (client *XenAPIClient) host_get_edition(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_edition", session_id, self)
}

func (client *XenAPIClient) host_get_external_auth_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_external_auth_configuration", session_id, self)
}

func (client *XenAPIClient) host_get_external_auth_service_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_external_auth_service_name", session_id, self)
}

func (client *XenAPIClient) host_get_external_auth_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_external_auth_type", session_id, self)
}

func (client *XenAPIClient) host_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_tags", session_id, self)
}

func (client *XenAPIClient) host_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_blobs", session_id, self)
}

func (client *XenAPIClient) host_get_ha_network_peers(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_ha_network_peers", session_id, self)
}

func (client *XenAPIClient) host_get_ha_statefiles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_ha_statefiles", session_id, self)
}

func (client *XenAPIClient) host_get_license_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_license_params", session_id, self)
}

func (client *XenAPIClient) host_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_metrics", session_id, self)
}

func (client *XenAPIClient) host_get_address(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_address", session_id, self)
}

func (client *XenAPIClient) host_get_hostname(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_hostname", session_id, self)
}

func (client *XenAPIClient) host_get_cpu_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_cpu_info", session_id, self)
}

func (client *XenAPIClient) host_get_host_CPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_host_CPUs", session_id, self)
}

func (client *XenAPIClient) host_get_PBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PBDs", session_id, self)
}

func (client *XenAPIClient) host_get_patches(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_patches", session_id, self)
}

func (client *XenAPIClient) host_get_crashdumps(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_crashdumps", session_id, self)
}

func (client *XenAPIClient) host_get_crash_dump_sr(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_crash_dump_sr", session_id, self)
}

func (client *XenAPIClient) host_get_suspend_image_sr(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_suspend_image_sr", session_id, self)
}

func (client *XenAPIClient) host_get_PIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PIFs", session_id, self)
}

func (client *XenAPIClient) host_get_logging(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_logging", session_id, self)
}

func (client *XenAPIClient) host_get_resident_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_resident_VMs", session_id, self)
}

func (client *XenAPIClient) host_get_supported_bootloaders(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_supported_bootloaders", session_id, self)
}

func (client *XenAPIClient) host_get_sched_policy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_sched_policy", session_id, self)
}

func (client *XenAPIClient) host_get_cpu_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_cpu_configuration", session_id, self)
}

func (client *XenAPIClient) host_get_capabilities(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_capabilities", session_id, self)
}

func (client *XenAPIClient) host_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_other_config", session_id, self)
}

func (client *XenAPIClient) host_get_software_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_software_version", session_id, self)
}

func (client *XenAPIClient) host_get_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_enabled", session_id, self)
}

func (client *XenAPIClient) host_get_API_version_vendor_implementation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_vendor_implementation", session_id, self)
}

func (client *XenAPIClient) host_get_API_version_vendor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_vendor", session_id, self)
}

func (client *XenAPIClient) host_get_API_version_minor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_minor", session_id, self)
}

func (client *XenAPIClient) host_get_API_version_major(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_major", session_id, self)
}

func (client *XenAPIClient) host_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_current_operations", session_id, self)
}

func (client *XenAPIClient) host_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) host_get_memory_overhead(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_memory_overhead", session_id, self)
}

func (client *XenAPIClient) host_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_name_description", session_id, self)
}

func (client *XenAPIClient) host_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_name_label", session_id, self)
}

func (client *XenAPIClient) host_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_uuid", session_id, self)
}

func (client *XenAPIClient) host_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("host.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) host_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("host.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) host_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_record", session_id, self)
}

// Represents a host crash dump

func (client *XenAPIClient) hostCrashdump_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_all_records", session_id)
}

func (client *XenAPIClient) hostCrashdump_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_all", session_id)
}

func (client *XenAPIClient) hostCrashdump_upload(session_id interface{}, self interface{}, url string, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.upload", session_id, self, url, options)
}

func (client *XenAPIClient) hostCrashdump_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.destroy", session_id, self)
}

func (client *XenAPIClient) hostCrashdump_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) hostCrashdump_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) hostCrashdump_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) hostCrashdump_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_other_config", session_id, self)
}

func (client *XenAPIClient) hostCrashdump_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_size", session_id, self)
}

func (client *XenAPIClient) hostCrashdump_get_timestamp(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_timestamp", session_id, self)
}

func (client *XenAPIClient) hostCrashdump_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_host", session_id, self)
}

func (client *XenAPIClient) hostCrashdump_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_uuid", session_id, self)
}

func (client *XenAPIClient) hostCrashdump_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) hostCrashdump_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_record", session_id, self)
}

// Represents a patch stored on a server

func (client *XenAPIClient) hostPatch_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_all_records", session_id)
}

func (client *XenAPIClient) hostPatch_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_all", session_id)
}

func (client *XenAPIClient) hostPatch_apply(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.apply", session_id, self)
}

func (client *XenAPIClient) hostPatch_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.destroy", session_id, self)
}

func (client *XenAPIClient) hostPatch_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) hostPatch_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) hostPatch_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) hostPatch_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_other_config", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_pool_patch(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_pool_patch", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_size", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_timestamp_applied(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_timestamp_applied", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_applied(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_applied", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_host", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_version", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_name_description", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_name_label", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_uuid", session_id, self)
}

func (client *XenAPIClient) hostPatch_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) hostPatch_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) hostPatch_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_record", session_id, self)
}

// The metrics associated with a host

func (client *XenAPIClient) host_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_all_records", session_id)
}

func (client *XenAPIClient) host_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_all", session_id)
}

func (client *XenAPIClient) host_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) host_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) host_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) host_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_other_config", session_id, self)
}

func (client *XenAPIClient) host_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_last_updated", session_id, self)
}

func (client *XenAPIClient) host_metrics_get_live(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_live", session_id, self)
}

func (client *XenAPIClient) host_metrics_get_memory_free(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_memory_free", session_id, self)
}

func (client *XenAPIClient) host_metrics_get_memory_total(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_memory_total", session_id, self)
}

func (client *XenAPIClient) host_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_uuid", session_id, self)
}

func (client *XenAPIClient) host_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) host_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_record", session_id, self)
}

// A physical CPU

func (client *XenAPIClient) hostCpu_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_all_records", session_id)
}

func (client *XenAPIClient) hostCpu_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_all", session_id)
}

func (client *XenAPIClient) hostCpu_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) hostCpu_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) hostCpu_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) hostCpu_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_other_config", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_utilisation", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_features(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_features", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_flags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_flags", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_stepping(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_stepping", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_model(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_model", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_family(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_family", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_modelname(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_modelname", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_speed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_speed", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_vendor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_vendor", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_number(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_number", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_host", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_uuid", session_id, self)
}

func (client *XenAPIClient) hostCpu_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) hostCpu_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_record", session_id, self)
}

// A virtual network

func (client *XenAPIClient) network_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_all_records", session_id)
}

func (client *XenAPIClient) network_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_all", session_id)
}

func (client *XenAPIClient) network_set_default_locking_mode(session_id interface{}, network interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("network.set_default_locking_mode", session_id, network, value)
}

func (client *XenAPIClient) network_create_new_blob(session_id interface{}, network interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("network.create_new_blob", session_id, network, name, mime_type, public)
}

func (client *XenAPIClient) network_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.remove_tags", session_id, self, value)
}

func (client *XenAPIClient) network_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.add_tags", session_id, self, value)
}

func (client *XenAPIClient) network_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("network.set_tags", session_id, self, value)
}

func (client *XenAPIClient) network_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("network.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) network_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("network.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) network_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("network.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) network_set_MTU(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("network.set_MTU", session_id, self, value)
}

func (client *XenAPIClient) network_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) network_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) network_get_assigned_ips(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_assigned_ips", session_id, self)
}

func (client *XenAPIClient) network_get_default_locking_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_default_locking_mode", session_id, self)
}

func (client *XenAPIClient) network_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_tags", session_id, self)
}

func (client *XenAPIClient) network_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_blobs", session_id, self)
}

func (client *XenAPIClient) network_get_bridge(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_bridge", session_id, self)
}

func (client *XenAPIClient) network_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_other_config", session_id, self)
}

func (client *XenAPIClient) network_get_MTU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_MTU", session_id, self)
}

func (client *XenAPIClient) network_get_PIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_PIFs", session_id, self)
}

func (client *XenAPIClient) network_get_VIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_VIFs", session_id, self)
}

func (client *XenAPIClient) network_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_current_operations", session_id, self)
}

func (client *XenAPIClient) network_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) network_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_name_description", session_id, self)
}

func (client *XenAPIClient) network_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_name_label", session_id, self)
}

func (client *XenAPIClient) network_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_uuid", session_id, self)
}

func (client *XenAPIClient) network_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("network.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) network_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.destroy", session_id, self)
}

func (client *XenAPIClient) network_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("network.create", session_id, args)
}

func (client *XenAPIClient) network_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("network.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) network_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_record", session_id, self)
}

// A virtual network interface

func (client *XenAPIClient) VIF_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_all_records", session_id)
}

func (client *XenAPIClient) VIF_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_all", session_id)
}

func (client *XenAPIClient) VIF_remove_ipv6_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_ipv6_allowed", session_id, self, value)
}

func (client *XenAPIClient) VIF_add_ipv6_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_ipv6_allowed", session_id, self, value)
}

func (client *XenAPIClient) VIF_set_ipv6_allowed(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.set_ipv6_allowed", session_id, self, value)
}

func (client *XenAPIClient) VIF_remove_ipv4_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_ipv4_allowed", session_id, self, value)
}

func (client *XenAPIClient) VIF_add_ipv4_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_ipv4_allowed", session_id, self, value)
}

func (client *XenAPIClient) VIF_set_ipv4_allowed(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.set_ipv4_allowed", session_id, self, value)
}

func (client *XenAPIClient) VIF_set_locking_mode(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.set_locking_mode", session_id, self, value)
}

func (client *XenAPIClient) VIF_unplug_force(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.unplug_force", session_id, self)
}

func (client *XenAPIClient) VIF_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.unplug", session_id, self)
}

func (client *XenAPIClient) VIF_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.plug", session_id, self)
}

func (client *XenAPIClient) VIF_remove_from_qos_algorithm_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_from_qos_algorithm_params", session_id, self, key)
}

func (client *XenAPIClient) VIF_add_to_qos_algorithm_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_to_qos_algorithm_params", session_id, self, key, value)
}

func (client *XenAPIClient) VIF_set_qos_algorithm_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VIF.set_qos_algorithm_params", session_id, self, value)
}

func (client *XenAPIClient) VIF_set_qos_algorithm_type(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.set_qos_algorithm_type", session_id, self, value)
}

func (client *XenAPIClient) VIF_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VIF_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VIF_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VIF.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VIF_get_ipv6_allowed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_ipv6_allowed", session_id, self)
}

func (client *XenAPIClient) VIF_get_ipv4_allowed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_ipv4_allowed", session_id, self)
}

func (client *XenAPIClient) VIF_get_locking_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_locking_mode", session_id, self)
}

func (client *XenAPIClient) VIF_get_MAC_autogenerated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_MAC_autogenerated", session_id, self)
}

func (client *XenAPIClient) VIF_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_metrics", session_id, self)
}

func (client *XenAPIClient) VIF_get_qos_supported_algorithms(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_qos_supported_algorithms", session_id, self)
}

func (client *XenAPIClient) VIF_get_qos_algorithm_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_qos_algorithm_params", session_id, self)
}

func (client *XenAPIClient) VIF_get_qos_algorithm_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_qos_algorithm_type", session_id, self)
}

func (client *XenAPIClient) VIF_get_runtime_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_runtime_properties", session_id, self)
}

func (client *XenAPIClient) VIF_get_status_detail(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_status_detail", session_id, self)
}

func (client *XenAPIClient) VIF_get_status_code(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_status_code", session_id, self)
}

func (client *XenAPIClient) VIF_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_currently_attached", session_id, self)
}

func (client *XenAPIClient) VIF_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_other_config", session_id, self)
}

func (client *XenAPIClient) VIF_get_MTU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_MTU", session_id, self)
}

func (client *XenAPIClient) VIF_get_MAC(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_MAC", session_id, self)
}

func (client *XenAPIClient) VIF_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_VM", session_id, self)
}

func (client *XenAPIClient) VIF_get_network(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_network", session_id, self)
}

func (client *XenAPIClient) VIF_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_device", session_id, self)
}

func (client *XenAPIClient) VIF_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_current_operations", session_id, self)
}

func (client *XenAPIClient) VIF_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) VIF_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_uuid", session_id, self)
}

func (client *XenAPIClient) VIF_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.destroy", session_id, self)
}

func (client *XenAPIClient) VIF_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.create", session_id, args)
}

func (client *XenAPIClient) VIF_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VIF.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VIF_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_record", session_id, self)
}

// The metrics associated with a virtual network device

func (client *XenAPIClient) VIF_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_all_records", session_id)
}

func (client *XenAPIClient) VIF_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_all", session_id)
}

func (client *XenAPIClient) VIF_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VIF_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VIF_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VIF_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_other_config", session_id, self)
}

func (client *XenAPIClient) VIF_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_last_updated", session_id, self)
}

func (client *XenAPIClient) VIF_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_io_write_kbs", session_id, self)
}

func (client *XenAPIClient) VIF_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_io_read_kbs", session_id, self)
}

func (client *XenAPIClient) VIF_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_uuid", session_id, self)
}

func (client *XenAPIClient) VIF_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VIF_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_record", session_id, self)
}

// A physical network interface (note separate VLANs are represented as several PIFs)

func (client *XenAPIClient) PIF_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_all_records", session_id)
}

func (client *XenAPIClient) PIF_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_all", session_id)
}

func (client *XenAPIClient) PIF_set_property(session_id interface{}, self interface{}, name string, value string) (i interface{}, err error) {
	return client.RPCCall("PIF.set_property", session_id, self, name, value)
}

func (client *XenAPIClient) PIF_db_forget(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.db_forget", session_id, self)
}

func (client *XenAPIClient) PIF_db_introduce(session_id interface{}, device string, network interface{}, host interface{}, MAC string, MTU interface{}, VLAN interface{}, physical bool, ip_configuration_mode interface{}, IP string, netmask string, gateway string, DNS string, bond_slave_of interface{}, VLAN_master_of interface{}, management bool, other_config map[string]string, disallow_unplug bool, ipv6_configuration_mode interface{}, IPv6 interface{}, ipv6_gateway string, primary_address_type interface{}, managed bool, properties map[string]string) (i interface{}, err error) {
	return client.RPCCall("PIF.db_introduce", session_id, device, network, host, MAC, MTU, VLAN, physical, ip_configuration_mode, IP, netmask, gateway, DNS, bond_slave_of, VLAN_master_of, management, other_config, disallow_unplug, ipv6_configuration_mode, IPv6, ipv6_gateway, primary_address_type, managed, properties)
}

func (client *XenAPIClient) PIF_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.plug", session_id, self)
}

func (client *XenAPIClient) PIF_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.unplug", session_id, self)
}

func (client *XenAPIClient) PIF_forget(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.forget", session_id, self)
}

func (client *XenAPIClient) PIF_introduce(session_id interface{}, host interface{}, MAC string, device string, managed bool) (i interface{}, err error) {
	return client.RPCCall("PIF.introduce", session_id, host, MAC, device, managed)
}

func (client *XenAPIClient) PIF_scan(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.scan", session_id, host)
}

func (client *XenAPIClient) PIF_set_primary_address_type(session_id interface{}, self interface{}, primary_address_type interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.set_primary_address_type", session_id, self, primary_address_type)
}

func (client *XenAPIClient) PIF_reconfigure_ipv6(session_id interface{}, self interface{}, mode interface{}, IPv6 string, gateway string, DNS string) (i interface{}, err error) {
	return client.RPCCall("PIF.reconfigure_ipv6", session_id, self, mode, IPv6, gateway, DNS)
}

func (client *XenAPIClient) PIF_reconfigure_ip(session_id interface{}, self interface{}, mode interface{}, IP string, netmask string, gateway string, DNS string) (i interface{}, err error) {
	return client.RPCCall("PIF.reconfigure_ip", session_id, self, mode, IP, netmask, gateway, DNS)
}

func (client *XenAPIClient) PIF_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.destroy", session_id, self)
}

func (client *XenAPIClient) PIF_create_VLAN(session_id interface{}, device string, network interface{}, host interface{}, VLAN interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.create_VLAN", session_id, device, network, host, VLAN)
}

func (client *XenAPIClient) PIF_set_disallow_unplug(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("PIF.set_disallow_unplug", session_id, self, value)
}

func (client *XenAPIClient) PIF_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PIF.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) PIF_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PIF.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) PIF_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PIF.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) PIF_get_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_properties", session_id, self)
}

func (client *XenAPIClient) PIF_get_managed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_managed", session_id, self)
}

func (client *XenAPIClient) PIF_get_primary_address_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_primary_address_type", session_id, self)
}

func (client *XenAPIClient) PIF_get_ipv6_gateway(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_ipv6_gateway", session_id, self)
}

func (client *XenAPIClient) PIF_get_IPv6(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_IPv6", session_id, self)
}

func (client *XenAPIClient) PIF_get_ipv6_configuration_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_ipv6_configuration_mode", session_id, self)
}

func (client *XenAPIClient) PIF_get_tunnel_transport_PIF_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_tunnel_transport_PIF_of", session_id, self)
}

func (client *XenAPIClient) PIF_get_tunnel_access_PIF_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_tunnel_access_PIF_of", session_id, self)
}

func (client *XenAPIClient) PIF_get_disallow_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_disallow_unplug", session_id, self)
}

func (client *XenAPIClient) PIF_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_other_config", session_id, self)
}

func (client *XenAPIClient) PIF_get_management(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_management", session_id, self)
}

func (client *XenAPIClient) PIF_get_VLAN_slave_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_VLAN_slave_of", session_id, self)
}

func (client *XenAPIClient) PIF_get_VLAN_master_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_VLAN_master_of", session_id, self)
}

func (client *XenAPIClient) PIF_get_bond_master_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_bond_master_of", session_id, self)
}

func (client *XenAPIClient) PIF_get_bond_slave_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_bond_slave_of", session_id, self)
}

func (client *XenAPIClient) PIF_get_DNS(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_DNS", session_id, self)
}

func (client *XenAPIClient) PIF_get_gateway(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_gateway", session_id, self)
}

func (client *XenAPIClient) PIF_get_netmask(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_netmask", session_id, self)
}

func (client *XenAPIClient) PIF_get_IP(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_IP", session_id, self)
}

func (client *XenAPIClient) PIF_get_ip_configuration_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_ip_configuration_mode", session_id, self)
}

func (client *XenAPIClient) PIF_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_currently_attached", session_id, self)
}

func (client *XenAPIClient) PIF_get_physical(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_physical", session_id, self)
}

func (client *XenAPIClient) PIF_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_metrics", session_id, self)
}

func (client *XenAPIClient) PIF_get_VLAN(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_VLAN", session_id, self)
}

func (client *XenAPIClient) PIF_get_MTU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_MTU", session_id, self)
}

func (client *XenAPIClient) PIF_get_MAC(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_MAC", session_id, self)
}

func (client *XenAPIClient) PIF_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_host", session_id, self)
}

func (client *XenAPIClient) PIF_get_network(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_network", session_id, self)
}

func (client *XenAPIClient) PIF_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_device", session_id, self)
}

func (client *XenAPIClient) PIF_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_uuid", session_id, self)
}

func (client *XenAPIClient) PIF_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PIF.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) PIF_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_record", session_id, self)
}

// The metrics associated with a physical network interface

func (client *XenAPIClient) PIF_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_all_records", session_id)
}

func (client *XenAPIClient) PIF_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_all", session_id)
}

func (client *XenAPIClient) PIF_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) PIF_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) PIF_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) PIF_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_other_config", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_last_updated", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_pci_bus_path(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_pci_bus_path", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_duplex(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_duplex", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_speed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_speed", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_device_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_device_name", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_device_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_device_id", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_vendor_name", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_vendor_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_vendor_id", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_carrier(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_carrier", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_io_write_kbs", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_io_read_kbs", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_uuid", session_id, self)
}

func (client *XenAPIClient) PIF_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) PIF_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_record", session_id, self)
}

//

func (client *XenAPIClient) Bond_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_all_records", session_id)
}

func (client *XenAPIClient) Bond_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_all", session_id)
}

func (client *XenAPIClient) Bond_set_property(session_id interface{}, self interface{}, name string, value string) (i interface{}, err error) {
	return client.RPCCall("Bond.set_property", session_id, self, name, value)
}

func (client *XenAPIClient) Bond_set_mode(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.set_mode", session_id, self, value)
}

func (client *XenAPIClient) Bond_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.destroy", session_id, self)
}

func (client *XenAPIClient) Bond_create(session_id interface{}, network interface{}, members interface{}, MAC string, mode interface{}, properties map[string]string) (i interface{}, err error) {
	return client.RPCCall("Bond.create", session_id, network, members, MAC, mode, properties)
}

func (client *XenAPIClient) Bond_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("Bond.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) Bond_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("Bond.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) Bond_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("Bond.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) Bond_get_links_up(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_links_up", session_id, self)
}

func (client *XenAPIClient) Bond_get_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_properties", session_id, self)
}

func (client *XenAPIClient) Bond_get_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_mode", session_id, self)
}

func (client *XenAPIClient) Bond_get_primary_slave(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_primary_slave", session_id, self)
}

func (client *XenAPIClient) Bond_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_other_config", session_id, self)
}

func (client *XenAPIClient) Bond_get_slaves(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_slaves", session_id, self)
}

func (client *XenAPIClient) Bond_get_master(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_master", session_id, self)
}

func (client *XenAPIClient) Bond_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_uuid", session_id, self)
}

func (client *XenAPIClient) Bond_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("Bond.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) Bond_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_record", session_id, self)
}

// A VLAN mux/demux

func (client *XenAPIClient) VLAN_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_all_records", session_id)
}

func (client *XenAPIClient) VLAN_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_all", session_id)
}

func (client *XenAPIClient) VLAN_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.destroy", session_id, self)
}

func (client *XenAPIClient) VLAN_create(session_id interface{}, tagged_PIF interface{}, tag interface{}, network interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.create", session_id, tagged_PIF, tag, network)
}

func (client *XenAPIClient) VLAN_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VLAN.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VLAN_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VLAN.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VLAN_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VLAN.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VLAN_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_other_config", session_id, self)
}

func (client *XenAPIClient) VLAN_get_tag(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_tag", session_id, self)
}

func (client *XenAPIClient) VLAN_get_untagged_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_untagged_PIF", session_id, self)
}

func (client *XenAPIClient) VLAN_get_tagged_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_tagged_PIF", session_id, self)
}

func (client *XenAPIClient) VLAN_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_uuid", session_id, self)
}

func (client *XenAPIClient) VLAN_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VLAN_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_record", session_id, self)
}

// A storage manager plugin

func (client *XenAPIClient) SM_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_all_records", session_id)
}

func (client *XenAPIClient) SM_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_all", session_id)
}

func (client *XenAPIClient) SM_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("SM.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) SM_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("SM.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) SM_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("SM.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) SM_get_driver_filename(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_driver_filename", session_id, self)
}

func (client *XenAPIClient) SM_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_other_config", session_id, self)
}

func (client *XenAPIClient) SM_get_features(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_features", session_id, self)
}

func (client *XenAPIClient) SM_get_capabilities(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_capabilities", session_id, self)
}

func (client *XenAPIClient) SM_get_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_configuration", session_id, self)
}

func (client *XenAPIClient) SM_get_required_api_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_required_api_version", session_id, self)
}

func (client *XenAPIClient) SM_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_version", session_id, self)
}

func (client *XenAPIClient) SM_get_copyright(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_copyright", session_id, self)
}

func (client *XenAPIClient) SM_get_vendor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_vendor", session_id, self)
}

func (client *XenAPIClient) SM_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_type", session_id, self)
}

func (client *XenAPIClient) SM_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_name_description", session_id, self)
}

func (client *XenAPIClient) SM_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_name_label", session_id, self)
}

func (client *XenAPIClient) SM_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_uuid", session_id, self)
}

func (client *XenAPIClient) SM_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("SM.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) SM_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("SM.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) SM_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_record", session_id, self)
}

// A storage repository

func (client *XenAPIClient) SR_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_all_records", session_id)
}

func (client *XenAPIClient) SR_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_all", session_id)
}

func (client *XenAPIClient) SR_disable_database_replication(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.disable_database_replication", session_id, sr)
}

func (client *XenAPIClient) SR_enable_database_replication(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.enable_database_replication", session_id, sr)
}

func (client *XenAPIClient) SR_assert_supports_database_replication(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.assert_supports_database_replication", session_id, sr)
}

func (client *XenAPIClient) SR_assert_can_host_ha_statefile(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.assert_can_host_ha_statefile", session_id, sr)
}

func (client *XenAPIClient) SR_set_physical_utilisation(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_physical_utilisation", session_id, self, value)
}

func (client *XenAPIClient) SR_set_virtual_allocation(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_virtual_allocation", session_id, self, value)
}

func (client *XenAPIClient) SR_set_physical_size(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_physical_size", session_id, self, value)
}

func (client *XenAPIClient) SR_create_new_blob(session_id interface{}, sr interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("SR.create_new_blob", session_id, sr, name, mime_type, public)
}

func (client *XenAPIClient) SR_set_name_description(session_id interface{}, sr interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.set_name_description", session_id, sr, value)
}

func (client *XenAPIClient) SR_set_name_label(session_id interface{}, sr interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.set_name_label", session_id, sr, value)
}

func (client *XenAPIClient) SR_set_shared(session_id interface{}, sr interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("SR.set_shared", session_id, sr, value)
}

func (client *XenAPIClient) SR_probe(session_id interface{}, host interface{}, device_config map[string]string, a_type string, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.probe", session_id, host, device_config, a_type, sm_config)
}

func (client *XenAPIClient) SR_scan(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.scan", session_id, sr)
}

func (client *XenAPIClient) SR_get_supported_types(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_supported_types", session_id)
}

func (client *XenAPIClient) SR_update(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.update", session_id, sr)
}

func (client *XenAPIClient) SR_forget(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.forget", session_id, sr)
}

func (client *XenAPIClient) SR_destroy(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.destroy", session_id, sr)
}

func (client *XenAPIClient) SR_make(session_id interface{}, host interface{}, device_config map[string]string, physical_size interface{}, name_label string, name_description string, a_type string, content_type string, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.make", session_id, host, device_config, physical_size, name_label, name_description, a_type, content_type, sm_config)
}

func (client *XenAPIClient) SR_introduce(session_id interface{}, uuid string, name_label string, name_description string, a_type string, content_type string, shared bool, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.introduce", session_id, uuid, name_label, name_description, a_type, content_type, shared, sm_config)
}

func (client *XenAPIClient) SR_create(session_id interface{}, host interface{}, device_config map[string]string, physical_size interface{}, name_label string, name_description string, a_type string, content_type string, shared bool, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.create", session_id, host, device_config, physical_size, name_label, name_description, a_type, content_type, shared, sm_config)
}

func (client *XenAPIClient) SR_remove_from_sm_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("SR.remove_from_sm_config", session_id, self, key)
}

func (client *XenAPIClient) SR_add_to_sm_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("SR.add_to_sm_config", session_id, self, key, value)
}

func (client *XenAPIClient) SR_set_sm_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.set_sm_config", session_id, self, value)
}

func (client *XenAPIClient) SR_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.remove_tags", session_id, self, value)
}

func (client *XenAPIClient) SR_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.add_tags", session_id, self, value)
}

func (client *XenAPIClient) SR_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_tags", session_id, self, value)
}

func (client *XenAPIClient) SR_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("SR.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) SR_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("SR.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) SR_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) SR_get_introduced_by(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_introduced_by", session_id, self)
}

func (client *XenAPIClient) SR_get_local_cache_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_local_cache_enabled", session_id, self)
}

func (client *XenAPIClient) SR_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_blobs", session_id, self)
}

func (client *XenAPIClient) SR_get_sm_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_sm_config", session_id, self)
}

func (client *XenAPIClient) SR_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_tags", session_id, self)
}

func (client *XenAPIClient) SR_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_other_config", session_id, self)
}

func (client *XenAPIClient) SR_get_shared(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_shared", session_id, self)
}

func (client *XenAPIClient) SR_get_content_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_content_type", session_id, self)
}

func (client *XenAPIClient) SR_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_type", session_id, self)
}

func (client *XenAPIClient) SR_get_physical_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_physical_size", session_id, self)
}

func (client *XenAPIClient) SR_get_physical_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_physical_utilisation", session_id, self)
}

func (client *XenAPIClient) SR_get_virtual_allocation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_virtual_allocation", session_id, self)
}

func (client *XenAPIClient) SR_get_PBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_PBDs", session_id, self)
}

func (client *XenAPIClient) SR_get_VDIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_VDIs", session_id, self)
}

func (client *XenAPIClient) SR_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_current_operations", session_id, self)
}

func (client *XenAPIClient) SR_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) SR_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_name_description", session_id, self)
}

func (client *XenAPIClient) SR_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_name_label", session_id, self)
}

func (client *XenAPIClient) SR_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_uuid", session_id, self)
}

func (client *XenAPIClient) SR_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("SR.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) SR_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("SR.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) SR_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_record", session_id, self)
}

// A virtual disk image

func (client *XenAPIClient) VDI_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_all_records", session_id)
}

func (client *XenAPIClient) VDI_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_all", session_id)
}

func (client *XenAPIClient) VDI_pool_migrate(session_id interface{}, vdi interface{}, sr interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.pool_migrate", session_id, vdi, sr, options)
}

func (client *XenAPIClient) VDI_read_database_pool_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.read_database_pool_uuid", session_id, self)
}

func (client *XenAPIClient) VDI_open_database(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.open_database", session_id, self)
}

func (client *XenAPIClient) VDI_set_allow_caching(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_allow_caching", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_on_boot(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_on_boot", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_metadata_of_pool(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_metadata_of_pool", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_snapshot_time(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_snapshot_time", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_snapshot_of(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_snapshot_of", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_is_a_snapshot(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_is_a_snapshot", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_physical_utilisation(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_physical_utilisation", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_virtual_size(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_virtual_size", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_missing(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_missing", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_read_only(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_read_only", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_sharable(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_sharable", session_id, self, value)
}

func (client *XenAPIClient) VDI_forget(session_id interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.forget", session_id, vdi)
}

func (client *XenAPIClient) VDI_set_managed(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_managed", session_id, self, value)
}

func (client *XenAPIClient) VDI_copy(session_id interface{}, vdi interface{}, sr interface{}, base_vdi interface{}, into_vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.copy", session_id, vdi, sr, base_vdi, into_vdi)
}

func (client *XenAPIClient) VDI_update(session_id interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.update", session_id, vdi)
}

func (client *XenAPIClient) VDI_db_forget(session_id interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.db_forget", session_id, vdi)
}

func (client *XenAPIClient) VDI_db_introduce(session_id interface{}, uuid string, name_label string, name_description string, SR interface{}, a_type interface{}, sharable bool, read_only bool, other_config map[string]string, location string, xenstore_data map[string]string, sm_config map[string]string, managed bool, virtual_size interface{}, physical_utilisation interface{}, metadata_of_pool interface{}, is_a_snapshot bool, snapshot_time interface{}, snapshot_of interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.db_introduce", session_id, uuid, name_label, name_description, SR, a_type, sharable, read_only, other_config, location, xenstore_data, sm_config, managed, virtual_size, physical_utilisation, metadata_of_pool, is_a_snapshot, snapshot_time, snapshot_of)
}

func (client *XenAPIClient) VDI_introduce(session_id interface{}, uuid string, name_label string, name_description string, SR interface{}, a_type interface{}, sharable bool, read_only bool, other_config map[string]string, location string, xenstore_data map[string]string, sm_config map[string]string, managed bool, virtual_size interface{}, physical_utilisation interface{}, metadata_of_pool interface{}, is_a_snapshot bool, snapshot_time interface{}, snapshot_of interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.introduce", session_id, uuid, name_label, name_description, SR, a_type, sharable, read_only, other_config, location, xenstore_data, sm_config, managed, virtual_size, physical_utilisation, metadata_of_pool, is_a_snapshot, snapshot_time, snapshot_of)
}

func (client *XenAPIClient) VDI_resize_online(session_id interface{}, vdi interface{}, size interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.resize_online", session_id, vdi, size)
}

func (client *XenAPIClient) VDI_resize(session_id interface{}, vdi interface{}, size interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.resize", session_id, vdi, size)
}

func (client *XenAPIClient) VDI_clone(session_id interface{}, vdi interface{}, driver_params map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.clone", session_id, vdi, driver_params)
}

func (client *XenAPIClient) VDI_snapshot(session_id interface{}, vdi interface{}, driver_params map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.snapshot", session_id, vdi, driver_params)
}

func (client *XenAPIClient) VDI_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_tags", session_id, self, value)
}

func (client *XenAPIClient) VDI_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_tags", session_id, self, value)
}

func (client *XenAPIClient) VDI_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_tags", session_id, self, value)
}

func (client *XenAPIClient) VDI_remove_from_sm_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_from_sm_config", session_id, self, key)
}

func (client *XenAPIClient) VDI_add_to_sm_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_to_sm_config", session_id, self, key, value)
}

func (client *XenAPIClient) VDI_set_sm_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_sm_config", session_id, self, value)
}

func (client *XenAPIClient) VDI_remove_from_xenstore_data(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_from_xenstore_data", session_id, self, key)
}

func (client *XenAPIClient) VDI_add_to_xenstore_data(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_to_xenstore_data", session_id, self, key, value)
}

func (client *XenAPIClient) VDI_set_xenstore_data(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_xenstore_data", session_id, self, value)
}

func (client *XenAPIClient) VDI_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VDI_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VDI_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VDI_get_metadata_latest(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_metadata_latest", session_id, self)
}

func (client *XenAPIClient) VDI_get_metadata_of_pool(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_metadata_of_pool", session_id, self)
}

func (client *XenAPIClient) VDI_get_on_boot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_on_boot", session_id, self)
}

func (client *XenAPIClient) VDI_get_allow_caching(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_allow_caching", session_id, self)
}

func (client *XenAPIClient) VDI_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_tags", session_id, self)
}

func (client *XenAPIClient) VDI_get_snapshot_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_snapshot_time", session_id, self)
}

func (client *XenAPIClient) VDI_get_snapshots(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_snapshots", session_id, self)
}

func (client *XenAPIClient) VDI_get_snapshot_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_snapshot_of", session_id, self)
}

func (client *XenAPIClient) VDI_get_is_a_snapshot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_is_a_snapshot", session_id, self)
}

func (client *XenAPIClient) VDI_get_sm_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_sm_config", session_id, self)
}

func (client *XenAPIClient) VDI_get_xenstore_data(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_xenstore_data", session_id, self)
}

func (client *XenAPIClient) VDI_get_parent(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_parent", session_id, self)
}

func (client *XenAPIClient) VDI_get_missing(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_missing", session_id, self)
}

func (client *XenAPIClient) VDI_get_managed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_managed", session_id, self)
}

func (client *XenAPIClient) VDI_get_location(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_location", session_id, self)
}

func (client *XenAPIClient) VDI_get_storage_lock(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_storage_lock", session_id, self)
}

func (client *XenAPIClient) VDI_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_other_config", session_id, self)
}

func (client *XenAPIClient) VDI_get_read_only(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_read_only", session_id, self)
}

func (client *XenAPIClient) VDI_get_sharable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_sharable", session_id, self)
}

func (client *XenAPIClient) VDI_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_type", session_id, self)
}

func (client *XenAPIClient) VDI_get_physical_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_physical_utilisation", session_id, self)
}

func (client *XenAPIClient) VDI_get_virtual_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_virtual_size", session_id, self)
}

func (client *XenAPIClient) VDI_get_crash_dumps(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_crash_dumps", session_id, self)
}

func (client *XenAPIClient) VDI_get_VBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_VBDs", session_id, self)
}

func (client *XenAPIClient) VDI_get_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_SR", session_id, self)
}

func (client *XenAPIClient) VDI_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_current_operations", session_id, self)
}

func (client *XenAPIClient) VDI_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) VDI_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_name_description", session_id, self)
}

func (client *XenAPIClient) VDI_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_name_label", session_id, self)
}

func (client *XenAPIClient) VDI_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_uuid", session_id, self)
}

func (client *XenAPIClient) VDI_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VDI.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) VDI_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.destroy", session_id, self)
}

func (client *XenAPIClient) VDI_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.create", session_id, args)
}

func (client *XenAPIClient) VDI_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VDI.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VDI_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_record", session_id, self)
}

// A virtual block device

func (client *XenAPIClient) VBD_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_all_records", session_id)
}

func (client *XenAPIClient) VBD_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_all", session_id)
}

func (client *XenAPIClient) VBD_assert_attachable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.assert_attachable", session_id, self)
}

func (client *XenAPIClient) VBD_unplug_force(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.unplug_force", session_id, self)
}

func (client *XenAPIClient) VBD_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.unplug", session_id, self)
}

func (client *XenAPIClient) VBD_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.plug", session_id, self)
}

func (client *XenAPIClient) VBD_insert(session_id interface{}, vbd interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.insert", session_id, vbd, vdi)
}

func (client *XenAPIClient) VBD_eject(session_id interface{}, vbd interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.eject", session_id, vbd)
}

func (client *XenAPIClient) VBD_remove_from_qos_algorithm_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VBD.remove_from_qos_algorithm_params", session_id, self, key)
}

func (client *XenAPIClient) VBD_add_to_qos_algorithm_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.add_to_qos_algorithm_params", session_id, self, key, value)
}

func (client *XenAPIClient) VBD_set_qos_algorithm_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_qos_algorithm_params", session_id, self, value)
}

func (client *XenAPIClient) VBD_set_qos_algorithm_type(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_qos_algorithm_type", session_id, self, value)
}

func (client *XenAPIClient) VBD_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VBD.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VBD_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VBD_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VBD_set_unpluggable(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VBD.set_unpluggable", session_id, self, value)
}

func (client *XenAPIClient) VBD_set_type(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.set_type", session_id, self, value)
}

func (client *XenAPIClient) VBD_set_mode(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.set_mode", session_id, self, value)
}

func (client *XenAPIClient) VBD_set_bootable(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VBD.set_bootable", session_id, self, value)
}

func (client *XenAPIClient) VBD_set_userdevice(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_userdevice", session_id, self, value)
}

func (client *XenAPIClient) VBD_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_metrics", session_id, self)
}

func (client *XenAPIClient) VBD_get_qos_supported_algorithms(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_qos_supported_algorithms", session_id, self)
}

func (client *XenAPIClient) VBD_get_qos_algorithm_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_qos_algorithm_params", session_id, self)
}

func (client *XenAPIClient) VBD_get_qos_algorithm_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_qos_algorithm_type", session_id, self)
}

func (client *XenAPIClient) VBD_get_runtime_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_runtime_properties", session_id, self)
}

func (client *XenAPIClient) VBD_get_status_detail(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_status_detail", session_id, self)
}

func (client *XenAPIClient) VBD_get_status_code(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_status_code", session_id, self)
}

func (client *XenAPIClient) VBD_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_currently_attached", session_id, self)
}

func (client *XenAPIClient) VBD_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_other_config", session_id, self)
}

func (client *XenAPIClient) VBD_get_empty(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_empty", session_id, self)
}

func (client *XenAPIClient) VBD_get_storage_lock(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_storage_lock", session_id, self)
}

func (client *XenAPIClient) VBD_get_unpluggable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_unpluggable", session_id, self)
}

func (client *XenAPIClient) VBD_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_type", session_id, self)
}

func (client *XenAPIClient) VBD_get_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_mode", session_id, self)
}

func (client *XenAPIClient) VBD_get_bootable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_bootable", session_id, self)
}

func (client *XenAPIClient) VBD_get_userdevice(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_userdevice", session_id, self)
}

func (client *XenAPIClient) VBD_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_device", session_id, self)
}

func (client *XenAPIClient) VBD_get_VDI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_VDI", session_id, self)
}

func (client *XenAPIClient) VBD_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_VM", session_id, self)
}

func (client *XenAPIClient) VBD_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_current_operations", session_id, self)
}

func (client *XenAPIClient) VBD_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_allowed_operations", session_id, self)
}

func (client *XenAPIClient) VBD_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_uuid", session_id, self)
}

func (client *XenAPIClient) VBD_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.destroy", session_id, self)
}

func (client *XenAPIClient) VBD_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.create", session_id, args)
}

func (client *XenAPIClient) VBD_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VBD.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VBD_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_record", session_id, self)
}

// The metrics associated with a virtual block device

func (client *XenAPIClient) VBD_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_all_records", session_id)
}

func (client *XenAPIClient) VBD_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_all", session_id)
}

func (client *XenAPIClient) VBD_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VBD_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VBD_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VBD_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_other_config", session_id, self)
}

func (client *XenAPIClient) VBD_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_last_updated", session_id, self)
}

func (client *XenAPIClient) VBD_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_io_write_kbs", session_id, self)
}

func (client *XenAPIClient) VBD_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_io_read_kbs", session_id, self)
}

func (client *XenAPIClient) VBD_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_uuid", session_id, self)
}

func (client *XenAPIClient) VBD_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VBD_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_record", session_id, self)
}

// The physical block devices through which hosts access SRs

func (client *XenAPIClient) PBD_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_all_records", session_id)
}

func (client *XenAPIClient) PBD_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_all", session_id)
}

func (client *XenAPIClient) PBD_set_device_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PBD.set_device_config", session_id, self, value)
}

func (client *XenAPIClient) PBD_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.unplug", session_id, self)
}

func (client *XenAPIClient) PBD_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.plug", session_id, self)
}

func (client *XenAPIClient) PBD_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PBD.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) PBD_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PBD.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) PBD_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PBD.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) PBD_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_other_config", session_id, self)
}

func (client *XenAPIClient) PBD_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_currently_attached", session_id, self)
}

func (client *XenAPIClient) PBD_get_device_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_device_config", session_id, self)
}

func (client *XenAPIClient) PBD_get_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_SR", session_id, self)
}

func (client *XenAPIClient) PBD_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_host", session_id, self)
}

func (client *XenAPIClient) PBD_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_uuid", session_id, self)
}

func (client *XenAPIClient) PBD_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.destroy", session_id, self)
}

func (client *XenAPIClient) PBD_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.create", session_id, args)
}

func (client *XenAPIClient) PBD_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PBD.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) PBD_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_record", session_id, self)
}

// A VM crashdump

func (client *XenAPIClient) crashdump_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_all_records", session_id)
}

func (client *XenAPIClient) crashdump_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_all", session_id)
}

func (client *XenAPIClient) crashdump_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.destroy", session_id, self)
}

func (client *XenAPIClient) crashdump_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("crashdump.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) crashdump_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("crashdump.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) crashdump_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("crashdump.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) crashdump_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_other_config", session_id, self)
}

func (client *XenAPIClient) crashdump_get_VDI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_VDI", session_id, self)
}

func (client *XenAPIClient) crashdump_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_VM", session_id, self)
}

func (client *XenAPIClient) crashdump_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_uuid", session_id, self)
}

func (client *XenAPIClient) crashdump_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) crashdump_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_record", session_id, self)
}

// A virtual TPM device

func (client *XenAPIClient) VTPM_get_backend(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_backend", session_id, self)
}

func (client *XenAPIClient) VTPM_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_VM", session_id, self)
}

func (client *XenAPIClient) VTPM_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_uuid", session_id, self)
}

func (client *XenAPIClient) VTPM_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.destroy", session_id, self)
}

func (client *XenAPIClient) VTPM_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.create", session_id, args)
}

func (client *XenAPIClient) VTPM_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VTPM_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_record", session_id, self)
}

// A console

func (client *XenAPIClient) console_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_all_records", session_id)
}

func (client *XenAPIClient) console_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_all", session_id)
}

func (client *XenAPIClient) console_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("console.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) console_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("console.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) console_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("console.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) console_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_other_config", session_id, self)
}

func (client *XenAPIClient) console_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_VM", session_id, self)
}

func (client *XenAPIClient) console_get_location(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_location", session_id, self)
}

func (client *XenAPIClient) console_get_protocol(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_protocol", session_id, self)
}

func (client *XenAPIClient) console_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_uuid", session_id, self)
}

func (client *XenAPIClient) console_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.destroy", session_id, self)
}

func (client *XenAPIClient) console_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("console.create", session_id, args)
}

func (client *XenAPIClient) console_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("console.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) console_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_record", session_id, self)
}

// A user of the system

func (client *XenAPIClient) user_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("user.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) user_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("user.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) user_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("user.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) user_set_fullname(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("user.set_fullname", session_id, self, value)
}

func (client *XenAPIClient) user_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_other_config", session_id, self)
}

func (client *XenAPIClient) user_get_fullname(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_fullname", session_id, self)
}

func (client *XenAPIClient) user_get_short_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_short_name", session_id, self)
}

func (client *XenAPIClient) user_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_uuid", session_id, self)
}

func (client *XenAPIClient) user_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.destroy", session_id, self)
}

func (client *XenAPIClient) user_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("user.create", session_id, args)
}

func (client *XenAPIClient) user_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("user.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) user_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_record", session_id, self)
}

// Data sources for logging in RRDs

// A placeholder for a binary blob

func (client *XenAPIClient) blob_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_all_records", session_id)
}

func (client *XenAPIClient) blob_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_all", session_id)
}

func (client *XenAPIClient) blob_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.destroy", session_id, self)
}

func (client *XenAPIClient) blob_create(session_id interface{}, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("blob.create", session_id, mime_type, public)
}

func (client *XenAPIClient) blob_set_public(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("blob.set_public", session_id, self, value)
}

func (client *XenAPIClient) blob_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("blob.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) blob_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("blob.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) blob_get_mime_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_mime_type", session_id, self)
}

func (client *XenAPIClient) blob_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_last_updated", session_id, self)
}

func (client *XenAPIClient) blob_get_public(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_public", session_id, self)
}

func (client *XenAPIClient) blob_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_size", session_id, self)
}

func (client *XenAPIClient) blob_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_name_description", session_id, self)
}

func (client *XenAPIClient) blob_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_name_label", session_id, self)
}

func (client *XenAPIClient) blob_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_uuid", session_id, self)
}

func (client *XenAPIClient) blob_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("blob.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) blob_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("blob.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) blob_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_record", session_id, self)
}

// An message for the attention of the administrator

func (client *XenAPIClient) message_get_all_records_where(session_id interface{}, expr string) (i interface{}, err error) {
	return client.RPCCall("message.get_all_records_where", session_id, expr)
}

func (client *XenAPIClient) message_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_all_records", session_id)
}

func (client *XenAPIClient) message_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("message.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) message_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_record", session_id, self)
}

func (client *XenAPIClient) message_get_since(session_id interface{}, since interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_since", session_id, since)
}

func (client *XenAPIClient) message_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_all", session_id)
}

func (client *XenAPIClient) message_get(session_id interface{}, cls interface{}, obj_uuid string, since interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get", session_id, cls, obj_uuid, since)
}

func (client *XenAPIClient) message_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("message.destroy", session_id, self)
}

func (client *XenAPIClient) message_create(session_id interface{}, name string, priority interface{}, cls interface{}, obj_uuid string, body string) (i interface{}, err error) {
	return client.RPCCall("message.create", session_id, name, priority, cls, obj_uuid, body)
}

// A secret

func (client *XenAPIClient) secret_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_all_records", session_id)
}

func (client *XenAPIClient) secret_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_all", session_id)
}

func (client *XenAPIClient) secret_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("secret.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) secret_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("secret.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) secret_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("secret.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) secret_set_value(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("secret.set_value", session_id, self, value)
}

func (client *XenAPIClient) secret_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_other_config", session_id, self)
}

func (client *XenAPIClient) secret_get_value(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_value", session_id, self)
}

func (client *XenAPIClient) secret_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_uuid", session_id, self)
}

func (client *XenAPIClient) secret_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.destroy", session_id, self)
}

func (client *XenAPIClient) secret_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.create", session_id, args)
}

func (client *XenAPIClient) secret_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("secret.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) secret_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_record", session_id, self)
}

// A tunnel for network traffic

func (client *XenAPIClient) tunnel_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_all_records", session_id)
}

func (client *XenAPIClient) tunnel_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_all", session_id)
}

func (client *XenAPIClient) tunnel_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.destroy", session_id, self)
}

func (client *XenAPIClient) tunnel_create(session_id interface{}, transport_PIF interface{}, network interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.create", session_id, transport_PIF, network)
}

func (client *XenAPIClient) tunnel_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("tunnel.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) tunnel_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("tunnel.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) tunnel_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("tunnel.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) tunnel_remove_from_status(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("tunnel.remove_from_status", session_id, self, key)
}

func (client *XenAPIClient) tunnel_add_to_status(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("tunnel.add_to_status", session_id, self, key, value)
}

func (client *XenAPIClient) tunnel_set_status(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("tunnel.set_status", session_id, self, value)
}

func (client *XenAPIClient) tunnel_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_other_config", session_id, self)
}

func (client *XenAPIClient) tunnel_get_status(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_status", session_id, self)
}

func (client *XenAPIClient) tunnel_get_transport_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_transport_PIF", session_id, self)
}

func (client *XenAPIClient) tunnel_get_access_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_access_PIF", session_id, self)
}

func (client *XenAPIClient) tunnel_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_uuid", session_id, self)
}

func (client *XenAPIClient) tunnel_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) tunnel_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_record", session_id, self)
}

// A PCI device

func (client *XenAPIClient) PCI_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_all_records", session_id)
}

func (client *XenAPIClient) PCI_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_all", session_id)
}

func (client *XenAPIClient) PCI_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PCI.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) PCI_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PCI.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) PCI_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PCI.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) PCI_get_subsystem_device_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_subsystem_device_name", session_id, self)
}

func (client *XenAPIClient) PCI_get_subsystem_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_subsystem_vendor_name", session_id, self)
}

func (client *XenAPIClient) PCI_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_other_config", session_id, self)
}

func (client *XenAPIClient) PCI_get_dependencies(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_dependencies", session_id, self)
}

func (client *XenAPIClient) PCI_get_pci_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_pci_id", session_id, self)
}

func (client *XenAPIClient) PCI_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_host", session_id, self)
}

func (client *XenAPIClient) PCI_get_device_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_device_name", session_id, self)
}

func (client *XenAPIClient) PCI_get_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_vendor_name", session_id, self)
}

func (client *XenAPIClient) PCI_get_class_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_class_name", session_id, self)
}

func (client *XenAPIClient) PCI_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_uuid", session_id, self)
}

func (client *XenAPIClient) PCI_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PCI.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) PCI_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_record", session_id, self)
}

// A physical GPU (pGPU)

func (client *XenAPIClient) PGPU_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_all_records", session_id)
}

func (client *XenAPIClient) PGPU_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_all", session_id)
}

func (client *XenAPIClient) PGPU_disable_dom0_access(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.disable_dom0_access", session_id, self)
}

func (client *XenAPIClient) PGPU_enable_dom0_access(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.enable_dom0_access", session_id, self)
}

func (client *XenAPIClient) PGPU_get_remaining_capacity(session_id interface{}, self interface{}, vgpu_type interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_remaining_capacity", session_id, self, vgpu_type)
}

func (client *XenAPIClient) PGPU_set_GPU_group(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.set_GPU_group", session_id, self, value)
}

func (client *XenAPIClient) PGPU_set_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.set_enabled_VGPU_types", session_id, self, value)
}

func (client *XenAPIClient) PGPU_remove_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.remove_enabled_VGPU_types", session_id, self, value)
}

func (client *XenAPIClient) PGPU_add_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.add_enabled_VGPU_types", session_id, self, value)
}

func (client *XenAPIClient) PGPU_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PGPU.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) PGPU_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PGPU.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) PGPU_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PGPU.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) PGPU_get_is_system_display_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_is_system_display_device", session_id, self)
}

func (client *XenAPIClient) PGPU_get_dom0_access(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_dom0_access", session_id, self)
}

func (client *XenAPIClient) PGPU_get_supported_VGPU_max_capacities(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_supported_VGPU_max_capacities", session_id, self)
}

func (client *XenAPIClient) PGPU_get_resident_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_resident_VGPUs", session_id, self)
}

func (client *XenAPIClient) PGPU_get_enabled_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_enabled_VGPU_types", session_id, self)
}

func (client *XenAPIClient) PGPU_get_supported_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_supported_VGPU_types", session_id, self)
}

func (client *XenAPIClient) PGPU_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_other_config", session_id, self)
}

func (client *XenAPIClient) PGPU_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_host", session_id, self)
}

func (client *XenAPIClient) PGPU_get_GPU_group(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_GPU_group", session_id, self)
}

func (client *XenAPIClient) PGPU_get_PCI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_PCI", session_id, self)
}

func (client *XenAPIClient) PGPU_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_uuid", session_id, self)
}

func (client *XenAPIClient) PGPU_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) PGPU_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_record", session_id, self)
}

// A group of compatible GPUs across the resource pool

func (client *XenAPIClient) GPU_group_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_all_records", session_id)
}

func (client *XenAPIClient) GPU_group_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_all", session_id)
}

func (client *XenAPIClient) GPU_group_get_remaining_capacity(session_id interface{}, self interface{}, vgpu_type interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_remaining_capacity", session_id, self, vgpu_type)
}

func (client *XenAPIClient) GPU_group_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.destroy", session_id, self)
}

func (client *XenAPIClient) GPU_group_create(session_id interface{}, name_label string, name_description string, other_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.create", session_id, name_label, name_description, other_config)
}

func (client *XenAPIClient) GPU_group_set_allocation_algorithm(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_allocation_algorithm", session_id, self, value)
}

func (client *XenAPIClient) GPU_group_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) GPU_group_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) GPU_group_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) GPU_group_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_name_description", session_id, self, value)
}

func (client *XenAPIClient) GPU_group_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_name_label", session_id, self, value)
}

func (client *XenAPIClient) GPU_group_get_enabled_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_enabled_VGPU_types", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_supported_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_supported_VGPU_types", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_allocation_algorithm(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_allocation_algorithm", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_other_config", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_GPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_GPU_types", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_VGPUs", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_PGPUs", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_name_description", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_name_label", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_uuid", session_id, self)
}

func (client *XenAPIClient) GPU_group_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_by_name_label", session_id, label)
}

func (client *XenAPIClient) GPU_group_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) GPU_group_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_record", session_id, self)
}

// A virtual GPU (vGPU)

func (client *XenAPIClient) VGPU_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_all_records", session_id)
}

func (client *XenAPIClient) VGPU_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_all", session_id)
}

func (client *XenAPIClient) VGPU_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.destroy", session_id, self)
}

func (client *XenAPIClient) VGPU_create(session_id interface{}, VM interface{}, GPU_group interface{}, device string, other_config map[string]string, a_type interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.create", session_id, VM, GPU_group, device, other_config, a_type)
}

func (client *XenAPIClient) VGPU_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VGPU.remove_from_other_config", session_id, self, key)
}

func (client *XenAPIClient) VGPU_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VGPU.add_to_other_config", session_id, self, key, value)
}

func (client *XenAPIClient) VGPU_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VGPU.set_other_config", session_id, self, value)
}

func (client *XenAPIClient) VGPU_get_resident_on(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_resident_on", session_id, self)
}

func (client *XenAPIClient) VGPU_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_type", session_id, self)
}

func (client *XenAPIClient) VGPU_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_other_config", session_id, self)
}

func (client *XenAPIClient) VGPU_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_currently_attached", session_id, self)
}

func (client *XenAPIClient) VGPU_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_device", session_id, self)
}

func (client *XenAPIClient) VGPU_get_GPU_group(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_GPU_group", session_id, self)
}

func (client *XenAPIClient) VGPU_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_VM", session_id, self)
}

func (client *XenAPIClient) VGPU_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_uuid", session_id, self)
}

func (client *XenAPIClient) VGPU_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VGPU_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_record", session_id, self)
}

// A type of virtual GPU

func (client *XenAPIClient) VGPUType_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_all_records", session_id)
}

func (client *XenAPIClient) VGPUType_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_all", session_id)
}

func (client *XenAPIClient) VGPUType_get_enabled_on_GPU_groups(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_enabled_on_GPU_groups", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_supported_on_GPU_groups(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_supported_on_GPU_groups", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_VGPUs", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_enabled_on_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_enabled_on_PGPUs", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_supported_on_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_supported_on_PGPUs", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_max_resolution_y(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_max_resolution_y", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_max_resolution_x(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_max_resolution_x", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_max_heads(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_max_heads", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_framebuffer_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_framebuffer_size", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_model_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_model_name", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_vendor_name", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_uuid", session_id, self)
}

func (client *XenAPIClient) VGPUType_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_by_uuid", session_id, uuid)
}

func (client *XenAPIClient) VGPUType_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_record", session_id, self)
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
