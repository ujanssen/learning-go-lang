package main

import (
	"fmt"
)


// A session
type session struct {
	
	Uuid     string // 
	
	This_host     string // 
	
	This_user     string // 
	
	Last_active     string // 
	
	Pool     string // 
	
	Other_config     string // 
	
	Is_local_superuser     string // 
	
	Subject     string // 
	
	Validation_time     string // 
	
	Auth_user_sid     string // 
	
	Auth_user_name     string // 
	
	Rbac_permissions     string // 
	
	Tasks     string // 
	
	Parent     string // 
	
	Originator     string // 
	
}

// Management of remote authentication services
type auth struct {
	
}

// A user or group that can log in xapi
type subject struct {
	
	Uuid     string // 
	
	Subject_identifier     string // 
	
	Other_config     string // 
	
	Roles     string // 
	
}

// A set of permissions associated with a subject
type role struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Subroles     string // 
	
}

// A long-running asynchronous task
type task struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	Created     string // 
	
	Finished     string // 
	
	Status     string // 
	
	Resident_on     string // 
	
	Progress     string // 
	
	Type     string // 
	
	Result     string // 
	
	Error_info     string // 
	
	Other_config     string // 
	
	Subtask_of     string // 
	
	Subtasks     string // 
	
	Backtrace     string // 
	
}

// Asynchronous event registration and handling
type event struct {
	
	Id     string // 
	
	Timestamp     string // 
	
	Class     string // 
	
	Operation     string // 
	
	Ref     string // 
	
	Obj_uuid     string // 
	
}

// Pool-wide information
type pool struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Master     string // 
	
	Default_SR     string // 
	
	Suspend_image_SR     string // 
	
	Crash_dump_SR     string // 
	
	Other_config     string // 
	
	Ha_enabled     string // 
	
	Ha_configuration     string // 
	
	Ha_statefiles     string // 
	
	Ha_host_failures_to_tolerate     string // 
	
	Ha_plan_exists_for     string // 
	
	Ha_allow_overcommit     string // 
	
	Ha_overcommitted     string // 
	
	Blobs     string // 
	
	Tags     string // 
	
	Gui_config     string // 
	
	Wlb_url     string // 
	
	Wlb_username     string // 
	
	Wlb_enabled     string // 
	
	Wlb_verify_cert     string // 
	
	Redo_log_enabled     string // 
	
	Redo_log_vdi     string // 
	
	Vswitch_controller     string // 
	
	Restrictions     string // 
	
	Metadata_VDIs     string // 
	
}

// Pool-wide patches
type pool_patch struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Version     string // 
	
	Size     string // 
	
	Pool_applied     string // 
	
	Host_patches     string // 
	
	After_apply_guidance     string // 
	
	Other_config     string // 
	
}

// A virtual machine (or 'guest').
type VM struct {
	
	Uuid     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	Power_state     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	User_version     string // 
	
	Is_a_template     string // 
	
	Suspend_VDI     string // 
	
	Resident_on     string // 
	
	Affinity     string // 
	
	Memory_overhead     string // 
	
	Memory_target     string // 
	
	Memory_static_max     string // 
	
	Memory_dynamic_max     string // 
	
	Memory_dynamic_min     string // 
	
	Memory_static_min     string // 
	
	VCPUs_params     string // 
	
	VCPUs_max     string // 
	
	VCPUs_at_startup     string // 
	
	Actions_after_shutdown     string // 
	
	Actions_after_reboot     string // 
	
	Actions_after_crash     string // 
	
	Consoles     string // 
	
	VIFs     string // 
	
	VBDs     string // 
	
	Crash_dumps     string // 
	
	VTPMs     string // 
	
	PV_bootloader     string // 
	
	PV_kernel     string // 
	
	PV_ramdisk     string // 
	
	PV_args     string // 
	
	PV_bootloader_args     string // 
	
	PV_legacy_args     string // 
	
	HVM_boot_policy     string // 
	
	HVM_boot_params     string // 
	
	HVM_shadow_multiplier     string // 
	
	Platform     string // 
	
	PCI_bus     string // 
	
	Other_config     string // 
	
	Domid     string // 
	
	Domarch     string // 
	
	Last_boot_CPU_flags     string // 
	
	Is_control_domain     string // 
	
	Metrics     string // 
	
	Guest_metrics     string // 
	
	Last_booted_record     string // 
	
	Recommendations     string // 
	
	Xenstore_data     string // 
	
	Ha_always_run     string // 
	
	Ha_restart_priority     string // 
	
	Is_a_snapshot     string // 
	
	Snapshot_of     string // 
	
	Snapshots     string // 
	
	Snapshot_time     string // 
	
	Transportable_snapshot_id     string // 
	
	Blobs     string // 
	
	Tags     string // 
	
	Blocked_operations     string // 
	
	Snapshot_info     string // 
	
	Snapshot_metadata     string // 
	
	Parent     string // 
	
	Children     string // 
	
	Bios_strings     string // 
	
	Protection_policy     string // 
	
	Is_snapshot_from_vmpp     string // 
	
	Appliance     string // 
	
	Start_delay     string // 
	
	Shutdown_delay     string // 
	
	Order     string // 
	
	VGPUs     string // 
	
	Attached_PCIs     string // 
	
	Suspend_SR     string // 
	
	Version     string // 
	
	Generation_id     string // 
	
}

// The metrics associated with a VM
type VM_metrics struct {
	
	Uuid     string // 
	
	Memory_actual     string // 
	
	VCPUs_number     string // 
	
	VCPUs_utilisation     string // 
	
	VCPUs_CPU     string // 
	
	VCPUs_params     string // 
	
	VCPUs_flags     string // 
	
	State     string // 
	
	Start_time     string // 
	
	Install_time     string // 
	
	Last_updated     string // 
	
	Other_config     string // 
	
}

// The metrics reported by the guest (as opposed to inferred from outside)
type VM_guest_metrics struct {
	
	Uuid     string // 
	
	Os_version     string // 
	
	PV_drivers_version     string // 
	
	PV_drivers_up_to_date     string // 
	
	Memory     string // 
	
	Disks     string // 
	
	Networks     string // 
	
	Other     string // 
	
	Last_updated     string // 
	
	Other_config     string // 
	
	Live     string // 
	
}

// VM Protection Policy
type VMPP struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Is_policy_enabled     string // 
	
	Backup_type     string // 
	
	Backup_retention_value     string // 
	
	Backup_frequency     string // 
	
	Backup_schedule     string // 
	
	Is_backup_running     string // 
	
	Backup_last_run_time     string // 
	
	Archive_target_type     string // 
	
	Archive_target_config     string // 
	
	Archive_frequency     string // 
	
	Archive_schedule     string // 
	
	Is_archive_running     string // 
	
	Archive_last_run_time     string // 
	
	VMs     string // 
	
	Is_alarm_enabled     string // 
	
	Alarm_config     string // 
	
	Recent_alerts     string // 
	
}

// VM appliance
type VM_appliance struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	VMs     string // 
	
}

// DR task
type DR_task struct {
	
	Uuid     string // 
	
	Introduced_SRs     string // 
	
}

// A physical host
type host struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Memory_overhead     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	API_version_major     string // 
	
	API_version_minor     string // 
	
	API_version_vendor     string // 
	
	API_version_vendor_implementation     string // 
	
	Enabled     string // 
	
	Software_version     string // 
	
	Other_config     string // 
	
	Capabilities     string // 
	
	Cpu_configuration     string // 
	
	Sched_policy     string // 
	
	Supported_bootloaders     string // 
	
	Resident_VMs     string // 
	
	Logging     string // 
	
	PIFs     string // 
	
	Suspend_image_sr     string // 
	
	Crash_dump_sr     string // 
	
	Crashdumps     string // 
	
	Patches     string // 
	
	PBDs     string // 
	
	Host_CPUs     string // 
	
	Cpu_info     string // 
	
	Hostname     string // 
	
	Address     string // 
	
	Metrics     string // 
	
	License_params     string // 
	
	Ha_statefiles     string // 
	
	Ha_network_peers     string // 
	
	Blobs     string // 
	
	Tags     string // 
	
	External_auth_type     string // 
	
	External_auth_service_name     string // 
	
	External_auth_configuration     string // 
	
	Edition     string // 
	
	License_server     string // 
	
	Bios_strings     string // 
	
	Power_on_mode     string // 
	
	Power_on_config     string // 
	
	Local_cache_sr     string // 
	
	Chipset_info     string // 
	
	PCIs     string // 
	
	PGPUs     string // 
	
	Guest_VCPUs_params     string // 
	
	Display     string // 
	
}

// Represents a host crash dump
type host_crashdump struct {
	
	Uuid     string // 
	
	Host     string // 
	
	Timestamp     string // 
	
	Size     string // 
	
	Other_config     string // 
	
}

// Represents a patch stored on a server
type host_patch struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Version     string // 
	
	Host     string // 
	
	Applied     string // 
	
	Timestamp_applied     string // 
	
	Size     string // 
	
	Pool_patch     string // 
	
	Other_config     string // 
	
}

// The metrics associated with a host
type host_metrics struct {
	
	Uuid     string // 
	
	Memory_total     string // 
	
	Memory_free     string // 
	
	Live     string // 
	
	Last_updated     string // 
	
	Other_config     string // 
	
}

// A physical CPU
type host_cpu struct {
	
	Uuid     string // 
	
	Host     string // 
	
	Number     string // 
	
	Vendor     string // 
	
	Speed     string // 
	
	Modelname     string // 
	
	Family     string // 
	
	Model     string // 
	
	Stepping     string // 
	
	Flags     string // 
	
	Features     string // 
	
	Utilisation     string // 
	
	Other_config     string // 
	
}

// A virtual network
type network struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	VIFs     string // 
	
	PIFs     string // 
	
	MTU     string // 
	
	Other_config     string // 
	
	Bridge     string // 
	
	Blobs     string // 
	
	Tags     string // 
	
	Default_locking_mode     string // 
	
	Assigned_ips     string // 
	
}

// A virtual network interface
type VIF struct {
	
	Uuid     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	Device     string // 
	
	Network     string // 
	
	VM     string // 
	
	MAC     string // 
	
	MTU     string // 
	
	Other_config     string // 
	
	Currently_attached     string // 
	
	Status_code     string // 
	
	Status_detail     string // 
	
	Runtime_properties     string // 
	
	Qos_algorithm_type     string // 
	
	Qos_algorithm_params     string // 
	
	Qos_supported_algorithms     string // 
	
	Metrics     string // 
	
	MAC_autogenerated     string // 
	
	Locking_mode     string // 
	
	Ipv4_allowed     string // 
	
	Ipv6_allowed     string // 
	
}

// The metrics associated with a virtual network device
type VIF_metrics struct {
	
	Uuid     string // 
	
	Io_read_kbs     string // 
	
	Io_write_kbs     string // 
	
	Last_updated     string // 
	
	Other_config     string // 
	
}

// A physical network interface (note separate VLANs are represented as several PIFs)
type PIF struct {
	
	Uuid     string // 
	
	Device     string // 
	
	Network     string // 
	
	Host     string // 
	
	MAC     string // 
	
	MTU     string // 
	
	VLAN     string // 
	
	Metrics     string // 
	
	Physical     string // 
	
	Currently_attached     string // 
	
	Ip_configuration_mode     string // 
	
	IP     string // 
	
	Netmask     string // 
	
	Gateway     string // 
	
	DNS     string // 
	
	Bond_slave_of     string // 
	
	Bond_master_of     string // 
	
	VLAN_master_of     string // 
	
	VLAN_slave_of     string // 
	
	Management     string // 
	
	Other_config     string // 
	
	Disallow_unplug     string // 
	
	Tunnel_access_PIF_of     string // 
	
	Tunnel_transport_PIF_of     string // 
	
	Ipv6_configuration_mode     string // 
	
	IPv6     string // 
	
	Ipv6_gateway     string // 
	
	Primary_address_type     string // 
	
	Managed     string // 
	
	Properties     string // 
	
}

// The metrics associated with a physical network interface
type PIF_metrics struct {
	
	Uuid     string // 
	
	Io_read_kbs     string // 
	
	Io_write_kbs     string // 
	
	Carrier     string // 
	
	Vendor_id     string // 
	
	Vendor_name     string // 
	
	Device_id     string // 
	
	Device_name     string // 
	
	Speed     string // 
	
	Duplex     string // 
	
	Pci_bus_path     string // 
	
	Last_updated     string // 
	
	Other_config     string // 
	
}

// 
type Bond struct {
	
	Uuid     string // 
	
	Master     string // 
	
	Slaves     string // 
	
	Other_config     string // 
	
	Primary_slave     string // 
	
	Mode     string // 
	
	Properties     string // 
	
	Links_up     string // 
	
}

// A VLAN mux/demux
type VLAN struct {
	
	Uuid     string // 
	
	Tagged_PIF     string // 
	
	Untagged_PIF     string // 
	
	Tag     string // 
	
	Other_config     string // 
	
}

// A storage manager plugin
type SM struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Type     string // 
	
	Vendor     string // 
	
	Copyright     string // 
	
	Version     string // 
	
	Required_api_version     string // 
	
	Configuration     string // 
	
	Capabilities     string // 
	
	Features     string // 
	
	Other_config     string // 
	
	Driver_filename     string // 
	
}

// A storage repository
type SR struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	VDIs     string // 
	
	PBDs     string // 
	
	Virtual_allocation     string // 
	
	Physical_utilisation     string // 
	
	Physical_size     string // 
	
	Type     string // 
	
	Content_type     string // 
	
	Shared     string // 
	
	Other_config     string // 
	
	Tags     string // 
	
	Sm_config     string // 
	
	Blobs     string // 
	
	Local_cache_enabled     string // 
	
	Introduced_by     string // 
	
}

// A virtual disk image
type VDI struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	SR     string // 
	
	VBDs     string // 
	
	Crash_dumps     string // 
	
	Virtual_size     string // 
	
	Physical_utilisation     string // 
	
	Type     string // 
	
	Sharable     string // 
	
	Read_only     string // 
	
	Other_config     string // 
	
	Storage_lock     string // 
	
	Location     string // 
	
	Managed     string // 
	
	Missing     string // 
	
	Parent     string // 
	
	Xenstore_data     string // 
	
	Sm_config     string // 
	
	Is_a_snapshot     string // 
	
	Snapshot_of     string // 
	
	Snapshots     string // 
	
	Snapshot_time     string // 
	
	Tags     string // 
	
	Allow_caching     string // 
	
	On_boot     string // 
	
	Metadata_of_pool     string // 
	
	Metadata_latest     string // 
	
}

// A virtual block device
type VBD struct {
	
	Uuid     string // 
	
	Allowed_operations     string // 
	
	Current_operations     string // 
	
	VM     string // 
	
	VDI     string // 
	
	Device     string // 
	
	Userdevice     string // 
	
	Bootable     string // 
	
	Mode     string // 
	
	Type     string // 
	
	Unpluggable     string // 
	
	Storage_lock     string // 
	
	Empty     string // 
	
	Other_config     string // 
	
	Currently_attached     string // 
	
	Status_code     string // 
	
	Status_detail     string // 
	
	Runtime_properties     string // 
	
	Qos_algorithm_type     string // 
	
	Qos_algorithm_params     string // 
	
	Qos_supported_algorithms     string // 
	
	Metrics     string // 
	
}

// The metrics associated with a virtual block device
type VBD_metrics struct {
	
	Uuid     string // 
	
	Io_read_kbs     string // 
	
	Io_write_kbs     string // 
	
	Last_updated     string // 
	
	Other_config     string // 
	
}

// The physical block devices through which hosts access SRs
type PBD struct {
	
	Uuid     string // 
	
	Host     string // 
	
	SR     string // 
	
	Device_config     string // 
	
	Currently_attached     string // 
	
	Other_config     string // 
	
}

// A VM crashdump
type crashdump struct {
	
	Uuid     string // 
	
	VM     string // 
	
	VDI     string // 
	
	Other_config     string // 
	
}

// A virtual TPM device
type VTPM struct {
	
	Uuid     string // 
	
	VM     string // 
	
	Backend     string // 
	
}

// A console
type console struct {
	
	Uuid     string // 
	
	Protocol     string // 
	
	Location     string // 
	
	VM     string // 
	
	Other_config     string // 
	
}

// A user of the system
type user struct {
	
	Uuid     string // 
	
	Short_name     string // 
	
	Fullname     string // 
	
	Other_config     string // 
	
}

// Data sources for logging in RRDs
type data_source struct {
	
	Name_label     string // 
	
	Name_description     string // 
	
	Enabled     string // 
	
	Standard     string // 
	
	Units     string // 
	
	Min     string // 
	
	Max     string // 
	
	Value     string // 
	
}

// A placeholder for a binary blob
type blob struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	Size     string // 
	
	Public     string // 
	
	Last_updated     string // 
	
	Mime_type     string // 
	
}

// An message for the attention of the administrator
type message struct {
	
	Uuid     string // 
	
	Name     string // 
	
	Priority     string // 
	
	Cls     string // 
	
	Obj_uuid     string // 
	
	Timestamp     string // 
	
	Body     string // 
	
}

// A secret
type secret struct {
	
	Uuid     string // 
	
	Value     string // 
	
	Other_config     string // 
	
}

// A tunnel for network traffic
type tunnel struct {
	
	Uuid     string // 
	
	Access_PIF     string // 
	
	Transport_PIF     string // 
	
	Status     string // 
	
	Other_config     string // 
	
}

// A PCI device
type PCI struct {
	
	Uuid     string // 
	
	Class_name     string // 
	
	Vendor_name     string // 
	
	Device_name     string // 
	
	Host     string // 
	
	Pci_id     string // 
	
	Dependencies     string // 
	
	Other_config     string // 
	
	Subsystem_vendor_name     string // 
	
	Subsystem_device_name     string // 
	
}

// A physical GPU (pGPU)
type PGPU struct {
	
	Uuid     string // 
	
	PCI     string // 
	
	GPU_group     string // 
	
	Host     string // 
	
	Other_config     string // 
	
	Supported_VGPU_types     string // 
	
	Enabled_VGPU_types     string // 
	
	Resident_VGPUs     string // 
	
	Supported_VGPU_max_capacities     string // 
	
	Dom0_access     string // 
	
	Is_system_display_device     string // 
	
}

// A group of compatible GPUs across the resource pool
type GPU_group struct {
	
	Uuid     string // 
	
	Name_label     string // 
	
	Name_description     string // 
	
	PGPUs     string // 
	
	VGPUs     string // 
	
	GPU_types     string // 
	
	Other_config     string // 
	
	Allocation_algorithm     string // 
	
	Supported_VGPU_types     string // 
	
	Enabled_VGPU_types     string // 
	
}

// A virtual GPU (vGPU)
type VGPU struct {
	
	Uuid     string // 
	
	VM     string // 
	
	GPU_group     string // 
	
	Device     string // 
	
	Currently_attached     string // 
	
	Other_config     string // 
	
	Type     string // 
	
	Resident_on     string // 
	
}

// A type of virtual GPU
type VGPU_type struct {
	
	Uuid     string // 
	
	Vendor_name     string // 
	
	Model_name     string // 
	
	Framebuffer_size     string // 
	
	Max_heads     string // 
	
	Max_resolution_x     string // 
	
	Max_resolution_y     string // 
	
	Supported_on_PGPUs     string // 
	
	Enabled_on_PGPUs     string // 
	
	VGPUs     string // 
	
	Supported_on_GPU_groups     string // 
	
	Enabled_on_GPU_groups     string // 
	
}


func main() {
		fmt.Println("it compiles")
}
