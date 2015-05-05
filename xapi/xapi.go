package main

import (
	"fmt"
	"github.com/nilshell/xmlrpc"
	"os"
)

// session_logout_subject_identifier
//
// Log out all sessions associated to a user subject-identifier, except the session associated with the context calling this function
//
// params:
// - session_id, session ref, Reference to a valid session
// - subject_identifier, string, User subject-identifier of the sessions to be destroyed
//
// returns:
// - void
func (client *XenAPIClient) session_logout_subject_identifier(session_id interface{}, subject_identifier string) (i interface{}, err error) {
	return client.RPCCall("session.logout_subject_identifier", session_id, subject_identifier)
}

// session_get_all_subject_identifiers
//
// Return a list of all the user subject-identifiers of all existing sessions
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - string set
// - The list of user subject-identifiers of all existing sessions
func (client *XenAPIClient) session_get_all_subject_identifiers(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_all_subject_identifiers", session_id)
}

// session_local_logout
//
// Log out of local session.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) session_local_logout(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("session.local_logout", session_id)
}

// session_slave_local_login_with_password
//
// Authenticate locally against a slave in emergency mode. Note the resulting sessions are only good for use on this host.
//
// params:
// - uname, string, Username for login.
// - pwd, string, Password for login.
//
// returns:
// - session ref
// - ID of newly created session
func (client *XenAPIClient) session_slave_local_login_with_password(uname string, pwd string) (i interface{}, err error) {
	return client.RPCCall("session.slave_local_login_with_password", uname, pwd)
}

// session_change_password
//
// Change the account password; if your session is authenticated with root priviledges then the old_pwd is validated and the new_pwd is set regardless
//
// params:
// - session_id, session ref, Reference to a valid session
// - old_pwd, string, Old password for account
// - new_pwd, string, New password for account
//
// returns:
// - void
func (client *XenAPIClient) session_change_password(session_id interface{}, old_pwd string, new_pwd string) (i interface{}, err error) {
	return client.RPCCall("session.change_password", session_id, old_pwd, new_pwd)
}

// session_logout
//
// Log out of a session
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) session_logout(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("session.logout", session_id)
}

// session_login_with_password
//
// Attempt to authenticate the user, returning a session reference if successful
//
// params:
// - uname, string, Username for login.
// - pwd, string, Password for login.
// - version, string, Client API version.
// - originator, string, Key string for distinguishing different API users sharing the same login name.
//
// returns:
// - session ref
// - reference of newly created session
func (client *XenAPIClient) session_login_with_password(uname string, pwd string, version string, originator string) (i interface{}, err error) {
	return client.RPCCall("session.login_with_password", uname, pwd, version, originator)
}

// session_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given session.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) session_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("session.remove_from_other_config", session_id, self, key)
}

// session_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) session_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("session.add_to_other_config", session_id, self, key, value)
}

// session_set_other_config
//
// Set the other_config field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) session_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("session.set_other_config", session_id, self, value)
}

// session_get_originator
//
// Get the originator field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) session_get_originator(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_originator", session_id, self)
}

// session_get_parent
//
// Get the parent field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - session ref
// - value of the field
func (client *XenAPIClient) session_get_parent(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_parent", session_id, self)
}

// session_get_tasks
//
// Get the tasks field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - task ref set
// - value of the field
func (client *XenAPIClient) session_get_tasks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_tasks", session_id, self)
}

// session_get_rbac_permissions
//
// Get the rbac_permissions field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) session_get_rbac_permissions(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_rbac_permissions", session_id, self)
}

// session_get_auth_user_name
//
// Get the auth_user_name field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) session_get_auth_user_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_auth_user_name", session_id, self)
}

// session_get_auth_user_sid
//
// Get the auth_user_sid field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) session_get_auth_user_sid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_auth_user_sid", session_id, self)
}

// session_get_validation_time
//
// Get the validation_time field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) session_get_validation_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_validation_time", session_id, self)
}

// session_get_subject
//
// Get the subject field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - subject ref
// - value of the field
func (client *XenAPIClient) session_get_subject(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_subject", session_id, self)
}

// session_get_is_local_superuser
//
// Get the is_local_superuser field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) session_get_is_local_superuser(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_is_local_superuser", session_id, self)
}

// session_get_other_config
//
// Get the other_config field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) session_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_other_config", session_id, self)
}

// session_get_pool
//
// Get the pool field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) session_get_pool(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_pool", session_id, self)
}

// session_get_last_active
//
// Get the last_active field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) session_get_last_active(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_last_active", session_id, self)
}

// session_get_this_user
//
// Get the this_user field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - user ref
// - value of the field
func (client *XenAPIClient) session_get_this_user(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_this_user", session_id, self)
}

// session_get_this_host
//
// Get the this_host field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) session_get_this_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_this_host", session_id, self)
}

// session_get_uuid
//
// Get the uuid field of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) session_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_uuid", session_id, self)
}

// session_get_by_uuid
//
// Get a reference to the session instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - session ref
// - reference to the object
func (client *XenAPIClient) session_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("session.get_by_uuid", session_id, uuid)
}

// session_get_record
//
// Get a record containing the current state of the given session.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, session ref, reference to the object
//
// returns:
// - session record
// - all fields from the object
func (client *XenAPIClient) session_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("session.get_record", session_id, self)
}

// auth_get_group_membership
//
// This calls queries the external directory service to obtain the transitively-closed set of groups that the the subject_identifier is member of.
//
// params:
// - session_id, session ref, Reference to a valid session
// - subject_identifier, string, A string containing the subject_identifier, unique in the external directory service
//
// returns:
// - string set
// - set of subject_identifiers that provides the group membership of subject_identifier passed as argument, it contains, recursively, all groups a subject_identifier is member of.
func (client *XenAPIClient) auth_get_group_membership(session_id interface{}, subject_identifier string) (i interface{}, err error) {
	return client.RPCCall("auth.get_group_membership", session_id, subject_identifier)
}

// auth_get_subject_information_from_identifier
//
// This call queries the external directory service to obtain the user information (e.g. username, organization etc) from the specified subject_identifier
//
// params:
// - session_id, session ref, Reference to a valid session
// - subject_identifier, string, A string containing the subject_identifier, unique in the external directory service
//
// returns:
// - (string -> string) map
// - key-value pairs containing at least a key called subject_name
func (client *XenAPIClient) auth_get_subject_information_from_identifier(session_id interface{}, subject_identifier string) (i interface{}, err error) {
	return client.RPCCall("auth.get_subject_information_from_identifier", session_id, subject_identifier)
}

// auth_get_subject_identifier
//
// This call queries the external directory service to obtain the subject_identifier as a string from the human-readable subject_name
//
// params:
// - session_id, session ref, Reference to a valid session
// - subject_name, string, The human-readable subject_name, such as a username or a groupname
//
// returns:
// - string
// - the subject_identifier obtained from the external directory service
func (client *XenAPIClient) auth_get_subject_identifier(session_id interface{}, subject_name string) (i interface{}, err error) {
	return client.RPCCall("auth.get_subject_identifier", session_id, subject_name)
}

// subject_get_all_records
//
// Return a map of subject references to subject records for all subjects known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (subject ref -> subject record) map
// - records of all objects
func (client *XenAPIClient) subject_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_all_records", session_id)
}

// subject_get_all
//
// Return a list of all the subjects known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - subject ref set
// - references to all objects
func (client *XenAPIClient) subject_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_all", session_id)
}

// subject_get_permissions_name_label
//
// This call returns a list of permission names given a subject
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, The subject whose permissions will be retrieved
//
// returns:
// - string set
// - a list of permission names
func (client *XenAPIClient) subject_get_permissions_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_permissions_name_label", session_id, self)
}

// subject_remove_from_roles
//
// This call removes a role from a subject
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, The subject from whom we want to remove the role
// - role, role ref, The unique role reference in the subject's roles field
//
// returns:
// - void
func (client *XenAPIClient) subject_remove_from_roles(session_id interface{}, self interface{}, role interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.remove_from_roles", session_id, self, role)
}

// subject_add_to_roles
//
// This call adds a new role to a subject
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, The subject who we want to add the role to
// - role, role ref, The unique role reference
//
// returns:
// - void
func (client *XenAPIClient) subject_add_to_roles(session_id interface{}, self interface{}, role interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.add_to_roles", session_id, self, role)
}

// subject_get_roles
//
// Get the roles field of the given subject.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, reference to the object
//
// returns:
// - role ref set
// - value of the field
func (client *XenAPIClient) subject_get_roles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_roles", session_id, self)
}

// subject_get_other_config
//
// Get the other_config field of the given subject.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) subject_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_other_config", session_id, self)
}

// subject_get_subject_identifier
//
// Get the subject_identifier field of the given subject.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) subject_get_subject_identifier(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_subject_identifier", session_id, self)
}

// subject_get_uuid
//
// Get the uuid field of the given subject.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) subject_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_uuid", session_id, self)
}

// subject_destroy
//
// Destroy the specified subject instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) subject_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.destroy", session_id, self)
}

// subject_create
//
// Create a new subject instance, and return its handle.
// The constructor args are: subject_identifier, other_config (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, subject record, All constructor arguments
//
// returns:
// - subject ref
// - reference to the newly created object
func (client *XenAPIClient) subject_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.create", session_id, args)
}

// subject_get_by_uuid
//
// Get a reference to the subject instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - subject ref
// - reference to the object
func (client *XenAPIClient) subject_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("subject.get_by_uuid", session_id, uuid)
}

// subject_get_record
//
// Get a record containing the current state of the given subject.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, subject ref, reference to the object
//
// returns:
// - subject record
// - all fields from the object
func (client *XenAPIClient) subject_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("subject.get_record", session_id, self)
}

// role_get_all_records
//
// Return a map of role references to role records for all roles known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (role ref -> role record) map
// - records of all objects
func (client *XenAPIClient) role_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_all_records", session_id)
}

// role_get_all
//
// Return a list of all the roles known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - role ref set
// - references to all objects
func (client *XenAPIClient) role_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_all", session_id)
}

// role_get_by_permission_name_label
//
// This call returns a list of roles given a permission name
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, The short friendly name of the role
//
// returns:
// - role ref set
// - a list of references to roles
func (client *XenAPIClient) role_get_by_permission_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("role.get_by_permission_name_label", session_id, label)
}

// role_get_by_permission
//
// This call returns a list of roles given a permission
//
// params:
// - session_id, session ref, Reference to a valid session
// - permission, role ref, a reference to a permission
//
// returns:
// - role ref set
// - a list of references to roles
func (client *XenAPIClient) role_get_by_permission(session_id interface{}, permission interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_by_permission", session_id, permission)
}

// role_get_permissions_name_label
//
// This call returns a list of permission names given a role
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, role ref, a reference to a role
//
// returns:
// - string set
// - a list of permission names
func (client *XenAPIClient) role_get_permissions_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_permissions_name_label", session_id, self)
}

// role_get_permissions
//
// This call returns a list of permissions given a role
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, role ref, a reference to a role
//
// returns:
// - role ref set
// - a list of permissions
func (client *XenAPIClient) role_get_permissions(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_permissions", session_id, self)
}

// role_get_subroles
//
// Get the subroles field of the given role.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, role ref, reference to the object
//
// returns:
// - role ref set
// - value of the field
func (client *XenAPIClient) role_get_subroles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_subroles", session_id, self)
}

// role_get_name_description
//
// Get the name/description field of the given role.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, role ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) role_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_name_description", session_id, self)
}

// role_get_name_label
//
// Get the name/label field of the given role.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, role ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) role_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_name_label", session_id, self)
}

// role_get_uuid
//
// Get the uuid field of the given role.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, role ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) role_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_uuid", session_id, self)
}

// role_get_by_name_label
//
// Get all the role instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - role ref set
// - references to objects with matching names
func (client *XenAPIClient) role_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("role.get_by_name_label", session_id, label)
}

// role_get_by_uuid
//
// Get a reference to the role instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - role ref
// - reference to the object
func (client *XenAPIClient) role_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("role.get_by_uuid", session_id, uuid)
}

// role_get_record
//
// Get a record containing the current state of the given role.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, role ref, reference to the object
//
// returns:
// - role record
// - all fields from the object
func (client *XenAPIClient) role_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("role.get_record", session_id, self)
}

// task_get_all_records
//
// Return a map of task references to task records for all tasks known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (task ref -> task record) map
// - records of all objects
func (client *XenAPIClient) task_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_all_records", session_id)
}

// task_get_all
//
// Return a list of all the tasks known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - task ref set
// - references to all objects
func (client *XenAPIClient) task_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_all", session_id)
}

// task_cancel
//
// Request that a task be cancelled. Note that a task may fail to be cancelled and may complete or fail normally and note that, even when a task does cancel, it might take an arbitrary amount of time.
//
// params:
// - session_id, session ref, Reference to a valid session
// - task, task ref, The task
//
// returns:
// - void
func (client *XenAPIClient) task_cancel(session_id interface{}, task interface{}) (i interface{}, err error) {
	return client.RPCCall("task.cancel", session_id, task)
}

// task_destroy
//
// Destroy the task object
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, Reference to the task object
//
// returns:
// - void
func (client *XenAPIClient) task_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.destroy", session_id, self)
}

// task_create
//
// Create a new task object which must be manually destroyed.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, short label for the new task
// - description, string, longer description for the new task
//
// returns:
// - task ref
// - The reference of the created task object
func (client *XenAPIClient) task_create(session_id interface{}, label string, description string) (i interface{}, err error) {
	return client.RPCCall("task.create", session_id, label, description)
}

// task_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given task.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) task_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("task.remove_from_other_config", session_id, self, key)
}

// task_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) task_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("task.add_to_other_config", session_id, self, key, value)
}

// task_set_other_config
//
// Set the other_config field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) task_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("task.set_other_config", session_id, self, value)
}

// task_get_backtrace
//
// Get the backtrace field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) task_get_backtrace(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_backtrace", session_id, self)
}

// task_get_subtasks
//
// Get the subtasks field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - task ref set
// - value of the field
func (client *XenAPIClient) task_get_subtasks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_subtasks", session_id, self)
}

// task_get_subtask_of
//
// Get the subtask_of field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - task ref
// - value of the field
func (client *XenAPIClient) task_get_subtask_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_subtask_of", session_id, self)
}

// task_get_other_config
//
// Get the other_config field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) task_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_other_config", session_id, self)
}

// task_get_error_info
//
// Get the error_info field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) task_get_error_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_error_info", session_id, self)
}

// task_get_result
//
// Get the result field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) task_get_result(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_result", session_id, self)
}

// task_get_type
//
// Get the type field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) task_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_type", session_id, self)
}

// task_get_progress
//
// Get the progress field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) task_get_progress(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_progress", session_id, self)
}

// task_get_resident_on
//
// Get the resident_on field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) task_get_resident_on(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_resident_on", session_id, self)
}

// task_get_status
//
// Get the status field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - enum task_status_type
// - value of the field
func (client *XenAPIClient) task_get_status(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_status", session_id, self)
}

// task_get_finished
//
// Get the finished field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) task_get_finished(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_finished", session_id, self)
}

// task_get_created
//
// Get the created field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) task_get_created(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_created", session_id, self)
}

// task_get_current_operations
//
// Get the current_operations field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - (string -> enum task_allowed_operations) map
// - value of the field
func (client *XenAPIClient) task_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_current_operations", session_id, self)
}

// task_get_allowed_operations
//
// Get the allowed_operations field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - enum task_allowed_operations set
// - value of the field
func (client *XenAPIClient) task_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_allowed_operations", session_id, self)
}

// task_get_name_description
//
// Get the name/description field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) task_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_name_description", session_id, self)
}

// task_get_name_label
//
// Get the name/label field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) task_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_name_label", session_id, self)
}

// task_get_uuid
//
// Get the uuid field of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) task_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_uuid", session_id, self)
}

// task_get_by_name_label
//
// Get all the task instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - task ref set
// - references to objects with matching names
func (client *XenAPIClient) task_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("task.get_by_name_label", session_id, label)
}

// task_get_by_uuid
//
// Get a reference to the task instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - task ref
// - reference to the object
func (client *XenAPIClient) task_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("task.get_by_uuid", session_id, uuid)
}

// task_get_record
//
// Get a record containing the current state of the given task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, task ref, reference to the object
//
// returns:
// - task record
// - all fields from the object
func (client *XenAPIClient) task_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("task.get_record", session_id, self)
}

// event_inject
//
// Injects an artificial event on the given object and return the corresponding ID
//
// params:
// - session_id, session ref, Reference to a valid session
// - class, string, class of the object
// - ref, string, A reference to the object that will be changed.
//
// returns:
// - string
// - the event ID
func (client *XenAPIClient) event_inject(session_id interface{}, class string, ref string) (i interface{}, err error) {
	return client.RPCCall("event.inject", session_id, class, ref)
}

// event_get_current_id
//
// Return the ID of the next event to be generated by the system
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - int
// - the event ID
func (client *XenAPIClient) event_get_current_id(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("event.get_current_id", session_id)
}

// event_from
//
// Blocking call which returns a new token and a (possibly empty) batch of events. The returned token can be used in subsequent calls to this function.
//
// params:
// - session_id, session ref, Reference to a valid session
// - classes, string set, register for events for the indicated classes
// - token, string, A token representing the point from which to generate database events. The empty string represents the beginning.
// - timeout, float, Return after this many seconds if no events match
//
// returns:
// - event record set
// - the batch of events
func (client *XenAPIClient) event_from(session_id interface{}, classes interface{}, token string, timeout interface{}) (i interface{}, err error) {
	return client.RPCCall("event.from", session_id, classes, token, timeout)
}

// event_next
//
// Blocking call which returns a (possibly empty) batch of events. This method is only recommended for legacy use. New development should use event.from which supercedes this method.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - event record set
// - the batch of events
func (client *XenAPIClient) event_next(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("event.next", session_id)
}

// event_unregister
//
// Unregisters this session with the event system
//
// params:
// - session_id, session ref, Reference to a valid session
// - classes, string set, remove this session's registration for the indicated classes
//
// returns:
// - void
func (client *XenAPIClient) event_unregister(session_id interface{}, classes interface{}) (i interface{}, err error) {
	return client.RPCCall("event.unregister", session_id, classes)
}

// event_register
//
// Registers this session with the event system.  Specifying * as the desired class will register for all classes.
//
// params:
// - session_id, session ref, Reference to a valid session
// - classes, string set, register for events for the indicated classes
//
// returns:
// - void
func (client *XenAPIClient) event_register(session_id interface{}, classes interface{}) (i interface{}, err error) {
	return client.RPCCall("event.register", session_id, classes)
}

// pool_get_all_records
//
// Return a map of pool references to pool records for all pools known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (pool ref -> pool record) map
// - records of all objects
func (client *XenAPIClient) pool_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_all_records", session_id)
}

// pool_get_all
//
// Return a list of all the pools known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - pool ref set
// - references to all objects
func (client *XenAPIClient) pool_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_all", session_id)
}

// pool_apply_edition
//
// Apply an edition to all hosts in the pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, Reference to the pool
// - edition, string, The requested edition
//
// returns:
// - void
func (client *XenAPIClient) pool_apply_edition(session_id interface{}, self interface{}, edition string) (i interface{}, err error) {
	return client.RPCCall("pool.apply_edition", session_id, self, edition)
}

// pool_get_license_state
//
// This call returns the license state for the pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, Reference to the pool
//
// returns:
// - (string -> string) map
// - The pool's license state
func (client *XenAPIClient) pool_get_license_state(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_license_state", session_id, self)
}

// pool_disable_local_storage_caching
//
// This call disables pool-wide local storage caching
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, Reference to the pool
//
// returns:
// - void
func (client *XenAPIClient) pool_disable_local_storage_caching(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.disable_local_storage_caching", session_id, self)
}

// pool_enable_local_storage_caching
//
// This call attempts to enable pool-wide local storage caching
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, Reference to the pool
//
// returns:
// - void
func (client *XenAPIClient) pool_enable_local_storage_caching(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.enable_local_storage_caching", session_id, self)
}

// pool_test_archive_target
//
// This call tests if a location is valid
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, Reference to the pool
// - config, (string -> string) map, Location config settings to test
//
// returns:
// - string
// - An XMLRPC result
func (client *XenAPIClient) pool_test_archive_target(session_id interface{}, self interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.test_archive_target", session_id, self, config)
}

// pool_set_vswitch_controller
//
// Set the IP address of the vswitch controller.
//
// params:
// - session_id, session ref, Reference to a valid session
// - address, string, IP address of the vswitch controller.
//
// returns:
// - void
func (client *XenAPIClient) pool_set_vswitch_controller(session_id interface{}, address string) (i interface{}, err error) {
	return client.RPCCall("pool.set_vswitch_controller", session_id, address)
}

// pool_disable_redo_log
//
// Disable the redo log if in use, unless HA is enabled.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) pool_disable_redo_log(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.disable_redo_log", session_id)
}

// pool_enable_redo_log
//
// Enable the redo log on the given SR and start using it, unless HA is enabled.
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, SR to hold the redo log.
//
// returns:
// - void
func (client *XenAPIClient) pool_enable_redo_log(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.enable_redo_log", session_id, sr)
}

// pool_certificate_sync
//
// Sync SSL certificates from master to slaves.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) pool_certificate_sync(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_sync", session_id)
}

// pool_crl_list
//
// List all installed SSL certificate revocation lists.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - string set
// - All installed CRLs
func (client *XenAPIClient) pool_crl_list(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.crl_list", session_id)
}

// pool_crl_uninstall
//
// Remove an SSL certificate revocation list.
//
// params:
// - session_id, session ref, Reference to a valid session
// - name, string, The CRL name
//
// returns:
// - void
func (client *XenAPIClient) pool_crl_uninstall(session_id interface{}, name string) (i interface{}, err error) {
	return client.RPCCall("pool.crl_uninstall", session_id, name)
}

// pool_crl_install
//
// Install an SSL certificate revocation list, pool-wide.
//
// params:
// - session_id, session ref, Reference to a valid session
// - name, string, A name to give the CRL
// - cert, string, The CRL
//
// returns:
// - void
func (client *XenAPIClient) pool_crl_install(session_id interface{}, name string, cert string) (i interface{}, err error) {
	return client.RPCCall("pool.crl_install", session_id, name, cert)
}

// pool_certificate_list
//
// List all installed SSL certificates.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - string set
// - All installed certificates
func (client *XenAPIClient) pool_certificate_list(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_list", session_id)
}

// pool_certificate_uninstall
//
// Remove an SSL certificate.
//
// params:
// - session_id, session ref, Reference to a valid session
// - name, string, The certificate name
//
// returns:
// - void
func (client *XenAPIClient) pool_certificate_uninstall(session_id interface{}, name string) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_uninstall", session_id, name)
}

// pool_certificate_install
//
// Install an SSL certificate pool-wide.
//
// params:
// - session_id, session ref, Reference to a valid session
// - name, string, A name to give the certificate
// - cert, string, The certificate
//
// returns:
// - void
func (client *XenAPIClient) pool_certificate_install(session_id interface{}, name string, cert string) (i interface{}, err error) {
	return client.RPCCall("pool.certificate_install", session_id, name, cert)
}

// pool_send_test_post
//
// Send the given body to the given host and port, using HTTPS, and print the response.  This is used for debugging the SSL layer.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, string,
// - port, int,
// - body, string,
//
// returns:
// - string
// - The response
func (client *XenAPIClient) pool_send_test_post(session_id interface{}, host string, port interface{}, body string) (i interface{}, err error) {
	return client.RPCCall("pool.send_test_post", session_id, host, port, body)
}

// pool_retrieve_wlb_recommendations
//
// Retrieves vm migrate recommendations for the pool from the workload balancing server
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VM ref -> string set) map
// - The list of vm migration recommendations
func (client *XenAPIClient) pool_retrieve_wlb_recommendations(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.retrieve_wlb_recommendations", session_id)
}

// pool_retrieve_wlb_configuration
//
// Retrieves the pool optimization criteria from the workload balancing server
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (string -> string) map
// - The configuration used in optimizing this pool
func (client *XenAPIClient) pool_retrieve_wlb_configuration(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.retrieve_wlb_configuration", session_id)
}

// pool_send_wlb_configuration
//
// Sets the pool optimization criteria for the workload balancing server
//
// params:
// - session_id, session ref, Reference to a valid session
// - config, (string -> string) map, The configuration to use in optimizing this pool
//
// returns:
// - void
func (client *XenAPIClient) pool_send_wlb_configuration(session_id interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.send_wlb_configuration", session_id, config)
}

// pool_deconfigure_wlb
//
// Permanently deconfigures workload balancing monitoring on this pool
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) pool_deconfigure_wlb(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.deconfigure_wlb", session_id)
}

// pool_initialize_wlb
//
// Initializes workload balancing monitoring on this pool with the specified wlb server
//
// params:
// - session_id, session ref, Reference to a valid session
// - wlb_url, string, The ip address and port to use when accessing the wlb server
// - wlb_username, string, The username used to authenticate with the wlb server
// - wlb_password, string, The password used to authenticate with the wlb server
// - xenserver_username, string, The usernamed used by the wlb server to authenticate with the xenserver
// - xenserver_password, string, The password used by the wlb server to authenticate with the xenserver
//
// returns:
// - void
func (client *XenAPIClient) pool_initialize_wlb(session_id interface{}, wlb_url string, wlb_username string, wlb_password string, xenserver_username string, xenserver_password string) (i interface{}, err error) {
	return client.RPCCall("pool.initialize_wlb", session_id, wlb_url, wlb_username, wlb_password, xenserver_username, xenserver_password)
}

// pool_detect_nonhomogeneous_external_auth
//
// This call asynchronously detects if the external authentication configuration in any slave is different from that in the master and raises appropriate alerts
//
// params:
// - session_id, session ref, Reference to a valid session
// - pool, pool ref, The pool where to detect non-homogeneous external authentication configuration
//
// returns:
// - void
func (client *XenAPIClient) pool_detect_nonhomogeneous_external_auth(session_id interface{}, pool interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.detect_nonhomogeneous_external_auth", session_id, pool)
}

// pool_disable_external_auth
//
// This call disables external authentication on all the hosts of the pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - pool, pool ref, The pool whose external authentication should be disabled
// - config, (string -> string) map, Optional parameters as a list of key-values containing the configuration data
//
// returns:
// - void
func (client *XenAPIClient) pool_disable_external_auth(session_id interface{}, pool interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.disable_external_auth", session_id, pool, config)
}

// pool_enable_external_auth
//
// This call enables external authentication on all the hosts of the pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - pool, pool ref, The pool whose external authentication should be enabled
// - config, (string -> string) map, A list of key-values containing the configuration data
// - service_name, string, The name of the service
// - auth_type, string, The type of authentication (e.g. AD for Active Directory)
//
// returns:
// - void
func (client *XenAPIClient) pool_enable_external_auth(session_id interface{}, pool interface{}, config map[string]string, service_name string, auth_type string) (i interface{}, err error) {
	return client.RPCCall("pool.enable_external_auth", session_id, pool, config, service_name, auth_type)
}

// pool_create_new_blob
//
// Create a placeholder for a named binary blob of data that is associated with this pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - pool, pool ref, The pool
// - name, string, The name associated with the blob
// - mime_type, string, The mime type for the data. Empty string translates to application/octet-stream
// - public, bool, True if the blob should be publicly available
//
// returns:
// - blob ref
// - The reference of the blob, needed for populating its data
func (client *XenAPIClient) pool_create_new_blob(session_id interface{}, pool interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("pool.create_new_blob", session_id, pool, name, mime_type, public)
}

// pool_set_ha_host_failures_to_tolerate
//
// Set the maximum number of host failures to consider in the HA VM restart planner
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, The pool
// - value, int, New number of host failures to consider
//
// returns:
// - void
func (client *XenAPIClient) pool_set_ha_host_failures_to_tolerate(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_ha_host_failures_to_tolerate", session_id, self, value)
}

// pool_ha_compute_vm_failover_plan
//
// Return a VM failover plan assuming a given subset of hosts fail
//
// params:
// - session_id, session ref, Reference to a valid session
// - failed_hosts, host ref set, The set of hosts to assume have failed
// - failed_vms, VM ref set, The set of VMs to restart
//
// returns:
// - (VM ref -> (string -> string) map) map
// - VM failover plan: a map of VM to host to restart the host on
func (client *XenAPIClient) pool_ha_compute_vm_failover_plan(session_id interface{}, failed_hosts interface{}, failed_vms interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_compute_vm_failover_plan", session_id, failed_hosts, failed_vms)
}

// pool_ha_compute_hypothetical_max_host_failures_to_tolerate
//
// Returns the maximum number of host failures we could tolerate before we would be unable to restart the provided VMs
//
// params:
// - session_id, session ref, Reference to a valid session
// - configuration, (VM ref -> string) map, Map of protected VM reference to restart priority
//
// returns:
// - int
// - maximum value for ha_host_failures_to_tolerate given provided configuration
func (client *XenAPIClient) pool_ha_compute_hypothetical_max_host_failures_to_tolerate(session_id interface{}, configuration interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_compute_hypothetical_max_host_failures_to_tolerate", session_id, configuration)
}

// pool_ha_compute_max_host_failures_to_tolerate
//
// Returns the maximum number of host failures we could tolerate before we would be unable to restart configured VMs
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - int
// - maximum value for ha_host_failures_to_tolerate given current configuration
func (client *XenAPIClient) pool_ha_compute_max_host_failures_to_tolerate(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_compute_max_host_failures_to_tolerate", session_id)
}

// pool_ha_failover_plan_exists
//
// Returns true if a VM failover plan exists for up to 'n' host failures
//
// params:
// - session_id, session ref, Reference to a valid session
// - n, int, The number of host failures to plan for
//
// returns:
// - bool
// - true if a failover plan exists for the supplied number of host failures
func (client *XenAPIClient) pool_ha_failover_plan_exists(session_id interface{}, n interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_failover_plan_exists", session_id, n)
}

// pool_ha_prevent_restarts_for
//
// When this call returns the VM restart logic will not run for the requested number of seconds. If the argument is zero then the restart thread is immediately unblocked
//
// params:
// - session_id, session ref, Reference to a valid session
// - seconds, int, The number of seconds to block the restart thread for
//
// returns:
// - void
func (client *XenAPIClient) pool_ha_prevent_restarts_for(session_id interface{}, seconds interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.ha_prevent_restarts_for", session_id, seconds)
}

// pool_designate_new_master
//
// Perform an orderly handover of the role of master to the referenced host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host who should become the new master
//
// returns:
// - void
func (client *XenAPIClient) pool_designate_new_master(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.designate_new_master", session_id, host)
}

// pool_sync_database
//
// Forcibly synchronise the database now
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) pool_sync_database(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.sync_database", session_id)
}

// pool_disable_ha
//
// Turn off High Availability mode
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) pool_disable_ha(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.disable_ha", session_id)
}

// pool_enable_ha
//
// Turn on High Availability mode
//
// params:
// - session_id, session ref, Reference to a valid session
// - heartbeat_srs, SR ref set, Set of SRs to use for storage heartbeating.
// - configuration, (string -> string) map, Detailed HA configuration to apply
//
// returns:
// - void
func (client *XenAPIClient) pool_enable_ha(session_id interface{}, heartbeat_srs interface{}, configuration map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.enable_ha", session_id, heartbeat_srs, configuration)
}

// pool_create_VLAN_from_PIF
//
// Create a pool-wide VLAN by taking the PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - pif, PIF ref, physical interface on any particular host, that identifies the PIF on which to create the (pool-wide) VLAN interface
// - network, network ref, network to which this interface should be connected
// - VLAN, int, VLAN tag for the new interface
//
// returns:
// - PIF ref set
// - The references of the created PIF objects
func (client *XenAPIClient) pool_create_VLAN_from_PIF(session_id interface{}, pif interface{}, network interface{}, VLAN interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.create_VLAN_from_PIF", session_id, pif, network, VLAN)
}

// pool_create_VLAN
//
// Create PIFs, mapping a network to the same physical interface/VLAN on each host. This call is deprecated: use Pool.create_VLAN_from_PIF instead.
//
// params:
// - session_id, session ref, Reference to a valid session
// - device, string, physical interface on which to create the VLAN interface
// - network, network ref, network to which this interface should be connected
// - VLAN, int, VLAN tag for the new interface
//
// returns:
// - PIF ref set
// - The references of the created PIF objects
func (client *XenAPIClient) pool_create_VLAN(session_id interface{}, device string, network interface{}, VLAN interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.create_VLAN", session_id, device, network, VLAN)
}

// pool_recover_slaves
//
// Instruct a pool master, M, to try and contact its slaves and, if slaves are in emergency mode, reset their master address to M.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - host ref set
// - list of hosts whose master address were successfully reset
func (client *XenAPIClient) pool_recover_slaves(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.recover_slaves", session_id)
}

// pool_emergency_reset_master
//
// Instruct a slave already in a pool that the master has changed
//
// params:
// - session_id, session ref, Reference to a valid session
// - master_address, string, The hostname of the master
//
// returns:
// - void
func (client *XenAPIClient) pool_emergency_reset_master(session_id interface{}, master_address string) (i interface{}, err error) {
	return client.RPCCall("pool.emergency_reset_master", session_id, master_address)
}

// pool_emergency_transition_to_master
//
// Instruct host that's currently a slave to transition to being master
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) pool_emergency_transition_to_master(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.emergency_transition_to_master", session_id)
}

// pool_eject
//
// Instruct a pool master to eject a host from the pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to eject
//
// returns:
// - void
func (client *XenAPIClient) pool_eject(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.eject", session_id, host)
}

// pool_join_force
//
// Instruct host to join a new pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - master_address, string, The hostname of the master of the pool to join
// - master_username, string, The username of the master (for initial authentication)
// - master_password, string, The password for the master (for initial authentication)
//
// returns:
// - void
func (client *XenAPIClient) pool_join_force(session_id interface{}, master_address string, master_username string, master_password string) (i interface{}, err error) {
	return client.RPCCall("pool.join_force", session_id, master_address, master_username, master_password)
}

// pool_join
//
// Instruct host to join a new pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - master_address, string, The hostname of the master of the pool to join
// - master_username, string, The username of the master (for initial authentication)
// - master_password, string, The password for the master (for initial authentication)
//
// returns:
// - void
func (client *XenAPIClient) pool_join(session_id interface{}, master_address string, master_username string, master_password string) (i interface{}, err error) {
	return client.RPCCall("pool.join", session_id, master_address, master_username, master_password)
}

// pool_set_wlb_verify_cert
//
// Set the wlb_verify_cert field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_wlb_verify_cert(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("pool.set_wlb_verify_cert", session_id, self, value)
}

// pool_set_wlb_enabled
//
// Set the wlb_enabled field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_wlb_enabled(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("pool.set_wlb_enabled", session_id, self, value)
}

// pool_remove_from_gui_config
//
// Remove the given key and its corresponding value from the gui_config field of the given pool.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) pool_remove_from_gui_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("pool.remove_from_gui_config", session_id, self, key)
}

// pool_add_to_gui_config
//
// Add the given key-value pair to the gui_config field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) pool_add_to_gui_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("pool.add_to_gui_config", session_id, self, key, value)
}

// pool_set_gui_config
//
// Set the gui_config field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_gui_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.set_gui_config", session_id, self, value)
}

// pool_remove_tags
//
// Remove the given value from the tags field of the given pool.  If the value is not in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, string, Value to remove
//
// returns:
// - void
func (client *XenAPIClient) pool_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.remove_tags", session_id, self, value)
}

// pool_add_tags
//
// Add the given value to the tags field of the given pool.  If the value is already in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, string, New value to add
//
// returns:
// - void
func (client *XenAPIClient) pool_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.add_tags", session_id, self, value)
}

// pool_set_tags
//
// Set the tags field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, string set, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_tags", session_id, self, value)
}

// pool_set_ha_allow_overcommit
//
// Set the ha_allow_overcommit field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_ha_allow_overcommit(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("pool.set_ha_allow_overcommit", session_id, self, value)
}

// pool_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given pool.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) pool_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("pool.remove_from_other_config", session_id, self, key)
}

// pool_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) pool_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("pool.add_to_other_config", session_id, self, key, value)
}

// pool_set_other_config
//
// Set the other_config field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("pool.set_other_config", session_id, self, value)
}

// pool_set_crash_dump_SR
//
// Set the crash_dump_SR field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, SR ref, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_crash_dump_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_crash_dump_SR", session_id, self, value)
}

// pool_set_suspend_image_SR
//
// Set the suspend_image_SR field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, SR ref, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_suspend_image_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_suspend_image_SR", session_id, self, value)
}

// pool_set_default_SR
//
// Set the default_SR field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, SR ref, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_default_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.set_default_SR", session_id, self, value)
}

// pool_set_name_description
//
// Set the name_description field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.set_name_description", session_id, self, value)
}

// pool_set_name_label
//
// Set the name_label field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) pool_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("pool.set_name_label", session_id, self, value)
}

// pool_get_metadata_VDIs
//
// Get the metadata_VDIs field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - VDI ref set
// - value of the field
func (client *XenAPIClient) pool_get_metadata_VDIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_metadata_VDIs", session_id, self)
}

// pool_get_restrictions
//
// Get the restrictions field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) pool_get_restrictions(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_restrictions", session_id, self)
}

// pool_get_vswitch_controller
//
// Get the vswitch_controller field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) pool_get_vswitch_controller(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_vswitch_controller", session_id, self)
}

// pool_get_redo_log_vdi
//
// Get the redo_log_vdi field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - VDI ref
// - value of the field
func (client *XenAPIClient) pool_get_redo_log_vdi(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_redo_log_vdi", session_id, self)
}

// pool_get_redo_log_enabled
//
// Get the redo_log_enabled field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) pool_get_redo_log_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_redo_log_enabled", session_id, self)
}

// pool_get_wlb_verify_cert
//
// Get the wlb_verify_cert field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) pool_get_wlb_verify_cert(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_verify_cert", session_id, self)
}

// pool_get_wlb_enabled
//
// Get the wlb_enabled field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) pool_get_wlb_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_enabled", session_id, self)
}

// pool_get_wlb_username
//
// Get the wlb_username field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) pool_get_wlb_username(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_username", session_id, self)
}

// pool_get_wlb_url
//
// Get the wlb_url field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) pool_get_wlb_url(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_wlb_url", session_id, self)
}

// pool_get_gui_config
//
// Get the gui_config field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) pool_get_gui_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_gui_config", session_id, self)
}

// pool_get_tags
//
// Get the tags field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) pool_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_tags", session_id, self)
}

// pool_get_blobs
//
// Get the blobs field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - (string -> blob ref) map
// - value of the field
func (client *XenAPIClient) pool_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_blobs", session_id, self)
}

// pool_get_ha_overcommitted
//
// Get the ha_overcommitted field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) pool_get_ha_overcommitted(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_overcommitted", session_id, self)
}

// pool_get_ha_allow_overcommit
//
// Get the ha_allow_overcommit field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) pool_get_ha_allow_overcommit(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_allow_overcommit", session_id, self)
}

// pool_get_ha_plan_exists_for
//
// Get the ha_plan_exists_for field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) pool_get_ha_plan_exists_for(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_plan_exists_for", session_id, self)
}

// pool_get_ha_host_failures_to_tolerate
//
// Get the ha_host_failures_to_tolerate field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) pool_get_ha_host_failures_to_tolerate(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_host_failures_to_tolerate", session_id, self)
}

// pool_get_ha_statefiles
//
// Get the ha_statefiles field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) pool_get_ha_statefiles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_statefiles", session_id, self)
}

// pool_get_ha_configuration
//
// Get the ha_configuration field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) pool_get_ha_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_configuration", session_id, self)
}

// pool_get_ha_enabled
//
// Get the ha_enabled field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) pool_get_ha_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_ha_enabled", session_id, self)
}

// pool_get_other_config
//
// Get the other_config field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) pool_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_other_config", session_id, self)
}

// pool_get_crash_dump_SR
//
// Get the crash_dump_SR field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) pool_get_crash_dump_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_crash_dump_SR", session_id, self)
}

// pool_get_suspend_image_SR
//
// Get the suspend_image_SR field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) pool_get_suspend_image_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_suspend_image_SR", session_id, self)
}

// pool_get_default_SR
//
// Get the default_SR field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) pool_get_default_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_default_SR", session_id, self)
}

// pool_get_master
//
// Get the master field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) pool_get_master(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_master", session_id, self)
}

// pool_get_name_description
//
// Get the name_description field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) pool_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_name_description", session_id, self)
}

// pool_get_name_label
//
// Get the name_label field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) pool_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_name_label", session_id, self)
}

// pool_get_uuid
//
// Get the uuid field of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) pool_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_uuid", session_id, self)
}

// pool_get_by_uuid
//
// Get a reference to the pool instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - pool ref
// - reference to the object
func (client *XenAPIClient) pool_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("pool.get_by_uuid", session_id, uuid)
}

// pool_get_record
//
// Get a record containing the current state of the given pool.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool ref, reference to the object
//
// returns:
// - pool record
// - all fields from the object
func (client *XenAPIClient) pool_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("pool.get_record", session_id, self)
}

// poolPatch_get_all_records
//
// Return a map of pool_patch references to pool_patch records for all pool_patchs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (pool_patch ref -> pool_patch record) map
// - records of all objects
func (client *XenAPIClient) poolPatch_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_all_records", session_id)
}

// poolPatch_get_all
//
// Return a list of all the pool_patchs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - pool_patch ref set
// - references to all objects
func (client *XenAPIClient) poolPatch_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_all", session_id)
}

// poolPatch_clean_on_host
//
// Removes the patch's files from the specified host
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, The patch to clean up
// - host, host ref, The host on which to clean the patch
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_clean_on_host(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.clean_on_host", session_id, self, host)
}

// poolPatch_destroy
//
// Removes the patch's files from all hosts in the pool, and removes the database entries.  Only works on unapplied patches.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, The patch to destroy
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.destroy", session_id, self)
}

// poolPatch_pool_clean
//
// Removes the patch's files from all hosts in the pool, but does not remove the database entries
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, The patch to clean up
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_pool_clean(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.pool_clean", session_id, self)
}

// poolPatch_clean
//
// Removes the patch's files from the server
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, The patch to clean up
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_clean(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.clean", session_id, self)
}

// poolPatch_precheck
//
// Execute the precheck stage of the selected patch on a host and return its output
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, The patch whose prechecks will be run
// - host, host ref, The host to run the prechecks on
//
// returns:
// - string
// - the output of the patch prechecks
func (client *XenAPIClient) poolPatch_precheck(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.precheck", session_id, self, host)
}

// poolPatch_pool_apply
//
// Apply the selected patch to all hosts in the pool and return a map of host_ref -> patch output
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, The patch to apply
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_pool_apply(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.pool_apply", session_id, self)
}

// poolPatch_apply
//
// Apply the selected patch to a host and return its output
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, The patch to apply
// - host, host ref, The host to apply the patch too
//
// returns:
// - string
// - the output of the patch application process
func (client *XenAPIClient) poolPatch_apply(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.apply", session_id, self, host)
}

// poolPatch_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given pool_patch.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.remove_from_other_config", session_id, self, key)
}

// poolPatch_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.add_to_other_config", session_id, self, key, value)
}

// poolPatch_set_other_config
//
// Set the other_config field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) poolPatch_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.set_other_config", session_id, self, value)
}

// poolPatch_get_other_config
//
// Get the other_config field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) poolPatch_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_other_config", session_id, self)
}

// poolPatch_get_after_apply_guidance
//
// Get the after_apply_guidance field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - enum after_apply_guidance set
// - value of the field
func (client *XenAPIClient) poolPatch_get_after_apply_guidance(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_after_apply_guidance", session_id, self)
}

// poolPatch_get_host_patches
//
// Get the host_patches field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - host_patch ref set
// - value of the field
func (client *XenAPIClient) poolPatch_get_host_patches(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_host_patches", session_id, self)
}

// poolPatch_get_pool_applied
//
// Get the pool_applied field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) poolPatch_get_pool_applied(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_pool_applied", session_id, self)
}

// poolPatch_get_size
//
// Get the size field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) poolPatch_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_size", session_id, self)
}

// poolPatch_get_version
//
// Get the version field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) poolPatch_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_version", session_id, self)
}

// poolPatch_get_name_description
//
// Get the name/description field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) poolPatch_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_name_description", session_id, self)
}

// poolPatch_get_name_label
//
// Get the name/label field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) poolPatch_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_name_label", session_id, self)
}

// poolPatch_get_uuid
//
// Get the uuid field of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) poolPatch_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_uuid", session_id, self)
}

// poolPatch_get_by_name_label
//
// Get all the pool_patch instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - pool_patch ref set
// - references to objects with matching names
func (client *XenAPIClient) poolPatch_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_by_name_label", session_id, label)
}

// poolPatch_get_by_uuid
//
// Get a reference to the pool_patch instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - pool_patch ref
// - reference to the object
func (client *XenAPIClient) poolPatch_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_by_uuid", session_id, uuid)
}

// poolPatch_get_record
//
// Get a record containing the current state of the given pool_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, pool_patch ref, reference to the object
//
// returns:
// - pool_patch record
// - all fields from the object
func (client *XenAPIClient) poolPatch_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("poolPatch.get_record", session_id, self)
}

// VM_get_all_records
//
// Return a map of VM references to VM records for all VMs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VM ref -> VM record) map
// - records of all objects
func (client *XenAPIClient) VM_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_all_records", session_id)
}

// VM_get_all
//
// Return a list of all the VMs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VM ref set
// - references to all objects
func (client *XenAPIClient) VM_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_all", session_id)
}

// VM_call_plugin
//
// Call a XenAPI plugin on this vm
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The vm
// - plugin, string, The name of the plugin
// - fn, string, The name of the function within the plugin
// - args, (string -> string) map, Arguments for the function
//
// returns:
// - string
// - Result from the plugin
func (client *XenAPIClient) VM_call_plugin(session_id interface{}, vm interface{}, plugin string, fn string, args map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.call_plugin", session_id, vm, plugin, fn, args)
}

// VM_query_services
//
// Query the system services advertised by this VM and register them. This can only be applied to a system domain.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
//
// returns:
// - (string -> string) map
// - map of service type to name
func (client *XenAPIClient) VM_query_services(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.query_services", session_id, self)
}

// VM_set_appliance
//
// Assign this VM to an appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to assign to an appliance.
// - value, VM_appliance ref, The appliance to which this VM should be assigned.
//
// returns:
// - void
func (client *XenAPIClient) VM_set_appliance(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_appliance", session_id, self, value)
}

// VM_import_convert
//
// Import using a conversion service.
//
// params:
// - session_id, session ref, Reference to a valid session
// - a_type, string, Type of the conversion
// - username, string, Admin username on the host
// - password, string, Password on the host
// - sr, SR ref, The destination SR
// - remote_config, (string -> string) map, Remote configuration options
//
// returns:
// - void
func (client *XenAPIClient) VM_import_convert(session_id interface{}, a_type string, username string, password string, sr interface{}, remote_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.import_convert", session_id, a_type, username, password, sr, remote_config)
}

// VM_recover
//
// Recover the VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to recover
// - session_to, session ref, The session to which the VM is to be recovered.
// - force, bool, Whether the VM should replace newer versions of itself.
//
// returns:
// - void
func (client *XenAPIClient) VM_recover(session_id interface{}, self interface{}, session_to interface{}, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.recover", session_id, self, session_to, force)
}

// VM_get_SRs_required_for_recovery
//
// List all the SR's that are required for the VM to be recovered
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM for which the SRs have to be recovered
// - session_to, session ref, The session to which the SRs of the VM have to be recovered.
//
// returns:
// - SR ref set
// - refs for SRs required to recover the VM
func (client *XenAPIClient) VM_get_SRs_required_for_recovery(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_SRs_required_for_recovery", session_id, self, session_to)
}

// VM_assert_can_be_recovered
//
// Assert whether all SRs required to recover this VM are available.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to recover
// - session_to, session ref, The session to which the VM is to be recovered.
//
// returns:
// - void
func (client *XenAPIClient) VM_assert_can_be_recovered(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_can_be_recovered", session_id, self, session_to)
}

// VM_set_suspend_VDI
//
// Set this VM's suspend VDI, which must be indentical to its current one
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, VDI ref, The suspend VDI uuid
//
// returns:
// - void
func (client *XenAPIClient) VM_set_suspend_VDI(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_suspend_VDI", session_id, self, value)
}

// VM_set_order
//
// Set this VM's boot order
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, int, This VM's boot order
//
// returns:
// - void
func (client *XenAPIClient) VM_set_order(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_order", session_id, self, value)
}

// VM_set_shutdown_delay
//
// Set this VM's shutdown delay in seconds
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, int, This VM's shutdown delay in seconds
//
// returns:
// - void
func (client *XenAPIClient) VM_set_shutdown_delay(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_shutdown_delay", session_id, self, value)
}

// VM_set_start_delay
//
// Set this VM's start delay in seconds
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, int, This VM's start delay in seconds
//
// returns:
// - void
func (client *XenAPIClient) VM_set_start_delay(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_start_delay", session_id, self, value)
}

// VM_set_protection_policy
//
// Set the value of the protection_policy field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, VMPP ref, The value
//
// returns:
// - void
func (client *XenAPIClient) VM_set_protection_policy(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_protection_policy", session_id, self, value)
}

// VM_copy_bios_strings
//
// Copy the BIOS strings from the given host to this VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to modify
// - host, host ref, The host to copy the BIOS strings from
//
// returns:
// - void
func (client *XenAPIClient) VM_copy_bios_strings(session_id interface{}, vm interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.copy_bios_strings", session_id, vm, host)
}

// VM_retrieve_wlb_recommendations
//
// Returns mapping of hosts to ratings, indicating the suitability of starting the VM at that location according to wlb. Rating is replaced with an error if the VM cannot boot there.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM
//
// returns:
// - (host ref -> string set) map
// - The potential hosts and their corresponding recommendations or errors
func (client *XenAPIClient) VM_retrieve_wlb_recommendations(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.retrieve_wlb_recommendations", session_id, vm)
}

// VM_assert_agile
//
// Returns an error if the VM is not considered agile e.g. because it is tied to a resource local to a host
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
//
// returns:
// - void
func (client *XenAPIClient) VM_assert_agile(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_agile", session_id, self)
}

// VM_create_new_blob
//
// Create a placeholder for a named binary blob of data that is associated with this VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM
// - name, string, The name associated with the blob
// - mime_type, string, The mime type for the data. Empty string translates to application/octet-stream
// - public, bool, True if the blob should be publicly available
//
// returns:
// - blob ref
// - The reference of the blob, needed for populating its data
func (client *XenAPIClient) VM_create_new_blob(session_id interface{}, vm interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("VM.create_new_blob", session_id, vm, name, mime_type, public)
}

// VM_assert_can_boot_here
//
// Returns an error if the VM could not boot on this host for some reason
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - host, host ref, The host
//
// returns:
// - void
func (client *XenAPIClient) VM_assert_can_boot_here(session_id interface{}, self interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_can_boot_here", session_id, self, host)
}

// VM_get_possible_hosts
//
// Return the list of hosts on which this VM may run.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM
//
// returns:
// - host ref set
// - The possible hosts
func (client *XenAPIClient) VM_get_possible_hosts(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_possible_hosts", session_id, vm)
}

// VM_get_allowed_VIF_devices
//
// Returns a list of the allowed values that a VIF device field can take
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to query
//
// returns:
// - string set
// - The allowed values
func (client *XenAPIClient) VM_get_allowed_VIF_devices(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_allowed_VIF_devices", session_id, vm)
}

// VM_get_allowed_VBD_devices
//
// Returns a list of the allowed values that a VBD device field can take
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to query
//
// returns:
// - string set
// - The allowed values
func (client *XenAPIClient) VM_get_allowed_VBD_devices(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_allowed_VBD_devices", session_id, vm)
}

// VM_update_allowed_operations
//
// Recomputes the list of acceptable operations
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VM_update_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.update_allowed_operations", session_id, self)
}

// VM_assert_operation_valid
//
// Check to see whether this operation is acceptable in the current state of the system, raising an error if the operation is invalid for some reason
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - op, enum vm_operations, proposed operation
//
// returns:
// - void
func (client *XenAPIClient) VM_assert_operation_valid(session_id interface{}, self interface{}, op interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.assert_operation_valid", session_id, self, op)
}

// VM_forget_data_source_archives
//
// Forget the recorded statistics related to the specified data source
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - data_source, string, The data source whose archives are to be forgotten
//
// returns:
// - void
func (client *XenAPIClient) VM_forget_data_source_archives(session_id interface{}, self interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("VM.forget_data_source_archives", session_id, self, data_source)
}

// VM_query_data_source
//
// Query the latest value of the specified data source
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - data_source, string, The data source to query
//
// returns:
// - float
// - The latest value, averaged over the last 5 seconds
func (client *XenAPIClient) VM_query_data_source(session_id interface{}, self interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("VM.query_data_source", session_id, self, data_source)
}

// VM_record_data_source
//
// Start recording the specified data source
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - data_source, string, The data source to record
//
// returns:
// - void
func (client *XenAPIClient) VM_record_data_source(session_id interface{}, self interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("VM.record_data_source", session_id, self, data_source)
}

// VM_get_data_sources
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to interrogate
//
// returns:
// - data_source record set
// - A set of data sources
func (client *XenAPIClient) VM_get_data_sources(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_data_sources", session_id, self)
}

// VM_get_boot_record
//
// Returns a record describing the VM's dynamic state, initialised when the VM boots and updated to reflect runtime configuration changes e.g. CPU hotplug
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM whose boot-time state to return
//
// returns:
// - VM record
// - A record describing the VM
func (client *XenAPIClient) VM_get_boot_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_boot_record", session_id, self)
}

// VM_assert_can_migrate
//
// Assert whether a VM can be migrated to the specified destination.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM
// - dest, (string -> string) map, The result of a VM.migrate_receive call.
// - live, bool, Live migration
// - vdi_map, (VDI ref -> SR ref) map, Map of source VDI to destination SR
// - vif_map, (VIF ref -> network ref) map, Map of source VIF to destination network
// - options, (string -> string) map, Other parameters
//
// returns:
// - void
func (client *XenAPIClient) VM_assert_can_migrate(session_id interface{}, vm interface{}, dest map[string]string, live bool, vdi_map interface{}, vif_map interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.assert_can_migrate", session_id, vm, dest, live, vdi_map, vif_map, options)
}

// VM_migrate_send
//
// Migrate the VM to another host.  This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM
// - dest, (string -> string) map, The result of a Host.migrate_receive call.
// - live, bool, Live migration
// - vdi_map, (VDI ref -> SR ref) map, Map of source VDI to destination SR
// - vif_map, (VIF ref -> network ref) map, Map of source VIF to destination network
// - options, (string -> string) map, Other parameters
//
// returns:
// - void
func (client *XenAPIClient) VM_migrate_send(session_id interface{}, vm interface{}, dest map[string]string, live bool, vdi_map interface{}, vif_map interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.migrate_send", session_id, vm, dest, live, vdi_map, vif_map, options)
}

// VM_maximise_memory
//
// Returns the maximum amount of guest memory which will fit, together with overheads, in the supplied amount of physical memory. If 'exact' is true then an exact calculation is performed using the VM's current settings. If 'exact' is false then a more conservative approximation is used
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - total, int, Total amount of physical RAM to fit within
// - approximate, bool, If false the limit is calculated with the guest's current exact configuration. Otherwise a more approximate calculation is performed
//
// returns:
// - int
// - The maximum possible static-max
func (client *XenAPIClient) VM_maximise_memory(session_id interface{}, self interface{}, total interface{}, approximate bool) (i interface{}, err error) {
	return client.RPCCall("VM.maximise_memory", session_id, self, total, approximate)
}

// VM_send_trigger
//
// Send the named trigger to this VM.  This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM
// - trigger, string, The trigger to send
//
// returns:
// - void
func (client *XenAPIClient) VM_send_trigger(session_id interface{}, vm interface{}, trigger string) (i interface{}, err error) {
	return client.RPCCall("VM.send_trigger", session_id, vm, trigger)
}

// VM_send_sysrq
//
// Send the given key as a sysrq to this VM.  The key is specified as a single character (a String of length 1).  This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM
// - key, string, The key to send
//
// returns:
// - void
func (client *XenAPIClient) VM_send_sysrq(session_id interface{}, vm interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.send_sysrq", session_id, vm, key)
}

// VM_set_VCPUs_at_startup
//
// Set the number of startup VCPUs for a halted VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, int, The new maximum number of VCPUs
//
// returns:
// - void
func (client *XenAPIClient) VM_set_VCPUs_at_startup(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_at_startup", session_id, self, value)
}

// VM_set_VCPUs_max
//
// Set the maximum number of VCPUs for a halted VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, int, The new maximum number of VCPUs
//
// returns:
// - void
func (client *XenAPIClient) VM_set_VCPUs_max(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_max", session_id, self, value)
}

// VM_set_shadow_multiplier_live
//
// Set the shadow memory multiplier on a running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - multiplier, float, The new shadow memory multiplier to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_shadow_multiplier_live(session_id interface{}, self interface{}, multiplier interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_shadow_multiplier_live", session_id, self, multiplier)
}

// VM_set_HVM_shadow_multiplier
//
// Set the shadow memory multiplier on a halted VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, float, The new shadow memory multiplier to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_HVM_shadow_multiplier(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_HVM_shadow_multiplier", session_id, self, value)
}

// VM_get_cooperative
//
// Return true if the VM is currently 'co-operative' i.e. is expected to reach a balloon target and actually has done
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
//
// returns:
// - bool
// - true if the VM is currently 'co-operative'; false otherwise
func (client *XenAPIClient) VM_get_cooperative(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_cooperative", session_id, self)
}

// VM_wait_memory_target_live
//
// Wait for a running VM to reach its current memory target
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
//
// returns:
// - void
func (client *XenAPIClient) VM_wait_memory_target_live(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.wait_memory_target_live", session_id, self)
}

// VM_set_memory_target_live
//
// Set the memory target for a running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - target, int, The target in bytes
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_target_live(session_id interface{}, self interface{}, target interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_target_live", session_id, self, target)
}

// VM_set_memory_limits
//
// Set the memory limits of this VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - static_min, int, The new value of memory_static_min.
// - static_max, int, The new value of memory_static_max.
// - dynamic_min, int, The new value of memory_dynamic_min.
// - dynamic_max, int, The new value of memory_dynamic_max.
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_limits(session_id interface{}, self interface{}, static_min interface{}, static_max interface{}, dynamic_min interface{}, dynamic_max interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_limits", session_id, self, static_min, static_max, dynamic_min, dynamic_max)
}

// VM_set_memory_static_range
//
// Set the static (ie boot-time) range of virtual memory that the VM is allowed to use.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - min, int, The new minimum value
// - max, int, The new maximum value
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_static_range(session_id interface{}, self interface{}, min interface{}, max interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_static_range", session_id, self, min, max)
}

// VM_set_memory_static_min
//
// Set the value of the memory_static_min field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to modify
// - value, int, The new value of memory_static_min
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_static_min(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_static_min", session_id, self, value)
}

// VM_set_memory_static_max
//
// Set the value of the memory_static_max field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to modify
// - value, int, The new value of memory_static_max
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_static_max(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_static_max", session_id, self, value)
}

// VM_set_memory_dynamic_range
//
// Set the minimum and maximum amounts of physical memory the VM is allowed to use.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - min, int, The new minimum value
// - max, int, The new maximum value
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_dynamic_range(session_id interface{}, self interface{}, min interface{}, max interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_dynamic_range", session_id, self, min, max)
}

// VM_set_memory_dynamic_min
//
// Set the value of the memory_dynamic_min field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to modify
// - value, int, The new value of memory_dynamic_min
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_dynamic_min(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_dynamic_min", session_id, self, value)
}

// VM_set_memory_dynamic_max
//
// Set the value of the memory_dynamic_max field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM to modify
// - value, int, The new value of memory_dynamic_max
//
// returns:
// - void
func (client *XenAPIClient) VM_set_memory_dynamic_max(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_memory_dynamic_max", session_id, self, value)
}

// VM_compute_memory_overhead
//
// Computes the virtualization memory overhead of a VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM for which to compute the memory overhead
//
// returns:
// - int
// - the virtualization memory overhead of the VM.
func (client *XenAPIClient) VM_compute_memory_overhead(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.compute_memory_overhead", session_id, vm)
}

// VM_set_ha_always_run
//
// Set the value of the ha_always_run
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, bool, The value
//
// returns:
// - void
func (client *XenAPIClient) VM_set_ha_always_run(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VM.set_ha_always_run", session_id, self, value)
}

// VM_set_ha_restart_priority
//
// Set the value of the ha_restart_priority field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - value, string, The value
//
// returns:
// - void
func (client *XenAPIClient) VM_set_ha_restart_priority(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_ha_restart_priority", session_id, self, value)
}

// VM_add_to_VCPUs_params_live
//
// Add the given key-value pair to VM.VCPUs_params, and apply that value on the running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - key, string, The key
// - value, string, The value
//
// returns:
// - void
func (client *XenAPIClient) VM_add_to_VCPUs_params_live(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_VCPUs_params_live", session_id, self, key, value)
}

// VM_set_VCPUs_number_live
//
// Set the number of VCPUs for a running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, The VM
// - nvcpu, int, The number of VCPUs
//
// returns:
// - void
func (client *XenAPIClient) VM_set_VCPUs_number_live(session_id interface{}, self interface{}, nvcpu interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_number_live", session_id, self, nvcpu)
}

// VM_pool_migrate
//
// Migrate a VM to another Host. This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to migrate
// - host, host ref, The target host
// - options, (string -> string) map, Extra configuration operations
//
// returns:
// - void
func (client *XenAPIClient) VM_pool_migrate(session_id interface{}, vm interface{}, host interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.pool_migrate", session_id, vm, host, options)
}

// VM_resume_on
//
// Awaken the specified VM and resume it on a particular Host.  This can only be called when the specified VM is in the Suspended state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to resume
// - host, host ref, The Host on which to resume the VM
// - start_paused, bool, Resume VM in paused state if set to true.
// - force, bool, Attempt to force the VM to resume. If this flag is false then the VM may fail pre-resume safety checks (e.g. if the CPU the VM was running on looks substantially different to the current one)
//
// returns:
// - void
func (client *XenAPIClient) VM_resume_on(session_id interface{}, vm interface{}, host interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.resume_on", session_id, vm, host, start_paused, force)
}

// VM_resume
//
// Awaken the specified VM and resume it.  This can only be called when the specified VM is in the Suspended state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to resume
// - start_paused, bool, Resume VM in paused state if set to true.
// - force, bool, Attempt to force the VM to resume. If this flag is false then the VM may fail pre-resume safety checks (e.g. if the CPU the VM was running on looks substantially different to the current one)
//
// returns:
// - void
func (client *XenAPIClient) VM_resume(session_id interface{}, vm interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.resume", session_id, vm, start_paused, force)
}

// VM_suspend
//
// Suspend the specified VM to disk.  This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to suspend
//
// returns:
// - void
func (client *XenAPIClient) VM_suspend(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.suspend", session_id, vm)
}

// VM_hard_reboot
//
// Stop executing the specified VM without attempting a clean shutdown and immediately restart the VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to reboot
//
// returns:
// - void
func (client *XenAPIClient) VM_hard_reboot(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.hard_reboot", session_id, vm)
}

// VM_power_state_reset
//
// Reset the power-state of the VM to halted in the database only. (Used to recover from slave failures in pooling scenarios by resetting the power-states of VMs running on dead slaves to halted.) This is a potentially dangerous operation; use with care.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to reset
//
// returns:
// - void
func (client *XenAPIClient) VM_power_state_reset(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.power_state_reset", session_id, vm)
}

// VM_hard_shutdown
//
// Stop executing the specified VM without attempting a clean shutdown.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to destroy
//
// returns:
// - void
func (client *XenAPIClient) VM_hard_shutdown(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.hard_shutdown", session_id, vm)
}

// VM_clean_reboot
//
// Attempt to cleanly shutdown the specified VM (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to shutdown
//
// returns:
// - void
func (client *XenAPIClient) VM_clean_reboot(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.clean_reboot", session_id, vm)
}

// VM_shutdown
//
// Attempts to first clean shutdown a VM and if it should fail then perform a hard shutdown on it.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to shutdown
//
// returns:
// - void
func (client *XenAPIClient) VM_shutdown(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.shutdown", session_id, vm)
}

// VM_clean_shutdown
//
// Attempt to cleanly shutdown the specified VM. (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to shutdown
//
// returns:
// - void
func (client *XenAPIClient) VM_clean_shutdown(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.clean_shutdown", session_id, vm)
}

// VM_unpause
//
// Resume the specified VM. This can only be called when the specified VM is in the Paused state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to unpause
//
// returns:
// - void
func (client *XenAPIClient) VM_unpause(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.unpause", session_id, vm)
}

// VM_pause
//
// Pause the specified VM. This can only be called when the specified VM is in the Running state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to pause
//
// returns:
// - void
func (client *XenAPIClient) VM_pause(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.pause", session_id, vm)
}

// VM_start_on
//
// Start the specified VM on a particular host.  This function can only be called with the VM is in the Halted State.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to start
// - host, host ref, The Host on which to start the VM
// - start_paused, bool, Instantiate VM in paused state if set to true.
// - force, bool, Attempt to force the VM to start. If this flag is false then the VM may fail pre-boot safety checks (e.g. if the CPU the VM last booted on looks substantially different to the current one)
//
// returns:
// - void
func (client *XenAPIClient) VM_start_on(session_id interface{}, vm interface{}, host interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.start_on", session_id, vm, host, start_paused, force)
}

// VM_start
//
// Start the specified VM.  This function can only be called with the VM is in the Halted State.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to start
// - start_paused, bool, Instantiate VM in paused state if set to true.
// - force, bool, Attempt to force the VM to start. If this flag is false then the VM may fail pre-boot safety checks (e.g. if the CPU the VM last booted on looks substantially different to the current one)
//
// returns:
// - void
func (client *XenAPIClient) VM_start(session_id interface{}, vm interface{}, start_paused bool, force bool) (i interface{}, err error) {
	return client.RPCCall("VM.start", session_id, vm, start_paused, force)
}

// VM_provision
//
// Inspects the disk configuration contained within the VM's other_config, creates VDIs and VBDs and then executes any applicable post-install script.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to be provisioned
//
// returns:
// - void
func (client *XenAPIClient) VM_provision(session_id interface{}, vm interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.provision", session_id, vm)
}

// VM_checkpoint
//
// Checkpoints the specified VM, making a new VM. Checkpoint automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write) and saves the memory image as well.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to be checkpointed
// - new_name, string, The name of the checkpointed VM
//
// returns:
// - VM ref
// - The reference of the newly created VM.
func (client *XenAPIClient) VM_checkpoint(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.checkpoint", session_id, vm, new_name)
}

// VM_revert
//
// Reverts the specified VM to a previous state.
//
// params:
// - session_id, session ref, Reference to a valid session
// - snapshot, VM ref, The snapshotted state that we revert to
//
// returns:
// - void
func (client *XenAPIClient) VM_revert(session_id interface{}, snapshot interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.revert", session_id, snapshot)
}

// VM_copy
//
// Copied the specified VM, making a new VM. Unlike clone, copy does not exploits the capabilities of the underlying storage repository in which the VM's disk images are stored. Instead, copy guarantees that the disk images of the newly created VM will be 'full disks' - i.e. not part of a CoW chain.  This function can only be called when the VM is in the Halted State.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to be copied
// - new_name, string, The name of the copied VM
// - sr, SR ref, An SR to copy all the VM's disks into (if an invalid reference then it uses the existing SRs)
//
// returns:
// - VM ref
// - The reference of the newly created VM.
func (client *XenAPIClient) VM_copy(session_id interface{}, vm interface{}, new_name string, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.copy", session_id, vm, new_name, sr)
}

// VM_clone
//
// Clones the specified VM, making a new VM. Clone automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).   This function can only be called when the VM is in the Halted State.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to be cloned
// - new_name, string, The name of the cloned VM
//
// returns:
// - VM ref
// - The reference of the newly created VM.
func (client *XenAPIClient) VM_clone(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.clone", session_id, vm, new_name)
}

// VM_snapshot_with_quiesce
//
// Snapshots the specified VM with quiesce, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to be snapshotted
// - new_name, string, The name of the snapshotted VM
//
// returns:
// - VM ref
// - The reference of the newly created VM.
func (client *XenAPIClient) VM_snapshot_with_quiesce(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.snapshot_with_quiesce", session_id, vm, new_name)
}

// VM_snapshot
//
// Snapshots the specified VM, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// params:
// - session_id, session ref, Reference to a valid session
// - vm, VM ref, The VM to be snapshotted
// - new_name, string, The name of the snapshotted VM
//
// returns:
// - VM ref
// - The reference of the newly created VM.
func (client *XenAPIClient) VM_snapshot(session_id interface{}, vm interface{}, new_name string) (i interface{}, err error) {
	return client.RPCCall("VM.snapshot", session_id, vm, new_name)
}

// VM_set_suspend_SR
//
// Set the suspend_SR field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, SR ref, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_suspend_SR(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_suspend_SR", session_id, self, value)
}

// VM_remove_from_blocked_operations
//
// Remove the given key and its corresponding value from the blocked_operations field of the given VM.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, enum vm_operations, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_remove_from_blocked_operations(session_id interface{}, self interface{}, key interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_blocked_operations", session_id, self, key)
}

// VM_add_to_blocked_operations
//
// Add the given key-value pair to the blocked_operations field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, enum vm_operations, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_add_to_blocked_operations(session_id interface{}, self interface{}, key interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_blocked_operations", session_id, self, key, value)
}

// VM_set_blocked_operations
//
// Set the blocked_operations field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, (enum vm_operations -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_blocked_operations(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_blocked_operations", session_id, self, value)
}

// VM_remove_tags
//
// Remove the given value from the tags field of the given VM.  If the value is not in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, Value to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_tags", session_id, self, value)
}

// VM_add_tags
//
// Add the given value to the tags field of the given VM.  If the value is already in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_tags", session_id, self, value)
}

// VM_set_tags
//
// Set the tags field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string set, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_tags", session_id, self, value)
}

// VM_remove_from_xenstore_data
//
// Remove the given key and its corresponding value from the xenstore_data field of the given VM.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_remove_from_xenstore_data(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_xenstore_data", session_id, self, key)
}

// VM_add_to_xenstore_data
//
// Add the given key-value pair to the xenstore_data field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_add_to_xenstore_data(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_xenstore_data", session_id, self, key, value)
}

// VM_set_xenstore_data
//
// Set the xenstore_data field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_xenstore_data(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_xenstore_data", session_id, self, value)
}

// VM_set_recommendations
//
// Set the recommendations field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_recommendations(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_recommendations", session_id, self, value)
}

// VM_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VM.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_other_config", session_id, self, key)
}

// VM_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_other_config", session_id, self, key, value)
}

// VM_set_other_config
//
// Set the other_config field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_other_config", session_id, self, value)
}

// VM_set_PCI_bus
//
// Set the PCI_bus field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_PCI_bus(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PCI_bus", session_id, self, value)
}

// VM_remove_from_platform
//
// Remove the given key and its corresponding value from the platform field of the given VM.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_remove_from_platform(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_platform", session_id, self, key)
}

// VM_add_to_platform
//
// Add the given key-value pair to the platform field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_add_to_platform(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_platform", session_id, self, key, value)
}

// VM_set_platform
//
// Set the platform field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_platform(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_platform", session_id, self, value)
}

// VM_remove_from_HVM_boot_params
//
// Remove the given key and its corresponding value from the HVM/boot_params field of the given VM.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_remove_from_HVM_boot_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_HVM_boot_params", session_id, self, key)
}

// VM_add_to_HVM_boot_params
//
// Add the given key-value pair to the HVM/boot_params field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_add_to_HVM_boot_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_HVM_boot_params", session_id, self, key, value)
}

// VM_set_HVM_boot_params
//
// Set the HVM/boot_params field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_HVM_boot_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_HVM_boot_params", session_id, self, value)
}

// VM_set_HVM_boot_policy
//
// Set the HVM/boot_policy field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_HVM_boot_policy(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_HVM_boot_policy", session_id, self, value)
}

// VM_set_PV_legacy_args
//
// Set the PV/legacy_args field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_PV_legacy_args(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_legacy_args", session_id, self, value)
}

// VM_set_PV_bootloader_args
//
// Set the PV/bootloader_args field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_PV_bootloader_args(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_bootloader_args", session_id, self, value)
}

// VM_set_PV_args
//
// Set the PV/args field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_PV_args(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_args", session_id, self, value)
}

// VM_set_PV_ramdisk
//
// Set the PV/ramdisk field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_PV_ramdisk(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_ramdisk", session_id, self, value)
}

// VM_set_PV_kernel
//
// Set the PV/kernel field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_PV_kernel(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_kernel", session_id, self, value)
}

// VM_set_PV_bootloader
//
// Set the PV/bootloader field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_PV_bootloader(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_PV_bootloader", session_id, self, value)
}

// VM_set_actions_after_crash
//
// Set the actions/after_crash field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, enum on_crash_behaviour, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_actions_after_crash(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_actions_after_crash", session_id, self, value)
}

// VM_set_actions_after_reboot
//
// Set the actions/after_reboot field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, enum on_normal_exit, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_actions_after_reboot(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_actions_after_reboot", session_id, self, value)
}

// VM_set_actions_after_shutdown
//
// Set the actions/after_shutdown field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, enum on_normal_exit, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_actions_after_shutdown(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_actions_after_shutdown", session_id, self, value)
}

// VM_remove_from_VCPUs_params
//
// Remove the given key and its corresponding value from the VCPUs/params field of the given VM.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_remove_from_VCPUs_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM.remove_from_VCPUs_params", session_id, self, key)
}

// VM_add_to_VCPUs_params
//
// Add the given key-value pair to the VCPUs/params field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_add_to_VCPUs_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM.add_to_VCPUs_params", session_id, self, key, value)
}

// VM_set_VCPUs_params
//
// Set the VCPUs/params field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_VCPUs_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM.set_VCPUs_params", session_id, self, value)
}

// VM_set_affinity
//
// Set the affinity field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, host ref, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_affinity(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_affinity", session_id, self, value)
}

// VM_set_is_a_template
//
// Set the is_a_template field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_is_a_template(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VM.set_is_a_template", session_id, self, value)
}

// VM_set_user_version
//
// Set the user_version field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, int, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_user_version(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.set_user_version", session_id, self, value)
}

// VM_set_name_description
//
// Set the name/description field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_name_description", session_id, self, value)
}

// VM_set_name_label
//
// Set the name/label field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VM.set_name_label", session_id, self, value)
}

// VM_get_generation_id
//
// Get the generation_id field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_generation_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_generation_id", session_id, self)
}

// VM_get_version
//
// Get the version field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_version", session_id, self)
}

// VM_get_suspend_SR
//
// Get the suspend_SR field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) VM_get_suspend_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_suspend_SR", session_id, self)
}

// VM_get_attached_PCIs
//
// Get the attached_PCIs field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - PCI ref set
// - value of the field
func (client *XenAPIClient) VM_get_attached_PCIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_attached_PCIs", session_id, self)
}

// VM_get_VGPUs
//
// Get the VGPUs field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VGPU ref set
// - value of the field
func (client *XenAPIClient) VM_get_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VGPUs", session_id, self)
}

// VM_get_order
//
// Get the order field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_order(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_order", session_id, self)
}

// VM_get_shutdown_delay
//
// Get the shutdown_delay field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_shutdown_delay(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_shutdown_delay", session_id, self)
}

// VM_get_start_delay
//
// Get the start_delay field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_start_delay(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_start_delay", session_id, self)
}

// VM_get_appliance
//
// Get the appliance field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM_appliance ref
// - value of the field
func (client *XenAPIClient) VM_get_appliance(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_appliance", session_id, self)
}

// VM_get_is_snapshot_from_vmpp
//
// Get the is_snapshot_from_vmpp field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VM_get_is_snapshot_from_vmpp(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_snapshot_from_vmpp", session_id, self)
}

// VM_get_protection_policy
//
// Get the protection_policy field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VMPP ref
// - value of the field
func (client *XenAPIClient) VM_get_protection_policy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_protection_policy", session_id, self)
}

// VM_get_bios_strings
//
// Get the bios_strings field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_bios_strings(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_bios_strings", session_id, self)
}

// VM_get_children
//
// Get the children field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM ref set
// - value of the field
func (client *XenAPIClient) VM_get_children(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_children", session_id, self)
}

// VM_get_parent
//
// Get the parent field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) VM_get_parent(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_parent", session_id, self)
}

// VM_get_snapshot_metadata
//
// Get the snapshot_metadata field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_snapshot_metadata(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_metadata", session_id, self)
}

// VM_get_snapshot_info
//
// Get the snapshot_info field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_snapshot_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_info", session_id, self)
}

// VM_get_blocked_operations
//
// Get the blocked_operations field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (enum vm_operations -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_blocked_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_blocked_operations", session_id, self)
}

// VM_get_tags
//
// Get the tags field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VM_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_tags", session_id, self)
}

// VM_get_blobs
//
// Get the blobs field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> blob ref) map
// - value of the field
func (client *XenAPIClient) VM_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_blobs", session_id, self)
}

// VM_get_transportable_snapshot_id
//
// Get the transportable_snapshot_id field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_transportable_snapshot_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_transportable_snapshot_id", session_id, self)
}

// VM_get_snapshot_time
//
// Get the snapshot_time field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VM_get_snapshot_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_time", session_id, self)
}

// VM_get_snapshots
//
// Get the snapshots field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM ref set
// - value of the field
func (client *XenAPIClient) VM_get_snapshots(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshots", session_id, self)
}

// VM_get_snapshot_of
//
// Get the snapshot_of field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) VM_get_snapshot_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_snapshot_of", session_id, self)
}

// VM_get_is_a_snapshot
//
// Get the is_a_snapshot field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VM_get_is_a_snapshot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_a_snapshot", session_id, self)
}

// VM_get_ha_restart_priority
//
// Get the ha_restart_priority field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_ha_restart_priority(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_ha_restart_priority", session_id, self)
}

// VM_get_ha_always_run
//
// Get the ha_always_run field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VM_get_ha_always_run(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_ha_always_run", session_id, self)
}

// VM_get_xenstore_data
//
// Get the xenstore_data field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_xenstore_data(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_xenstore_data", session_id, self)
}

// VM_get_recommendations
//
// Get the recommendations field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_recommendations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_recommendations", session_id, self)
}

// VM_get_last_booted_record
//
// Get the last_booted_record field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_last_booted_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_last_booted_record", session_id, self)
}

// VM_get_guest_metrics
//
// Get the guest_metrics field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM_guest_metrics ref
// - value of the field
func (client *XenAPIClient) VM_get_guest_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_guest_metrics", session_id, self)
}

// VM_get_metrics
//
// Get the metrics field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM_metrics ref
// - value of the field
func (client *XenAPIClient) VM_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_metrics", session_id, self)
}

// VM_get_is_control_domain
//
// Get the is_control_domain field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VM_get_is_control_domain(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_control_domain", session_id, self)
}

// VM_get_last_boot_CPU_flags
//
// Get the last_boot_CPU_flags field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_last_boot_CPU_flags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_last_boot_CPU_flags", session_id, self)
}

// VM_get_domarch
//
// Get the domarch field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_domarch(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_domarch", session_id, self)
}

// VM_get_domid
//
// Get the domid field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_domid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_domid", session_id, self)
}

// VM_get_other_config
//
// Get the other_config field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_other_config", session_id, self)
}

// VM_get_PCI_bus
//
// Get the PCI_bus field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_PCI_bus(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PCI_bus", session_id, self)
}

// VM_get_platform
//
// Get the platform field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_platform(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_platform", session_id, self)
}

// VM_get_HVM_shadow_multiplier
//
// Get the HVM/shadow_multiplier field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) VM_get_HVM_shadow_multiplier(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_HVM_shadow_multiplier", session_id, self)
}

// VM_get_HVM_boot_params
//
// Get the HVM/boot_params field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_HVM_boot_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_HVM_boot_params", session_id, self)
}

// VM_get_HVM_boot_policy
//
// Get the HVM/boot_policy field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_HVM_boot_policy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_HVM_boot_policy", session_id, self)
}

// VM_get_PV_legacy_args
//
// Get the PV/legacy_args field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_PV_legacy_args(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_legacy_args", session_id, self)
}

// VM_get_PV_bootloader_args
//
// Get the PV/bootloader_args field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_PV_bootloader_args(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_bootloader_args", session_id, self)
}

// VM_get_PV_args
//
// Get the PV/args field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_PV_args(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_args", session_id, self)
}

// VM_get_PV_ramdisk
//
// Get the PV/ramdisk field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_PV_ramdisk(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_ramdisk", session_id, self)
}

// VM_get_PV_kernel
//
// Get the PV/kernel field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_PV_kernel(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_kernel", session_id, self)
}

// VM_get_PV_bootloader
//
// Get the PV/bootloader field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_PV_bootloader(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_PV_bootloader", session_id, self)
}

// VM_get_VTPMs
//
// Get the VTPMs field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VTPM ref set
// - value of the field
func (client *XenAPIClient) VM_get_VTPMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VTPMs", session_id, self)
}

// VM_get_crash_dumps
//
// Get the crash_dumps field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - crashdump ref set
// - value of the field
func (client *XenAPIClient) VM_get_crash_dumps(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_crash_dumps", session_id, self)
}

// VM_get_VBDs
//
// Get the VBDs field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VBD ref set
// - value of the field
func (client *XenAPIClient) VM_get_VBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VBDs", session_id, self)
}

// VM_get_VIFs
//
// Get the VIFs field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VIF ref set
// - value of the field
func (client *XenAPIClient) VM_get_VIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VIFs", session_id, self)
}

// VM_get_consoles
//
// Get the consoles field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - console ref set
// - value of the field
func (client *XenAPIClient) VM_get_consoles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_consoles", session_id, self)
}

// VM_get_actions_after_crash
//
// Get the actions/after_crash field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - enum on_crash_behaviour
// - value of the field
func (client *XenAPIClient) VM_get_actions_after_crash(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_actions_after_crash", session_id, self)
}

// VM_get_actions_after_reboot
//
// Get the actions/after_reboot field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - enum on_normal_exit
// - value of the field
func (client *XenAPIClient) VM_get_actions_after_reboot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_actions_after_reboot", session_id, self)
}

// VM_get_actions_after_shutdown
//
// Get the actions/after_shutdown field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - enum on_normal_exit
// - value of the field
func (client *XenAPIClient) VM_get_actions_after_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_actions_after_shutdown", session_id, self)
}

// VM_get_VCPUs_at_startup
//
// Get the VCPUs/at_startup field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_VCPUs_at_startup(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VCPUs_at_startup", session_id, self)
}

// VM_get_VCPUs_max
//
// Get the VCPUs/max field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_VCPUs_max(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VCPUs_max", session_id, self)
}

// VM_get_VCPUs_params
//
// Get the VCPUs/params field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_get_VCPUs_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_VCPUs_params", session_id, self)
}

// VM_get_memory_static_min
//
// Get the memory/static_min field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_memory_static_min(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_static_min", session_id, self)
}

// VM_get_memory_dynamic_min
//
// Get the memory/dynamic_min field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_memory_dynamic_min(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_dynamic_min", session_id, self)
}

// VM_get_memory_dynamic_max
//
// Get the memory/dynamic_max field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_memory_dynamic_max(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_dynamic_max", session_id, self)
}

// VM_get_memory_static_max
//
// Get the memory/static_max field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_memory_static_max(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_static_max", session_id, self)
}

// VM_get_memory_target
//
// Get the memory/target field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_memory_target(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_target", session_id, self)
}

// VM_get_memory_overhead
//
// Get the memory/overhead field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_memory_overhead(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_memory_overhead", session_id, self)
}

// VM_get_affinity
//
// Get the affinity field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) VM_get_affinity(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_affinity", session_id, self)
}

// VM_get_resident_on
//
// Get the resident_on field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) VM_get_resident_on(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_resident_on", session_id, self)
}

// VM_get_suspend_VDI
//
// Get the suspend_VDI field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VDI ref
// - value of the field
func (client *XenAPIClient) VM_get_suspend_VDI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_suspend_VDI", session_id, self)
}

// VM_get_is_a_template
//
// Get the is_a_template field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VM_get_is_a_template(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_is_a_template", session_id, self)
}

// VM_get_user_version
//
// Get the user_version field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_get_user_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_user_version", session_id, self)
}

// VM_get_name_description
//
// Get the name/description field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_name_description", session_id, self)
}

// VM_get_name_label
//
// Get the name/label field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_name_label", session_id, self)
}

// VM_get_power_state
//
// Get the power_state field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - enum vm_power_state
// - value of the field
func (client *XenAPIClient) VM_get_power_state(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_power_state", session_id, self)
}

// VM_get_current_operations
//
// Get the current_operations field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - (string -> enum vm_operations) map
// - value of the field
func (client *XenAPIClient) VM_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_current_operations", session_id, self)
}

// VM_get_allowed_operations
//
// Get the allowed_operations field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - enum vm_operations set
// - value of the field
func (client *XenAPIClient) VM_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_allowed_operations", session_id, self)
}

// VM_get_uuid
//
// Get the uuid field of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_uuid", session_id, self)
}

// VM_get_by_name_label
//
// Get all the VM instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - VM ref set
// - references to objects with matching names
func (client *XenAPIClient) VM_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VM.get_by_name_label", session_id, label)
}

// VM_destroy
//
// Destroy the specified VM.  The VM is completely removed from the system.  This function can only be called when the VM is in the Halted State.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VM_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.destroy", session_id, self)
}

// VM_create
//
// Create a new VM instance, and return its handle.
// The constructor args are: name_label, name_description, user_version*, is_a_template*, affinity*, memory_target, memory_static_max*, memory_dynamic_max*, memory_dynamic_min*, memory_static_min*, VCPUs_params*, VCPUs_max*, VCPUs_at_startup*, actions_after_shutdown*, actions_after_reboot*, actions_after_crash*, PV_bootloader*, PV_kernel*, PV_ramdisk*, PV_args*, PV_bootloader_args*, PV_legacy_args*, HVM_boot_policy*, HVM_boot_params*, HVM_shadow_multiplier, platform*, PCI_bus*, other_config*, recommendations*, xenstore_data, ha_always_run, ha_restart_priority, tags, blocked_operations, protection_policy, is_snapshot_from_vmpp, appliance, start_delay, shutdown_delay, order, suspend_SR, version, generation_id (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, VM record, All constructor arguments
//
// returns:
// - VM ref
// - reference to the newly created object
func (client *XenAPIClient) VM_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.create", session_id, args)
}

// VM_get_by_uuid
//
// Get a reference to the VM instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VM ref
// - reference to the object
func (client *XenAPIClient) VM_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VM.get_by_uuid", session_id, uuid)
}

// VM_get_record
//
// Get a record containing the current state of the given VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM ref, reference to the object
//
// returns:
// - VM record
// - all fields from the object
func (client *XenAPIClient) VM_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM.get_record", session_id, self)
}

// VM_metrics_get_all_records
//
// Return a map of VM_metrics references to VM_metrics records for all VM_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VM_metrics ref -> VM_metrics record) map
// - records of all objects
func (client *XenAPIClient) VM_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_all_records", session_id)
}

// VM_metrics_get_all
//
// Return a list of all the VM_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VM_metrics ref set
// - references to all objects
func (client *XenAPIClient) VM_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_all", session_id)
}

// VM_metrics_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VM_metrics.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.remove_from_other_config", session_id, self, key)
}

// VM_metrics_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.add_to_other_config", session_id, self, key, value)
}

// VM_metrics_set_other_config
//
// Set the other_config field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.set_other_config", session_id, self, value)
}

// VM_metrics_get_other_config
//
// Get the other_config field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_other_config", session_id, self)
}

// VM_metrics_get_last_updated
//
// Get the last_updated field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VM_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_last_updated", session_id, self)
}

// VM_metrics_get_install_time
//
// Get the install_time field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VM_metrics_get_install_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_install_time", session_id, self)
}

// VM_metrics_get_start_time
//
// Get the start_time field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VM_metrics_get_start_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_start_time", session_id, self)
}

// VM_metrics_get_state
//
// Get the state field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VM_metrics_get_state(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_state", session_id, self)
}

// VM_metrics_get_VCPUs_flags
//
// Get the VCPUs/flags field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - (int -> string set) map
// - value of the field
func (client *XenAPIClient) VM_metrics_get_VCPUs_flags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_flags", session_id, self)
}

// VM_metrics_get_VCPUs_params
//
// Get the VCPUs/params field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_metrics_get_VCPUs_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_params", session_id, self)
}

// VM_metrics_get_VCPUs_CPU
//
// Get the VCPUs/CPU field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - (int -> int) map
// - value of the field
func (client *XenAPIClient) VM_metrics_get_VCPUs_CPU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_CPU", session_id, self)
}

// VM_metrics_get_VCPUs_utilisation
//
// Get the VCPUs/utilisation field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - (int -> float) map
// - value of the field
func (client *XenAPIClient) VM_metrics_get_VCPUs_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_utilisation", session_id, self)
}

// VM_metrics_get_VCPUs_number
//
// Get the VCPUs/number field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_metrics_get_VCPUs_number(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_VCPUs_number", session_id, self)
}

// VM_metrics_get_memory_actual
//
// Get the memory/actual field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VM_metrics_get_memory_actual(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_memory_actual", session_id, self)
}

// VM_metrics_get_uuid
//
// Get the uuid field of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_uuid", session_id, self)
}

// VM_metrics_get_by_uuid
//
// Get a reference to the VM_metrics instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VM_metrics ref
// - reference to the object
func (client *XenAPIClient) VM_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_by_uuid", session_id, uuid)
}

// VM_metrics_get_record
//
// Get a record containing the current state of the given VM_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_metrics ref, reference to the object
//
// returns:
// - VM_metrics record
// - all fields from the object
func (client *XenAPIClient) VM_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_metrics.get_record", session_id, self)
}

// VM_guest_metrics_get_all_records
//
// Return a map of VM_guest_metrics references to VM_guest_metrics records for all VM_guest_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VM_guest_metrics ref -> VM_guest_metrics record) map
// - records of all objects
func (client *XenAPIClient) VM_guest_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_all_records", session_id)
}

// VM_guest_metrics_get_all
//
// Return a list of all the VM_guest_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VM_guest_metrics ref set
// - references to all objects
func (client *XenAPIClient) VM_guest_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_all", session_id)
}

// VM_guest_metrics_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VM_guest_metrics.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VM_guest_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.remove_from_other_config", session_id, self, key)
}

// VM_guest_metrics_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VM_guest_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.add_to_other_config", session_id, self, key, value)
}

// VM_guest_metrics_set_other_config
//
// Set the other_config field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VM_guest_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.set_other_config", session_id, self, value)
}

// VM_guest_metrics_get_live
//
// Get the live field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_live(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_live", session_id, self)
}

// VM_guest_metrics_get_other_config
//
// Get the other_config field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_other_config", session_id, self)
}

// VM_guest_metrics_get_last_updated
//
// Get the last_updated field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_last_updated", session_id, self)
}

// VM_guest_metrics_get_other
//
// Get the other field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_other(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_other", session_id, self)
}

// VM_guest_metrics_get_networks
//
// Get the networks field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_networks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_networks", session_id, self)
}

// VM_guest_metrics_get_disks
//
// Get the disks field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_disks(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_disks", session_id, self)
}

// VM_guest_metrics_get_memory
//
// Get the memory field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_memory(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_memory", session_id, self)
}

// VM_guest_metrics_get_PV_drivers_up_to_date
//
// Get the PV_drivers_up_to_date field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_PV_drivers_up_to_date(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_PV_drivers_up_to_date", session_id, self)
}

// VM_guest_metrics_get_PV_drivers_version
//
// Get the PV_drivers_version field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_PV_drivers_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_PV_drivers_version", session_id, self)
}

// VM_guest_metrics_get_os_version
//
// Get the os_version field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_os_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_os_version", session_id, self)
}

// VM_guest_metrics_get_uuid
//
// Get the uuid field of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VM_guest_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_uuid", session_id, self)
}

// VM_guest_metrics_get_by_uuid
//
// Get a reference to the VM_guest_metrics instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VM_guest_metrics ref
// - reference to the object
func (client *XenAPIClient) VM_guest_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_by_uuid", session_id, uuid)
}

// VM_guest_metrics_get_record
//
// Get a record containing the current state of the given VM_guest_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_guest_metrics ref, reference to the object
//
// returns:
// - VM_guest_metrics record
// - all fields from the object
func (client *XenAPIClient) VM_guest_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VM_guest_metrics.get_record", session_id, self)
}

// VMPP_get_all_records
//
// Return a map of VMPP references to VMPP records for all VMPPs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VMPP ref -> VMPP record) map
// - records of all objects
func (client *XenAPIClient) VMPP_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_all_records", session_id)
}

// VMPP_get_all
//
// Return a list of all the VMPPs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VMPP ref set
// - references to all objects
func (client *XenAPIClient) VMPP_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_all", session_id)
}

// VMPP_set_archive_last_run_time
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, datetime, the value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_archive_last_run_time(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_last_run_time", session_id, self, value)
}

// VMPP_set_backup_last_run_time
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, datetime, the value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_backup_last_run_time(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_last_run_time", session_id, self, value)
}

// VMPP_remove_from_alarm_config
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to remove
//
// returns:
// - void
func (client *XenAPIClient) VMPP_remove_from_alarm_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_alarm_config", session_id, self, key)
}

// VMPP_remove_from_archive_schedule
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to remove
//
// returns:
// - void
func (client *XenAPIClient) VMPP_remove_from_archive_schedule(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_archive_schedule", session_id, self, key)
}

// VMPP_remove_from_archive_target_config
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to remove
//
// returns:
// - void
func (client *XenAPIClient) VMPP_remove_from_archive_target_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_archive_target_config", session_id, self, key)
}

// VMPP_remove_from_backup_schedule
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to remove
//
// returns:
// - void
func (client *XenAPIClient) VMPP_remove_from_backup_schedule(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VMPP.remove_from_backup_schedule", session_id, self, key)
}

// VMPP_add_to_alarm_config
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to add
// - value, string, the value to add
//
// returns:
// - void
func (client *XenAPIClient) VMPP_add_to_alarm_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_alarm_config", session_id, self, key, value)
}

// VMPP_add_to_archive_schedule
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to add
// - value, string, the value to add
//
// returns:
// - void
func (client *XenAPIClient) VMPP_add_to_archive_schedule(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_archive_schedule", session_id, self, key, value)
}

// VMPP_add_to_archive_target_config
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to add
// - value, string, the value to add
//
// returns:
// - void
func (client *XenAPIClient) VMPP_add_to_archive_target_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_archive_target_config", session_id, self, key, value)
}

// VMPP_add_to_backup_schedule
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - key, string, the key to add
// - value, string, the value to add
//
// returns:
// - void
func (client *XenAPIClient) VMPP_add_to_backup_schedule(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.add_to_backup_schedule", session_id, self, key, value)
}

// VMPP_set_alarm_config
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, (string -> string) map, the value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_alarm_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_alarm_config", session_id, self, value)
}

// VMPP_set_is_alarm_enabled
//
// Set the value of the is_alarm_enabled field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, bool, true if alarm is enabled for this policy
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_is_alarm_enabled(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_is_alarm_enabled", session_id, self, value)
}

// VMPP_set_archive_target_config
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, (string -> string) map, the value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_archive_target_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_target_config", session_id, self, value)
}

// VMPP_set_archive_target_type
//
// Set the value of the archive_target_config_type field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, enum vmpp_archive_target_type, the archive target config type
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_archive_target_type(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_target_type", session_id, self, value)
}

// VMPP_set_archive_schedule
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, (string -> string) map, the value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_archive_schedule(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_schedule", session_id, self, value)
}

// VMPP_set_archive_frequency
//
// Set the value of the archive_frequency field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, enum vmpp_archive_frequency, the archive frequency
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_archive_frequency(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_archive_frequency", session_id, self, value)
}

// VMPP_set_backup_schedule
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, (string -> string) map, the value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_backup_schedule(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_schedule", session_id, self, value)
}

// VMPP_set_backup_frequency
//
// Set the value of the backup_frequency field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, enum vmpp_backup_frequency, the backup frequency
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_backup_frequency(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_frequency", session_id, self, value)
}

// VMPP_set_backup_retention_value
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, The protection policy
// - value, int, the value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_backup_retention_value(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_retention_value", session_id, self, value)
}

// VMPP_get_alerts
//
// This call fetches a history of alerts for a given protection policy
//
// params:
// - session_id, session ref, Reference to a valid session
// - vmpp, VMPP ref, The protection policy
// - hours_from_now, int, how many hours in the past the oldest record to fetch is
//
// returns:
// - string set
// - A list of alerts encoded in xml
func (client *XenAPIClient) VMPP_get_alerts(session_id interface{}, vmpp interface{}, hours_from_now interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_alerts", session_id, vmpp, hours_from_now)
}

// VMPP_archive_now
//
// This call archives the snapshot provided as a parameter
//
// params:
// - session_id, session ref, Reference to a valid session
// - snapshot, VM ref, The snapshot to archive
//
// returns:
// - string
// - An XMLRPC result
func (client *XenAPIClient) VMPP_archive_now(session_id interface{}, snapshot interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.archive_now", session_id, snapshot)
}

// VMPP_protect_now
//
// This call executes the protection policy immediately
//
// params:
// - session_id, session ref, Reference to a valid session
// - vmpp, VMPP ref, The protection policy to execute
//
// returns:
// - string
// - An XMLRPC result
func (client *XenAPIClient) VMPP_protect_now(session_id interface{}, vmpp interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.protect_now", session_id, vmpp)
}

// VMPP_set_backup_type
//
// Set the backup_type field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
// - value, enum vmpp_backup_type, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_backup_type(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_backup_type", session_id, self, value)
}

// VMPP_set_is_policy_enabled
//
// Set the is_policy_enabled field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_is_policy_enabled(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_is_policy_enabled", session_id, self, value)
}

// VMPP_set_name_description
//
// Set the name/description field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_name_description", session_id, self, value)
}

// VMPP_set_name_label
//
// Set the name/label field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VMPP_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMPP.set_name_label", session_id, self, value)
}

// VMPP_get_recent_alerts
//
// Get the recent_alerts field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VMPP_get_recent_alerts(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_recent_alerts", session_id, self)
}

// VMPP_get_alarm_config
//
// Get the alarm_config field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VMPP_get_alarm_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_alarm_config", session_id, self)
}

// VMPP_get_is_alarm_enabled
//
// Get the is_alarm_enabled field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VMPP_get_is_alarm_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_alarm_enabled", session_id, self)
}

// VMPP_get_VMs
//
// Get the VMs field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - VM ref set
// - value of the field
func (client *XenAPIClient) VMPP_get_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_VMs", session_id, self)
}

// VMPP_get_archive_last_run_time
//
// Get the archive_last_run_time field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VMPP_get_archive_last_run_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_last_run_time", session_id, self)
}

// VMPP_get_is_archive_running
//
// Get the is_archive_running field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VMPP_get_is_archive_running(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_archive_running", session_id, self)
}

// VMPP_get_archive_schedule
//
// Get the archive_schedule field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VMPP_get_archive_schedule(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_schedule", session_id, self)
}

// VMPP_get_archive_frequency
//
// Get the archive_frequency field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - enum vmpp_archive_frequency
// - value of the field
func (client *XenAPIClient) VMPP_get_archive_frequency(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_frequency", session_id, self)
}

// VMPP_get_archive_target_config
//
// Get the archive_target_config field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VMPP_get_archive_target_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_target_config", session_id, self)
}

// VMPP_get_archive_target_type
//
// Get the archive_target_type field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - enum vmpp_archive_target_type
// - value of the field
func (client *XenAPIClient) VMPP_get_archive_target_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_archive_target_type", session_id, self)
}

// VMPP_get_backup_last_run_time
//
// Get the backup_last_run_time field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VMPP_get_backup_last_run_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_last_run_time", session_id, self)
}

// VMPP_get_is_backup_running
//
// Get the is_backup_running field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VMPP_get_is_backup_running(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_backup_running", session_id, self)
}

// VMPP_get_backup_schedule
//
// Get the backup_schedule field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VMPP_get_backup_schedule(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_schedule", session_id, self)
}

// VMPP_get_backup_frequency
//
// Get the backup_frequency field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - enum vmpp_backup_frequency
// - value of the field
func (client *XenAPIClient) VMPP_get_backup_frequency(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_frequency", session_id, self)
}

// VMPP_get_backup_retention_value
//
// Get the backup_retention_value field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VMPP_get_backup_retention_value(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_retention_value", session_id, self)
}

// VMPP_get_backup_type
//
// Get the backup_type field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - enum vmpp_backup_type
// - value of the field
func (client *XenAPIClient) VMPP_get_backup_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_backup_type", session_id, self)
}

// VMPP_get_is_policy_enabled
//
// Get the is_policy_enabled field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VMPP_get_is_policy_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_is_policy_enabled", session_id, self)
}

// VMPP_get_name_description
//
// Get the name/description field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VMPP_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_name_description", session_id, self)
}

// VMPP_get_name_label
//
// Get the name/label field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VMPP_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_name_label", session_id, self)
}

// VMPP_get_uuid
//
// Get the uuid field of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VMPP_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_uuid", session_id, self)
}

// VMPP_get_by_name_label
//
// Get all the VMPP instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - VMPP ref set
// - references to objects with matching names
func (client *XenAPIClient) VMPP_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_by_name_label", session_id, label)
}

// VMPP_destroy
//
// Destroy the specified VMPP instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VMPP_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.destroy", session_id, self)
}

// VMPP_create
//
// Create a new VMPP instance, and return its handle.
// The constructor args are: name_label, name_description, is_policy_enabled, backup_type, backup_retention_value, backup_frequency, backup_schedule, archive_target_type, archive_target_config, archive_frequency, archive_schedule, is_alarm_enabled, alarm_config (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, VMPP record, All constructor arguments
//
// returns:
// - VMPP ref
// - reference to the newly created object
func (client *XenAPIClient) VMPP_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.create", session_id, args)
}

// VMPP_get_by_uuid
//
// Get a reference to the VMPP instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VMPP ref
// - reference to the object
func (client *XenAPIClient) VMPP_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_by_uuid", session_id, uuid)
}

// VMPP_get_record
//
// Get a record containing the current state of the given VMPP.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VMPP ref, reference to the object
//
// returns:
// - VMPP record
// - all fields from the object
func (client *XenAPIClient) VMPP_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMPP.get_record", session_id, self)
}

// VMAppliance_get_all_records
//
// Return a map of VM_appliance references to VM_appliance records for all VM_appliances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VM_appliance ref -> VM_appliance record) map
// - records of all objects
func (client *XenAPIClient) VMAppliance_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_all_records", session_id)
}

// VMAppliance_get_all
//
// Return a list of all the VM_appliances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VM_appliance ref set
// - references to all objects
func (client *XenAPIClient) VMAppliance_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_all", session_id)
}

// VMAppliance_recover
//
// Recover the VM appliance
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, The VM appliance to recover
// - session_to, session ref, The session to which the VM appliance is to be recovered.
// - force, bool, Whether the VMs should replace newer versions of themselves.
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_recover(session_id interface{}, self interface{}, session_to interface{}, force bool) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.recover", session_id, self, session_to, force)
}

// VMAppliance_get_SRs_required_for_recovery
//
// Get the list of SRs required by the VM appliance to recover.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, The VM appliance for which the required list of SRs has to be recovered.
// - session_to, session ref, The session to which the list of SRs have to be recovered .
//
// returns:
// - SR ref set
// - refs for SRs required to recover the VM
func (client *XenAPIClient) VMAppliance_get_SRs_required_for_recovery(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_SRs_required_for_recovery", session_id, self, session_to)
}

// VMAppliance_assert_can_be_recovered
//
// Assert whether all SRs required to recover this VM appliance are available.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, The VM appliance to recover
// - session_to, session ref, The session to which the VM appliance is to be recovered.
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_assert_can_be_recovered(session_id interface{}, self interface{}, session_to interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.assert_can_be_recovered", session_id, self, session_to)
}

// VMAppliance_shutdown
//
// For each VM in the appliance, try to shut it down cleanly. If this fails, perform a hard shutdown of the VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, The VM appliance
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.shutdown", session_id, self)
}

// VMAppliance_hard_shutdown
//
// Perform a hard shutdown of all the VMs in the appliance
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, The VM appliance
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_hard_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.hard_shutdown", session_id, self)
}

// VMAppliance_clean_shutdown
//
// Perform a clean shutdown of all the VMs in the appliance
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, The VM appliance
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_clean_shutdown(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.clean_shutdown", session_id, self)
}

// VMAppliance_start
//
// Start all VMs in the appliance
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, The VM appliance
// - paused, bool, Instantiate all VMs belonging to this appliance in paused state if set to true.
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_start(session_id interface{}, self interface{}, paused bool) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.start", session_id, self, paused)
}

// VMAppliance_set_name_description
//
// Set the name/description field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.set_name_description", session_id, self, value)
}

// VMAppliance_set_name_label
//
// Set the name/label field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.set_name_label", session_id, self, value)
}

// VMAppliance_get_VMs
//
// Get the VMs field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - VM ref set
// - value of the field
func (client *XenAPIClient) VMAppliance_get_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_VMs", session_id, self)
}

// VMAppliance_get_current_operations
//
// Get the current_operations field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - (string -> enum vm_appliance_operation) map
// - value of the field
func (client *XenAPIClient) VMAppliance_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_current_operations", session_id, self)
}

// VMAppliance_get_allowed_operations
//
// Get the allowed_operations field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - enum vm_appliance_operation set
// - value of the field
func (client *XenAPIClient) VMAppliance_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_allowed_operations", session_id, self)
}

// VMAppliance_get_name_description
//
// Get the name/description field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VMAppliance_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_name_description", session_id, self)
}

// VMAppliance_get_name_label
//
// Get the name/label field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VMAppliance_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_name_label", session_id, self)
}

// VMAppliance_get_uuid
//
// Get the uuid field of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VMAppliance_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_uuid", session_id, self)
}

// VMAppliance_get_by_name_label
//
// Get all the VM_appliance instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - VM_appliance ref set
// - references to objects with matching names
func (client *XenAPIClient) VMAppliance_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_by_name_label", session_id, label)
}

// VMAppliance_destroy
//
// Destroy the specified VM_appliance instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VMAppliance_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.destroy", session_id, self)
}

// VMAppliance_create
//
// Create a new VM_appliance instance, and return its handle.
// The constructor args are: name_label, name_description (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, VM_appliance record, All constructor arguments
//
// returns:
// - VM_appliance ref
// - reference to the newly created object
func (client *XenAPIClient) VMAppliance_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.create", session_id, args)
}

// VMAppliance_get_by_uuid
//
// Get a reference to the VM_appliance instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VM_appliance ref
// - reference to the object
func (client *XenAPIClient) VMAppliance_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_by_uuid", session_id, uuid)
}

// VMAppliance_get_record
//
// Get a record containing the current state of the given VM_appliance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VM_appliance ref, reference to the object
//
// returns:
// - VM_appliance record
// - all fields from the object
func (client *XenAPIClient) VMAppliance_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VMAppliance.get_record", session_id, self)
}

// DRTask_get_all_records
//
// Return a map of DR_task references to DR_task records for all DR_tasks known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (DR_task ref -> DR_task record) map
// - records of all objects
func (client *XenAPIClient) DRTask_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_all_records", session_id)
}

// DRTask_get_all
//
// Return a list of all the DR_tasks known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - DR_task ref set
// - references to all objects
func (client *XenAPIClient) DRTask_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_all", session_id)
}

// DRTask_destroy
//
// Destroy the disaster recovery task, detaching and forgetting any SRs introduced which are no longer required
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, DR_task ref, The disaster recovery task to destroy
//
// returns:
// - void
func (client *XenAPIClient) DRTask_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.destroy", session_id, self)
}

// DRTask_create
//
// Create a disaster recovery task which will query the supplied list of devices
//
// params:
// - session_id, session ref, Reference to a valid session
// - a_type, string, The SR driver type of the SRs to introduce
// - device_config, (string -> string) map, The device configuration of the SRs to introduce
// - whitelist, string set, The devices to use for disaster recovery
//
// returns:
// - DR_task ref
// - The reference to the created task
func (client *XenAPIClient) DRTask_create(session_id interface{}, a_type string, device_config map[string]string, whitelist interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.create", session_id, a_type, device_config, whitelist)
}

// DRTask_get_introduced_SRs
//
// Get the introduced_SRs field of the given DR_task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, DR_task ref, reference to the object
//
// returns:
// - SR ref set
// - value of the field
func (client *XenAPIClient) DRTask_get_introduced_SRs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_introduced_SRs", session_id, self)
}

// DRTask_get_uuid
//
// Get the uuid field of the given DR_task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, DR_task ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) DRTask_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_uuid", session_id, self)
}

// DRTask_get_by_uuid
//
// Get a reference to the DR_task instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - DR_task ref
// - reference to the object
func (client *XenAPIClient) DRTask_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_by_uuid", session_id, uuid)
}

// DRTask_get_record
//
// Get a record containing the current state of the given DR_task.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, DR_task ref, reference to the object
//
// returns:
// - DR_task record
// - all fields from the object
func (client *XenAPIClient) DRTask_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("DRTask.get_record", session_id, self)
}

// host_get_all_records
//
// Return a map of host references to host records for all hosts known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (host ref -> host record) map
// - records of all objects
func (client *XenAPIClient) host_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_all_records", session_id)
}

// host_get_all
//
// Return a list of all the hosts known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - host ref set
// - references to all objects
func (client *XenAPIClient) host_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_all", session_id)
}

// host_disable_display
//
// Disable console output to the physical display device next time this host boots
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
//
// returns:
// - enum host_display
// - This host's physical display usage
func (client *XenAPIClient) host_disable_display(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.disable_display", session_id, host)
}

// host_enable_display
//
// Enable console output to the physical display device next time this host boots
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
//
// returns:
// - enum host_display
// - This host's physical display usage
func (client *XenAPIClient) host_enable_display(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.enable_display", session_id, host)
}

// host_declare_dead
//
// Declare that a host is dead. This is a dangerous operation, and should only be called if the administrator is absolutely sure the host is definitely dead
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to declare is dead
//
// returns:
// - void
func (client *XenAPIClient) host_declare_dead(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.declare_dead", session_id, host)
}

// host_migrate_receive
//
// Prepare to receive a VM, returning a token which can be passed to VM.migrate.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The target host
// - network, network ref, The network through which migration traffic should be received.
// - options, (string -> string) map, Extra configuration operations
//
// returns:
// - (string -> string) map
// - A value which should be passed to VM.migrate
func (client *XenAPIClient) host_migrate_receive(session_id interface{}, host interface{}, network interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.migrate_receive", session_id, host, network, options)
}

// host_disable_local_storage_caching
//
// Disable the use of a local SR for caching purposes
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
//
// returns:
// - void
func (client *XenAPIClient) host_disable_local_storage_caching(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.disable_local_storage_caching", session_id, host)
}

// host_enable_local_storage_caching
//
// Enable the use of a local SR for caching purposes
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - sr, SR ref, The SR to use as a local cache
//
// returns:
// - void
func (client *XenAPIClient) host_enable_local_storage_caching(session_id interface{}, host interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("host.enable_local_storage_caching", session_id, host, sr)
}

// host_reset_cpu_features
//
// Remove the feature mask, such that after a reboot all features of the CPU are enabled.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
//
// returns:
// - void
func (client *XenAPIClient) host_reset_cpu_features(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.reset_cpu_features", session_id, host)
}

// host_set_cpu_features
//
// Set the CPU features to be used after a reboot, if the given features string is valid.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - features, string, The features string (32 hexadecimal digits)
//
// returns:
// - void
func (client *XenAPIClient) host_set_cpu_features(session_id interface{}, host interface{}, features string) (i interface{}, err error) {
	return client.RPCCall("host.set_cpu_features", session_id, host, features)
}

// host_set_power_on_mode
//
// Set the power-on-mode, host, user and password
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, The host
// - power_on_mode, string, power-on-mode can be empty,iLO,wake-on-lan, DRAC or other
// - power_on_config, (string -> string) map, Power on config
//
// returns:
// - void
func (client *XenAPIClient) host_set_power_on_mode(session_id interface{}, self interface{}, power_on_mode string, power_on_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_power_on_mode", session_id, self, power_on_mode, power_on_config)
}

// host_refresh_pack_info
//
// Refresh the list of installed Supplemental Packs.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to modify
//
// returns:
// - void
func (client *XenAPIClient) host_refresh_pack_info(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.refresh_pack_info", session_id, host)
}

// host_apply_edition
//
// Change to another edition, or reactivate the current edition after a license has expired. This may be subject to the successful checkout of an appropriate license.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - edition, string, The requested edition
// - force, bool, Update the license params even if the apply call fails
//
// returns:
// - void
func (client *XenAPIClient) host_apply_edition(session_id interface{}, host interface{}, edition string, force bool) (i interface{}, err error) {
	return client.RPCCall("host.apply_edition", session_id, host, edition, force)
}

// host_get_server_certificate
//
// Get the installed server SSL certificate.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
//
// returns:
// - string
// - The installed server SSL certificate, in PEM form.
func (client *XenAPIClient) host_get_server_certificate(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_server_certificate", session_id, host)
}

// host_retrieve_wlb_evacuate_recommendations
//
// Retrieves recommended host migrations to perform when evacuating the host from the wlb server. If a VM cannot be migrated from the host the reason is listed instead of a recommendation.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, The host to query
//
// returns:
// - (VM ref -> string set) map
// - VMs and the reasons why they would block evacuation, or their target host recommended by the wlb server
func (client *XenAPIClient) host_retrieve_wlb_evacuate_recommendations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.retrieve_wlb_evacuate_recommendations", session_id, self)
}

// host_disable_external_auth
//
// This call disables external authentication on the local host
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host whose external authentication should be disabled
// - config, (string -> string) map, Optional parameters as a list of key-values containing the configuration data
//
// returns:
// - void
func (client *XenAPIClient) host_disable_external_auth(session_id interface{}, host interface{}, config map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.disable_external_auth", session_id, host, config)
}

// host_enable_external_auth
//
// This call enables external authentication on a host
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host whose external authentication should be enabled
// - config, (string -> string) map, A list of key-values containing the configuration data
// - service_name, string, The name of the service
// - auth_type, string, The type of authentication (e.g. AD for Active Directory)
//
// returns:
// - void
func (client *XenAPIClient) host_enable_external_auth(session_id interface{}, host interface{}, config map[string]string, service_name string, auth_type string) (i interface{}, err error) {
	return client.RPCCall("host.enable_external_auth", session_id, host, config, service_name, auth_type)
}

// host_get_server_localtime
//
// This call queries the host's clock for the current time in the host's local timezone
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host whose clock should be queried
//
// returns:
// - datetime
// - The current local time
func (client *XenAPIClient) host_get_server_localtime(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_server_localtime", session_id, host)
}

// host_get_servertime
//
// This call queries the host's clock for the current time
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host whose clock should be queried
//
// returns:
// - datetime
// - The current time
func (client *XenAPIClient) host_get_servertime(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_servertime", session_id, host)
}

// host_call_plugin
//
// Call a XenAPI plugin on this host
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - plugin, string, The name of the plugin
// - fn, string, The name of the function within the plugin
// - args, (string -> string) map, Arguments for the function
//
// returns:
// - string
// - Result from the plugin
func (client *XenAPIClient) host_call_plugin(session_id interface{}, host interface{}, plugin string, fn string, args map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.call_plugin", session_id, host, plugin, fn, args)
}

// host_create_new_blob
//
// Create a placeholder for a named binary blob of data that is associated with this host
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - name, string, The name associated with the blob
// - mime_type, string, The mime type for the data. Empty string translates to application/octet-stream
// - public, bool, True if the blob should be publicly available
//
// returns:
// - blob ref
// - The reference of the blob, needed for populating its data
func (client *XenAPIClient) host_create_new_blob(session_id interface{}, host interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("host.create_new_blob", session_id, host, name, mime_type, public)
}

// host_backup_rrds
//
// This causes the RRDs to be backed up to the master
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, Schedule a backup of the RRDs of this host
// - delay, float, Delay in seconds from when the call is received to perform the backup
//
// returns:
// - void
func (client *XenAPIClient) host_backup_rrds(session_id interface{}, host interface{}, delay interface{}) (i interface{}, err error) {
	return client.RPCCall("host.backup_rrds", session_id, host, delay)
}

// host_sync_data
//
// This causes the synchronisation of the non-database data (messages, RRDs and so on) stored on the master to be synchronised with the host
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to whom the data should be sent
//
// returns:
// - void
func (client *XenAPIClient) host_sync_data(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.sync_data", session_id, host)
}

// host_compute_memory_overhead
//
// Computes the virtualization memory overhead of a host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host for which to compute the memory overhead
//
// returns:
// - int
// - the virtualization memory overhead of the host.
func (client *XenAPIClient) host_compute_memory_overhead(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.compute_memory_overhead", session_id, host)
}

// host_compute_free_memory
//
// Computes the amount of free memory on the host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to send the request to
//
// returns:
// - int
// - the amount of free memory on the host.
func (client *XenAPIClient) host_compute_free_memory(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.compute_free_memory", session_id, host)
}

// host_set_hostname_live
//
// Sets the host name to the specified string.  Both the API and lower-level system hostname are changed immediately.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host whose host name to set
// - hostname, string, The new host name
//
// returns:
// - void
func (client *XenAPIClient) host_set_hostname_live(session_id interface{}, host interface{}, hostname string) (i interface{}, err error) {
	return client.RPCCall("host.set_hostname_live", session_id, host, hostname)
}

// host_shutdown_agent
//
// Shuts the agent down after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) host_shutdown_agent(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.shutdown_agent", session_id)
}

// host_restart_agent
//
// Restarts the agent after a 10 second pause. WARNING: this is a dangerous operation. Any operations in progress will be aborted, and unrecoverable data loss may occur. The caller is responsible for ensuring that there are no operations in progress when this method is called.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host on which you want to restart the agent
//
// returns:
// - void
func (client *XenAPIClient) host_restart_agent(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.restart_agent", session_id, host)
}

// host_get_system_status_capabilities
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to interrogate
//
// returns:
// - string
// - An XML fragment containing the system status capabilities.
func (client *XenAPIClient) host_get_system_status_capabilities(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_system_status_capabilities", session_id, host)
}

// host_get_management_interface
//
// Returns the management interface for the specified host
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, Which host's management interface is required
//
// returns:
// - PIF ref
// - The management interface for the host
func (client *XenAPIClient) host_get_management_interface(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_management_interface", session_id, host)
}

// host_management_disable
//
// Disable the management network interface
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) host_management_disable(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.management_disable", session_id)
}

// host_local_management_reconfigure
//
// Reconfigure the management network interface. Should only be used if Host.management_reconfigure is impossible because the network configuration is broken.
//
// params:
// - session_id, session ref, Reference to a valid session
// - a_interface, string, name of the interface to use as a management interface
//
// returns:
// - void
func (client *XenAPIClient) host_local_management_reconfigure(session_id interface{}, a_interface string) (i interface{}, err error) {
	return client.RPCCall("host.local_management_reconfigure", session_id, a_interface)
}

// host_management_reconfigure
//
// Reconfigure the management network interface
//
// params:
// - session_id, session ref, Reference to a valid session
// - pif, PIF ref, reference to a PIF object corresponding to the management interface
//
// returns:
// - void
func (client *XenAPIClient) host_management_reconfigure(session_id interface{}, pif interface{}) (i interface{}, err error) {
	return client.RPCCall("host.management_reconfigure", session_id, pif)
}

// host_syslog_reconfigure
//
// Re-configure syslog logging
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, Tell the host to reread its Host.logging parameters and reconfigure itself accordingly
//
// returns:
// - void
func (client *XenAPIClient) host_syslog_reconfigure(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.syslog_reconfigure", session_id, host)
}

// host_evacuate
//
// Migrate all VMs off of this host, where possible.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to evacuate
//
// returns:
// - void
func (client *XenAPIClient) host_evacuate(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.evacuate", session_id, host)
}

// host_get_uncooperative_resident_VMs
//
// Return a set of VMs which are not co-operating with the host's memory control system
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, The host to query
//
// returns:
// - VM ref set
// - VMs which are not co-operating
func (client *XenAPIClient) host_get_uncooperative_resident_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_uncooperative_resident_VMs", session_id, self)
}

// host_get_vms_which_prevent_evacuation
//
// Return a set of VMs which prevent the host being evacuated, with per-VM error codes
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, The host to query
//
// returns:
// - (VM ref -> string set) map
// - VMs which block evacuation together with reasons
func (client *XenAPIClient) host_get_vms_which_prevent_evacuation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_vms_which_prevent_evacuation", session_id, self)
}

// host_assert_can_evacuate
//
// Check this host can be evacuated.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to evacuate
//
// returns:
// - void
func (client *XenAPIClient) host_assert_can_evacuate(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.assert_can_evacuate", session_id, host)
}

// host_forget_data_source_archives
//
// Forget the recorded statistics related to the specified data source
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - data_source, string, The data source whose archives are to be forgotten
//
// returns:
// - void
func (client *XenAPIClient) host_forget_data_source_archives(session_id interface{}, host interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("host.forget_data_source_archives", session_id, host, data_source)
}

// host_query_data_source
//
// Query the latest value of the specified data source
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - data_source, string, The data source to query
//
// returns:
// - float
// - The latest value, averaged over the last 5 seconds
func (client *XenAPIClient) host_query_data_source(session_id interface{}, host interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("host.query_data_source", session_id, host, data_source)
}

// host_record_data_source
//
// Start recording the specified data source
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - data_source, string, The data source to record
//
// returns:
// - void
func (client *XenAPIClient) host_record_data_source(session_id interface{}, host interface{}, data_source string) (i interface{}, err error) {
	return client.RPCCall("host.record_data_source", session_id, host, data_source)
}

// host_get_data_sources
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to interrogate
//
// returns:
// - data_source record set
// - A set of data sources
func (client *XenAPIClient) host_get_data_sources(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_data_sources", session_id, host)
}

// host_emergency_ha_disable
//
// This call disables HA on the local host. This should only be used with extreme care.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - void
func (client *XenAPIClient) host_emergency_ha_disable(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.emergency_ha_disable", session_id)
}

// host_power_on
//
// Attempt to power-on the host (if the capability exists).
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to power on
//
// returns:
// - void
func (client *XenAPIClient) host_power_on(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.power_on", session_id, host)
}

// host_destroy
//
// Destroy specified host record in database
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, The host record to remove
//
// returns:
// - void
func (client *XenAPIClient) host_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.destroy", session_id, self)
}

// host_license_apply
//
// Apply a new license to a host
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to upload the license to
// - contents, string, The contents of the license file, base64 encoded
//
// returns:
// - void
func (client *XenAPIClient) host_license_apply(session_id interface{}, host interface{}, contents string) (i interface{}, err error) {
	return client.RPCCall("host.license_apply", session_id, host, contents)
}

// host_list_methods
//
// List all supported methods
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - string set
// - The name of every supported method.
func (client *XenAPIClient) host_list_methods(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host.list_methods", session_id)
}

// host_bugreport_upload
//
// Run xen-bugtool --yestoall and upload the output to Citrix support
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host on which to run xen-bugtool
// - url, string, The URL to upload to
// - options, (string -> string) map, Extra configuration operations
//
// returns:
// - void
func (client *XenAPIClient) host_bugreport_upload(session_id interface{}, host interface{}, url string, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.bugreport_upload", session_id, host, url, options)
}

// host_send_debug_keys
//
// Inject the given string as debugging keys into Xen
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host
// - keys, string, The keys to send
//
// returns:
// - void
func (client *XenAPIClient) host_send_debug_keys(session_id interface{}, host interface{}, keys string) (i interface{}, err error) {
	return client.RPCCall("host.send_debug_keys", session_id, host, keys)
}

// host_get_log
//
// Get the host's log file
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to query
//
// returns:
// - string
// - The contents of the host's primary log file
func (client *XenAPIClient) host_get_log(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_log", session_id, host)
}

// host_dmesg_clear
//
// Get the host xen dmesg, and clear the buffer.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to query
//
// returns:
// - string
// - dmesg string
func (client *XenAPIClient) host_dmesg_clear(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.dmesg_clear", session_id, host)
}

// host_dmesg
//
// Get the host xen dmesg.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to query
//
// returns:
// - string
// - dmesg string
func (client *XenAPIClient) host_dmesg(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.dmesg", session_id, host)
}

// host_reboot
//
// Reboot the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to reboot
//
// returns:
// - void
func (client *XenAPIClient) host_reboot(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.reboot", session_id, host)
}

// host_shutdown
//
// Shutdown the host. (This function can only be called if there are no currently running VMs on the host and it is disabled.)
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to shutdown
//
// returns:
// - void
func (client *XenAPIClient) host_shutdown(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.shutdown", session_id, host)
}

// host_enable
//
// Puts the host into a state in which new VMs can be started.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to enable
//
// returns:
// - void
func (client *XenAPIClient) host_enable(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.enable", session_id, host)
}

// host_disable
//
// Puts the host into a state in which no new VMs can be started. Currently active VMs on the host continue to execute.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The Host to disable
//
// returns:
// - void
func (client *XenAPIClient) host_disable(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("host.disable", session_id, host)
}

// host_set_display
//
// Set the display field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, enum host_display, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_display(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_display", session_id, self, value)
}

// host_remove_from_guest_VCPUs_params
//
// Remove the given key and its corresponding value from the guest_VCPUs_params field of the given host.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) host_remove_from_guest_VCPUs_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_guest_VCPUs_params", session_id, self, key)
}

// host_add_to_guest_VCPUs_params
//
// Add the given key-value pair to the guest_VCPUs_params field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) host_add_to_guest_VCPUs_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_guest_VCPUs_params", session_id, self, key, value)
}

// host_set_guest_VCPUs_params
//
// Set the guest_VCPUs_params field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_guest_VCPUs_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_guest_VCPUs_params", session_id, self, value)
}

// host_remove_from_license_server
//
// Remove the given key and its corresponding value from the license_server field of the given host.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) host_remove_from_license_server(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_license_server", session_id, self, key)
}

// host_add_to_license_server
//
// Add the given key-value pair to the license_server field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) host_add_to_license_server(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_license_server", session_id, self, key, value)
}

// host_set_license_server
//
// Set the license_server field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_license_server(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_license_server", session_id, self, value)
}

// host_remove_tags
//
// Remove the given value from the tags field of the given host.  If the value is not in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, string, Value to remove
//
// returns:
// - void
func (client *XenAPIClient) host_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.remove_tags", session_id, self, value)
}

// host_add_tags
//
// Add the given value to the tags field of the given host.  If the value is already in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, string, New value to add
//
// returns:
// - void
func (client *XenAPIClient) host_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_tags", session_id, self, value)
}

// host_set_tags
//
// Set the tags field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, string set, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_tags", session_id, self, value)
}

// host_set_address
//
// Set the address field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_address(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_address", session_id, self, value)
}

// host_set_hostname
//
// Set the hostname field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_hostname(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_hostname", session_id, self, value)
}

// host_set_crash_dump_sr
//
// Set the crash_dump_sr field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, SR ref, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_crash_dump_sr(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_crash_dump_sr", session_id, self, value)
}

// host_set_suspend_image_sr
//
// Set the suspend_image_sr field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, SR ref, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_suspend_image_sr(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("host.set_suspend_image_sr", session_id, self, value)
}

// host_remove_from_logging
//
// Remove the given key and its corresponding value from the logging field of the given host.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) host_remove_from_logging(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_logging", session_id, self, key)
}

// host_add_to_logging
//
// Add the given key-value pair to the logging field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) host_add_to_logging(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_logging", session_id, self, key, value)
}

// host_set_logging
//
// Set the logging field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_logging(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_logging", session_id, self, value)
}

// host_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given host.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) host_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host.remove_from_other_config", session_id, self, key)
}

// host_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) host_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host.add_to_other_config", session_id, self, key, value)
}

// host_set_other_config
//
// Set the other_config field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host.set_other_config", session_id, self, value)
}

// host_set_name_description
//
// Set the name/description field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_name_description", session_id, self, value)
}

// host_set_name_label
//
// Set the name/label field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("host.set_name_label", session_id, self, value)
}

// host_get_display
//
// Get the display field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - enum host_display
// - value of the field
func (client *XenAPIClient) host_get_display(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_display", session_id, self)
}

// host_get_guest_VCPUs_params
//
// Get the guest_VCPUs_params field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_guest_VCPUs_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_guest_VCPUs_params", session_id, self)
}

// host_get_PGPUs
//
// Get the PGPUs field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - PGPU ref set
// - value of the field
func (client *XenAPIClient) host_get_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PGPUs", session_id, self)
}

// host_get_PCIs
//
// Get the PCIs field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - PCI ref set
// - value of the field
func (client *XenAPIClient) host_get_PCIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PCIs", session_id, self)
}

// host_get_chipset_info
//
// Get the chipset_info field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_chipset_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_chipset_info", session_id, self)
}

// host_get_local_cache_sr
//
// Get the local_cache_sr field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) host_get_local_cache_sr(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_local_cache_sr", session_id, self)
}

// host_get_power_on_config
//
// Get the power_on_config field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_power_on_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_power_on_config", session_id, self)
}

// host_get_power_on_mode
//
// Get the power_on_mode field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_power_on_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_power_on_mode", session_id, self)
}

// host_get_bios_strings
//
// Get the bios_strings field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_bios_strings(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_bios_strings", session_id, self)
}

// host_get_license_server
//
// Get the license_server field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_license_server(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_license_server", session_id, self)
}

// host_get_edition
//
// Get the edition field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_edition(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_edition", session_id, self)
}

// host_get_external_auth_configuration
//
// Get the external_auth_configuration field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_external_auth_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_external_auth_configuration", session_id, self)
}

// host_get_external_auth_service_name
//
// Get the external_auth_service_name field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_external_auth_service_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_external_auth_service_name", session_id, self)
}

// host_get_external_auth_type
//
// Get the external_auth_type field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_external_auth_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_external_auth_type", session_id, self)
}

// host_get_tags
//
// Get the tags field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) host_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_tags", session_id, self)
}

// host_get_blobs
//
// Get the blobs field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> blob ref) map
// - value of the field
func (client *XenAPIClient) host_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_blobs", session_id, self)
}

// host_get_ha_network_peers
//
// Get the ha_network_peers field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) host_get_ha_network_peers(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_ha_network_peers", session_id, self)
}

// host_get_ha_statefiles
//
// Get the ha_statefiles field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) host_get_ha_statefiles(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_ha_statefiles", session_id, self)
}

// host_get_license_params
//
// Get the license_params field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_license_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_license_params", session_id, self)
}

// host_get_metrics
//
// Get the metrics field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - host_metrics ref
// - value of the field
func (client *XenAPIClient) host_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_metrics", session_id, self)
}

// host_get_address
//
// Get the address field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_address(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_address", session_id, self)
}

// host_get_hostname
//
// Get the hostname field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_hostname(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_hostname", session_id, self)
}

// host_get_cpu_info
//
// Get the cpu_info field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_cpu_info(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_cpu_info", session_id, self)
}

// host_get_host_CPUs
//
// Get the host_CPUs field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - host_cpu ref set
// - value of the field
func (client *XenAPIClient) host_get_host_CPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_host_CPUs", session_id, self)
}

// host_get_PBDs
//
// Get the PBDs field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - PBD ref set
// - value of the field
func (client *XenAPIClient) host_get_PBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PBDs", session_id, self)
}

// host_get_patches
//
// Get the patches field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - host_patch ref set
// - value of the field
func (client *XenAPIClient) host_get_patches(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_patches", session_id, self)
}

// host_get_crashdumps
//
// Get the crashdumps field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - host_crashdump ref set
// - value of the field
func (client *XenAPIClient) host_get_crashdumps(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_crashdumps", session_id, self)
}

// host_get_crash_dump_sr
//
// Get the crash_dump_sr field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) host_get_crash_dump_sr(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_crash_dump_sr", session_id, self)
}

// host_get_suspend_image_sr
//
// Get the suspend_image_sr field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) host_get_suspend_image_sr(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_suspend_image_sr", session_id, self)
}

// host_get_PIFs
//
// Get the PIFs field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - PIF ref set
// - value of the field
func (client *XenAPIClient) host_get_PIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_PIFs", session_id, self)
}

// host_get_logging
//
// Get the logging field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_logging(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_logging", session_id, self)
}

// host_get_resident_VMs
//
// Get the resident_VMs field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - VM ref set
// - value of the field
func (client *XenAPIClient) host_get_resident_VMs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_resident_VMs", session_id, self)
}

// host_get_supported_bootloaders
//
// Get the supported_bootloaders field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) host_get_supported_bootloaders(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_supported_bootloaders", session_id, self)
}

// host_get_sched_policy
//
// Get the sched_policy field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_sched_policy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_sched_policy", session_id, self)
}

// host_get_cpu_configuration
//
// Get the cpu_configuration field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_cpu_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_cpu_configuration", session_id, self)
}

// host_get_capabilities
//
// Get the capabilities field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) host_get_capabilities(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_capabilities", session_id, self)
}

// host_get_other_config
//
// Get the other_config field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_other_config", session_id, self)
}

// host_get_software_version
//
// Get the software_version field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_software_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_software_version", session_id, self)
}

// host_get_enabled
//
// Get the enabled field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) host_get_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_enabled", session_id, self)
}

// host_get_API_version_vendor_implementation
//
// Get the API_version/vendor_implementation field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_get_API_version_vendor_implementation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_vendor_implementation", session_id, self)
}

// host_get_API_version_vendor
//
// Get the API_version/vendor field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_API_version_vendor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_vendor", session_id, self)
}

// host_get_API_version_minor
//
// Get the API_version/minor field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) host_get_API_version_minor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_minor", session_id, self)
}

// host_get_API_version_major
//
// Get the API_version/major field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) host_get_API_version_major(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_API_version_major", session_id, self)
}

// host_get_current_operations
//
// Get the current_operations field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - (string -> enum host_allowed_operations) map
// - value of the field
func (client *XenAPIClient) host_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_current_operations", session_id, self)
}

// host_get_allowed_operations
//
// Get the allowed_operations field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - enum host_allowed_operations set
// - value of the field
func (client *XenAPIClient) host_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_allowed_operations", session_id, self)
}

// host_get_memory_overhead
//
// Get the memory/overhead field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) host_get_memory_overhead(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_memory_overhead", session_id, self)
}

// host_get_name_description
//
// Get the name/description field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_name_description", session_id, self)
}

// host_get_name_label
//
// Get the name/label field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_name_label", session_id, self)
}

// host_get_uuid
//
// Get the uuid field of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_uuid", session_id, self)
}

// host_get_by_name_label
//
// Get all the host instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - host ref set
// - references to objects with matching names
func (client *XenAPIClient) host_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("host.get_by_name_label", session_id, label)
}

// host_get_by_uuid
//
// Get a reference to the host instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - host ref
// - reference to the object
func (client *XenAPIClient) host_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("host.get_by_uuid", session_id, uuid)
}

// host_get_record
//
// Get a record containing the current state of the given host.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host ref, reference to the object
//
// returns:
// - host record
// - all fields from the object
func (client *XenAPIClient) host_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host.get_record", session_id, self)
}

// hostCrashdump_get_all_records
//
// Return a map of host_crashdump references to host_crashdump records for all host_crashdumps known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (host_crashdump ref -> host_crashdump record) map
// - records of all objects
func (client *XenAPIClient) hostCrashdump_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_all_records", session_id)
}

// hostCrashdump_get_all
//
// Return a list of all the host_crashdumps known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - host_crashdump ref set
// - references to all objects
func (client *XenAPIClient) hostCrashdump_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_all", session_id)
}

// hostCrashdump_upload
//
// Upload the specified host crash dump to a specified URL
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, The host crashdump to upload
// - url, string, The URL to upload to
// - options, (string -> string) map, Extra configuration operations
//
// returns:
// - void
func (client *XenAPIClient) hostCrashdump_upload(session_id interface{}, self interface{}, url string, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.upload", session_id, self, url, options)
}

// hostCrashdump_destroy
//
// Destroy specified host crash dump, removing it from the disk.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, The host crashdump to destroy
//
// returns:
// - void
func (client *XenAPIClient) hostCrashdump_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.destroy", session_id, self)
}

// hostCrashdump_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given host_crashdump.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) hostCrashdump_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.remove_from_other_config", session_id, self, key)
}

// hostCrashdump_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) hostCrashdump_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.add_to_other_config", session_id, self, key, value)
}

// hostCrashdump_set_other_config
//
// Set the other_config field of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) hostCrashdump_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.set_other_config", session_id, self, value)
}

// hostCrashdump_get_other_config
//
// Get the other_config field of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) hostCrashdump_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_other_config", session_id, self)
}

// hostCrashdump_get_size
//
// Get the size field of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) hostCrashdump_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_size", session_id, self)
}

// hostCrashdump_get_timestamp
//
// Get the timestamp field of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) hostCrashdump_get_timestamp(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_timestamp", session_id, self)
}

// hostCrashdump_get_host
//
// Get the host field of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) hostCrashdump_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_host", session_id, self)
}

// hostCrashdump_get_uuid
//
// Get the uuid field of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostCrashdump_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_uuid", session_id, self)
}

// hostCrashdump_get_by_uuid
//
// Get a reference to the host_crashdump instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - host_crashdump ref
// - reference to the object
func (client *XenAPIClient) hostCrashdump_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_by_uuid", session_id, uuid)
}

// hostCrashdump_get_record
//
// Get a record containing the current state of the given host_crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_crashdump ref, reference to the object
//
// returns:
// - host_crashdump record
// - all fields from the object
func (client *XenAPIClient) hostCrashdump_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCrashdump.get_record", session_id, self)
}

// hostPatch_get_all_records
//
// Return a map of host_patch references to host_patch records for all host_patchs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (host_patch ref -> host_patch record) map
// - records of all objects
func (client *XenAPIClient) hostPatch_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_all_records", session_id)
}

// hostPatch_get_all
//
// Return a list of all the host_patchs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - host_patch ref set
// - references to all objects
func (client *XenAPIClient) hostPatch_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_all", session_id)
}

// hostPatch_apply
//
// Apply the selected patch and return its output
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, The patch to apply
//
// returns:
// - string
// - the output of the patch application process
func (client *XenAPIClient) hostPatch_apply(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.apply", session_id, self)
}

// hostPatch_destroy
//
// Destroy the specified host patch, removing it from the disk. This does NOT reverse the patch
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, The patch to destroy
//
// returns:
// - void
func (client *XenAPIClient) hostPatch_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.destroy", session_id, self)
}

// hostPatch_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given host_patch.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) hostPatch_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.remove_from_other_config", session_id, self, key)
}

// hostPatch_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) hostPatch_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.add_to_other_config", session_id, self, key, value)
}

// hostPatch_set_other_config
//
// Set the other_config field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) hostPatch_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.set_other_config", session_id, self, value)
}

// hostPatch_get_other_config
//
// Get the other_config field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) hostPatch_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_other_config", session_id, self)
}

// hostPatch_get_pool_patch
//
// Get the pool_patch field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - pool_patch ref
// - value of the field
func (client *XenAPIClient) hostPatch_get_pool_patch(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_pool_patch", session_id, self)
}

// hostPatch_get_size
//
// Get the size field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) hostPatch_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_size", session_id, self)
}

// hostPatch_get_timestamp_applied
//
// Get the timestamp_applied field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) hostPatch_get_timestamp_applied(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_timestamp_applied", session_id, self)
}

// hostPatch_get_applied
//
// Get the applied field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) hostPatch_get_applied(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_applied", session_id, self)
}

// hostPatch_get_host
//
// Get the host field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) hostPatch_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_host", session_id, self)
}

// hostPatch_get_version
//
// Get the version field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostPatch_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_version", session_id, self)
}

// hostPatch_get_name_description
//
// Get the name/description field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostPatch_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_name_description", session_id, self)
}

// hostPatch_get_name_label
//
// Get the name/label field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostPatch_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_name_label", session_id, self)
}

// hostPatch_get_uuid
//
// Get the uuid field of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostPatch_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_uuid", session_id, self)
}

// hostPatch_get_by_name_label
//
// Get all the host_patch instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - host_patch ref set
// - references to objects with matching names
func (client *XenAPIClient) hostPatch_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_by_name_label", session_id, label)
}

// hostPatch_get_by_uuid
//
// Get a reference to the host_patch instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - host_patch ref
// - reference to the object
func (client *XenAPIClient) hostPatch_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_by_uuid", session_id, uuid)
}

// hostPatch_get_record
//
// Get a record containing the current state of the given host_patch.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_patch ref, reference to the object
//
// returns:
// - host_patch record
// - all fields from the object
func (client *XenAPIClient) hostPatch_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostPatch.get_record", session_id, self)
}

// host_metrics_get_all_records
//
// Return a map of host_metrics references to host_metrics records for all host_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (host_metrics ref -> host_metrics record) map
// - records of all objects
func (client *XenAPIClient) host_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_all_records", session_id)
}

// host_metrics_get_all
//
// Return a list of all the host_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - host_metrics ref set
// - references to all objects
func (client *XenAPIClient) host_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_all", session_id)
}

// host_metrics_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given host_metrics.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) host_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.remove_from_other_config", session_id, self, key)
}

// host_metrics_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) host_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.add_to_other_config", session_id, self, key, value)
}

// host_metrics_set_other_config
//
// Set the other_config field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) host_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.set_other_config", session_id, self, value)
}

// host_metrics_get_other_config
//
// Get the other_config field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) host_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_other_config", session_id, self)
}

// host_metrics_get_last_updated
//
// Get the last_updated field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) host_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_last_updated", session_id, self)
}

// host_metrics_get_live
//
// Get the live field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) host_metrics_get_live(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_live", session_id, self)
}

// host_metrics_get_memory_free
//
// Get the memory/free field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) host_metrics_get_memory_free(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_memory_free", session_id, self)
}

// host_metrics_get_memory_total
//
// Get the memory/total field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) host_metrics_get_memory_total(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_memory_total", session_id, self)
}

// host_metrics_get_uuid
//
// Get the uuid field of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) host_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_uuid", session_id, self)
}

// host_metrics_get_by_uuid
//
// Get a reference to the host_metrics instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - host_metrics ref
// - reference to the object
func (client *XenAPIClient) host_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_by_uuid", session_id, uuid)
}

// host_metrics_get_record
//
// Get a record containing the current state of the given host_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_metrics ref, reference to the object
//
// returns:
// - host_metrics record
// - all fields from the object
func (client *XenAPIClient) host_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("host_metrics.get_record", session_id, self)
}

// hostCpu_get_all_records
//
// Return a map of host_cpu references to host_cpu records for all host_cpus known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (host_cpu ref -> host_cpu record) map
// - records of all objects
func (client *XenAPIClient) hostCpu_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_all_records", session_id)
}

// hostCpu_get_all
//
// Return a list of all the host_cpus known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - host_cpu ref set
// - references to all objects
func (client *XenAPIClient) hostCpu_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_all", session_id)
}

// hostCpu_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given host_cpu.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) hostCpu_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.remove_from_other_config", session_id, self, key)
}

// hostCpu_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) hostCpu_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.add_to_other_config", session_id, self, key, value)
}

// hostCpu_set_other_config
//
// Set the other_config field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) hostCpu_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.set_other_config", session_id, self, value)
}

// hostCpu_get_other_config
//
// Get the other_config field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) hostCpu_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_other_config", session_id, self)
}

// hostCpu_get_utilisation
//
// Get the utilisation field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) hostCpu_get_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_utilisation", session_id, self)
}

// hostCpu_get_features
//
// Get the features field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostCpu_get_features(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_features", session_id, self)
}

// hostCpu_get_flags
//
// Get the flags field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostCpu_get_flags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_flags", session_id, self)
}

// hostCpu_get_stepping
//
// Get the stepping field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostCpu_get_stepping(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_stepping", session_id, self)
}

// hostCpu_get_model
//
// Get the model field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) hostCpu_get_model(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_model", session_id, self)
}

// hostCpu_get_family
//
// Get the family field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) hostCpu_get_family(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_family", session_id, self)
}

// hostCpu_get_modelname
//
// Get the modelname field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostCpu_get_modelname(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_modelname", session_id, self)
}

// hostCpu_get_speed
//
// Get the speed field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) hostCpu_get_speed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_speed", session_id, self)
}

// hostCpu_get_vendor
//
// Get the vendor field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostCpu_get_vendor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_vendor", session_id, self)
}

// hostCpu_get_number
//
// Get the number field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) hostCpu_get_number(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_number", session_id, self)
}

// hostCpu_get_host
//
// Get the host field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) hostCpu_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_host", session_id, self)
}

// hostCpu_get_uuid
//
// Get the uuid field of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) hostCpu_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_uuid", session_id, self)
}

// hostCpu_get_by_uuid
//
// Get a reference to the host_cpu instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - host_cpu ref
// - reference to the object
func (client *XenAPIClient) hostCpu_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_by_uuid", session_id, uuid)
}

// hostCpu_get_record
//
// Get a record containing the current state of the given host_cpu.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, host_cpu ref, reference to the object
//
// returns:
// - host_cpu record
// - all fields from the object
func (client *XenAPIClient) hostCpu_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("hostCpu.get_record", session_id, self)
}

// network_get_all_records
//
// Return a map of network references to network records for all networks known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (network ref -> network record) map
// - records of all objects
func (client *XenAPIClient) network_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_all_records", session_id)
}

// network_get_all
//
// Return a list of all the networks known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - network ref set
// - references to all objects
func (client *XenAPIClient) network_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_all", session_id)
}

// network_set_default_locking_mode
//
// Set the default locking mode for VIFs attached to this network
//
// params:
// - session_id, session ref, Reference to a valid session
// - network, network ref, The network
// - value, enum network_default_locking_mode, The default locking mode for VIFs attached to this network.
//
// returns:
// - void
func (client *XenAPIClient) network_set_default_locking_mode(session_id interface{}, network interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("network.set_default_locking_mode", session_id, network, value)
}

// network_create_new_blob
//
// Create a placeholder for a named binary blob of data that is associated with this pool
//
// params:
// - session_id, session ref, Reference to a valid session
// - network, network ref, The network
// - name, string, The name associated with the blob
// - mime_type, string, The mime type for the data. Empty string translates to application/octet-stream
// - public, bool, True if the blob should be publicly available
//
// returns:
// - blob ref
// - The reference of the blob, needed for populating its data
func (client *XenAPIClient) network_create_new_blob(session_id interface{}, network interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("network.create_new_blob", session_id, network, name, mime_type, public)
}

// network_remove_tags
//
// Remove the given value from the tags field of the given network.  If the value is not in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - value, string, Value to remove
//
// returns:
// - void
func (client *XenAPIClient) network_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.remove_tags", session_id, self, value)
}

// network_add_tags
//
// Add the given value to the tags field of the given network.  If the value is already in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - value, string, New value to add
//
// returns:
// - void
func (client *XenAPIClient) network_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.add_tags", session_id, self, value)
}

// network_set_tags
//
// Set the tags field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - value, string set, New value to set
//
// returns:
// - void
func (client *XenAPIClient) network_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("network.set_tags", session_id, self, value)
}

// network_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given network.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) network_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("network.remove_from_other_config", session_id, self, key)
}

// network_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) network_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("network.add_to_other_config", session_id, self, key, value)
}

// network_set_other_config
//
// Set the other_config field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) network_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("network.set_other_config", session_id, self, value)
}

// network_set_MTU
//
// Set the MTU field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - value, int, New value to set
//
// returns:
// - void
func (client *XenAPIClient) network_set_MTU(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("network.set_MTU", session_id, self, value)
}

// network_set_name_description
//
// Set the name/description field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) network_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.set_name_description", session_id, self, value)
}

// network_set_name_label
//
// Set the name/label field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) network_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("network.set_name_label", session_id, self, value)
}

// network_get_assigned_ips
//
// Get the assigned_ips field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - (VIF ref -> string) map
// - value of the field
func (client *XenAPIClient) network_get_assigned_ips(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_assigned_ips", session_id, self)
}

// network_get_default_locking_mode
//
// Get the default_locking_mode field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - enum network_default_locking_mode
// - value of the field
func (client *XenAPIClient) network_get_default_locking_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_default_locking_mode", session_id, self)
}

// network_get_tags
//
// Get the tags field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) network_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_tags", session_id, self)
}

// network_get_blobs
//
// Get the blobs field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - (string -> blob ref) map
// - value of the field
func (client *XenAPIClient) network_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_blobs", session_id, self)
}

// network_get_bridge
//
// Get the bridge field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) network_get_bridge(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_bridge", session_id, self)
}

// network_get_other_config
//
// Get the other_config field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) network_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_other_config", session_id, self)
}

// network_get_MTU
//
// Get the MTU field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) network_get_MTU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_MTU", session_id, self)
}

// network_get_PIFs
//
// Get the PIFs field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - PIF ref set
// - value of the field
func (client *XenAPIClient) network_get_PIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_PIFs", session_id, self)
}

// network_get_VIFs
//
// Get the VIFs field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - VIF ref set
// - value of the field
func (client *XenAPIClient) network_get_VIFs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_VIFs", session_id, self)
}

// network_get_current_operations
//
// Get the current_operations field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - (string -> enum network_operations) map
// - value of the field
func (client *XenAPIClient) network_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_current_operations", session_id, self)
}

// network_get_allowed_operations
//
// Get the allowed_operations field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - enum network_operations set
// - value of the field
func (client *XenAPIClient) network_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_allowed_operations", session_id, self)
}

// network_get_name_description
//
// Get the name/description field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) network_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_name_description", session_id, self)
}

// network_get_name_label
//
// Get the name/label field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) network_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_name_label", session_id, self)
}

// network_get_uuid
//
// Get the uuid field of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) network_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_uuid", session_id, self)
}

// network_get_by_name_label
//
// Get all the network instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - network ref set
// - references to objects with matching names
func (client *XenAPIClient) network_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("network.get_by_name_label", session_id, label)
}

// network_destroy
//
// Destroy the specified network instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) network_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.destroy", session_id, self)
}

// network_create
//
// Create a new network instance, and return its handle.
// The constructor args are: name_label, name_description, MTU, other_config*, tags (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, network record, All constructor arguments
//
// returns:
// - network ref
// - reference to the newly created object
func (client *XenAPIClient) network_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("network.create", session_id, args)
}

// network_get_by_uuid
//
// Get a reference to the network instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - network ref
// - reference to the object
func (client *XenAPIClient) network_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("network.get_by_uuid", session_id, uuid)
}

// network_get_record
//
// Get a record containing the current state of the given network.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, network ref, reference to the object
//
// returns:
// - network record
// - all fields from the object
func (client *XenAPIClient) network_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("network.get_record", session_id, self)
}

// VIF_get_all_records
//
// Return a map of VIF references to VIF records for all VIFs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VIF ref -> VIF record) map
// - records of all objects
func (client *XenAPIClient) VIF_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_all_records", session_id)
}

// VIF_get_all
//
// Return a list of all the VIFs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VIF ref set
// - references to all objects
func (client *XenAPIClient) VIF_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_all", session_id)
}

// VIF_remove_ipv6_allowed
//
// Removes an IPv6 address from this VIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF from which the IP address will be removed
// - value, string, The IP address which will be removed from the VIF
//
// returns:
// - void
func (client *XenAPIClient) VIF_remove_ipv6_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_ipv6_allowed", session_id, self, value)
}

// VIF_add_ipv6_allowed
//
// Associates an IPv6 address with this VIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF which the IP address will be associated with
// - value, string, The IP address which will be associated with the VIF
//
// returns:
// - void
func (client *XenAPIClient) VIF_add_ipv6_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_ipv6_allowed", session_id, self, value)
}

// VIF_set_ipv6_allowed
//
// Set the IPv6 addresses to which traffic on this VIF can be restricted
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF which the IP addresses will be associated with
// - value, string set, The IP addresses which will be associated with the VIF
//
// returns:
// - void
func (client *XenAPIClient) VIF_set_ipv6_allowed(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.set_ipv6_allowed", session_id, self, value)
}

// VIF_remove_ipv4_allowed
//
// Removes an IPv4 address from this VIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF from which the IP address will be removed
// - value, string, The IP address which will be removed from the VIF
//
// returns:
// - void
func (client *XenAPIClient) VIF_remove_ipv4_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_ipv4_allowed", session_id, self, value)
}

// VIF_add_ipv4_allowed
//
// Associates an IPv4 address with this VIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF which the IP address will be associated with
// - value, string, The IP address which will be associated with the VIF
//
// returns:
// - void
func (client *XenAPIClient) VIF_add_ipv4_allowed(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_ipv4_allowed", session_id, self, value)
}

// VIF_set_ipv4_allowed
//
// Set the IPv4 addresses to which traffic on this VIF can be restricted
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF which the IP addresses will be associated with
// - value, string set, The IP addresses which will be associated with the VIF
//
// returns:
// - void
func (client *XenAPIClient) VIF_set_ipv4_allowed(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.set_ipv4_allowed", session_id, self, value)
}

// VIF_set_locking_mode
//
// Set the locking mode for this VIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF whose locking mode will be set
// - value, enum vif_locking_mode, The new locking mode for the VIF
//
// returns:
// - void
func (client *XenAPIClient) VIF_set_locking_mode(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.set_locking_mode", session_id, self, value)
}

// VIF_unplug_force
//
// Forcibly unplug the specified VIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF to forcibly unplug
//
// returns:
// - void
func (client *XenAPIClient) VIF_unplug_force(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.unplug_force", session_id, self)
}

// VIF_unplug
//
// Hot-unplug the specified VIF, dynamically unattaching it from the running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF to hot-unplug
//
// returns:
// - void
func (client *XenAPIClient) VIF_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.unplug", session_id, self)
}

// VIF_plug
//
// Hotplug the specified VIF, dynamically attaching it to the running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, The VIF to hotplug
//
// returns:
// - void
func (client *XenAPIClient) VIF_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.plug", session_id, self)
}

// VIF_remove_from_qos_algorithm_params
//
// Remove the given key and its corresponding value from the qos/algorithm_params field of the given VIF.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VIF_remove_from_qos_algorithm_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_from_qos_algorithm_params", session_id, self, key)
}

// VIF_add_to_qos_algorithm_params
//
// Add the given key-value pair to the qos/algorithm_params field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VIF_add_to_qos_algorithm_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_to_qos_algorithm_params", session_id, self, key, value)
}

// VIF_set_qos_algorithm_params
//
// Set the qos/algorithm_params field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VIF_set_qos_algorithm_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VIF.set_qos_algorithm_params", session_id, self, value)
}

// VIF_set_qos_algorithm_type
//
// Set the qos/algorithm_type field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VIF_set_qos_algorithm_type(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.set_qos_algorithm_type", session_id, self, value)
}

// VIF_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VIF.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VIF_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VIF.remove_from_other_config", session_id, self, key)
}

// VIF_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VIF_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VIF.add_to_other_config", session_id, self, key, value)
}

// VIF_set_other_config
//
// Set the other_config field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VIF_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VIF.set_other_config", session_id, self, value)
}

// VIF_get_ipv6_allowed
//
// Get the ipv6_allowed field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VIF_get_ipv6_allowed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_ipv6_allowed", session_id, self)
}

// VIF_get_ipv4_allowed
//
// Get the ipv4_allowed field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VIF_get_ipv4_allowed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_ipv4_allowed", session_id, self)
}

// VIF_get_locking_mode
//
// Get the locking_mode field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - enum vif_locking_mode
// - value of the field
func (client *XenAPIClient) VIF_get_locking_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_locking_mode", session_id, self)
}

// VIF_get_MAC_autogenerated
//
// Get the MAC_autogenerated field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VIF_get_MAC_autogenerated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_MAC_autogenerated", session_id, self)
}

// VIF_get_metrics
//
// Get the metrics field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - VIF_metrics ref
// - value of the field
func (client *XenAPIClient) VIF_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_metrics", session_id, self)
}

// VIF_get_qos_supported_algorithms
//
// Get the qos/supported_algorithms field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VIF_get_qos_supported_algorithms(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_qos_supported_algorithms", session_id, self)
}

// VIF_get_qos_algorithm_params
//
// Get the qos/algorithm_params field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VIF_get_qos_algorithm_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_qos_algorithm_params", session_id, self)
}

// VIF_get_qos_algorithm_type
//
// Get the qos/algorithm_type field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VIF_get_qos_algorithm_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_qos_algorithm_type", session_id, self)
}

// VIF_get_runtime_properties
//
// Get the runtime_properties field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VIF_get_runtime_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_runtime_properties", session_id, self)
}

// VIF_get_status_detail
//
// Get the status_detail field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VIF_get_status_detail(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_status_detail", session_id, self)
}

// VIF_get_status_code
//
// Get the status_code field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VIF_get_status_code(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_status_code", session_id, self)
}

// VIF_get_currently_attached
//
// Get the currently_attached field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VIF_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_currently_attached", session_id, self)
}

// VIF_get_other_config
//
// Get the other_config field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VIF_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_other_config", session_id, self)
}

// VIF_get_MTU
//
// Get the MTU field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VIF_get_MTU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_MTU", session_id, self)
}

// VIF_get_MAC
//
// Get the MAC field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VIF_get_MAC(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_MAC", session_id, self)
}

// VIF_get_VM
//
// Get the VM field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) VIF_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_VM", session_id, self)
}

// VIF_get_network
//
// Get the network field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - network ref
// - value of the field
func (client *XenAPIClient) VIF_get_network(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_network", session_id, self)
}

// VIF_get_device
//
// Get the device field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VIF_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_device", session_id, self)
}

// VIF_get_current_operations
//
// Get the current_operations field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - (string -> enum vif_operations) map
// - value of the field
func (client *XenAPIClient) VIF_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_current_operations", session_id, self)
}

// VIF_get_allowed_operations
//
// Get the allowed_operations field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - enum vif_operations set
// - value of the field
func (client *XenAPIClient) VIF_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_allowed_operations", session_id, self)
}

// VIF_get_uuid
//
// Get the uuid field of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VIF_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_uuid", session_id, self)
}

// VIF_destroy
//
// Destroy the specified VIF instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VIF_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.destroy", session_id, self)
}

// VIF_create
//
// Create a new VIF instance, and return its handle.
// The constructor args are: device*, network*, VM*, MAC*, MTU*, other_config*, qos_algorithm_type*, qos_algorithm_params*, locking_mode, ipv4_allowed, ipv6_allowed (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, VIF record, All constructor arguments
//
// returns:
// - VIF ref
// - reference to the newly created object
func (client *XenAPIClient) VIF_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.create", session_id, args)
}

// VIF_get_by_uuid
//
// Get a reference to the VIF instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VIF ref
// - reference to the object
func (client *XenAPIClient) VIF_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VIF.get_by_uuid", session_id, uuid)
}

// VIF_get_record
//
// Get a record containing the current state of the given VIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF ref, reference to the object
//
// returns:
// - VIF record
// - all fields from the object
func (client *XenAPIClient) VIF_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF.get_record", session_id, self)
}

// VIF_metrics_get_all_records
//
// Return a map of VIF_metrics references to VIF_metrics records for all VIF_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VIF_metrics ref -> VIF_metrics record) map
// - records of all objects
func (client *XenAPIClient) VIF_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_all_records", session_id)
}

// VIF_metrics_get_all
//
// Return a list of all the VIF_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VIF_metrics ref set
// - references to all objects
func (client *XenAPIClient) VIF_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_all", session_id)
}

// VIF_metrics_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VIF_metrics.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VIF_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.remove_from_other_config", session_id, self, key)
}

// VIF_metrics_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VIF_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.add_to_other_config", session_id, self, key, value)
}

// VIF_metrics_set_other_config
//
// Set the other_config field of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VIF_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.set_other_config", session_id, self, value)
}

// VIF_metrics_get_other_config
//
// Get the other_config field of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VIF_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_other_config", session_id, self)
}

// VIF_metrics_get_last_updated
//
// Get the last_updated field of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VIF_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_last_updated", session_id, self)
}

// VIF_metrics_get_io_write_kbs
//
// Get the io/write_kbs field of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) VIF_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_io_write_kbs", session_id, self)
}

// VIF_metrics_get_io_read_kbs
//
// Get the io/read_kbs field of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) VIF_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_io_read_kbs", session_id, self)
}

// VIF_metrics_get_uuid
//
// Get the uuid field of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VIF_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_uuid", session_id, self)
}

// VIF_metrics_get_by_uuid
//
// Get a reference to the VIF_metrics instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VIF_metrics ref
// - reference to the object
func (client *XenAPIClient) VIF_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_by_uuid", session_id, uuid)
}

// VIF_metrics_get_record
//
// Get a record containing the current state of the given VIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VIF_metrics ref, reference to the object
//
// returns:
// - VIF_metrics record
// - all fields from the object
func (client *XenAPIClient) VIF_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VIF_metrics.get_record", session_id, self)
}

// PIF_get_all_records
//
// Return a map of PIF references to PIF records for all PIFs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (PIF ref -> PIF record) map
// - records of all objects
func (client *XenAPIClient) PIF_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_all_records", session_id)
}

// PIF_get_all
//
// Return a list of all the PIFs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - PIF ref set
// - references to all objects
func (client *XenAPIClient) PIF_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_all", session_id)
}

// PIF_set_property
//
// Set the value of a property of the PIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, The PIF
// - name, string, The property name
// - value, string, The property value
//
// returns:
// - void
func (client *XenAPIClient) PIF_set_property(session_id interface{}, self interface{}, name string, value string) (i interface{}, err error) {
	return client.RPCCall("PIF.set_property", session_id, self, name, value)
}

// PIF_db_forget
//
// Destroy a PIF database record.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, The ref of the PIF whose database record should be destroyed
//
// returns:
// - void
func (client *XenAPIClient) PIF_db_forget(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.db_forget", session_id, self)
}

// PIF_db_introduce
//
// Create a new PIF record in the database only
//
// params:
// - session_id, session ref, Reference to a valid session
// - device, string,
// - network, network ref,
// - host, host ref,
// - MAC, string,
// - MTU, int,
// - VLAN, int,
// - physical, bool,
// - ip_configuration_mode, enum ip_configuration_mode,
// - IP, string,
// - netmask, string,
// - gateway, string,
// - DNS, string,
// - bond_slave_of, Bond ref,
// - VLAN_master_of, VLAN ref,
// - management, bool,
// - other_config, (string -> string) map,
// - disallow_unplug, bool,
// - ipv6_configuration_mode, enum ipv6_configuration_mode,
// - IPv6, string set,
// - ipv6_gateway, string,
// - primary_address_type, enum primary_address_type,
// - managed, bool,
// - properties, (string -> string) map,
//
// returns:
// - PIF ref
// - The ref of the newly created PIF record.
func (client *XenAPIClient) PIF_db_introduce(session_id interface{}, device string, network interface{}, host interface{}, MAC string, MTU interface{}, VLAN interface{}, physical bool, ip_configuration_mode interface{}, IP string, netmask string, gateway string, DNS string, bond_slave_of interface{}, VLAN_master_of interface{}, management bool, other_config map[string]string, disallow_unplug bool, ipv6_configuration_mode interface{}, IPv6 interface{}, ipv6_gateway string, primary_address_type interface{}, managed bool, properties map[string]string) (i interface{}, err error) {
	return client.RPCCall("PIF.db_introduce", session_id, device, network, host, MAC, MTU, VLAN, physical, ip_configuration_mode, IP, netmask, gateway, DNS, bond_slave_of, VLAN_master_of, management, other_config, disallow_unplug, ipv6_configuration_mode, IPv6, ipv6_gateway, primary_address_type, managed, properties)
}

// PIF_plug
//
// Attempt to bring up a physical interface
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, the PIF object to plug
//
// returns:
// - void
func (client *XenAPIClient) PIF_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.plug", session_id, self)
}

// PIF_unplug
//
// Attempt to bring down a physical interface
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, the PIF object to unplug
//
// returns:
// - void
func (client *XenAPIClient) PIF_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.unplug", session_id, self)
}

// PIF_forget
//
// Destroy the PIF object matching a particular network interface
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, The PIF object to destroy
//
// returns:
// - void
func (client *XenAPIClient) PIF_forget(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.forget", session_id, self)
}

// PIF_introduce
//
// Create a PIF object matching a particular network interface
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host on which the interface exists
// - MAC, string, The MAC address of the interface
// - device, string, The device name to use for the interface
// - managed, bool, Indicates whether the interface is managed by xapi (defaults to "true")
//
// returns:
// - PIF ref
// - The reference of the created PIF object
func (client *XenAPIClient) PIF_introduce(session_id interface{}, host interface{}, MAC string, device string, managed bool) (i interface{}, err error) {
	return client.RPCCall("PIF.introduce", session_id, host, MAC, device, managed)
}

// PIF_scan
//
// Scan for physical interfaces on a host and create PIF objects to represent them
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host on which to scan
//
// returns:
// - void
func (client *XenAPIClient) PIF_scan(session_id interface{}, host interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.scan", session_id, host)
}

// PIF_set_primary_address_type
//
// Change the primary address type used by this PIF
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, the PIF object to reconfigure
// - primary_address_type, enum primary_address_type, Whether to prefer IPv4 or IPv6 connections
//
// returns:
// - void
func (client *XenAPIClient) PIF_set_primary_address_type(session_id interface{}, self interface{}, primary_address_type interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.set_primary_address_type", session_id, self, primary_address_type)
}

// PIF_reconfigure_ipv6
//
// Reconfigure the IPv6 address settings for this interface
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, the PIF object to reconfigure
// - mode, enum ipv6_configuration_mode, whether to use dynamic/static/no-assignment
// - IPv6, string, the new IPv6 address (in <addr>/<prefix length> format)
// - gateway, string, the new gateway
// - DNS, string, the new DNS settings
//
// returns:
// - void
func (client *XenAPIClient) PIF_reconfigure_ipv6(session_id interface{}, self interface{}, mode interface{}, IPv6 string, gateway string, DNS string) (i interface{}, err error) {
	return client.RPCCall("PIF.reconfigure_ipv6", session_id, self, mode, IPv6, gateway, DNS)
}

// PIF_reconfigure_ip
//
// Reconfigure the IP address settings for this interface
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, the PIF object to reconfigure
// - mode, enum ip_configuration_mode, whether to use dynamic/static/no-assignment
// - IP, string, the new IP address
// - netmask, string, the new netmask
// - gateway, string, the new gateway
// - DNS, string, the new DNS settings
//
// returns:
// - void
func (client *XenAPIClient) PIF_reconfigure_ip(session_id interface{}, self interface{}, mode interface{}, IP string, netmask string, gateway string, DNS string) (i interface{}, err error) {
	return client.RPCCall("PIF.reconfigure_ip", session_id, self, mode, IP, netmask, gateway, DNS)
}

// PIF_destroy
//
// Destroy the PIF object (provided it is a VLAN interface). This call is deprecated: use VLAN.destroy or Bond.destroy instead
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, the PIF object to destroy
//
// returns:
// - void
func (client *XenAPIClient) PIF_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.destroy", session_id, self)
}

// PIF_create_VLAN
//
// Create a VLAN interface from an existing physical interface. This call is deprecated: use VLAN.create instead
//
// params:
// - session_id, session ref, Reference to a valid session
// - device, string, physical interface on which to create the VLAN interface
// - network, network ref, network to which this interface should be connected
// - host, host ref, physical machine to which this PIF is connected
// - VLAN, int, VLAN tag for the new interface
//
// returns:
// - PIF ref
// - The reference of the created PIF object
func (client *XenAPIClient) PIF_create_VLAN(session_id interface{}, device string, network interface{}, host interface{}, VLAN interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.create_VLAN", session_id, device, network, host, VLAN)
}

// PIF_set_disallow_unplug
//
// Set the disallow_unplug field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) PIF_set_disallow_unplug(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("PIF.set_disallow_unplug", session_id, self, value)
}

// PIF_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given PIF.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) PIF_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PIF.remove_from_other_config", session_id, self, key)
}

// PIF_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) PIF_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PIF.add_to_other_config", session_id, self, key, value)
}

// PIF_set_other_config
//
// Set the other_config field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) PIF_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PIF.set_other_config", session_id, self, value)
}

// PIF_get_properties
//
// Get the properties field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) PIF_get_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_properties", session_id, self)
}

// PIF_get_managed
//
// Get the managed field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PIF_get_managed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_managed", session_id, self)
}

// PIF_get_primary_address_type
//
// Get the primary_address_type field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - enum primary_address_type
// - value of the field
func (client *XenAPIClient) PIF_get_primary_address_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_primary_address_type", session_id, self)
}

// PIF_get_ipv6_gateway
//
// Get the ipv6_gateway field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_ipv6_gateway(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_ipv6_gateway", session_id, self)
}

// PIF_get_IPv6
//
// Get the IPv6 field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) PIF_get_IPv6(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_IPv6", session_id, self)
}

// PIF_get_ipv6_configuration_mode
//
// Get the ipv6_configuration_mode field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - enum ipv6_configuration_mode
// - value of the field
func (client *XenAPIClient) PIF_get_ipv6_configuration_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_ipv6_configuration_mode", session_id, self)
}

// PIF_get_tunnel_transport_PIF_of
//
// Get the tunnel_transport_PIF_of field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - tunnel ref set
// - value of the field
func (client *XenAPIClient) PIF_get_tunnel_transport_PIF_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_tunnel_transport_PIF_of", session_id, self)
}

// PIF_get_tunnel_access_PIF_of
//
// Get the tunnel_access_PIF_of field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - tunnel ref set
// - value of the field
func (client *XenAPIClient) PIF_get_tunnel_access_PIF_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_tunnel_access_PIF_of", session_id, self)
}

// PIF_get_disallow_unplug
//
// Get the disallow_unplug field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PIF_get_disallow_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_disallow_unplug", session_id, self)
}

// PIF_get_other_config
//
// Get the other_config field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) PIF_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_other_config", session_id, self)
}

// PIF_get_management
//
// Get the management field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PIF_get_management(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_management", session_id, self)
}

// PIF_get_VLAN_slave_of
//
// Get the VLAN_slave_of field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - VLAN ref set
// - value of the field
func (client *XenAPIClient) PIF_get_VLAN_slave_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_VLAN_slave_of", session_id, self)
}

// PIF_get_VLAN_master_of
//
// Get the VLAN_master_of field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - VLAN ref
// - value of the field
func (client *XenAPIClient) PIF_get_VLAN_master_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_VLAN_master_of", session_id, self)
}

// PIF_get_bond_master_of
//
// Get the bond_master_of field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - Bond ref set
// - value of the field
func (client *XenAPIClient) PIF_get_bond_master_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_bond_master_of", session_id, self)
}

// PIF_get_bond_slave_of
//
// Get the bond_slave_of field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - Bond ref
// - value of the field
func (client *XenAPIClient) PIF_get_bond_slave_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_bond_slave_of", session_id, self)
}

// PIF_get_DNS
//
// Get the DNS field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_DNS(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_DNS", session_id, self)
}

// PIF_get_gateway
//
// Get the gateway field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_gateway(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_gateway", session_id, self)
}

// PIF_get_netmask
//
// Get the netmask field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_netmask(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_netmask", session_id, self)
}

// PIF_get_IP
//
// Get the IP field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_IP(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_IP", session_id, self)
}

// PIF_get_ip_configuration_mode
//
// Get the ip_configuration_mode field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - enum ip_configuration_mode
// - value of the field
func (client *XenAPIClient) PIF_get_ip_configuration_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_ip_configuration_mode", session_id, self)
}

// PIF_get_currently_attached
//
// Get the currently_attached field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PIF_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_currently_attached", session_id, self)
}

// PIF_get_physical
//
// Get the physical field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PIF_get_physical(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_physical", session_id, self)
}

// PIF_get_metrics
//
// Get the metrics field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - PIF_metrics ref
// - value of the field
func (client *XenAPIClient) PIF_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_metrics", session_id, self)
}

// PIF_get_VLAN
//
// Get the VLAN field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) PIF_get_VLAN(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_VLAN", session_id, self)
}

// PIF_get_MTU
//
// Get the MTU field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) PIF_get_MTU(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_MTU", session_id, self)
}

// PIF_get_MAC
//
// Get the MAC field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_MAC(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_MAC", session_id, self)
}

// PIF_get_host
//
// Get the host field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) PIF_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_host", session_id, self)
}

// PIF_get_network
//
// Get the network field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - network ref
// - value of the field
func (client *XenAPIClient) PIF_get_network(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_network", session_id, self)
}

// PIF_get_device
//
// Get the device field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_device", session_id, self)
}

// PIF_get_uuid
//
// Get the uuid field of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_uuid", session_id, self)
}

// PIF_get_by_uuid
//
// Get a reference to the PIF instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - PIF ref
// - reference to the object
func (client *XenAPIClient) PIF_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PIF.get_by_uuid", session_id, uuid)
}

// PIF_get_record
//
// Get a record containing the current state of the given PIF.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF ref, reference to the object
//
// returns:
// - PIF record
// - all fields from the object
func (client *XenAPIClient) PIF_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF.get_record", session_id, self)
}

// PIF_metrics_get_all_records
//
// Return a map of PIF_metrics references to PIF_metrics records for all PIF_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (PIF_metrics ref -> PIF_metrics record) map
// - records of all objects
func (client *XenAPIClient) PIF_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_all_records", session_id)
}

// PIF_metrics_get_all
//
// Return a list of all the PIF_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - PIF_metrics ref set
// - references to all objects
func (client *XenAPIClient) PIF_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_all", session_id)
}

// PIF_metrics_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given PIF_metrics.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) PIF_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.remove_from_other_config", session_id, self, key)
}

// PIF_metrics_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) PIF_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.add_to_other_config", session_id, self, key, value)
}

// PIF_metrics_set_other_config
//
// Set the other_config field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) PIF_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.set_other_config", session_id, self, value)
}

// PIF_metrics_get_other_config
//
// Get the other_config field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_other_config", session_id, self)
}

// PIF_metrics_get_last_updated
//
// Get the last_updated field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_last_updated", session_id, self)
}

// PIF_metrics_get_pci_bus_path
//
// Get the pci_bus_path field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_pci_bus_path(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_pci_bus_path", session_id, self)
}

// PIF_metrics_get_duplex
//
// Get the duplex field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_duplex(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_duplex", session_id, self)
}

// PIF_metrics_get_speed
//
// Get the speed field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_speed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_speed", session_id, self)
}

// PIF_metrics_get_device_name
//
// Get the device_name field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_device_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_device_name", session_id, self)
}

// PIF_metrics_get_device_id
//
// Get the device_id field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_device_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_device_id", session_id, self)
}

// PIF_metrics_get_vendor_name
//
// Get the vendor_name field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_vendor_name", session_id, self)
}

// PIF_metrics_get_vendor_id
//
// Get the vendor_id field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_vendor_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_vendor_id", session_id, self)
}

// PIF_metrics_get_carrier
//
// Get the carrier field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_carrier(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_carrier", session_id, self)
}

// PIF_metrics_get_io_write_kbs
//
// Get the io/write_kbs field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_io_write_kbs", session_id, self)
}

// PIF_metrics_get_io_read_kbs
//
// Get the io/read_kbs field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_io_read_kbs", session_id, self)
}

// PIF_metrics_get_uuid
//
// Get the uuid field of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PIF_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_uuid", session_id, self)
}

// PIF_metrics_get_by_uuid
//
// Get a reference to the PIF_metrics instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - PIF_metrics ref
// - reference to the object
func (client *XenAPIClient) PIF_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_by_uuid", session_id, uuid)
}

// PIF_metrics_get_record
//
// Get a record containing the current state of the given PIF_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PIF_metrics ref, reference to the object
//
// returns:
// - PIF_metrics record
// - all fields from the object
func (client *XenAPIClient) PIF_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PIF_metrics.get_record", session_id, self)
}

// Bond_get_all_records
//
// Return a map of Bond references to Bond records for all Bonds known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (Bond ref -> Bond record) map
// - records of all objects
func (client *XenAPIClient) Bond_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_all_records", session_id)
}

// Bond_get_all
//
// Return a list of all the Bonds known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - Bond ref set
// - references to all objects
func (client *XenAPIClient) Bond_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_all", session_id)
}

// Bond_set_property
//
// Set the value of a property of the bond
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, The bond
// - name, string, The property name
// - value, string, The property value
//
// returns:
// - void
func (client *XenAPIClient) Bond_set_property(session_id interface{}, self interface{}, name string, value string) (i interface{}, err error) {
	return client.RPCCall("Bond.set_property", session_id, self, name, value)
}

// Bond_set_mode
//
// Change the bond mode
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, The bond
// - value, enum bond_mode, The new bond mode
//
// returns:
// - void
func (client *XenAPIClient) Bond_set_mode(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.set_mode", session_id, self, value)
}

// Bond_destroy
//
// Destroy an interface bond
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, Bond to destroy
//
// returns:
// - void
func (client *XenAPIClient) Bond_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.destroy", session_id, self)
}

// Bond_create
//
// Create an interface bond
//
// params:
// - session_id, session ref, Reference to a valid session
// - network, network ref, Network to add the bonded PIF to
// - members, PIF ref set, PIFs to add to this bond
// - MAC, string, The MAC address to use on the bond itself. If this parameter is the empty string then the bond will inherit its MAC address from the primary slave.
// - mode, enum bond_mode, Bonding mode to use for the new bond
// - properties, (string -> string) map, Additional configuration parameters specific to the bond mode
//
// returns:
// - Bond ref
// - The reference of the created Bond object
func (client *XenAPIClient) Bond_create(session_id interface{}, network interface{}, members interface{}, MAC string, mode interface{}, properties map[string]string) (i interface{}, err error) {
	return client.RPCCall("Bond.create", session_id, network, members, MAC, mode, properties)
}

// Bond_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given Bond.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) Bond_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("Bond.remove_from_other_config", session_id, self, key)
}

// Bond_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) Bond_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("Bond.add_to_other_config", session_id, self, key, value)
}

// Bond_set_other_config
//
// Set the other_config field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) Bond_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("Bond.set_other_config", session_id, self, value)
}

// Bond_get_links_up
//
// Get the links_up field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) Bond_get_links_up(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_links_up", session_id, self)
}

// Bond_get_properties
//
// Get the properties field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) Bond_get_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_properties", session_id, self)
}

// Bond_get_mode
//
// Get the mode field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - enum bond_mode
// - value of the field
func (client *XenAPIClient) Bond_get_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_mode", session_id, self)
}

// Bond_get_primary_slave
//
// Get the primary_slave field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - PIF ref
// - value of the field
func (client *XenAPIClient) Bond_get_primary_slave(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_primary_slave", session_id, self)
}

// Bond_get_other_config
//
// Get the other_config field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) Bond_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_other_config", session_id, self)
}

// Bond_get_slaves
//
// Get the slaves field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - PIF ref set
// - value of the field
func (client *XenAPIClient) Bond_get_slaves(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_slaves", session_id, self)
}

// Bond_get_master
//
// Get the master field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - PIF ref
// - value of the field
func (client *XenAPIClient) Bond_get_master(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_master", session_id, self)
}

// Bond_get_uuid
//
// Get the uuid field of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) Bond_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_uuid", session_id, self)
}

// Bond_get_by_uuid
//
// Get a reference to the Bond instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - Bond ref
// - reference to the object
func (client *XenAPIClient) Bond_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("Bond.get_by_uuid", session_id, uuid)
}

// Bond_get_record
//
// Get a record containing the current state of the given Bond.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, Bond ref, reference to the object
//
// returns:
// - Bond record
// - all fields from the object
func (client *XenAPIClient) Bond_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("Bond.get_record", session_id, self)
}

// VLAN_get_all_records
//
// Return a map of VLAN references to VLAN records for all VLANs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VLAN ref -> VLAN record) map
// - records of all objects
func (client *XenAPIClient) VLAN_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_all_records", session_id)
}

// VLAN_get_all
//
// Return a list of all the VLANs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VLAN ref set
// - references to all objects
func (client *XenAPIClient) VLAN_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_all", session_id)
}

// VLAN_destroy
//
// Destroy a VLAN mux/demuxer
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, VLAN mux/demuxer to destroy
//
// returns:
// - void
func (client *XenAPIClient) VLAN_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.destroy", session_id, self)
}

// VLAN_create
//
// Create a VLAN mux/demuxer
//
// params:
// - session_id, session ref, Reference to a valid session
// - tagged_PIF, PIF ref, PIF which receives the tagged traffic
// - tag, int, VLAN tag to use
// - network, network ref, Network to receive the untagged traffic
//
// returns:
// - VLAN ref
// - The reference of the created VLAN object
func (client *XenAPIClient) VLAN_create(session_id interface{}, tagged_PIF interface{}, tag interface{}, network interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.create", session_id, tagged_PIF, tag, network)
}

// VLAN_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VLAN.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VLAN_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VLAN.remove_from_other_config", session_id, self, key)
}

// VLAN_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VLAN_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VLAN.add_to_other_config", session_id, self, key, value)
}

// VLAN_set_other_config
//
// Set the other_config field of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VLAN_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VLAN.set_other_config", session_id, self, value)
}

// VLAN_get_other_config
//
// Get the other_config field of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VLAN_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_other_config", session_id, self)
}

// VLAN_get_tag
//
// Get the tag field of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VLAN_get_tag(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_tag", session_id, self)
}

// VLAN_get_untagged_PIF
//
// Get the untagged_PIF field of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
//
// returns:
// - PIF ref
// - value of the field
func (client *XenAPIClient) VLAN_get_untagged_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_untagged_PIF", session_id, self)
}

// VLAN_get_tagged_PIF
//
// Get the tagged_PIF field of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
//
// returns:
// - PIF ref
// - value of the field
func (client *XenAPIClient) VLAN_get_tagged_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_tagged_PIF", session_id, self)
}

// VLAN_get_uuid
//
// Get the uuid field of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VLAN_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_uuid", session_id, self)
}

// VLAN_get_by_uuid
//
// Get a reference to the VLAN instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VLAN ref
// - reference to the object
func (client *XenAPIClient) VLAN_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_by_uuid", session_id, uuid)
}

// VLAN_get_record
//
// Get a record containing the current state of the given VLAN.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VLAN ref, reference to the object
//
// returns:
// - VLAN record
// - all fields from the object
func (client *XenAPIClient) VLAN_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VLAN.get_record", session_id, self)
}

// SM_get_all_records
//
// Return a map of SM references to SM records for all SMs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (SM ref -> SM record) map
// - records of all objects
func (client *XenAPIClient) SM_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_all_records", session_id)
}

// SM_get_all
//
// Return a list of all the SMs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - SM ref set
// - references to all objects
func (client *XenAPIClient) SM_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_all", session_id)
}

// SM_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given SM.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) SM_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("SM.remove_from_other_config", session_id, self, key)
}

// SM_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) SM_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("SM.add_to_other_config", session_id, self, key, value)
}

// SM_set_other_config
//
// Set the other_config field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) SM_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("SM.set_other_config", session_id, self, value)
}

// SM_get_driver_filename
//
// Get the driver_filename field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_driver_filename(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_driver_filename", session_id, self)
}

// SM_get_other_config
//
// Get the other_config field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) SM_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_other_config", session_id, self)
}

// SM_get_features
//
// Get the features field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - (string -> int) map
// - value of the field
func (client *XenAPIClient) SM_get_features(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_features", session_id, self)
}

// SM_get_capabilities
//
// Get the capabilities field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) SM_get_capabilities(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_capabilities", session_id, self)
}

// SM_get_configuration
//
// Get the configuration field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) SM_get_configuration(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_configuration", session_id, self)
}

// SM_get_required_api_version
//
// Get the required_api_version field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_required_api_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_required_api_version", session_id, self)
}

// SM_get_version
//
// Get the version field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_version(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_version", session_id, self)
}

// SM_get_copyright
//
// Get the copyright field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_copyright(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_copyright", session_id, self)
}

// SM_get_vendor
//
// Get the vendor field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_vendor(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_vendor", session_id, self)
}

// SM_get_type
//
// Get the type field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_type", session_id, self)
}

// SM_get_name_description
//
// Get the name/description field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_name_description", session_id, self)
}

// SM_get_name_label
//
// Get the name/label field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_name_label", session_id, self)
}

// SM_get_uuid
//
// Get the uuid field of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SM_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_uuid", session_id, self)
}

// SM_get_by_name_label
//
// Get all the SM instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - SM ref set
// - references to objects with matching names
func (client *XenAPIClient) SM_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("SM.get_by_name_label", session_id, label)
}

// SM_get_by_uuid
//
// Get a reference to the SM instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - SM ref
// - reference to the object
func (client *XenAPIClient) SM_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("SM.get_by_uuid", session_id, uuid)
}

// SM_get_record
//
// Get a record containing the current state of the given SM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SM ref, reference to the object
//
// returns:
// - SM record
// - all fields from the object
func (client *XenAPIClient) SM_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SM.get_record", session_id, self)
}

// SR_get_all_records
//
// Return a map of SR references to SR records for all SRs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (SR ref -> SR record) map
// - records of all objects
func (client *XenAPIClient) SR_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_all_records", session_id)
}

// SR_get_all
//
// Return a list of all the SRs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - SR ref set
// - references to all objects
func (client *XenAPIClient) SR_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_all", session_id)
}

// SR_disable_database_replication
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR to which metadata should be no longer replicated
//
// returns:
// - void
func (client *XenAPIClient) SR_disable_database_replication(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.disable_database_replication", session_id, sr)
}

// SR_enable_database_replication
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR to which metadata should be replicated
//
// returns:
// - void
func (client *XenAPIClient) SR_enable_database_replication(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.enable_database_replication", session_id, sr)
}

// SR_assert_supports_database_replication
//
// Returns successfully if the given SR supports database replication. Otherwise returns an error to explain why not.
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR to query
//
// returns:
// - void
func (client *XenAPIClient) SR_assert_supports_database_replication(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.assert_supports_database_replication", session_id, sr)
}

// SR_assert_can_host_ha_statefile
//
// Returns successfully if the given SR can host an HA statefile. Otherwise returns an error to explain why not
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR to query
//
// returns:
// - void
func (client *XenAPIClient) SR_assert_can_host_ha_statefile(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.assert_can_host_ha_statefile", session_id, sr)
}

// SR_set_physical_utilisation
//
// Sets the SR's physical_utilisation field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, The SR to modify
// - value, int, The new value of the SR's physical utilisation
//
// returns:
// - void
func (client *XenAPIClient) SR_set_physical_utilisation(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_physical_utilisation", session_id, self, value)
}

// SR_set_virtual_allocation
//
// Sets the SR's virtual_allocation field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, The SR to modify
// - value, int, The new value of the SR's virtual_allocation
//
// returns:
// - void
func (client *XenAPIClient) SR_set_virtual_allocation(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_virtual_allocation", session_id, self, value)
}

// SR_set_physical_size
//
// Sets the SR's physical_size field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, The SR to modify
// - value, int, The new value of the SR's physical_size
//
// returns:
// - void
func (client *XenAPIClient) SR_set_physical_size(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_physical_size", session_id, self, value)
}

// SR_create_new_blob
//
// Create a placeholder for a named binary blob of data that is associated with this SR
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR
// - name, string, The name associated with the blob
// - mime_type, string, The mime type for the data. Empty string translates to application/octet-stream
// - public, bool, True if the blob should be publicly available
//
// returns:
// - blob ref
// - The reference of the blob, needed for populating its data
func (client *XenAPIClient) SR_create_new_blob(session_id interface{}, sr interface{}, name string, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("SR.create_new_blob", session_id, sr, name, mime_type, public)
}

// SR_set_name_description
//
// Set the name description of the SR
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR
// - value, string, The name description for the SR
//
// returns:
// - void
func (client *XenAPIClient) SR_set_name_description(session_id interface{}, sr interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.set_name_description", session_id, sr, value)
}

// SR_set_name_label
//
// Set the name label of the SR
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR
// - value, string, The name label for the SR
//
// returns:
// - void
func (client *XenAPIClient) SR_set_name_label(session_id interface{}, sr interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.set_name_label", session_id, sr, value)
}

// SR_set_shared
//
// Sets the shared flag on the SR
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR
// - value, bool, True if the SR is shared
//
// returns:
// - void
func (client *XenAPIClient) SR_set_shared(session_id interface{}, sr interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("SR.set_shared", session_id, sr, value)
}

// SR_probe
//
// Perform a backend-specific scan, using the given device_config.  If the device_config is complete, then this will return a list of the SRs present of this type on the device, if any.  If the device_config is partial, then a backend-specific scan will be performed, returning results that will guide the user in improving the device_config.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to create/make the SR on
// - device_config, (string -> string) map, The device config string that will be passed to backend SR driver
// - a_type, string, The type of the SR; used to specify the SR backend driver to use
// - sm_config, (string -> string) map, Storage backend specific configuration options
//
// returns:
// - string
// - An XML fragment containing the scan results.  These are specific to the scan being performed, and the backend.
func (client *XenAPIClient) SR_probe(session_id interface{}, host interface{}, device_config map[string]string, a_type string, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.probe", session_id, host, device_config, a_type, sm_config)
}

// SR_scan
//
// Refreshes the list of VDIs associated with an SR
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR to scan
//
// returns:
// - void
func (client *XenAPIClient) SR_scan(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.scan", session_id, sr)
}

// SR_get_supported_types
//
// Return a set of all the SR types supported by the system
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - string set
// - the supported SR types
func (client *XenAPIClient) SR_get_supported_types(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_supported_types", session_id)
}

// SR_update
//
// Refresh the fields on the SR object
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR whose fields should be refreshed
//
// returns:
// - void
func (client *XenAPIClient) SR_update(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.update", session_id, sr)
}

// SR_forget
//
// Removing specified SR-record from database, without attempting to remove SR from disk
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR to destroy
//
// returns:
// - void
func (client *XenAPIClient) SR_forget(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.forget", session_id, sr)
}

// SR_destroy
//
// Destroy specified SR, removing SR-record from database and remove SR from disk. (In order to affect this operation the appropriate device_config is read from the specified SR's PBD on current host)
//
// params:
// - session_id, session ref, Reference to a valid session
// - sr, SR ref, The SR to destroy
//
// returns:
// - void
func (client *XenAPIClient) SR_destroy(session_id interface{}, sr interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.destroy", session_id, sr)
}

// SR_make
//
// Create a new Storage Repository on disk. This call is deprecated: use SR.create instead.
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to create/make the SR on
// - device_config, (string -> string) map, The device config string that will be passed to backend SR driver
// - physical_size, int, The physical size of the new storage repository
// - name_label, string, The name of the new storage repository
// - name_description, string, The description of the new storage repository
// - a_type, string, The type of the SR; used to specify the SR backend driver to use
// - content_type, string, The type of the new SRs content, if required (e.g. ISOs)
// - sm_config, (string -> string) map, Storage backend specific configuration options
//
// returns:
// - string
// - The uuid of the newly created Storage Repository.
func (client *XenAPIClient) SR_make(session_id interface{}, host interface{}, device_config map[string]string, physical_size interface{}, name_label string, name_description string, a_type string, content_type string, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.make", session_id, host, device_config, physical_size, name_label, name_description, a_type, content_type, sm_config)
}

// SR_introduce
//
// Introduce a new Storage Repository into the managed system
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, The uuid assigned to the introduced SR
// - name_label, string, The name of the new storage repository
// - name_description, string, The description of the new storage repository
// - a_type, string, The type of the SR; used to specify the SR backend driver to use
// - content_type, string, The type of the new SRs content, if required (e.g. ISOs)
// - shared, bool, True if the SR (is capable of) being shared by multiple hosts
// - sm_config, (string -> string) map, Storage backend specific configuration options
//
// returns:
// - SR ref
// - The reference of the newly introduced Storage Repository.
func (client *XenAPIClient) SR_introduce(session_id interface{}, uuid string, name_label string, name_description string, a_type string, content_type string, shared bool, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.introduce", session_id, uuid, name_label, name_description, a_type, content_type, shared, sm_config)
}

// SR_create
//
// Create a new Storage Repository and introduce it into the managed system, creating both SR record and PBD record to attach it to current host (with specified device_config parameters)
//
// params:
// - session_id, session ref, Reference to a valid session
// - host, host ref, The host to create/make the SR on
// - device_config, (string -> string) map, The device config string that will be passed to backend SR driver
// - physical_size, int, The physical size of the new storage repository
// - name_label, string, The name of the new storage repository
// - name_description, string, The description of the new storage repository
// - a_type, string, The type of the SR; used to specify the SR backend driver to use
// - content_type, string, The type of the new SRs content, if required (e.g. ISOs)
// - shared, bool, True if the SR (is capable of) being shared by multiple hosts
// - sm_config, (string -> string) map, Storage backend specific configuration options
//
// returns:
// - SR ref
// - The reference of the newly created Storage Repository.
func (client *XenAPIClient) SR_create(session_id interface{}, host interface{}, device_config map[string]string, physical_size interface{}, name_label string, name_description string, a_type string, content_type string, shared bool, sm_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.create", session_id, host, device_config, physical_size, name_label, name_description, a_type, content_type, shared, sm_config)
}

// SR_remove_from_sm_config
//
// Remove the given key and its corresponding value from the sm_config field of the given SR.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) SR_remove_from_sm_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("SR.remove_from_sm_config", session_id, self, key)
}

// SR_add_to_sm_config
//
// Add the given key-value pair to the sm_config field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) SR_add_to_sm_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("SR.add_to_sm_config", session_id, self, key, value)
}

// SR_set_sm_config
//
// Set the sm_config field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) SR_set_sm_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.set_sm_config", session_id, self, value)
}

// SR_remove_tags
//
// Remove the given value from the tags field of the given SR.  If the value is not in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - value, string, Value to remove
//
// returns:
// - void
func (client *XenAPIClient) SR_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.remove_tags", session_id, self, value)
}

// SR_add_tags
//
// Add the given value to the tags field of the given SR.  If the value is already in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - value, string, New value to add
//
// returns:
// - void
func (client *XenAPIClient) SR_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("SR.add_tags", session_id, self, value)
}

// SR_set_tags
//
// Set the tags field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - value, string set, New value to set
//
// returns:
// - void
func (client *XenAPIClient) SR_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.set_tags", session_id, self, value)
}

// SR_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given SR.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) SR_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("SR.remove_from_other_config", session_id, self, key)
}

// SR_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) SR_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("SR.add_to_other_config", session_id, self, key, value)
}

// SR_set_other_config
//
// Set the other_config field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) SR_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("SR.set_other_config", session_id, self, value)
}

// SR_get_introduced_by
//
// Get the introduced_by field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - DR_task ref
// - value of the field
func (client *XenAPIClient) SR_get_introduced_by(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_introduced_by", session_id, self)
}

// SR_get_local_cache_enabled
//
// Get the local_cache_enabled field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) SR_get_local_cache_enabled(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_local_cache_enabled", session_id, self)
}

// SR_get_blobs
//
// Get the blobs field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - (string -> blob ref) map
// - value of the field
func (client *XenAPIClient) SR_get_blobs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_blobs", session_id, self)
}

// SR_get_sm_config
//
// Get the sm_config field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) SR_get_sm_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_sm_config", session_id, self)
}

// SR_get_tags
//
// Get the tags field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) SR_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_tags", session_id, self)
}

// SR_get_other_config
//
// Get the other_config field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) SR_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_other_config", session_id, self)
}

// SR_get_shared
//
// Get the shared field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) SR_get_shared(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_shared", session_id, self)
}

// SR_get_content_type
//
// Get the content_type field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SR_get_content_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_content_type", session_id, self)
}

// SR_get_type
//
// Get the type field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SR_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_type", session_id, self)
}

// SR_get_physical_size
//
// Get the physical_size field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) SR_get_physical_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_physical_size", session_id, self)
}

// SR_get_physical_utilisation
//
// Get the physical_utilisation field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) SR_get_physical_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_physical_utilisation", session_id, self)
}

// SR_get_virtual_allocation
//
// Get the virtual_allocation field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) SR_get_virtual_allocation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_virtual_allocation", session_id, self)
}

// SR_get_PBDs
//
// Get the PBDs field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - PBD ref set
// - value of the field
func (client *XenAPIClient) SR_get_PBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_PBDs", session_id, self)
}

// SR_get_VDIs
//
// Get the VDIs field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - VDI ref set
// - value of the field
func (client *XenAPIClient) SR_get_VDIs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_VDIs", session_id, self)
}

// SR_get_current_operations
//
// Get the current_operations field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - (string -> enum storage_operations) map
// - value of the field
func (client *XenAPIClient) SR_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_current_operations", session_id, self)
}

// SR_get_allowed_operations
//
// Get the allowed_operations field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - enum storage_operations set
// - value of the field
func (client *XenAPIClient) SR_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_allowed_operations", session_id, self)
}

// SR_get_name_description
//
// Get the name/description field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SR_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_name_description", session_id, self)
}

// SR_get_name_label
//
// Get the name/label field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SR_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_name_label", session_id, self)
}

// SR_get_uuid
//
// Get the uuid field of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) SR_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_uuid", session_id, self)
}

// SR_get_by_name_label
//
// Get all the SR instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - SR ref set
// - references to objects with matching names
func (client *XenAPIClient) SR_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("SR.get_by_name_label", session_id, label)
}

// SR_get_by_uuid
//
// Get a reference to the SR instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - SR ref
// - reference to the object
func (client *XenAPIClient) SR_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("SR.get_by_uuid", session_id, uuid)
}

// SR_get_record
//
// Get a record containing the current state of the given SR.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, SR ref, reference to the object
//
// returns:
// - SR record
// - all fields from the object
func (client *XenAPIClient) SR_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("SR.get_record", session_id, self)
}

// VDI_get_all_records
//
// Return a map of VDI references to VDI records for all VDIs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VDI ref -> VDI record) map
// - records of all objects
func (client *XenAPIClient) VDI_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_all_records", session_id)
}

// VDI_get_all
//
// Return a list of all the VDIs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VDI ref set
// - references to all objects
func (client *XenAPIClient) VDI_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_all", session_id)
}

// VDI_pool_migrate
//
// Migrate a VDI, which may be attached to a running guest, to a different SR. The destination SR must be visible to the guest.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to migrate
// - sr, SR ref, The destination SR
// - options, (string -> string) map, Other parameters
//
// returns:
// - VDI ref
// - The new reference of the migrated VDI.
func (client *XenAPIClient) VDI_pool_migrate(session_id interface{}, vdi interface{}, sr interface{}, options map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.pool_migrate", session_id, vdi, sr, options)
}

// VDI_read_database_pool_uuid
//
// Check the VDI cache for the pool UUID of the database on this VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The metadata VDI to look up in the cache.
//
// returns:
// - string
// - The cached pool UUID of the database on the VDI.
func (client *XenAPIClient) VDI_read_database_pool_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.read_database_pool_uuid", session_id, self)
}

// VDI_open_database
//
// Load the metadata found on the supplied VDI and return a session reference which can be used in XenAPI calls to query its contents.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI which contains the database to open
//
// returns:
// - session ref
// - A session which can be used to query the database
func (client *XenAPIClient) VDI_open_database(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.open_database", session_id, self)
}

// VDI_set_allow_caching
//
// Set the value of the allow_caching parameter. This value can only be changed when the VDI is not attached to a running VM. The caching behaviour is only affected by this flag for VHD-based VDIs that have one parent and no child VHDs. Moreover, caching only takes place when the host running the VM containing this VDI has a nominated SR for local caching.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, bool, The value to set
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_allow_caching(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_allow_caching", session_id, self, value)
}

// VDI_set_on_boot
//
// Set the value of the on_boot parameter. This value can only be changed when the VDI is not attached to a running VM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, enum on_boot, The value to set
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_on_boot(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_on_boot", session_id, self, value)
}

// VDI_set_name_description
//
// Set the name description of the VDI. This can only happen when its SR is currently attached.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, string, The name description for the VDI
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_name_description", session_id, self, value)
}

// VDI_set_name_label
//
// Set the name label of the VDI. This can only happen when then its SR is currently attached.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, string, The name lable for the VDI
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_name_label", session_id, self, value)
}

// VDI_set_metadata_of_pool
//
// Records the pool whose metadata is contained by this VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, pool ref, The pool whose metadata is contained by this VDI
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_metadata_of_pool(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_metadata_of_pool", session_id, self, value)
}

// VDI_set_snapshot_time
//
// Sets the snapshot time of this VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, datetime, The snapshot time of this VDI.
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_snapshot_time(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_snapshot_time", session_id, self, value)
}

// VDI_set_snapshot_of
//
// Sets the VDI of which this VDI is a snapshot
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, VDI ref, The VDI of which this VDI is a snapshot
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_snapshot_of(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_snapshot_of", session_id, self, value)
}

// VDI_set_is_a_snapshot
//
// Sets whether this VDI is a snapshot
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, bool, The new value indicating whether this VDI is a snapshot
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_is_a_snapshot(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_is_a_snapshot", session_id, self, value)
}

// VDI_set_physical_utilisation
//
// Sets the VDI's physical_utilisation field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, int, The new value of the VDI's physical utilisation
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_physical_utilisation(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_physical_utilisation", session_id, self, value)
}

// VDI_set_virtual_size
//
// Sets the VDI's virtual_size field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, int, The new value of the VDI's virtual size
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_virtual_size(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_virtual_size", session_id, self, value)
}

// VDI_set_missing
//
// Sets the VDI's missing field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, bool, The new value of the VDI's missing field
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_missing(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_missing", session_id, self, value)
}

// VDI_set_read_only
//
// Sets the VDI's read_only field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, bool, The new value of the VDI's read_only field
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_read_only(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_read_only", session_id, self, value)
}

// VDI_set_sharable
//
// Sets the VDI's sharable field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, bool, The new value of the VDI's sharable field
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_sharable(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_sharable", session_id, self, value)
}

// VDI_forget
//
// Removes a VDI record from the database
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to forget about
//
// returns:
// - void
func (client *XenAPIClient) VDI_forget(session_id interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.forget", session_id, vdi)
}

// VDI_set_managed
//
// Sets the VDI's managed field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, The VDI to modify
// - value, bool, The new value of the VDI's managed field
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_managed(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VDI.set_managed", session_id, self, value)
}

// VDI_copy
//
// Copy either a full VDI or the block differences between two VDIs into either a fresh VDI or an existing VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to copy
// - sr, SR ref, The destination SR (only required if the destination VDI is not specified
// - base_vdi, VDI ref, The base VDI (only required if copying only changed blocks, by default all blocks will be copied)
// - into_vdi, VDI ref, The destination VDI to copy blocks into (if omitted then a destination SR must be provided and a fresh VDI will be created)
//
// returns:
// - VDI ref
// - The reference of the VDI where the blocks were written.
func (client *XenAPIClient) VDI_copy(session_id interface{}, vdi interface{}, sr interface{}, base_vdi interface{}, into_vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.copy", session_id, vdi, sr, base_vdi, into_vdi)
}

// VDI_update
//
// Ask the storage backend to refresh the fields in the VDI object
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI whose stats (eg size) should be updated
//
// returns:
// - void
func (client *XenAPIClient) VDI_update(session_id interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.update", session_id, vdi)
}

// VDI_db_forget
//
// Removes a VDI record from the database
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to forget about
//
// returns:
// - void
func (client *XenAPIClient) VDI_db_forget(session_id interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.db_forget", session_id, vdi)
}

// VDI_db_introduce
//
// Create a new VDI record in the database only
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, The uuid of the disk to introduce
// - name_label, string, The name of the disk record
// - name_description, string, The description of the disk record
// - SR, SR ref, The SR that the VDI is in
// - a_type, enum vdi_type, The type of the VDI
// - sharable, bool, true if this disk may be shared
// - read_only, bool, true if this disk may ONLY be mounted read-only
// - other_config, (string -> string) map, additional configuration
// - location, string, location information
// - xenstore_data, (string -> string) map, Data to insert into xenstore
// - sm_config, (string -> string) map, Storage-specific config
// - managed, bool, Storage-specific config
// - virtual_size, int, Storage-specific config
// - physical_utilisation, int, Storage-specific config
// - metadata_of_pool, pool ref, Storage-specific config
// - is_a_snapshot, bool, Storage-specific config
// - snapshot_time, datetime, Storage-specific config
// - snapshot_of, VDI ref, Storage-specific config
//
// returns:
// - VDI ref
// - The ref of the newly created VDI record.
func (client *XenAPIClient) VDI_db_introduce(session_id interface{}, uuid string, name_label string, name_description string, SR interface{}, a_type interface{}, sharable bool, read_only bool, other_config map[string]string, location string, xenstore_data map[string]string, sm_config map[string]string, managed bool, virtual_size interface{}, physical_utilisation interface{}, metadata_of_pool interface{}, is_a_snapshot bool, snapshot_time interface{}, snapshot_of interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.db_introduce", session_id, uuid, name_label, name_description, SR, a_type, sharable, read_only, other_config, location, xenstore_data, sm_config, managed, virtual_size, physical_utilisation, metadata_of_pool, is_a_snapshot, snapshot_time, snapshot_of)
}

// VDI_introduce
//
// Create a new VDI record in the database only
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, The uuid of the disk to introduce
// - name_label, string, The name of the disk record
// - name_description, string, The description of the disk record
// - SR, SR ref, The SR that the VDI is in
// - a_type, enum vdi_type, The type of the VDI
// - sharable, bool, true if this disk may be shared
// - read_only, bool, true if this disk may ONLY be mounted read-only
// - other_config, (string -> string) map, additional configuration
// - location, string, location information
// - xenstore_data, (string -> string) map, Data to insert into xenstore
// - sm_config, (string -> string) map, Storage-specific config
// - managed, bool, Storage-specific config
// - virtual_size, int, Storage-specific config
// - physical_utilisation, int, Storage-specific config
// - metadata_of_pool, pool ref, Storage-specific config
// - is_a_snapshot, bool, Storage-specific config
// - snapshot_time, datetime, Storage-specific config
// - snapshot_of, VDI ref, Storage-specific config
//
// returns:
// - VDI ref
// - The ref of the newly created VDI record.
func (client *XenAPIClient) VDI_introduce(session_id interface{}, uuid string, name_label string, name_description string, SR interface{}, a_type interface{}, sharable bool, read_only bool, other_config map[string]string, location string, xenstore_data map[string]string, sm_config map[string]string, managed bool, virtual_size interface{}, physical_utilisation interface{}, metadata_of_pool interface{}, is_a_snapshot bool, snapshot_time interface{}, snapshot_of interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.introduce", session_id, uuid, name_label, name_description, SR, a_type, sharable, read_only, other_config, location, xenstore_data, sm_config, managed, virtual_size, physical_utilisation, metadata_of_pool, is_a_snapshot, snapshot_time, snapshot_of)
}

// VDI_resize_online
//
// Resize the VDI which may or may not be attached to running guests.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to resize
// - size, int, The new size of the VDI
//
// returns:
// - void
func (client *XenAPIClient) VDI_resize_online(session_id interface{}, vdi interface{}, size interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.resize_online", session_id, vdi, size)
}

// VDI_resize
//
// Resize the VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to resize
// - size, int, The new size of the VDI
//
// returns:
// - void
func (client *XenAPIClient) VDI_resize(session_id interface{}, vdi interface{}, size interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.resize", session_id, vdi, size)
}

// VDI_clone
//
// Take an exact copy of the VDI and return a reference to the new disk. If any driver_params are specified then these are passed through to the storage-specific substrate driver that implements the clone operation. NB the clone lives in the same Storage Repository as its parent.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to clone
// - driver_params, (string -> string) map, Optional parameters that are passed through to the backend driver in order to specify storage-type-specific clone options
//
// returns:
// - VDI ref
// - The ID of the newly created VDI.
func (client *XenAPIClient) VDI_clone(session_id interface{}, vdi interface{}, driver_params map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.clone", session_id, vdi, driver_params)
}

// VDI_snapshot
//
// Take a read-only snapshot of the VDI, returning a reference to the snapshot. If any driver_params are specified then these are passed through to the storage-specific substrate driver that takes the snapshot. NB the snapshot lives in the same Storage Repository as its parent.
//
// params:
// - session_id, session ref, Reference to a valid session
// - vdi, VDI ref, The VDI to snapshot
// - driver_params, (string -> string) map, Optional parameters that can be passed through to backend driver in order to specify storage-type-specific snapshot options
//
// returns:
// - VDI ref
// - The ID of the newly created VDI.
func (client *XenAPIClient) VDI_snapshot(session_id interface{}, vdi interface{}, driver_params map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.snapshot", session_id, vdi, driver_params)
}

// VDI_remove_tags
//
// Remove the given value from the tags field of the given VDI.  If the value is not in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - value, string, Value to remove
//
// returns:
// - void
func (client *XenAPIClient) VDI_remove_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_tags", session_id, self, value)
}

// VDI_add_tags
//
// Add the given value to the tags field of the given VDI.  If the value is already in that Set, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - value, string, New value to add
//
// returns:
// - void
func (client *XenAPIClient) VDI_add_tags(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_tags", session_id, self, value)
}

// VDI_set_tags
//
// Set the tags field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - value, string set, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_tags(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.set_tags", session_id, self, value)
}

// VDI_remove_from_sm_config
//
// Remove the given key and its corresponding value from the sm_config field of the given VDI.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VDI_remove_from_sm_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_from_sm_config", session_id, self, key)
}

// VDI_add_to_sm_config
//
// Add the given key-value pair to the sm_config field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VDI_add_to_sm_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_to_sm_config", session_id, self, key, value)
}

// VDI_set_sm_config
//
// Set the sm_config field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_sm_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_sm_config", session_id, self, value)
}

// VDI_remove_from_xenstore_data
//
// Remove the given key and its corresponding value from the xenstore_data field of the given VDI.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VDI_remove_from_xenstore_data(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_from_xenstore_data", session_id, self, key)
}

// VDI_add_to_xenstore_data
//
// Add the given key-value pair to the xenstore_data field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VDI_add_to_xenstore_data(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_to_xenstore_data", session_id, self, key, value)
}

// VDI_set_xenstore_data
//
// Set the xenstore_data field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_xenstore_data(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_xenstore_data", session_id, self, value)
}

// VDI_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VDI.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VDI_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VDI.remove_from_other_config", session_id, self, key)
}

// VDI_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VDI_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VDI.add_to_other_config", session_id, self, key, value)
}

// VDI_set_other_config
//
// Set the other_config field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VDI_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VDI.set_other_config", session_id, self, value)
}

// VDI_get_metadata_latest
//
// Get the metadata_latest field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_metadata_latest(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_metadata_latest", session_id, self)
}

// VDI_get_metadata_of_pool
//
// Get the metadata_of_pool field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - pool ref
// - value of the field
func (client *XenAPIClient) VDI_get_metadata_of_pool(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_metadata_of_pool", session_id, self)
}

// VDI_get_on_boot
//
// Get the on_boot field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - enum on_boot
// - value of the field
func (client *XenAPIClient) VDI_get_on_boot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_on_boot", session_id, self)
}

// VDI_get_allow_caching
//
// Get the allow_caching field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_allow_caching(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_allow_caching", session_id, self)
}

// VDI_get_tags
//
// Get the tags field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VDI_get_tags(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_tags", session_id, self)
}

// VDI_get_snapshot_time
//
// Get the snapshot_time field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VDI_get_snapshot_time(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_snapshot_time", session_id, self)
}

// VDI_get_snapshots
//
// Get the snapshots field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - VDI ref set
// - value of the field
func (client *XenAPIClient) VDI_get_snapshots(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_snapshots", session_id, self)
}

// VDI_get_snapshot_of
//
// Get the snapshot_of field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - VDI ref
// - value of the field
func (client *XenAPIClient) VDI_get_snapshot_of(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_snapshot_of", session_id, self)
}

// VDI_get_is_a_snapshot
//
// Get the is_a_snapshot field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_is_a_snapshot(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_is_a_snapshot", session_id, self)
}

// VDI_get_sm_config
//
// Get the sm_config field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VDI_get_sm_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_sm_config", session_id, self)
}

// VDI_get_xenstore_data
//
// Get the xenstore_data field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VDI_get_xenstore_data(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_xenstore_data", session_id, self)
}

// VDI_get_parent
//
// Get the parent field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - VDI ref
// - value of the field
func (client *XenAPIClient) VDI_get_parent(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_parent", session_id, self)
}

// VDI_get_missing
//
// Get the missing field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_missing(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_missing", session_id, self)
}

// VDI_get_managed
//
// Get the managed field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_managed(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_managed", session_id, self)
}

// VDI_get_location
//
// Get the location field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VDI_get_location(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_location", session_id, self)
}

// VDI_get_storage_lock
//
// Get the storage_lock field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_storage_lock(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_storage_lock", session_id, self)
}

// VDI_get_other_config
//
// Get the other_config field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VDI_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_other_config", session_id, self)
}

// VDI_get_read_only
//
// Get the read_only field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_read_only(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_read_only", session_id, self)
}

// VDI_get_sharable
//
// Get the sharable field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VDI_get_sharable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_sharable", session_id, self)
}

// VDI_get_type
//
// Get the type field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - enum vdi_type
// - value of the field
func (client *XenAPIClient) VDI_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_type", session_id, self)
}

// VDI_get_physical_utilisation
//
// Get the physical_utilisation field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VDI_get_physical_utilisation(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_physical_utilisation", session_id, self)
}

// VDI_get_virtual_size
//
// Get the virtual_size field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VDI_get_virtual_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_virtual_size", session_id, self)
}

// VDI_get_crash_dumps
//
// Get the crash_dumps field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - crashdump ref set
// - value of the field
func (client *XenAPIClient) VDI_get_crash_dumps(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_crash_dumps", session_id, self)
}

// VDI_get_VBDs
//
// Get the VBDs field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - VBD ref set
// - value of the field
func (client *XenAPIClient) VDI_get_VBDs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_VBDs", session_id, self)
}

// VDI_get_SR
//
// Get the SR field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) VDI_get_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_SR", session_id, self)
}

// VDI_get_current_operations
//
// Get the current_operations field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - (string -> enum vdi_operations) map
// - value of the field
func (client *XenAPIClient) VDI_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_current_operations", session_id, self)
}

// VDI_get_allowed_operations
//
// Get the allowed_operations field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - enum vdi_operations set
// - value of the field
func (client *XenAPIClient) VDI_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_allowed_operations", session_id, self)
}

// VDI_get_name_description
//
// Get the name/description field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VDI_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_name_description", session_id, self)
}

// VDI_get_name_label
//
// Get the name/label field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VDI_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_name_label", session_id, self)
}

// VDI_get_uuid
//
// Get the uuid field of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VDI_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_uuid", session_id, self)
}

// VDI_get_by_name_label
//
// Get all the VDI instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - VDI ref set
// - references to objects with matching names
func (client *XenAPIClient) VDI_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("VDI.get_by_name_label", session_id, label)
}

// VDI_destroy
//
// Destroy the specified VDI instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VDI_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.destroy", session_id, self)
}

// VDI_create
//
// Create a new VDI instance, and return its handle.
// The constructor args are: name_label, name_description, SR*, virtual_size*, type*, sharable*, read_only*, other_config*, xenstore_data, sm_config, tags (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, VDI record, All constructor arguments
//
// returns:
// - VDI ref
// - reference to the newly created object
func (client *XenAPIClient) VDI_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.create", session_id, args)
}

// VDI_get_by_uuid
//
// Get a reference to the VDI instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VDI ref
// - reference to the object
func (client *XenAPIClient) VDI_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VDI.get_by_uuid", session_id, uuid)
}

// VDI_get_record
//
// Get a record containing the current state of the given VDI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VDI ref, reference to the object
//
// returns:
// - VDI record
// - all fields from the object
func (client *XenAPIClient) VDI_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VDI.get_record", session_id, self)
}

// VBD_get_all_records
//
// Return a map of VBD references to VBD records for all VBDs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VBD ref -> VBD record) map
// - records of all objects
func (client *XenAPIClient) VBD_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_all_records", session_id)
}

// VBD_get_all
//
// Return a list of all the VBDs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VBD ref set
// - references to all objects
func (client *XenAPIClient) VBD_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_all", session_id)
}

// VBD_assert_attachable
//
// Throws an error if this VBD could not be attached to this VM if the VM were running. Intended for debugging.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, The VBD to query
//
// returns:
// - void
func (client *XenAPIClient) VBD_assert_attachable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.assert_attachable", session_id, self)
}

// VBD_unplug_force
//
// Forcibly unplug the specified VBD
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, The VBD to forcibly unplug
//
// returns:
// - void
func (client *XenAPIClient) VBD_unplug_force(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.unplug_force", session_id, self)
}

// VBD_unplug
//
// Hot-unplug the specified VBD, dynamically unattaching it from the running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, The VBD to hot-unplug
//
// returns:
// - void
func (client *XenAPIClient) VBD_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.unplug", session_id, self)
}

// VBD_plug
//
// Hotplug the specified VBD, dynamically attaching it to the running VM
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, The VBD to hotplug
//
// returns:
// - void
func (client *XenAPIClient) VBD_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.plug", session_id, self)
}

// VBD_insert
//
// Insert new media into the device
//
// params:
// - session_id, session ref, Reference to a valid session
// - vbd, VBD ref, The vbd representing the CDROM-like device
// - vdi, VDI ref, The new VDI to 'insert'
//
// returns:
// - void
func (client *XenAPIClient) VBD_insert(session_id interface{}, vbd interface{}, vdi interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.insert", session_id, vbd, vdi)
}

// VBD_eject
//
// Remove the media from the device and leave it empty
//
// params:
// - session_id, session ref, Reference to a valid session
// - vbd, VBD ref, The vbd representing the CDROM-like device
//
// returns:
// - void
func (client *XenAPIClient) VBD_eject(session_id interface{}, vbd interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.eject", session_id, vbd)
}

// VBD_remove_from_qos_algorithm_params
//
// Remove the given key and its corresponding value from the qos/algorithm_params field of the given VBD.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VBD_remove_from_qos_algorithm_params(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VBD.remove_from_qos_algorithm_params", session_id, self, key)
}

// VBD_add_to_qos_algorithm_params
//
// Add the given key-value pair to the qos/algorithm_params field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VBD_add_to_qos_algorithm_params(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.add_to_qos_algorithm_params", session_id, self, key, value)
}

// VBD_set_qos_algorithm_params
//
// Set the qos/algorithm_params field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_qos_algorithm_params(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_qos_algorithm_params", session_id, self, value)
}

// VBD_set_qos_algorithm_type
//
// Set the qos/algorithm_type field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_qos_algorithm_type(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_qos_algorithm_type", session_id, self, value)
}

// VBD_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VBD.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VBD_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VBD.remove_from_other_config", session_id, self, key)
}

// VBD_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VBD_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.add_to_other_config", session_id, self, key, value)
}

// VBD_set_other_config
//
// Set the other_config field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_other_config", session_id, self, value)
}

// VBD_set_unpluggable
//
// Set the unpluggable field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_unpluggable(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VBD.set_unpluggable", session_id, self, value)
}

// VBD_set_type
//
// Set the type field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, enum vbd_type, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_type(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.set_type", session_id, self, value)
}

// VBD_set_mode
//
// Set the mode field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, enum vbd_mode, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_mode(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.set_mode", session_id, self, value)
}

// VBD_set_bootable
//
// Set the bootable field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_bootable(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("VBD.set_bootable", session_id, self, value)
}

// VBD_set_userdevice
//
// Set the userdevice field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_set_userdevice(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("VBD.set_userdevice", session_id, self, value)
}

// VBD_get_metrics
//
// Get the metrics field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - VBD_metrics ref
// - value of the field
func (client *XenAPIClient) VBD_get_metrics(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_metrics", session_id, self)
}

// VBD_get_qos_supported_algorithms
//
// Get the qos/supported_algorithms field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) VBD_get_qos_supported_algorithms(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_qos_supported_algorithms", session_id, self)
}

// VBD_get_qos_algorithm_params
//
// Get the qos/algorithm_params field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VBD_get_qos_algorithm_params(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_qos_algorithm_params", session_id, self)
}

// VBD_get_qos_algorithm_type
//
// Get the qos/algorithm_type field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VBD_get_qos_algorithm_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_qos_algorithm_type", session_id, self)
}

// VBD_get_runtime_properties
//
// Get the runtime_properties field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VBD_get_runtime_properties(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_runtime_properties", session_id, self)
}

// VBD_get_status_detail
//
// Get the status_detail field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VBD_get_status_detail(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_status_detail", session_id, self)
}

// VBD_get_status_code
//
// Get the status_code field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VBD_get_status_code(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_status_code", session_id, self)
}

// VBD_get_currently_attached
//
// Get the currently_attached field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VBD_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_currently_attached", session_id, self)
}

// VBD_get_other_config
//
// Get the other_config field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VBD_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_other_config", session_id, self)
}

// VBD_get_empty
//
// Get the empty field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VBD_get_empty(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_empty", session_id, self)
}

// VBD_get_storage_lock
//
// Get the storage_lock field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VBD_get_storage_lock(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_storage_lock", session_id, self)
}

// VBD_get_unpluggable
//
// Get the unpluggable field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VBD_get_unpluggable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_unpluggable", session_id, self)
}

// VBD_get_type
//
// Get the type field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - enum vbd_type
// - value of the field
func (client *XenAPIClient) VBD_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_type", session_id, self)
}

// VBD_get_mode
//
// Get the mode field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - enum vbd_mode
// - value of the field
func (client *XenAPIClient) VBD_get_mode(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_mode", session_id, self)
}

// VBD_get_bootable
//
// Get the bootable field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VBD_get_bootable(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_bootable", session_id, self)
}

// VBD_get_userdevice
//
// Get the userdevice field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VBD_get_userdevice(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_userdevice", session_id, self)
}

// VBD_get_device
//
// Get the device field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VBD_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_device", session_id, self)
}

// VBD_get_VDI
//
// Get the VDI field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - VDI ref
// - value of the field
func (client *XenAPIClient) VBD_get_VDI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_VDI", session_id, self)
}

// VBD_get_VM
//
// Get the VM field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) VBD_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_VM", session_id, self)
}

// VBD_get_current_operations
//
// Get the current_operations field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - (string -> enum vbd_operations) map
// - value of the field
func (client *XenAPIClient) VBD_get_current_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_current_operations", session_id, self)
}

// VBD_get_allowed_operations
//
// Get the allowed_operations field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - enum vbd_operations set
// - value of the field
func (client *XenAPIClient) VBD_get_allowed_operations(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_allowed_operations", session_id, self)
}

// VBD_get_uuid
//
// Get the uuid field of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VBD_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_uuid", session_id, self)
}

// VBD_destroy
//
// Destroy the specified VBD instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VBD_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.destroy", session_id, self)
}

// VBD_create
//
// Create a new VBD instance, and return its handle.
// The constructor args are: VM*, VDI*, userdevice*, bootable*, mode*, type*, unpluggable, empty*, other_config*, qos_algorithm_type*, qos_algorithm_params* (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, VBD record, All constructor arguments
//
// returns:
// - VBD ref
// - reference to the newly created object
func (client *XenAPIClient) VBD_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.create", session_id, args)
}

// VBD_get_by_uuid
//
// Get a reference to the VBD instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VBD ref
// - reference to the object
func (client *XenAPIClient) VBD_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VBD.get_by_uuid", session_id, uuid)
}

// VBD_get_record
//
// Get a record containing the current state of the given VBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD ref, reference to the object
//
// returns:
// - VBD record
// - all fields from the object
func (client *XenAPIClient) VBD_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD.get_record", session_id, self)
}

// VBD_metrics_get_all_records
//
// Return a map of VBD_metrics references to VBD_metrics records for all VBD_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VBD_metrics ref -> VBD_metrics record) map
// - records of all objects
func (client *XenAPIClient) VBD_metrics_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_all_records", session_id)
}

// VBD_metrics_get_all
//
// Return a list of all the VBD_metrics instances known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VBD_metrics ref set
// - references to all objects
func (client *XenAPIClient) VBD_metrics_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_all", session_id)
}

// VBD_metrics_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VBD_metrics.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VBD_metrics_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.remove_from_other_config", session_id, self, key)
}

// VBD_metrics_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VBD_metrics_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.add_to_other_config", session_id, self, key, value)
}

// VBD_metrics_set_other_config
//
// Set the other_config field of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VBD_metrics_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.set_other_config", session_id, self, value)
}

// VBD_metrics_get_other_config
//
// Get the other_config field of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VBD_metrics_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_other_config", session_id, self)
}

// VBD_metrics_get_last_updated
//
// Get the last_updated field of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) VBD_metrics_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_last_updated", session_id, self)
}

// VBD_metrics_get_io_write_kbs
//
// Get the io/write_kbs field of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) VBD_metrics_get_io_write_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_io_write_kbs", session_id, self)
}

// VBD_metrics_get_io_read_kbs
//
// Get the io/read_kbs field of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
//
// returns:
// - float
// - value of the field
func (client *XenAPIClient) VBD_metrics_get_io_read_kbs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_io_read_kbs", session_id, self)
}

// VBD_metrics_get_uuid
//
// Get the uuid field of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VBD_metrics_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_uuid", session_id, self)
}

// VBD_metrics_get_by_uuid
//
// Get a reference to the VBD_metrics instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VBD_metrics ref
// - reference to the object
func (client *XenAPIClient) VBD_metrics_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_by_uuid", session_id, uuid)
}

// VBD_metrics_get_record
//
// Get a record containing the current state of the given VBD_metrics.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VBD_metrics ref, reference to the object
//
// returns:
// - VBD_metrics record
// - all fields from the object
func (client *XenAPIClient) VBD_metrics_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VBD_metrics.get_record", session_id, self)
}

// PBD_get_all_records
//
// Return a map of PBD references to PBD records for all PBDs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (PBD ref -> PBD record) map
// - records of all objects
func (client *XenAPIClient) PBD_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_all_records", session_id)
}

// PBD_get_all
//
// Return a list of all the PBDs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - PBD ref set
// - references to all objects
func (client *XenAPIClient) PBD_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_all", session_id)
}

// PBD_set_device_config
//
// Sets the PBD's device_config field
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, The PBD to modify
// - value, (string -> string) map, The new value of the PBD's device_config
//
// returns:
// - void
func (client *XenAPIClient) PBD_set_device_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PBD.set_device_config", session_id, self, value)
}

// PBD_unplug
//
// Deactivate the specified PBD, causing the referenced SR to be detached and nolonger scanned
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, The PBD to deactivate
//
// returns:
// - void
func (client *XenAPIClient) PBD_unplug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.unplug", session_id, self)
}

// PBD_plug
//
// Activate the specified PBD, causing the referenced SR to be attached and scanned
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, The PBD to activate
//
// returns:
// - void
func (client *XenAPIClient) PBD_plug(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.plug", session_id, self)
}

// PBD_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given PBD.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) PBD_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PBD.remove_from_other_config", session_id, self, key)
}

// PBD_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) PBD_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PBD.add_to_other_config", session_id, self, key, value)
}

// PBD_set_other_config
//
// Set the other_config field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) PBD_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PBD.set_other_config", session_id, self, value)
}

// PBD_get_other_config
//
// Get the other_config field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) PBD_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_other_config", session_id, self)
}

// PBD_get_currently_attached
//
// Get the currently_attached field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PBD_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_currently_attached", session_id, self)
}

// PBD_get_device_config
//
// Get the device_config field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) PBD_get_device_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_device_config", session_id, self)
}

// PBD_get_SR
//
// Get the SR field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - SR ref
// - value of the field
func (client *XenAPIClient) PBD_get_SR(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_SR", session_id, self)
}

// PBD_get_host
//
// Get the host field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) PBD_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_host", session_id, self)
}

// PBD_get_uuid
//
// Get the uuid field of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PBD_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_uuid", session_id, self)
}

// PBD_destroy
//
// Destroy the specified PBD instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) PBD_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.destroy", session_id, self)
}

// PBD_create
//
// Create a new PBD instance, and return its handle.
// The constructor args are: host*, SR*, device_config*, other_config (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, PBD record, All constructor arguments
//
// returns:
// - PBD ref
// - reference to the newly created object
func (client *XenAPIClient) PBD_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.create", session_id, args)
}

// PBD_get_by_uuid
//
// Get a reference to the PBD instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - PBD ref
// - reference to the object
func (client *XenAPIClient) PBD_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PBD.get_by_uuid", session_id, uuid)
}

// PBD_get_record
//
// Get a record containing the current state of the given PBD.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PBD ref, reference to the object
//
// returns:
// - PBD record
// - all fields from the object
func (client *XenAPIClient) PBD_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PBD.get_record", session_id, self)
}

// crashdump_get_all_records
//
// Return a map of crashdump references to crashdump records for all crashdumps known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (crashdump ref -> crashdump record) map
// - records of all objects
func (client *XenAPIClient) crashdump_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_all_records", session_id)
}

// crashdump_get_all
//
// Return a list of all the crashdumps known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - crashdump ref set
// - references to all objects
func (client *XenAPIClient) crashdump_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_all", session_id)
}

// crashdump_destroy
//
// Destroy the specified crashdump
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, The crashdump to destroy
//
// returns:
// - void
func (client *XenAPIClient) crashdump_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.destroy", session_id, self)
}

// crashdump_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given crashdump.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) crashdump_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("crashdump.remove_from_other_config", session_id, self, key)
}

// crashdump_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) crashdump_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("crashdump.add_to_other_config", session_id, self, key, value)
}

// crashdump_set_other_config
//
// Set the other_config field of the given crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) crashdump_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("crashdump.set_other_config", session_id, self, value)
}

// crashdump_get_other_config
//
// Get the other_config field of the given crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) crashdump_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_other_config", session_id, self)
}

// crashdump_get_VDI
//
// Get the VDI field of the given crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
//
// returns:
// - VDI ref
// - value of the field
func (client *XenAPIClient) crashdump_get_VDI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_VDI", session_id, self)
}

// crashdump_get_VM
//
// Get the VM field of the given crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) crashdump_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_VM", session_id, self)
}

// crashdump_get_uuid
//
// Get the uuid field of the given crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) crashdump_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_uuid", session_id, self)
}

// crashdump_get_by_uuid
//
// Get a reference to the crashdump instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - crashdump ref
// - reference to the object
func (client *XenAPIClient) crashdump_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_by_uuid", session_id, uuid)
}

// crashdump_get_record
//
// Get a record containing the current state of the given crashdump.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, crashdump ref, reference to the object
//
// returns:
// - crashdump record
// - all fields from the object
func (client *XenAPIClient) crashdump_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("crashdump.get_record", session_id, self)
}

// VTPM_get_backend
//
// Get the backend field of the given VTPM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VTPM ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) VTPM_get_backend(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_backend", session_id, self)
}

// VTPM_get_VM
//
// Get the VM field of the given VTPM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VTPM ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) VTPM_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_VM", session_id, self)
}

// VTPM_get_uuid
//
// Get the uuid field of the given VTPM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VTPM ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VTPM_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_uuid", session_id, self)
}

// VTPM_destroy
//
// Destroy the specified VTPM instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VTPM ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) VTPM_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.destroy", session_id, self)
}

// VTPM_create
//
// Create a new VTPM instance, and return its handle.
// The constructor args are: VM*, backend* (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, VTPM record, All constructor arguments
//
// returns:
// - VTPM ref
// - reference to the newly created object
func (client *XenAPIClient) VTPM_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.create", session_id, args)
}

// VTPM_get_by_uuid
//
// Get a reference to the VTPM instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VTPM ref
// - reference to the object
func (client *XenAPIClient) VTPM_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_by_uuid", session_id, uuid)
}

// VTPM_get_record
//
// Get a record containing the current state of the given VTPM.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VTPM ref, reference to the object
//
// returns:
// - VTPM record
// - all fields from the object
func (client *XenAPIClient) VTPM_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VTPM.get_record", session_id, self)
}

// console_get_all_records
//
// Return a map of console references to console records for all consoles known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (console ref -> console record) map
// - records of all objects
func (client *XenAPIClient) console_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_all_records", session_id)
}

// console_get_all
//
// Return a list of all the consoles known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - console ref set
// - references to all objects
func (client *XenAPIClient) console_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_all", session_id)
}

// console_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given console.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) console_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("console.remove_from_other_config", session_id, self, key)
}

// console_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) console_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("console.add_to_other_config", session_id, self, key, value)
}

// console_set_other_config
//
// Set the other_config field of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) console_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("console.set_other_config", session_id, self, value)
}

// console_get_other_config
//
// Get the other_config field of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) console_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_other_config", session_id, self)
}

// console_get_VM
//
// Get the VM field of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) console_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_VM", session_id, self)
}

// console_get_location
//
// Get the location field of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) console_get_location(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_location", session_id, self)
}

// console_get_protocol
//
// Get the protocol field of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
//
// returns:
// - enum console_protocol
// - value of the field
func (client *XenAPIClient) console_get_protocol(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_protocol", session_id, self)
}

// console_get_uuid
//
// Get the uuid field of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) console_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_uuid", session_id, self)
}

// console_destroy
//
// Destroy the specified console instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) console_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.destroy", session_id, self)
}

// console_create
//
// Create a new console instance, and return its handle.
// The constructor args are: other_config* (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, console record, All constructor arguments
//
// returns:
// - console ref
// - reference to the newly created object
func (client *XenAPIClient) console_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("console.create", session_id, args)
}

// console_get_by_uuid
//
// Get a reference to the console instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - console ref
// - reference to the object
func (client *XenAPIClient) console_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("console.get_by_uuid", session_id, uuid)
}

// console_get_record
//
// Get a record containing the current state of the given console.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, console ref, reference to the object
//
// returns:
// - console record
// - all fields from the object
func (client *XenAPIClient) console_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("console.get_record", session_id, self)
}

// user_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given user.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) user_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("user.remove_from_other_config", session_id, self, key)
}

// user_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) user_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("user.add_to_other_config", session_id, self, key, value)
}

// user_set_other_config
//
// Set the other_config field of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) user_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("user.set_other_config", session_id, self, value)
}

// user_set_fullname
//
// Set the fullname field of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) user_set_fullname(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("user.set_fullname", session_id, self, value)
}

// user_get_other_config
//
// Get the other_config field of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) user_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_other_config", session_id, self)
}

// user_get_fullname
//
// Get the fullname field of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) user_get_fullname(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_fullname", session_id, self)
}

// user_get_short_name
//
// Get the short_name field of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) user_get_short_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_short_name", session_id, self)
}

// user_get_uuid
//
// Get the uuid field of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) user_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_uuid", session_id, self)
}

// user_destroy
//
// Destroy the specified user instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) user_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.destroy", session_id, self)
}

// user_create
//
// Create a new user instance, and return its handle.
// The constructor args are: short_name*, fullname*, other_config (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, user record, All constructor arguments
//
// returns:
// - user ref
// - reference to the newly created object
func (client *XenAPIClient) user_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("user.create", session_id, args)
}

// user_get_by_uuid
//
// Get a reference to the user instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - user ref
// - reference to the object
func (client *XenAPIClient) user_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("user.get_by_uuid", session_id, uuid)
}

// user_get_record
//
// Get a record containing the current state of the given user.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, user ref, reference to the object
//
// returns:
// - user record
// - all fields from the object
func (client *XenAPIClient) user_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("user.get_record", session_id, self)
}

// blob_get_all_records
//
// Return a map of blob references to blob records for all blobs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (blob ref -> blob record) map
// - records of all objects
func (client *XenAPIClient) blob_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_all_records", session_id)
}

// blob_get_all
//
// Return a list of all the blobs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - blob ref set
// - references to all objects
func (client *XenAPIClient) blob_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_all", session_id)
}

// blob_destroy
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, The reference of the blob to destroy
//
// returns:
// - void
func (client *XenAPIClient) blob_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.destroy", session_id, self)
}

// blob_create
//
// Create a placeholder for a binary blob
//
// params:
// - session_id, session ref, Reference to a valid session
// - mime_type, string, The mime-type of the blob. Defaults to 'application/octet-stream' if the empty string is supplied
// - public, bool, True if the blob should be publicly available
//
// returns:
// - blob ref
// - The reference to the created blob
func (client *XenAPIClient) blob_create(session_id interface{}, mime_type string, public bool) (i interface{}, err error) {
	return client.RPCCall("blob.create", session_id, mime_type, public)
}

// blob_set_public
//
// Set the public field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
// - value, bool, New value to set
//
// returns:
// - void
func (client *XenAPIClient) blob_set_public(session_id interface{}, self interface{}, value bool) (i interface{}, err error) {
	return client.RPCCall("blob.set_public", session_id, self, value)
}

// blob_set_name_description
//
// Set the name/description field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) blob_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("blob.set_name_description", session_id, self, value)
}

// blob_set_name_label
//
// Set the name/label field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) blob_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("blob.set_name_label", session_id, self, value)
}

// blob_get_mime_type
//
// Get the mime_type field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) blob_get_mime_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_mime_type", session_id, self)
}

// blob_get_last_updated
//
// Get the last_updated field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - datetime
// - value of the field
func (client *XenAPIClient) blob_get_last_updated(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_last_updated", session_id, self)
}

// blob_get_public
//
// Get the public field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) blob_get_public(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_public", session_id, self)
}

// blob_get_size
//
// Get the size field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) blob_get_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_size", session_id, self)
}

// blob_get_name_description
//
// Get the name/description field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) blob_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_name_description", session_id, self)
}

// blob_get_name_label
//
// Get the name/label field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) blob_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_name_label", session_id, self)
}

// blob_get_uuid
//
// Get the uuid field of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) blob_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_uuid", session_id, self)
}

// blob_get_by_name_label
//
// Get all the blob instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - blob ref set
// - references to objects with matching names
func (client *XenAPIClient) blob_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("blob.get_by_name_label", session_id, label)
}

// blob_get_by_uuid
//
// Get a reference to the blob instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - blob ref
// - reference to the object
func (client *XenAPIClient) blob_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("blob.get_by_uuid", session_id, uuid)
}

// blob_get_record
//
// Get a record containing the current state of the given blob.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, blob ref, reference to the object
//
// returns:
// - blob record
// - all fields from the object
func (client *XenAPIClient) blob_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("blob.get_record", session_id, self)
}

// message_get_all_records_where
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - expr, string, The expression to match (not currently used)
//
// returns:
// - (message ref -> message record) map
// - The messages
func (client *XenAPIClient) message_get_all_records_where(session_id interface{}, expr string) (i interface{}, err error) {
	return client.RPCCall("message.get_all_records_where", session_id, expr)
}

// message_get_all_records
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (message ref -> message record) map
// - The messages
func (client *XenAPIClient) message_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_all_records", session_id)
}

// message_get_by_uuid
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, The uuid of the message
//
// returns:
// - message ref
// - The message reference
func (client *XenAPIClient) message_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("message.get_by_uuid", session_id, uuid)
}

// message_get_record
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, message ref, The reference to the message
//
// returns:
// - message record
// - The message record
func (client *XenAPIClient) message_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_record", session_id, self)
}

// message_get_since
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - since, datetime, The cutoff time
//
// returns:
// - (message ref -> message record) map
// - The relevant messages
func (client *XenAPIClient) message_get_since(session_id interface{}, since interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_since", session_id, since)
}

// message_get_all
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - message ref set
// - The references to the messages
func (client *XenAPIClient) message_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get_all", session_id)
}

// message_get
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - cls, enum cls, The class of object
// - obj_uuid, string, The uuid of the object
// - since, datetime, The cutoff time
//
// returns:
// - (message ref -> message record) map
// - The relevant messages
func (client *XenAPIClient) message_get(session_id interface{}, cls interface{}, obj_uuid string, since interface{}) (i interface{}, err error) {
	return client.RPCCall("message.get", session_id, cls, obj_uuid, since)
}

// message_destroy
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, message ref, The reference of the message to destroy
//
// returns:
// - void
func (client *XenAPIClient) message_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("message.destroy", session_id, self)
}

// message_create
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - name, string, The name of the message
// - priority, int, The priority of the message
// - cls, enum cls, The class of object this message is associated with
// - obj_uuid, string, The uuid of the object this message is associated with
// - body, string, The body of the message
//
// returns:
// - message ref
// - The reference of the created message
func (client *XenAPIClient) message_create(session_id interface{}, name string, priority interface{}, cls interface{}, obj_uuid string, body string) (i interface{}, err error) {
	return client.RPCCall("message.create", session_id, name, priority, cls, obj_uuid, body)
}

// secret_get_all_records
//
// Return a map of secret references to secret records for all secrets known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (secret ref -> secret record) map
// - records of all objects
func (client *XenAPIClient) secret_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_all_records", session_id)
}

// secret_get_all
//
// Return a list of all the secrets known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - secret ref set
// - references to all objects
func (client *XenAPIClient) secret_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_all", session_id)
}

// secret_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given secret.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) secret_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("secret.remove_from_other_config", session_id, self, key)
}

// secret_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given secret.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) secret_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("secret.add_to_other_config", session_id, self, key, value)
}

// secret_set_other_config
//
// Set the other_config field of the given secret.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) secret_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("secret.set_other_config", session_id, self, value)
}

// secret_set_value
//
// Set the value field of the given secret.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) secret_set_value(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("secret.set_value", session_id, self, value)
}

// secret_get_other_config
//
// Get the other_config field of the given secret.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) secret_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_other_config", session_id, self)
}

// secret_get_value
//
// Get the value field of the given secret.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) secret_get_value(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_value", session_id, self)
}

// secret_get_uuid
//
// Get the uuid field of the given secret.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) secret_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_uuid", session_id, self)
}

// secret_destroy
//
// Destroy the specified secret instance.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
//
// returns:
// - void
func (client *XenAPIClient) secret_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.destroy", session_id, self)
}

// secret_create
//
// Create a new secret instance, and return its handle.
// The constructor args are: value*, other_config (* = non-optional).
//
// params:
// - session_id, session ref, Reference to a valid session
// - args, secret record, All constructor arguments
//
// returns:
// - secret ref
// - reference to the newly created object
func (client *XenAPIClient) secret_create(session_id interface{}, args interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.create", session_id, args)
}

// secret_get_by_uuid
//
// Get a reference to the secret instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - secret ref
// - reference to the object
func (client *XenAPIClient) secret_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("secret.get_by_uuid", session_id, uuid)
}

// secret_get_record
//
// Get a record containing the current state of the given secret.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, secret ref, reference to the object
//
// returns:
// - secret record
// - all fields from the object
func (client *XenAPIClient) secret_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("secret.get_record", session_id, self)
}

// tunnel_get_all_records
//
// Return a map of tunnel references to tunnel records for all tunnels known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (tunnel ref -> tunnel record) map
// - records of all objects
func (client *XenAPIClient) tunnel_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_all_records", session_id)
}

// tunnel_get_all
//
// Return a list of all the tunnels known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - tunnel ref set
// - references to all objects
func (client *XenAPIClient) tunnel_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_all", session_id)
}

// tunnel_destroy
//
// Destroy a tunnel
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, tunnel to destroy
//
// returns:
// - void
func (client *XenAPIClient) tunnel_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.destroy", session_id, self)
}

// tunnel_create
//
// Create a tunnel
//
// params:
// - session_id, session ref, Reference to a valid session
// - transport_PIF, PIF ref, PIF which receives the tagged traffic
// - network, network ref, Network to receive the tunnelled traffic
//
// returns:
// - tunnel ref
// - The reference of the created tunnel object
func (client *XenAPIClient) tunnel_create(session_id interface{}, transport_PIF interface{}, network interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.create", session_id, transport_PIF, network)
}

// tunnel_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given tunnel.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) tunnel_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("tunnel.remove_from_other_config", session_id, self, key)
}

// tunnel_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) tunnel_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("tunnel.add_to_other_config", session_id, self, key, value)
}

// tunnel_set_other_config
//
// Set the other_config field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) tunnel_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("tunnel.set_other_config", session_id, self, value)
}

// tunnel_remove_from_status
//
// Remove the given key and its corresponding value from the status field of the given tunnel.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) tunnel_remove_from_status(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("tunnel.remove_from_status", session_id, self, key)
}

// tunnel_add_to_status
//
// Add the given key-value pair to the status field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) tunnel_add_to_status(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("tunnel.add_to_status", session_id, self, key, value)
}

// tunnel_set_status
//
// Set the status field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) tunnel_set_status(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("tunnel.set_status", session_id, self, value)
}

// tunnel_get_other_config
//
// Get the other_config field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) tunnel_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_other_config", session_id, self)
}

// tunnel_get_status
//
// Get the status field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) tunnel_get_status(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_status", session_id, self)
}

// tunnel_get_transport_PIF
//
// Get the transport_PIF field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
//
// returns:
// - PIF ref
// - value of the field
func (client *XenAPIClient) tunnel_get_transport_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_transport_PIF", session_id, self)
}

// tunnel_get_access_PIF
//
// Get the access_PIF field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
//
// returns:
// - PIF ref
// - value of the field
func (client *XenAPIClient) tunnel_get_access_PIF(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_access_PIF", session_id, self)
}

// tunnel_get_uuid
//
// Get the uuid field of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) tunnel_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_uuid", session_id, self)
}

// tunnel_get_by_uuid
//
// Get a reference to the tunnel instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - tunnel ref
// - reference to the object
func (client *XenAPIClient) tunnel_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_by_uuid", session_id, uuid)
}

// tunnel_get_record
//
// Get a record containing the current state of the given tunnel.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, tunnel ref, reference to the object
//
// returns:
// - tunnel record
// - all fields from the object
func (client *XenAPIClient) tunnel_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("tunnel.get_record", session_id, self)
}

// PCI_get_all_records
//
// Return a map of PCI references to PCI records for all PCIs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (PCI ref -> PCI record) map
// - records of all objects
func (client *XenAPIClient) PCI_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_all_records", session_id)
}

// PCI_get_all
//
// Return a list of all the PCIs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - PCI ref set
// - references to all objects
func (client *XenAPIClient) PCI_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_all", session_id)
}

// PCI_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given PCI.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) PCI_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PCI.remove_from_other_config", session_id, self, key)
}

// PCI_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) PCI_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PCI.add_to_other_config", session_id, self, key, value)
}

// PCI_set_other_config
//
// Set the other_config field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) PCI_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PCI.set_other_config", session_id, self, value)
}

// PCI_get_subsystem_device_name
//
// Get the subsystem_device_name field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PCI_get_subsystem_device_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_subsystem_device_name", session_id, self)
}

// PCI_get_subsystem_vendor_name
//
// Get the subsystem_vendor_name field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PCI_get_subsystem_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_subsystem_vendor_name", session_id, self)
}

// PCI_get_other_config
//
// Get the other_config field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) PCI_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_other_config", session_id, self)
}

// PCI_get_dependencies
//
// Get the dependencies field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - PCI ref set
// - value of the field
func (client *XenAPIClient) PCI_get_dependencies(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_dependencies", session_id, self)
}

// PCI_get_pci_id
//
// Get the pci_id field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PCI_get_pci_id(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_pci_id", session_id, self)
}

// PCI_get_host
//
// Get the host field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) PCI_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_host", session_id, self)
}

// PCI_get_device_name
//
// Get the device_name field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PCI_get_device_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_device_name", session_id, self)
}

// PCI_get_vendor_name
//
// Get the vendor_name field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PCI_get_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_vendor_name", session_id, self)
}

// PCI_get_class_name
//
// Get the class_name field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PCI_get_class_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_class_name", session_id, self)
}

// PCI_get_uuid
//
// Get the uuid field of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PCI_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_uuid", session_id, self)
}

// PCI_get_by_uuid
//
// Get a reference to the PCI instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - PCI ref
// - reference to the object
func (client *XenAPIClient) PCI_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PCI.get_by_uuid", session_id, uuid)
}

// PCI_get_record
//
// Get a record containing the current state of the given PCI.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PCI ref, reference to the object
//
// returns:
// - PCI record
// - all fields from the object
func (client *XenAPIClient) PCI_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PCI.get_record", session_id, self)
}

// PGPU_get_all_records
//
// Return a map of PGPU references to PGPU records for all PGPUs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (PGPU ref -> PGPU record) map
// - records of all objects
func (client *XenAPIClient) PGPU_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_all_records", session_id)
}

// PGPU_get_all
//
// Return a list of all the PGPUs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - PGPU ref set
// - references to all objects
func (client *XenAPIClient) PGPU_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_all", session_id)
}

// PGPU_disable_dom0_access
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, The PGPU to which dom0 will be denied access
//
// returns:
// - enum pgpu_dom0_access
// - The accessibility of this PGPU from dom0
func (client *XenAPIClient) PGPU_disable_dom0_access(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.disable_dom0_access", session_id, self)
}

// PGPU_enable_dom0_access
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, The PGPU to which dom0 will be granted access
//
// returns:
// - enum pgpu_dom0_access
// - The accessibility of this PGPU from dom0
func (client *XenAPIClient) PGPU_enable_dom0_access(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.enable_dom0_access", session_id, self)
}

// PGPU_get_remaining_capacity
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, The PGPU to query
// - vgpu_type, VGPU_type ref, The VGPU type for which we want to find the number of VGPUs which can still be started on this PGPU
//
// returns:
// - int
// - The number of VGPUs of the specified type which can still be started on this PGPU
func (client *XenAPIClient) PGPU_get_remaining_capacity(session_id interface{}, self interface{}, vgpu_type interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_remaining_capacity", session_id, self, vgpu_type)
}

// PGPU_set_GPU_group
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, The PGPU to move to a new group
// - value, GPU_group ref, The group to which the PGPU will be moved
//
// returns:
// - void
func (client *XenAPIClient) PGPU_set_GPU_group(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.set_GPU_group", session_id, self, value)
}

// PGPU_set_enabled_VGPU_types
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, The PGPU on which we are enabling a set of VGPU types
// - value, VGPU_type ref set, The VGPU types to enable
//
// returns:
// - void
func (client *XenAPIClient) PGPU_set_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.set_enabled_VGPU_types", session_id, self, value)
}

// PGPU_remove_enabled_VGPU_types
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, The PGPU from which we are removing an enabled VGPU type
// - value, VGPU_type ref, The VGPU type to disable
//
// returns:
// - void
func (client *XenAPIClient) PGPU_remove_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.remove_enabled_VGPU_types", session_id, self, value)
}

// PGPU_add_enabled_VGPU_types
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, The PGPU to which we are adding an enabled VGPU type
// - value, VGPU_type ref, The VGPU type to enable
//
// returns:
// - void
func (client *XenAPIClient) PGPU_add_enabled_VGPU_types(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.add_enabled_VGPU_types", session_id, self, value)
}

// PGPU_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given PGPU.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) PGPU_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("PGPU.remove_from_other_config", session_id, self, key)
}

// PGPU_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) PGPU_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("PGPU.add_to_other_config", session_id, self, key, value)
}

// PGPU_set_other_config
//
// Set the other_config field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) PGPU_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("PGPU.set_other_config", session_id, self, value)
}

// PGPU_get_is_system_display_device
//
// Get the is_system_display_device field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) PGPU_get_is_system_display_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_is_system_display_device", session_id, self)
}

// PGPU_get_dom0_access
//
// Get the dom0_access field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - enum pgpu_dom0_access
// - value of the field
func (client *XenAPIClient) PGPU_get_dom0_access(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_dom0_access", session_id, self)
}

// PGPU_get_supported_VGPU_max_capacities
//
// Get the supported_VGPU_max_capacities field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - (VGPU_type ref -> int) map
// - value of the field
func (client *XenAPIClient) PGPU_get_supported_VGPU_max_capacities(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_supported_VGPU_max_capacities", session_id, self)
}

// PGPU_get_resident_VGPUs
//
// Get the resident_VGPUs field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - VGPU ref set
// - value of the field
func (client *XenAPIClient) PGPU_get_resident_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_resident_VGPUs", session_id, self)
}

// PGPU_get_enabled_VGPU_types
//
// Get the enabled_VGPU_types field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - VGPU_type ref set
// - value of the field
func (client *XenAPIClient) PGPU_get_enabled_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_enabled_VGPU_types", session_id, self)
}

// PGPU_get_supported_VGPU_types
//
// Get the supported_VGPU_types field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - VGPU_type ref set
// - value of the field
func (client *XenAPIClient) PGPU_get_supported_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_supported_VGPU_types", session_id, self)
}

// PGPU_get_other_config
//
// Get the other_config field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) PGPU_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_other_config", session_id, self)
}

// PGPU_get_host
//
// Get the host field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - host ref
// - value of the field
func (client *XenAPIClient) PGPU_get_host(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_host", session_id, self)
}

// PGPU_get_GPU_group
//
// Get the GPU_group field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - GPU_group ref
// - value of the field
func (client *XenAPIClient) PGPU_get_GPU_group(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_GPU_group", session_id, self)
}

// PGPU_get_PCI
//
// Get the PCI field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - PCI ref
// - value of the field
func (client *XenAPIClient) PGPU_get_PCI(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_PCI", session_id, self)
}

// PGPU_get_uuid
//
// Get the uuid field of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) PGPU_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_uuid", session_id, self)
}

// PGPU_get_by_uuid
//
// Get a reference to the PGPU instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - PGPU ref
// - reference to the object
func (client *XenAPIClient) PGPU_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_by_uuid", session_id, uuid)
}

// PGPU_get_record
//
// Get a record containing the current state of the given PGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, PGPU ref, reference to the object
//
// returns:
// - PGPU record
// - all fields from the object
func (client *XenAPIClient) PGPU_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("PGPU.get_record", session_id, self)
}

// GPU_group_get_all_records
//
// Return a map of GPU_group references to GPU_group records for all GPU_groups known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (GPU_group ref -> GPU_group record) map
// - records of all objects
func (client *XenAPIClient) GPU_group_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_all_records", session_id)
}

// GPU_group_get_all
//
// Return a list of all the GPU_groups known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - GPU_group ref set
// - references to all objects
func (client *XenAPIClient) GPU_group_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_all", session_id)
}

// GPU_group_get_remaining_capacity
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, The GPU group to query
// - vgpu_type, VGPU_type ref, The VGPU_type for which the remaining capacity will be calculated
//
// returns:
// - int
// - The number of VGPUs of the given type which can still be started on the PGPUs in the group
func (client *XenAPIClient) GPU_group_get_remaining_capacity(session_id interface{}, self interface{}, vgpu_type interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_remaining_capacity", session_id, self, vgpu_type)
}

// GPU_group_destroy
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, The vGPU to destroy
//
// returns:
// - void
func (client *XenAPIClient) GPU_group_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.destroy", session_id, self)
}

// GPU_group_create
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - name_label, string,
// - name_description, string,
// - other_config, (string -> string) map,
//
// returns:
// - GPU_group ref
// -
func (client *XenAPIClient) GPU_group_create(session_id interface{}, name_label string, name_description string, other_config map[string]string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.create", session_id, name_label, name_description, other_config)
}

// GPU_group_set_allocation_algorithm
//
// Set the allocation_algorithm field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
// - value, enum allocation_algorithm, New value to set
//
// returns:
// - void
func (client *XenAPIClient) GPU_group_set_allocation_algorithm(session_id interface{}, self interface{}, value interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_allocation_algorithm", session_id, self, value)
}

// GPU_group_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given GPU_group.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) GPU_group_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.remove_from_other_config", session_id, self, key)
}

// GPU_group_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) GPU_group_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.add_to_other_config", session_id, self, key, value)
}

// GPU_group_set_other_config
//
// Set the other_config field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) GPU_group_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_other_config", session_id, self, value)
}

// GPU_group_set_name_description
//
// Set the name/description field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) GPU_group_set_name_description(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_name_description", session_id, self, value)
}

// GPU_group_set_name_label
//
// Set the name/label field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
// - value, string, New value to set
//
// returns:
// - void
func (client *XenAPIClient) GPU_group_set_name_label(session_id interface{}, self interface{}, value string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.set_name_label", session_id, self, value)
}

// GPU_group_get_enabled_VGPU_types
//
// Get the enabled_VGPU_types field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - VGPU_type ref set
// - value of the field
func (client *XenAPIClient) GPU_group_get_enabled_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_enabled_VGPU_types", session_id, self)
}

// GPU_group_get_supported_VGPU_types
//
// Get the supported_VGPU_types field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - VGPU_type ref set
// - value of the field
func (client *XenAPIClient) GPU_group_get_supported_VGPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_supported_VGPU_types", session_id, self)
}

// GPU_group_get_allocation_algorithm
//
// Get the allocation_algorithm field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - enum allocation_algorithm
// - value of the field
func (client *XenAPIClient) GPU_group_get_allocation_algorithm(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_allocation_algorithm", session_id, self)
}

// GPU_group_get_other_config
//
// Get the other_config field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) GPU_group_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_other_config", session_id, self)
}

// GPU_group_get_GPU_types
//
// Get the GPU_types field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - string set
// - value of the field
func (client *XenAPIClient) GPU_group_get_GPU_types(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_GPU_types", session_id, self)
}

// GPU_group_get_VGPUs
//
// Get the VGPUs field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - VGPU ref set
// - value of the field
func (client *XenAPIClient) GPU_group_get_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_VGPUs", session_id, self)
}

// GPU_group_get_PGPUs
//
// Get the PGPUs field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - PGPU ref set
// - value of the field
func (client *XenAPIClient) GPU_group_get_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_PGPUs", session_id, self)
}

// GPU_group_get_name_description
//
// Get the name/description field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) GPU_group_get_name_description(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_name_description", session_id, self)
}

// GPU_group_get_name_label
//
// Get the name/label field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) GPU_group_get_name_label(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_name_label", session_id, self)
}

// GPU_group_get_uuid
//
// Get the uuid field of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) GPU_group_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_uuid", session_id, self)
}

// GPU_group_get_by_name_label
//
// Get all the GPU_group instances with the given label.
//
// params:
// - session_id, session ref, Reference to a valid session
// - label, string, label of object to return
//
// returns:
// - GPU_group ref set
// - references to objects with matching names
func (client *XenAPIClient) GPU_group_get_by_name_label(session_id interface{}, label string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_by_name_label", session_id, label)
}

// GPU_group_get_by_uuid
//
// Get a reference to the GPU_group instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - GPU_group ref
// - reference to the object
func (client *XenAPIClient) GPU_group_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_by_uuid", session_id, uuid)
}

// GPU_group_get_record
//
// Get a record containing the current state of the given GPU_group.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, GPU_group ref, reference to the object
//
// returns:
// - GPU_group record
// - all fields from the object
func (client *XenAPIClient) GPU_group_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("GPU_group.get_record", session_id, self)
}

// VGPU_get_all_records
//
// Return a map of VGPU references to VGPU records for all VGPUs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VGPU ref -> VGPU record) map
// - records of all objects
func (client *XenAPIClient) VGPU_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_all_records", session_id)
}

// VGPU_get_all
//
// Return a list of all the VGPUs known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VGPU ref set
// - references to all objects
func (client *XenAPIClient) VGPU_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_all", session_id)
}

// VGPU_destroy
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, The vGPU to destroy
//
// returns:
// - void
func (client *XenAPIClient) VGPU_destroy(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.destroy", session_id, self)
}

// VGPU_create
//
//
//
// params:
// - session_id, session ref, Reference to a valid session
// - VM, VM ref,
// - GPU_group, GPU_group ref,
// - device, string,
// - other_config, (string -> string) map,
// - a_type, VGPU_type ref,
//
// returns:
// - VGPU ref
// - reference to the newly created object
func (client *XenAPIClient) VGPU_create(session_id interface{}, VM interface{}, GPU_group interface{}, device string, other_config map[string]string, a_type interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.create", session_id, VM, GPU_group, device, other_config, a_type)
}

// VGPU_remove_from_other_config
//
// Remove the given key and its corresponding value from the other_config field of the given VGPU.  If the key is not in that Map, then do nothing.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
// - key, string, Key to remove
//
// returns:
// - void
func (client *XenAPIClient) VGPU_remove_from_other_config(session_id interface{}, self interface{}, key string) (i interface{}, err error) {
	return client.RPCCall("VGPU.remove_from_other_config", session_id, self, key)
}

// VGPU_add_to_other_config
//
// Add the given key-value pair to the other_config field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
// - key, string, Key to add
// - value, string, Value to add
//
// returns:
// - void
func (client *XenAPIClient) VGPU_add_to_other_config(session_id interface{}, self interface{}, key string, value string) (i interface{}, err error) {
	return client.RPCCall("VGPU.add_to_other_config", session_id, self, key, value)
}

// VGPU_set_other_config
//
// Set the other_config field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
// - value, (string -> string) map, New value to set
//
// returns:
// - void
func (client *XenAPIClient) VGPU_set_other_config(session_id interface{}, self interface{}, value map[string]string) (i interface{}, err error) {
	return client.RPCCall("VGPU.set_other_config", session_id, self, value)
}

// VGPU_get_resident_on
//
// Get the resident_on field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - PGPU ref
// - value of the field
func (client *XenAPIClient) VGPU_get_resident_on(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_resident_on", session_id, self)
}

// VGPU_get_type
//
// Get the type field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - VGPU_type ref
// - value of the field
func (client *XenAPIClient) VGPU_get_type(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_type", session_id, self)
}

// VGPU_get_other_config
//
// Get the other_config field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - (string -> string) map
// - value of the field
func (client *XenAPIClient) VGPU_get_other_config(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_other_config", session_id, self)
}

// VGPU_get_currently_attached
//
// Get the currently_attached field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - bool
// - value of the field
func (client *XenAPIClient) VGPU_get_currently_attached(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_currently_attached", session_id, self)
}

// VGPU_get_device
//
// Get the device field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VGPU_get_device(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_device", session_id, self)
}

// VGPU_get_GPU_group
//
// Get the GPU_group field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - GPU_group ref
// - value of the field
func (client *XenAPIClient) VGPU_get_GPU_group(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_GPU_group", session_id, self)
}

// VGPU_get_VM
//
// Get the VM field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - VM ref
// - value of the field
func (client *XenAPIClient) VGPU_get_VM(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_VM", session_id, self)
}

// VGPU_get_uuid
//
// Get the uuid field of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VGPU_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_uuid", session_id, self)
}

// VGPU_get_by_uuid
//
// Get a reference to the VGPU instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VGPU ref
// - reference to the object
func (client *XenAPIClient) VGPU_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_by_uuid", session_id, uuid)
}

// VGPU_get_record
//
// Get a record containing the current state of the given VGPU.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU ref, reference to the object
//
// returns:
// - VGPU record
// - all fields from the object
func (client *XenAPIClient) VGPU_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPU.get_record", session_id, self)
}

// VGPUType_get_all_records
//
// Return a map of VGPU_type references to VGPU_type records for all VGPU_types known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - (VGPU_type ref -> VGPU_type record) map
// - records of all objects
func (client *XenAPIClient) VGPUType_get_all_records(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_all_records", session_id)
}

// VGPUType_get_all
//
// Return a list of all the VGPU_types known to the system.
//
// params:
// - session_id, session ref, Reference to a valid session
//
// returns:
// - VGPU_type ref set
// - references to all objects
func (client *XenAPIClient) VGPUType_get_all(session_id interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_all", session_id)
}

// VGPUType_get_enabled_on_GPU_groups
//
// Get the enabled_on_GPU_groups field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - GPU_group ref set
// - value of the field
func (client *XenAPIClient) VGPUType_get_enabled_on_GPU_groups(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_enabled_on_GPU_groups", session_id, self)
}

// VGPUType_get_supported_on_GPU_groups
//
// Get the supported_on_GPU_groups field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - GPU_group ref set
// - value of the field
func (client *XenAPIClient) VGPUType_get_supported_on_GPU_groups(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_supported_on_GPU_groups", session_id, self)
}

// VGPUType_get_VGPUs
//
// Get the VGPUs field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - VGPU ref set
// - value of the field
func (client *XenAPIClient) VGPUType_get_VGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_VGPUs", session_id, self)
}

// VGPUType_get_enabled_on_PGPUs
//
// Get the enabled_on_PGPUs field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - PGPU ref set
// - value of the field
func (client *XenAPIClient) VGPUType_get_enabled_on_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_enabled_on_PGPUs", session_id, self)
}

// VGPUType_get_supported_on_PGPUs
//
// Get the supported_on_PGPUs field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - PGPU ref set
// - value of the field
func (client *XenAPIClient) VGPUType_get_supported_on_PGPUs(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_supported_on_PGPUs", session_id, self)
}

// VGPUType_get_max_resolution_y
//
// Get the max_resolution_y field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VGPUType_get_max_resolution_y(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_max_resolution_y", session_id, self)
}

// VGPUType_get_max_resolution_x
//
// Get the max_resolution_x field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VGPUType_get_max_resolution_x(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_max_resolution_x", session_id, self)
}

// VGPUType_get_max_heads
//
// Get the max_heads field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VGPUType_get_max_heads(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_max_heads", session_id, self)
}

// VGPUType_get_framebuffer_size
//
// Get the framebuffer_size field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - int
// - value of the field
func (client *XenAPIClient) VGPUType_get_framebuffer_size(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_framebuffer_size", session_id, self)
}

// VGPUType_get_model_name
//
// Get the model_name field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VGPUType_get_model_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_model_name", session_id, self)
}

// VGPUType_get_vendor_name
//
// Get the vendor_name field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VGPUType_get_vendor_name(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_vendor_name", session_id, self)
}

// VGPUType_get_uuid
//
// Get the uuid field of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - string
// - value of the field
func (client *XenAPIClient) VGPUType_get_uuid(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_uuid", session_id, self)
}

// VGPUType_get_by_uuid
//
// Get a reference to the VGPU_type instance with the specified UUID.
//
// params:
// - session_id, session ref, Reference to a valid session
// - uuid, string, UUID of object to return
//
// returns:
// - VGPU_type ref
// - reference to the object
func (client *XenAPIClient) VGPUType_get_by_uuid(session_id interface{}, uuid string) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_by_uuid", session_id, uuid)
}

// VGPUType_get_record
//
// Get a record containing the current state of the given VGPU_type.
//
// params:
// - session_id, session ref, Reference to a valid session
// - self, VGPU_type ref, reference to the object
//
// returns:
// - VGPU_type record
// - all fields from the object
func (client *XenAPIClient) VGPUType_get_record(session_id interface{}, self interface{}) (i interface{}, err error) {
	return client.RPCCall("VGPUType.get_record", session_id, self)
}

const (
	XS_HOST   = "XS_HOST"
	XS_USER   = "XS_USER"
	XS_PASSWD = "XS_PASSWD"
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

func main() {
	host, user, passwd := os.Getenv(XS_HOST), os.Getenv(XS_USER), os.Getenv(XS_PASSWD)
	client := NewXenAPIClient(host)

	fmt.Println("NewXenAPIClient for: " + client.Host)
	fmt.Println("login_with_password: " + user + "/" + passwd)

	session, err := client.session_login_with_password(user, passwd, "", "")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	vms, err := client.VM_get_all(session)

	// list running VMs
	for _, vm := range vms.([]interface{}) {
		ps, err := client.VM_get_power_state(session, vm)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if ps.(string) == "Running" {
			fmt.Printf("%+v VM_get_power_state: %+v\n", vm, ps)
		}
	}

	_, err = client.session_local_logout(session)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
