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

func (Session *Session) logout_subject_identifier() {
}
func (Session *Session) get_all_subject_identifiers() {
}
func (Session *Session) local_logout() {
}
func (Session *Session) slave_local_login_with_password() {
}
func (Session *Session) change_password() {
}
func (Session *Session) logout() {
}
func (Session *Session) login_with_password() {
}
func (Session *Session) remove_from_other_config() {
}
func (Session *Session) add_to_other_config() {
}
func (Session *Session) set_other_config() {
}
func (Session *Session) get_originator() {
}
func (Session *Session) get_parent() {
}
func (Session *Session) get_tasks() {
}
func (Session *Session) get_rbac_permissions() {
}
func (Session *Session) get_auth_user_name() {
}
func (Session *Session) get_auth_user_sid() {
}
func (Session *Session) get_validation_time() {
}
func (Session *Session) get_subject() {
}
func (Session *Session) get_is_local_superuser() {
}
func (Session *Session) get_other_config() {
}
func (Session *Session) get_pool() {
}
func (Session *Session) get_last_active() {
}
func (Session *Session) get_this_user() {
}
func (Session *Session) get_this_host() {
}
func (Session *Session) get_uuid() {
}
func (Session *Session) get_by_uuid() {
}
func (Session *Session) get_record() {
}

func main() {
	fmt.Println("it compiles")
}
