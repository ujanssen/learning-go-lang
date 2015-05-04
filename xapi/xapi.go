package main

import (
	"fmt"
)


// A session
type Session struct {
	
	Uuid     string
	
	This_host     string // host_ref
	
	This_user     string // user_ref
	
	Last_active     string // datetime
	
	Pool     bool
	
	Other_config     string // string__string_map
	
	Is_local_superuser     bool
	
	Subject     string // subject_ref
	
	Validation_time     string // datetime
	
	Auth_user_sid     string
	
	Auth_user_name     string
	
	Rbac_permissions     string // string_set
	
	Tasks     string // task_ref_set
	
	Parent     string // session_ref
	
	Originator     string
	
}

// Management of remote authentication services
type Auth struct {
	
}

// A user or group that can log in xapi
type Subject struct {
	
	Uuid     string
	
	Subject_identifier     string
	
	Other_config     string // string__string_map
	
	Roles     string // role_ref_set
	
}

// A set of permissions associated with a subject
type Role struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Subroles     string // role_ref_set
	
}

// A long-running asynchronous task
type Task struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Allowed_operations     string // enum_task_allowed_operations_set
	
	Current_operations     string // string__enum_task_allowed_operations_map
	
	Created     string // datetime
	
	Finished     string // datetime
	
	Status     string // enum_task_status_type
	
	Resident_on     string // host_ref
	
	Progress     string // float
	
	Type     string
	
	Result     string
	
	Error_info     string // string_set
	
	Other_config     string // string__string_map
	
	Subtask_of     string // task_ref
	
	Subtasks     string // task_ref_set
	
	Backtrace     string
	
}

// Asynchronous event registration and handling
type Event struct {
	
	Id     string // int
	
	Timestamp     string // datetime
	
	Class     string
	
	Operation     string // enum_event_operation
	
	Ref     string
	
	Obj_uuid     string
	
}

// Pool-wide information
type Pool struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Master     string // host_ref
	
	Default_SR     string // SR_ref
	
	Suspend_image_SR     string // SR_ref
	
	Crash_dump_SR     string // SR_ref
	
	Other_config     string // string__string_map
	
	Ha_enabled     bool
	
	Ha_configuration     string // string__string_map
	
	Ha_statefiles     string // string_set
	
	Ha_host_failures_to_tolerate     string // int
	
	Ha_plan_exists_for     string // int
	
	Ha_allow_overcommit     bool
	
	Ha_overcommitted     bool
	
	Blobs     string // string__blob_ref_map
	
	Tags     string // string_set
	
	Gui_config     string // string__string_map
	
	Wlb_url     string
	
	Wlb_username     string
	
	Wlb_enabled     bool
	
	Wlb_verify_cert     bool
	
	Redo_log_enabled     bool
	
	Redo_log_vdi     string // VDI_ref
	
	Vswitch_controller     string
	
	Restrictions     string // string__string_map
	
	Metadata_VDIs     string // VDI_ref_set
	
}

// Pool-wide patches
type Pool_patch struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Version     string
	
	Size     string // int
	
	Pool_applied     bool
	
	Host_patches     string // host_patch_ref_set
	
	After_apply_guidance     string // enum_after_apply_guidance_set
	
	Other_config     string // string__string_map
	
}

// A virtual machine (or 'guest').
type VM struct {
	
	Uuid     string
	
	Allowed_operations     string // enum_vm_operations_set
	
	Current_operations     string // string__enum_vm_operations_map
	
	Power_state     string // enum_vm_power_state
	
	Name_label     string
	
	Name_description     string
	
	User_version     string // int
	
	Is_a_template     bool
	
	Suspend_VDI     string // VDI_ref
	
	Resident_on     string // host_ref
	
	Affinity     string // host_ref
	
	Memory_overhead     string // int
	
	Memory_target     string // int
	
	Memory_static_max     string // int
	
	Memory_dynamic_max     string // int
	
	Memory_dynamic_min     string // int
	
	Memory_static_min     string // int
	
	VCPUs_params     string // string__string_map
	
	VCPUs_max     string // int
	
	VCPUs_at_startup     string // int
	
	Actions_after_shutdown     string // enum_on_normal_exit
	
	Actions_after_reboot     string // enum_on_normal_exit
	
	Actions_after_crash     string // enum_on_crash_behaviour
	
	Consoles     string // console_ref_set
	
	VIFs     string // VIF_ref_set
	
	VBDs     string // VBD_ref_set
	
	Crash_dumps     string // crashdump_ref_set
	
	VTPMs     string // VTPM_ref_set
	
	PV_bootloader     string
	
	PV_kernel     string
	
	PV_ramdisk     string
	
	PV_args     string
	
	PV_bootloader_args     string
	
	PV_legacy_args     string
	
	HVM_boot_policy     string
	
	HVM_boot_params     string // string__string_map
	
	HVM_shadow_multiplier     string // float
	
	Platform     string // string__string_map
	
	PCI_bus     string
	
	Other_config     string // string__string_map
	
	Domid     string // int
	
	Domarch     string
	
	Last_boot_CPU_flags     string // string__string_map
	
	Is_control_domain     bool
	
	Metrics     string // VM_metrics_ref
	
	Guest_metrics     string // VM_guest_metrics_ref
	
	Last_booted_record     string
	
	Recommendations     string
	
	Xenstore_data     string // string__string_map
	
	Ha_always_run     bool
	
	Ha_restart_priority     string
	
	Is_a_snapshot     bool
	
	Snapshot_of     string // VM_ref
	
	Snapshots     string // VM_ref_set
	
	Snapshot_time     string // datetime
	
	Transportable_snapshot_id     string
	
	Blobs     string // string__blob_ref_map
	
	Tags     string // string_set
	
	Blocked_operations     string // enum_vm_operations__string_map
	
	Snapshot_info     string // string__string_map
	
	Snapshot_metadata     string
	
	Parent     string // VM_ref
	
	Children     string // VM_ref_set
	
	Bios_strings     string // string__string_map
	
	Protection_policy     string // VMPP_ref
	
	Is_snapshot_from_vmpp     bool
	
	Appliance     string // VM_appliance_ref
	
	Start_delay     string // int
	
	Shutdown_delay     string // int
	
	Order     string // int
	
	VGPUs     string // VGPU_ref_set
	
	Attached_PCIs     string // PCI_ref_set
	
	Suspend_SR     string // SR_ref
	
	Version     string // int
	
	Generation_id     string
	
}

// The metrics associated with a VM
type VM_metrics struct {
	
	Uuid     string
	
	Memory_actual     string // int
	
	VCPUs_number     string // int
	
	VCPUs_utilisation     string // int__float_map
	
	VCPUs_CPU     string // int__int_map
	
	VCPUs_params     string // string__string_map
	
	VCPUs_flags     string // int__string_set_map
	
	State     string // string_set
	
	Start_time     string // datetime
	
	Install_time     string // datetime
	
	Last_updated     string // datetime
	
	Other_config     string // string__string_map
	
}

// The metrics reported by the guest (as opposed to inferred from outside)
type VM_guest_metrics struct {
	
	Uuid     string
	
	Os_version     string // string__string_map
	
	PV_drivers_version     string // string__string_map
	
	PV_drivers_up_to_date     bool
	
	Memory     string // string__string_map
	
	Disks     string // string__string_map
	
	Networks     string // string__string_map
	
	Other     string // string__string_map
	
	Last_updated     string // datetime
	
	Other_config     string // string__string_map
	
	Live     bool
	
}

// VM Protection Policy
type VMPP struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Is_policy_enabled     bool
	
	Backup_type     string // enum_vmpp_backup_type
	
	Backup_retention_value     string // int
	
	Backup_frequency     string // enum_vmpp_backup_frequency
	
	Backup_schedule     string // string__string_map
	
	Is_backup_running     bool
	
	Backup_last_run_time     string // datetime
	
	Archive_target_type     string // enum_vmpp_archive_target_type
	
	Archive_target_config     string // string__string_map
	
	Archive_frequency     string // enum_vmpp_archive_frequency
	
	Archive_schedule     string // string__string_map
	
	Is_archive_running     bool
	
	Archive_last_run_time     string // datetime
	
	VMs     string // VM_ref_set
	
	Is_alarm_enabled     bool
	
	Alarm_config     string // string__string_map
	
	Recent_alerts     string // string_set
	
}

// VM appliance
type VM_appliance struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Allowed_operations     string // enum_vm_appliance_operation_set
	
	Current_operations     string // string__enum_vm_appliance_operation_map
	
	VMs     string // VM_ref_set
	
}

// DR task
type DR_task struct {
	
	Uuid     string
	
	Introduced_SRs     string // SR_ref_set
	
}

// A physical host
type Host struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Memory_overhead     string // int
	
	Allowed_operations     string // enum_host_allowed_operations_set
	
	Current_operations     string // string__enum_host_allowed_operations_map
	
	API_version_major     string // int
	
	API_version_minor     string // int
	
	API_version_vendor     string
	
	API_version_vendor_implementation     string // string__string_map
	
	Enabled     bool
	
	Software_version     string // string__string_map
	
	Other_config     string // string__string_map
	
	Capabilities     string // string_set
	
	Cpu_configuration     string // string__string_map
	
	Sched_policy     string
	
	Supported_bootloaders     string // string_set
	
	Resident_VMs     string // VM_ref_set
	
	Logging     string // string__string_map
	
	PIFs     string // PIF_ref_set
	
	Suspend_image_sr     string // SR_ref
	
	Crash_dump_sr     string // SR_ref
	
	Crashdumps     string // host_crashdump_ref_set
	
	Patches     string // host_patch_ref_set
	
	PBDs     string // PBD_ref_set
	
	Host_CPUs     string // host_cpu_ref_set
	
	Cpu_info     string // string__string_map
	
	Hostname     string
	
	Address     string
	
	Metrics     string // host_metrics_ref
	
	License_params     string // string__string_map
	
	Ha_statefiles     string // string_set
	
	Ha_network_peers     string // string_set
	
	Blobs     string // string__blob_ref_map
	
	Tags     string // string_set
	
	External_auth_type     string
	
	External_auth_service_name     string
	
	External_auth_configuration     string // string__string_map
	
	Edition     string
	
	License_server     string // string__string_map
	
	Bios_strings     string // string__string_map
	
	Power_on_mode     string
	
	Power_on_config     string // string__string_map
	
	Local_cache_sr     string // SR_ref
	
	Chipset_info     string // string__string_map
	
	PCIs     string // PCI_ref_set
	
	PGPUs     string // PGPU_ref_set
	
	Guest_VCPUs_params     string // string__string_map
	
	Display     string // enum_host_display
	
}

// Represents a host crash dump
type Host_crashdump struct {
	
	Uuid     string
	
	Host     string // host_ref
	
	Timestamp     string // datetime
	
	Size     string // int
	
	Other_config     string // string__string_map
	
}

// Represents a patch stored on a server
type Host_patch struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Version     string
	
	Host     string // host_ref
	
	Applied     bool
	
	Timestamp_applied     string // datetime
	
	Size     string // int
	
	Pool_patch     string // pool_patch_ref
	
	Other_config     string // string__string_map
	
}

// The metrics associated with a host
type Host_metrics struct {
	
	Uuid     string
	
	Memory_total     string // int
	
	Memory_free     string // int
	
	Live     bool
	
	Last_updated     string // datetime
	
	Other_config     string // string__string_map
	
}

// A physical CPU
type Host_cpu struct {
	
	Uuid     string
	
	Host     string // host_ref
	
	Number     string // int
	
	Vendor     string
	
	Speed     string // int
	
	Modelname     string
	
	Family     string // int
	
	Model     string // int
	
	Stepping     string
	
	Flags     string
	
	Features     string
	
	Utilisation     string // float
	
	Other_config     string // string__string_map
	
}

// A virtual network
type Network struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Allowed_operations     string // enum_network_operations_set
	
	Current_operations     string // string__enum_network_operations_map
	
	VIFs     string // VIF_ref_set
	
	PIFs     string // PIF_ref_set
	
	MTU     string // int
	
	Other_config     string // string__string_map
	
	Bridge     string
	
	Blobs     string // string__blob_ref_map
	
	Tags     string // string_set
	
	Default_locking_mode     string // enum_network_default_locking_mode
	
	Assigned_ips     string // VIF_ref__string_map
	
}

// A virtual network interface
type VIF struct {
	
	Uuid     string
	
	Allowed_operations     string // enum_vif_operations_set
	
	Current_operations     string // string__enum_vif_operations_map
	
	Device     string
	
	Network     string // network_ref
	
	VM     string // VM_ref
	
	MAC     string
	
	MTU     string // int
	
	Other_config     string // string__string_map
	
	Currently_attached     bool
	
	Status_code     string // int
	
	Status_detail     string
	
	Runtime_properties     string // string__string_map
	
	Qos_algorithm_type     string
	
	Qos_algorithm_params     string // string__string_map
	
	Qos_supported_algorithms     string // string_set
	
	Metrics     string // VIF_metrics_ref
	
	MAC_autogenerated     bool
	
	Locking_mode     string // enum_vif_locking_mode
	
	Ipv4_allowed     string // string_set
	
	Ipv6_allowed     string // string_set
	
}

// The metrics associated with a virtual network device
type VIF_metrics struct {
	
	Uuid     string
	
	Io_read_kbs     string // float
	
	Io_write_kbs     string // float
	
	Last_updated     string // datetime
	
	Other_config     string // string__string_map
	
}

// A physical network interface (note separate VLANs are represented as several PIFs)
type PIF struct {
	
	Uuid     string
	
	Device     string
	
	Network     string // network_ref
	
	Host     string // host_ref
	
	MAC     string
	
	MTU     string // int
	
	VLAN     string // int
	
	Metrics     string // PIF_metrics_ref
	
	Physical     bool
	
	Currently_attached     bool
	
	Ip_configuration_mode     string // enum_ip_configuration_mode
	
	IP     string
	
	Netmask     string
	
	Gateway     string
	
	DNS     string
	
	Bond_slave_of     string // Bond_ref
	
	Bond_master_of     string // Bond_ref_set
	
	VLAN_master_of     string // VLAN_ref
	
	VLAN_slave_of     string // VLAN_ref_set
	
	Management     bool
	
	Other_config     string // string__string_map
	
	Disallow_unplug     bool
	
	Tunnel_access_PIF_of     string // tunnel_ref_set
	
	Tunnel_transport_PIF_of     string // tunnel_ref_set
	
	Ipv6_configuration_mode     string // enum_ipv6_configuration_mode
	
	IPv6     string // string_set
	
	Ipv6_gateway     string
	
	Primary_address_type     string // enum_primary_address_type
	
	Managed     bool
	
	Properties     string // string__string_map
	
}

// The metrics associated with a physical network interface
type PIF_metrics struct {
	
	Uuid     string
	
	Io_read_kbs     string // float
	
	Io_write_kbs     string // float
	
	Carrier     bool
	
	Vendor_id     string
	
	Vendor_name     string
	
	Device_id     string
	
	Device_name     string
	
	Speed     string // int
	
	Duplex     bool
	
	Pci_bus_path     string
	
	Last_updated     string // datetime
	
	Other_config     string // string__string_map
	
}

// 
type Bond struct {
	
	Uuid     string
	
	Master     string // PIF_ref
	
	Slaves     string // PIF_ref_set
	
	Other_config     string // string__string_map
	
	Primary_slave     string // PIF_ref
	
	Mode     string // enum_bond_mode
	
	Properties     string // string__string_map
	
	Links_up     string // int
	
}

// A VLAN mux/demux
type VLAN struct {
	
	Uuid     string
	
	Tagged_PIF     string // PIF_ref
	
	Untagged_PIF     string // PIF_ref
	
	Tag     string // int
	
	Other_config     string // string__string_map
	
}

// A storage manager plugin
type SM struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Type     string
	
	Vendor     string
	
	Copyright     string
	
	Version     string
	
	Required_api_version     string
	
	Configuration     string // string__string_map
	
	Capabilities     string // string_set
	
	Features     string // string__int_map
	
	Other_config     string // string__string_map
	
	Driver_filename     string
	
}

// A storage repository
type SR struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Allowed_operations     string // enum_storage_operations_set
	
	Current_operations     string // string__enum_storage_operations_map
	
	VDIs     string // VDI_ref_set
	
	PBDs     string // PBD_ref_set
	
	Virtual_allocation     string // int
	
	Physical_utilisation     string // int
	
	Physical_size     string // int
	
	Type     string
	
	Content_type     string
	
	Shared     bool
	
	Other_config     string // string__string_map
	
	Tags     string // string_set
	
	Sm_config     string // string__string_map
	
	Blobs     string // string__blob_ref_map
	
	Local_cache_enabled     bool
	
	Introduced_by     string // DR_task_ref
	
}

// A virtual disk image
type VDI struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Allowed_operations     string // enum_vdi_operations_set
	
	Current_operations     string // string__enum_vdi_operations_map
	
	SR     string // SR_ref
	
	VBDs     string // VBD_ref_set
	
	Crash_dumps     string // crashdump_ref_set
	
	Virtual_size     string // int
	
	Physical_utilisation     string // int
	
	Type     string // enum_vdi_type
	
	Sharable     bool
	
	Read_only     bool
	
	Other_config     string // string__string_map
	
	Storage_lock     bool
	
	Location     string
	
	Managed     bool
	
	Missing     bool
	
	Parent     string // VDI_ref
	
	Xenstore_data     string // string__string_map
	
	Sm_config     string // string__string_map
	
	Is_a_snapshot     bool
	
	Snapshot_of     string // VDI_ref
	
	Snapshots     string // VDI_ref_set
	
	Snapshot_time     string // datetime
	
	Tags     string // string_set
	
	Allow_caching     bool
	
	On_boot     string // enum_on_boot
	
	Metadata_of_pool     string // pool_ref
	
	Metadata_latest     bool
	
}

// A virtual block device
type VBD struct {
	
	Uuid     string
	
	Allowed_operations     string // enum_vbd_operations_set
	
	Current_operations     string // string__enum_vbd_operations_map
	
	VM     string // VM_ref
	
	VDI     string // VDI_ref
	
	Device     string
	
	Userdevice     string
	
	Bootable     bool
	
	Mode     string // enum_vbd_mode
	
	Type     string // enum_vbd_type
	
	Unpluggable     bool
	
	Storage_lock     bool
	
	Empty     bool
	
	Other_config     string // string__string_map
	
	Currently_attached     bool
	
	Status_code     string // int
	
	Status_detail     string
	
	Runtime_properties     string // string__string_map
	
	Qos_algorithm_type     string
	
	Qos_algorithm_params     string // string__string_map
	
	Qos_supported_algorithms     string // string_set
	
	Metrics     string // VBD_metrics_ref
	
}

// The metrics associated with a virtual block device
type VBD_metrics struct {
	
	Uuid     string
	
	Io_read_kbs     string // float
	
	Io_write_kbs     string // float
	
	Last_updated     string // datetime
	
	Other_config     string // string__string_map
	
}

// The physical block devices through which hosts access SRs
type PBD struct {
	
	Uuid     string
	
	Host     string // host_ref
	
	SR     string // SR_ref
	
	Device_config     string // string__string_map
	
	Currently_attached     bool
	
	Other_config     string // string__string_map
	
}

// A VM crashdump
type Crashdump struct {
	
	Uuid     string
	
	VM     string // VM_ref
	
	VDI     string // VDI_ref
	
	Other_config     string // string__string_map
	
}

// A virtual TPM device
type VTPM struct {
	
	Uuid     string
	
	VM     string // VM_ref
	
	Backend     string // VM_ref
	
}

// A console
type Console struct {
	
	Uuid     string
	
	Protocol     string // enum_console_protocol
	
	Location     string
	
	VM     string // VM_ref
	
	Other_config     string // string__string_map
	
}

// A user of the system
type User struct {
	
	Uuid     string
	
	Short_name     string
	
	Fullname     string
	
	Other_config     string // string__string_map
	
}

// Data sources for logging in RRDs
type Data_source struct {
	
	Name_label     string
	
	Name_description     string
	
	Enabled     bool
	
	Standard     bool
	
	Units     string
	
	Min     string // float
	
	Max     string // float
	
	Value     string // float
	
}

// A placeholder for a binary blob
type Blob struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	Size     string // int
	
	Public     bool
	
	Last_updated     string // datetime
	
	Mime_type     string
	
}

// An message for the attention of the administrator
type Message struct {
	
	Uuid     string
	
	Name     string
	
	Priority     string // int
	
	Cls     string // enum_cls
	
	Obj_uuid     string
	
	Timestamp     string // datetime
	
	Body     string
	
}

// A secret
type Secret struct {
	
	Uuid     string
	
	Value     string
	
	Other_config     string // string__string_map
	
}

// A tunnel for network traffic
type Tunnel struct {
	
	Uuid     string
	
	Access_PIF     string // PIF_ref
	
	Transport_PIF     string // PIF_ref
	
	Status     string // string__string_map
	
	Other_config     string // string__string_map
	
}

// A PCI device
type PCI struct {
	
	Uuid     string
	
	Class_name     string
	
	Vendor_name     string
	
	Device_name     string
	
	Host     string // host_ref
	
	Pci_id     string
	
	Dependencies     string // PCI_ref_set
	
	Other_config     string // string__string_map
	
	Subsystem_vendor_name     string
	
	Subsystem_device_name     string
	
}

// A physical GPU (pGPU)
type PGPU struct {
	
	Uuid     string
	
	PCI     string // PCI_ref
	
	GPU_group     string // GPU_group_ref
	
	Host     string // host_ref
	
	Other_config     string // string__string_map
	
	Supported_VGPU_types     string // VGPU_type_ref_set
	
	Enabled_VGPU_types     string // VGPU_type_ref_set
	
	Resident_VGPUs     string // VGPU_ref_set
	
	Supported_VGPU_max_capacities     string // VGPU_type_ref__int_map
	
	Dom0_access     string // enum_pgpu_dom0_access
	
	Is_system_display_device     bool
	
}

// A group of compatible GPUs across the resource pool
type GPU_group struct {
	
	Uuid     string
	
	Name_label     string
	
	Name_description     string
	
	PGPUs     string // PGPU_ref_set
	
	VGPUs     string // VGPU_ref_set
	
	GPU_types     string // string_set
	
	Other_config     string // string__string_map
	
	Allocation_algorithm     string // enum_allocation_algorithm
	
	Supported_VGPU_types     string // VGPU_type_ref_set
	
	Enabled_VGPU_types     string // VGPU_type_ref_set
	
}

// A virtual GPU (vGPU)
type VGPU struct {
	
	Uuid     string
	
	VM     string // VM_ref
	
	GPU_group     string // GPU_group_ref
	
	Device     string
	
	Currently_attached     bool
	
	Other_config     string // string__string_map
	
	Type     string // VGPU_type_ref
	
	Resident_on     string // PGPU_ref
	
}

// A type of virtual GPU
type VGPU_type struct {
	
	Uuid     string
	
	Vendor_name     string
	
	Model_name     string
	
	Framebuffer_size     string // int
	
	Max_heads     string // int
	
	Max_resolution_x     string // int
	
	Max_resolution_y     string // int
	
	Supported_on_PGPUs     string // PGPU_ref_set
	
	Enabled_on_PGPUs     string // PGPU_ref_set
	
	VGPUs     string // VGPU_ref_set
	
	Supported_on_GPU_groups     string // GPU_group_ref_set
	
	Enabled_on_GPU_groups     string // GPU_group_ref_set
	
}


func main() {
		fmt.Println("it compiles")
}
