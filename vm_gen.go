//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenAPI

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

type VMPowerState string

const (
  // VM is offline and not using any resources
	VMPowerStateHalted VMPowerState = "Halted"
  // All resources have been allocated but the VM itself is paused and its vCPUs are not running
	VMPowerStatePaused VMPowerState = "Paused"
  // Running
	VMPowerStateRunning VMPowerState = "Running"
  // VM state has been saved to disk and it is nolonger running. Note that disks remain in-use while the VM is suspended.
	VMPowerStateSuspended VMPowerState = "Suspended"
)

type OnNormalExit string

const (
  // destroy the VM state
	OnNormalExitDestroy OnNormalExit = "destroy"
  // restart the VM
	OnNormalExitRestart OnNormalExit = "restart"
)

type OnCrashBehaviour string

const (
  // destroy the VM state
	OnCrashBehaviourDestroy OnCrashBehaviour = "destroy"
  // record a coredump and then destroy the VM state
	OnCrashBehaviourCoredumpAndDestroy OnCrashBehaviour = "coredump_and_destroy"
  // restart the VM
	OnCrashBehaviourRestart OnCrashBehaviour = "restart"
  // record a coredump and then restart the VM
	OnCrashBehaviourCoredumpAndRestart OnCrashBehaviour = "coredump_and_restart"
  // leave the crashed VM paused
	OnCrashBehaviourPreserve OnCrashBehaviour = "preserve"
  // rename the crashed VM and start a new copy
	OnCrashBehaviourRenameRestart OnCrashBehaviour = "rename_restart"
)

type VMOperations string

const (
  // refers to the operation "snapshot"
	VMOperationsSnapshot VMOperations = "snapshot"
  // refers to the operation "clone"
	VMOperationsClone VMOperations = "clone"
  // refers to the operation "copy"
	VMOperationsCopy VMOperations = "copy"
  // refers to the operation "create_template"
	VMOperationsCreateTemplate VMOperations = "create_template"
  // refers to the operation "revert"
	VMOperationsRevert VMOperations = "revert"
  // refers to the operation "checkpoint"
	VMOperationsCheckpoint VMOperations = "checkpoint"
  // refers to the operation "snapshot_with_quiesce"
	VMOperationsSnapshotWithQuiesce VMOperations = "snapshot_with_quiesce"
  // refers to the operation "provision"
	VMOperationsProvision VMOperations = "provision"
  // refers to the operation "start"
	VMOperationsStart VMOperations = "start"
  // refers to the operation "start_on"
	VMOperationsStartOn VMOperations = "start_on"
  // refers to the operation "pause"
	VMOperationsPause VMOperations = "pause"
  // refers to the operation "unpause"
	VMOperationsUnpause VMOperations = "unpause"
  // refers to the operation "clean_shutdown"
	VMOperationsCleanShutdown VMOperations = "clean_shutdown"
  // refers to the operation "clean_reboot"
	VMOperationsCleanReboot VMOperations = "clean_reboot"
  // refers to the operation "hard_shutdown"
	VMOperationsHardShutdown VMOperations = "hard_shutdown"
  // refers to the operation "power_state_reset"
	VMOperationsPowerStateReset VMOperations = "power_state_reset"
  // refers to the operation "hard_reboot"
	VMOperationsHardReboot VMOperations = "hard_reboot"
  // refers to the operation "suspend"
	VMOperationsSuspend VMOperations = "suspend"
  // refers to the operation "csvm"
	VMOperationsCsvm VMOperations = "csvm"
  // refers to the operation "resume"
	VMOperationsResume VMOperations = "resume"
  // refers to the operation "resume_on"
	VMOperationsResumeOn VMOperations = "resume_on"
  // refers to the operation "pool_migrate"
	VMOperationsPoolMigrate VMOperations = "pool_migrate"
  // refers to the operation "migrate_send"
	VMOperationsMigrateSend VMOperations = "migrate_send"
  // refers to the operation "get_boot_record"
	VMOperationsGetBootRecord VMOperations = "get_boot_record"
  // refers to the operation "send_sysrq"
	VMOperationsSendSysrq VMOperations = "send_sysrq"
  // refers to the operation "send_trigger"
	VMOperationsSendTrigger VMOperations = "send_trigger"
  // refers to the operation "query_services"
	VMOperationsQueryServices VMOperations = "query_services"
  // refers to the operation "shutdown"
	VMOperationsShutdown VMOperations = "shutdown"
  // refers to the operation "call_plugin"
	VMOperationsCallPlugin VMOperations = "call_plugin"
  // Changing the memory settings
	VMOperationsChangingMemoryLive VMOperations = "changing_memory_live"
  // Waiting for the memory settings to change
	VMOperationsAwaitingMemoryLive VMOperations = "awaiting_memory_live"
  // Changing the memory dynamic range
	VMOperationsChangingDynamicRange VMOperations = "changing_dynamic_range"
  // Changing the memory static range
	VMOperationsChangingStaticRange VMOperations = "changing_static_range"
  // Changing the memory limits
	VMOperationsChangingMemoryLimits VMOperations = "changing_memory_limits"
  // Changing the shadow memory for a halted VM.
	VMOperationsChangingShadowMemory VMOperations = "changing_shadow_memory"
  // Changing the shadow memory for a running VM.
	VMOperationsChangingShadowMemoryLive VMOperations = "changing_shadow_memory_live"
  // Changing VCPU settings for a halted VM.
	VMOperationsChangingVCPUs VMOperations = "changing_VCPUs"
  // Changing VCPU settings for a running VM.
	VMOperationsChangingVCPUsLive VMOperations = "changing_VCPUs_live"
  // 
	VMOperationsAssertOperationValid VMOperations = "assert_operation_valid"
  // Add, remove, query or list data sources
	VMOperationsDataSourceOp VMOperations = "data_source_op"
  // 
	VMOperationsUpdateAllowedOperations VMOperations = "update_allowed_operations"
  // Turning this VM into a template
	VMOperationsMakeIntoTemplate VMOperations = "make_into_template"
  // importing a VM from a network stream
	VMOperationsImport VMOperations = "import"
  // exporting a VM to a network stream
	VMOperationsExport VMOperations = "export"
  // exporting VM metadata to a network stream
	VMOperationsMetadataExport VMOperations = "metadata_export"
  // Reverting the VM to a previous snapshotted state
	VMOperationsReverting VMOperations = "reverting"
  // refers to the act of uninstalling the VM
	VMOperationsDestroy VMOperations = "destroy"
)

type VMRecord struct {
  // Unique identifier/object reference
	UUID string
  // list of the operations allowed in this state. This list is advisory only and the server state may have changed by the time this field is read by a client.
	AllowedOperations []VMOperations
  // links each of the running tasks using this object (by reference) to a current_operation enum which describes the nature of the task.
	CurrentOperations map[string]VMOperations
  // Current power state of the machine
	PowerState VMPowerState
  // a human-readable name
	NameLabel string
  // a notes field containing human-readable description
	NameDescription string
  // a user version number for this machine
	UserVersion int
  // true if this is a template. Template VMs can never be started, they are used only for cloning other VMs
	IsATemplate bool
  // The VDI that a suspend image is stored on. (Only has meaning if VM is currently suspended)
	SuspendVDI VDIRef
  // the host the VM is currently resident on
	ResidentOn HostRef
  // a host which the VM has some affinity for (or NULL). This is used as a hint to the start call when it decides where to run the VM. Implementations are free to ignore this field.
	Affinity HostRef
  // Virtualization memory overhead (bytes).
	MemoryOverhead int
  // Dynamically-set memory target (bytes). The value of this field indicates the current target for memory available to this VM.
	MemoryTarget int
  // Statically-set (i.e. absolute) maximum (bytes). The value of this field at VM start time acts as a hard limit of the amount of memory a guest can use. New values only take effect on reboot.
	MemoryStaticMax int
  // Dynamic maximum (bytes)
	MemoryDynamicMax int
  // Dynamic minimum (bytes)
	MemoryDynamicMin int
  // Statically-set (i.e. absolute) mininum (bytes). The value of this field indicates the least amount of memory this VM can boot with without crashing.
	MemoryStaticMin int
  // configuration parameters for the selected VCPU policy
	VCPUsParams map[string]string
  // Max number of VCPUs
	VCPUsMax int
  // Boot number of VCPUs
	VCPUsAtStartup int
  // action to take after the guest has shutdown itself
	ActionsAfterShutdown OnNormalExit
  // action to take after the guest has rebooted itself
	ActionsAfterReboot OnNormalExit
  // action to take if the guest crashes
	ActionsAfterCrash OnCrashBehaviour
  // virtual console devices
	Consoles []ConsoleRef
  // virtual network interfaces
	VIFs []VIFRef
  // virtual block devices
	VBDs []VBDRef
  // crash dumps associated with this VM
	CrashDumps []CrashdumpRef
  // virtual TPMs
	VTPMs []VTPMRef
  // name of or path to bootloader
	PVBootloader string
  // path to the kernel
	PVKernel string
  // path to the initrd
	PVRamdisk string
  // kernel command-line arguments
	PVArgs string
  // miscellaneous arguments for the bootloader
	PVBootloaderArgs string
  // to make Zurich guests boot
	PVLegacyArgs string
  // HVM boot policy
	HVMBootPolicy string
  // HVM boot params
	HVMBootParams map[string]string
  // multiplier applied to the amount of shadow that will be made available to the guest
	HVMShadowMultiplier float64
  // platform-specific configuration
	Platform map[string]string
  // PCI bus path for pass-through devices
	PCIBus string
  // additional configuration
	OtherConfig map[string]string
  // domain ID (if available, -1 otherwise)
	Domid int
  // Domain architecture (if available, null string otherwise)
	Domarch string
  // describes the CPU flags on which the VM was last booted
	LastBootCPUFlags map[string]string
  // true if this is a control domain (domain 0 or a driver domain)
	IsControlDomain bool
  // metrics associated with this VM
	Metrics VMMetricsRef
  // metrics associated with the running guest
	GuestMetrics VMGuestMetricsRef
  // marshalled value containing VM record at time of last boot, updated dynamically to reflect the runtime state of the domain
	LastBootedRecord string
  // An XML specification of recommended values and ranges for properties of this VM
	Recommendations string
  // data to be inserted into the xenstore tree (/local/domain/<domid>/vm-data) after the VM is created.
	XenstoreData map[string]string
  // if true then the system will attempt to keep the VM running as much as possible.
	HaAlwaysRun bool
  // has possible values: "best-effort" meaning "try to restart this VM if possible but don't consider the Pool to be overcommitted if this is not possible"; "restart" meaning "this VM should be restarted"; "" meaning "do not try to restart this VM"
	HaRestartPriority string
  // true if this is a snapshot. Snapshotted VMs can never be started, they are used only for cloning other VMs
	IsASnapshot bool
  // Ref pointing to the VM this snapshot is of.
	SnapshotOf VMRef
  // List pointing to all the VM snapshots.
	Snapshots []VMRef
  // Date/time when this snapshot was created.
	SnapshotTime time.Time
  // Transportable ID of the snapshot VM
	TransportableSnapshotID string
  // Binary blobs associated with this VM
	Blobs map[string]BlobRef
  // user-specified tags for categorization purposes
	Tags []string
  // List of operations which have been explicitly blocked and an error code
	BlockedOperations map[VMOperations]string
  // Human-readable information concerning this snapshot
	SnapshotInfo map[string]string
  // Encoded information about the VM's metadata this is a snapshot of
	SnapshotMetadata string
  // Ref pointing to the parent of this VM
	Parent VMRef
  // List pointing to all the children of this VM
	Children []VMRef
  // BIOS strings
	BiosStrings map[string]string
  // Ref pointing to a protection policy for this VM
	ProtectionPolicy VMPPRef
  // true if this snapshot was created by the protection policy
	IsSnapshotFromVmpp bool
  // the appliance to which this VM belongs
	Appliance VMApplianceRef
  // The delay to wait before proceeding to the next order in the startup sequence (seconds)
	StartDelay int
  // The delay to wait before proceeding to the next order in the shutdown sequence (seconds)
	ShutdownDelay int
  // The point in the startup or shutdown sequence at which this VM will be started
	Order int
  // Virtual GPUs
	VGPUs []VGPURef
  // Currently passed-through PCI devices
	AttachedPCIs []PCIRef
  // The SR on which a suspend image is stored
	SuspendSR SRRef
  // The number of times this VM has been recovered
	Version int
  // Generation ID of the VM
	GenerationID string
  // The host virtual hardware platform version the VM can run on
	HardwarePlatformVersion int
  // True if the Windows Update feature is enabled on the VM; false otherwise
	AutoUpdateDrivers bool
}

type VMRef string

// A virtual machine (or 'guest').
type VMClass struct {
	client *Client
}

// A virtual machine (or 'guest').
type AsyncVMClass struct {
	client *Client
}

// Return a map of VM references to VM records for all VMs known to the system.
func (_class VMClass) GetAllRecords(sessionID SessionRef) (_retval map[VMRef]VMRecord, _err error) {
	_method := "VM.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToVMRecordMapToGo(_method + " -> ", _result.Value)
	return
}

// Return a list of all the VMs known to the system.
func (_class VMClass) GetAll(sessionID SessionRef) (_retval []VMRef, _err error) {
	_method := "VM.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Start the 'xenprep' process on the VM; the process will remove any tools and drivers for XenServer and then set auto update drivers true.
func (_class VMClass) XenprepStart(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.xenprep_start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Import an XVA from a URI
func (_class VMClass) Import(sessionID SessionRef, url string, sr SRRef, fullRestore bool, force bool) (_retval []VMRef, _err error) {
	_method := "VM.import"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_urlArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "url"), url)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_fullRestoreArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "full_restore"), fullRestore)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _urlArg, _srArg, _fullRestoreArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Check if PV auto update can be set on Windows vm
func (_class VMClass) AssertCanSetAutoUpdateDrivers(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.assert_can_set_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Enable or disable PV auto update on Windows vm
func (_class VMClass) SetAutoUpdateDrivers(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.set_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Call a XenAPI plugin on this vm
func (_class VMClass) CallPlugin(sessionID SessionRef, vm VMRef, plugin string, fn string, args map[string]string) (_retval string, _err error) {
	_method := "VM.call_plugin"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_pluginArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "plugin"), plugin)
	if _err != nil {
		return
	}
	_fnArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "fn"), fn)
	if _err != nil {
		return
	}
	_argsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _pluginArg, _fnArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Query the system services advertised by this VM and register them. This can only be applied to a system domain.
func (_class VMClass) QueryServices(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.query_services"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Assign this VM to an appliance.
func (_class VMClass) SetAppliance(sessionID SessionRef, self VMRef, value VMApplianceRef) (_err error) {
	_method := "VM.set_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Import using a conversion service.
func (_class VMClass) ImportConvert(sessionID SessionRef, atype string, username string, password string, sr SRRef, remoteConfig map[string]string) (_err error) {
	_method := "VM.import_convert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_usernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "username"), username)
	if _err != nil {
		return
	}
	_passwordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "password"), password)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_remoteConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "remote_config"), remoteConfig)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _atypeArg, _usernameArg, _passwordArg, _srArg, _remoteConfigArg)
	return
}

// Recover the VM
func (_class VMClass) Recover(sessionID SessionRef, self VMRef, sessionTo SessionRef, force bool) (_err error) {
	_method := "VM.recover"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg, _forceArg)
	return
}

// List all the SR's that are required for the VM to be recovered
func (_class VMClass) GetSRsRequiredForRecovery(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_retval []SRRef, _err error) {
	_method := "VM.get_SRs_required_for_recovery"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Assert whether all SRs required to recover this VM are available.
//
// Errors:
//  VM_IS_PART_OF_AN_APPLIANCE - This operation is not allowed as the VM is part of an appliance.
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
func (_class VMClass) AssertCanBeRecovered(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_err error) {
	_method := "VM.assert_can_be_recovered"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	return
}

// Set this VM's suspend VDI, which must be indentical to its current one
func (_class VMClass) SetSuspendVDI(sessionID SessionRef, self VMRef, value VDIRef) (_err error) {
	_method := "VM.set_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set this VM's boot order
func (_class VMClass) SetOrder(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set this VM's shutdown delay in seconds
func (_class VMClass) SetShutdownDelay(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set this VM's start delay in seconds
func (_class VMClass) SetStartDelay(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the value of the protection_policy field
func (_class VMClass) SetProtectionPolicy(sessionID SessionRef, self VMRef, value VMPPRef) (_err error) {
	_method := "VM.set_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Copy the BIOS strings from the given host to this VM
func (_class VMClass) CopyBiosStrings(sessionID SessionRef, vm VMRef, host HostRef) (_err error) {
	_method := "VM.copy_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg)
	return
}

// Returns mapping of hosts to ratings, indicating the suitability of starting the VM at that location according to wlb. Rating is replaced with an error if the VM cannot boot there.
func (_class VMClass) RetrieveWlbRecommendations(sessionID SessionRef, vm VMRef) (_retval map[HostRef][]string, _err error) {
	_method := "VM.retrieve_wlb_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToStringSetMapToGo(_method + " -> ", _result.Value)
	return
}

// Returns an error if the VM is not considered agile e.g. because it is tied to a resource local to a host
func (_class VMClass) AssertAgile(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.assert_agile"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create a placeholder for a named binary blob of data that is associated with this VM
func (_class VMClass) CreateNewBlob(sessionID SessionRef, vm VMRef, name string, mimeType string, public bool) (_retval BlobRef, _err error) {
	_method := "VM.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_mimeTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "mime_type"), mimeType)
	if _err != nil {
		return
	}
	_publicArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "public"), public)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBlobRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns an error if the VM could not boot on this host for some reason
//
// Errors:
//  HOST_NOT_ENOUGH_FREE_MEMORY - Not enough host memory is available to perform this operation
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
//  VM_HOST_INCOMPATIBLE_VERSION - This VM operation cannot be performed on an older-versioned host during an upgrade.
//  VM_HOST_INCOMPATIBLE_VIRTUAL_HARDWARE_PLATFORM_VERSION - You attempted to run a VM on a host that cannot provide the VM's required Virtual Hardware Platform version.
func (_class VMClass) AssertCanBootHere(sessionID SessionRef, self VMRef, host HostRef) (_err error) {
	_method := "VM.assert_can_boot_here"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _hostArg)
	return
}

// Return the list of hosts on which this VM may run.
func (_class VMClass) GetPossibleHosts(sessionID SessionRef, vm VMRef) (_retval []HostRef, _err error) {
	_method := "VM.get_possible_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Returns a list of the allowed values that a VIF device field can take
func (_class VMClass) GetAllowedVIFDevices(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	_method := "VM.get_allowed_VIF_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// Returns a list of the allowed values that a VBD device field can take
func (_class VMClass) GetAllowedVBDDevices(sessionID SessionRef, vm VMRef) (_retval []string, _err error) {
	_method := "VM.get_allowed_VBD_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// Recomputes the list of acceptable operations
func (_class VMClass) UpdateAllowedOperations(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.update_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Check to see whether this operation is acceptable in the current state of the system, raising an error if the operation is invalid for some reason
func (_class VMClass) AssertOperationValid(sessionID SessionRef, self VMRef, op VMOperations) (_err error) {
	_method := "VM.assert_operation_valid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_opArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "op"), op)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _opArg)
	return
}

// Forget the recorded statistics related to the specified data source
func (_class VMClass) ForgetDataSourceArchives(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	_method := "VM.forget_data_source_archives"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	return
}

// Query the latest value of the specified data source
func (_class VMClass) QueryDataSource(sessionID SessionRef, self VMRef, dataSource string) (_retval float64, _err error) {
	_method := "VM.query_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}

// Start recording the specified data source
func (_class VMClass) RecordDataSource(sessionID SessionRef, self VMRef, dataSource string) (_err error) {
	_method := "VM.record_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	return
}

// 
func (_class VMClass) GetDataSources(sessionID SessionRef, self VMRef) (_retval []DataSourceRecord, _err error) {
	_method := "VM.get_data_sources"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertDataSourceRecordSetToGo(_method + " -> ", _result.Value)
	return
}

// Returns a record describing the VM's dynamic state, initialised when the VM boots and updated to reflect runtime configuration changes e.g. CPU hotplug
func (_class VMClass) GetBootRecord(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	_method := "VM.get_boot_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRecordToGo(_method + " -> ", _result.Value)
	return
}

// Assert whether a VM can be migrated to the specified destination.
func (_class VMClass) AssertCanMigrate(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string) (_err error) {
	_method := "VM.assert_can_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg)
	return
}

// Migrate the VM to another host.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class VMClass) MigrateSend(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string) (_retval VMRef, _err error) {
	_method := "VM.migrate_send"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns the maximum amount of guest memory which will fit, together with overheads, in the supplied amount of physical memory. If 'exact' is true then an exact calculation is performed using the VM's current settings. If 'exact' is false then a more conservative approximation is used
func (_class VMClass) MaximiseMemory(sessionID SessionRef, self VMRef, total int, approximate bool) (_retval int, _err error) {
	_method := "VM.maximise_memory"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_totalArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "total"), total)
	if _err != nil {
		return
	}
	_approximateArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "approximate"), approximate)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _totalArg, _approximateArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Send the named trigger to this VM.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class VMClass) SendTrigger(sessionID SessionRef, vm VMRef, trigger string) (_err error) {
	_method := "VM.send_trigger"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_triggerArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "trigger"), trigger)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _triggerArg)
	return
}

// Send the given key as a sysrq to this VM.  The key is specified as a single character (a String of length 1).  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class VMClass) SendSysrq(sessionID SessionRef, vm VMRef, key string) (_err error) {
	_method := "VM.send_sysrq"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _keyArg)
	return
}

// Set the number of startup VCPUs for a halted VM
func (_class VMClass) SetVCPUsAtStartup(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the maximum number of VCPUs for a halted VM
func (_class VMClass) SetVCPUsMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the shadow memory multiplier on a running VM
func (_class VMClass) SetShadowMultiplierLive(sessionID SessionRef, self VMRef, multiplier float64) (_err error) {
	_method := "VM.set_shadow_multiplier_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_multiplierArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "multiplier"), multiplier)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _multiplierArg)
	return
}

// Set the shadow memory multiplier on a halted VM
func (_class VMClass) SetHVMShadowMultiplier(sessionID SessionRef, self VMRef, value float64) (_err error) {
	_method := "VM.set_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Return true if the VM is currently 'co-operative' i.e. is expected to reach a balloon target and actually has done
func (_class VMClass) GetCooperative(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_cooperative"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// Wait for a running VM to reach its current memory target
func (_class VMClass) WaitMemoryTargetLive(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.wait_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Set the memory target for a running VM
func (_class VMClass) SetMemoryTargetLive(sessionID SessionRef, self VMRef, target int) (_err error) {
	_method := "VM.set_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_targetArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "target"), target)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _targetArg)
	return
}

// Set the memory limits of this VM.
func (_class VMClass) SetMemoryLimits(sessionID SessionRef, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (_err error) {
	_method := "VM.set_memory_limits"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_staticMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_min"), staticMin)
	if _err != nil {
		return
	}
	_staticMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_max"), staticMax)
	if _err != nil {
		return
	}
	_dynamicMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_min"), dynamicMin)
	if _err != nil {
		return
	}
	_dynamicMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_max"), dynamicMax)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _staticMinArg, _staticMaxArg, _dynamicMinArg, _dynamicMaxArg)
	return
}

// Set the static (ie boot-time) range of virtual memory that the VM is allowed to use.
func (_class VMClass) SetMemoryStaticRange(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	_method := "VM.set_memory_static_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	return
}

// Set the value of the memory_static_min field
func (_class VMClass) SetMemoryStaticMin(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the value of the memory_static_max field
//
// Errors:
//  HA_OPERATION_WOULD_BREAK_FAILOVER_PLAN - This operation cannot be performed because it would invalidate VM failover planning such that the system would be unable to guarantee to restart protected VMs after a Host failure.
func (_class VMClass) SetMemoryStaticMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the minimum and maximum amounts of physical memory the VM is allowed to use.
func (_class VMClass) SetMemoryDynamicRange(sessionID SessionRef, self VMRef, min int, max int) (_err error) {
	_method := "VM.set_memory_dynamic_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	return
}

// Set the value of the memory_dynamic_min field
func (_class VMClass) SetMemoryDynamicMin(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the value of the memory_dynamic_max field
func (_class VMClass) SetMemoryDynamicMax(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Computes the virtualization memory overhead of a VM.
func (_class VMClass) ComputeMemoryOverhead(sessionID SessionRef, vm VMRef) (_retval int, _err error) {
	_method := "VM.compute_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the ha_always_run
func (_class VMClass) SetHaAlwaysRun(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.set_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the value of the ha_restart_priority field
func (_class VMClass) SetHaRestartPriority(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Add the given key-value pair to VM.VCPUs_params, and apply that value on the running VM
func (_class VMClass) AddToVCPUsParamsLive(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_VCPUs_params_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the number of VCPUs for a running VM
func (_class VMClass) SetVCPUsNumberLive(sessionID SessionRef, self VMRef, nvcpu int) (_err error) {
	_method := "VM.set_VCPUs_number_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nvcpuArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "nvcpu"), nvcpu)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _nvcpuArg)
	return
}

// Migrate a VM to another Host.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_MIGRATE_FAILED - An error occurred during the migration process.
func (_class VMClass) PoolMigrate(sessionID SessionRef, vm VMRef, host HostRef, options map[string]string) (_err error) {
	_method := "VM.pool_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _optionsArg)
	return
}

// Awaken the specified VM and resume it on a particular Host.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) ResumeOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	_method := "VM.resume_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	return
}

// Awaken the specified VM and resume it.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Resume(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	_method := "VM.resume"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	return
}

// Suspend the specified VM to disk.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Suspend(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.suspend"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Stop executing the specified VM without attempting a clean shutdown and immediately restart the VM.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) HardReboot(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.hard_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Reset the power-state of the VM to halted in the database only. (Used to recover from slave failures in pooling scenarios by resetting the power-states of VMs running on dead slaves to halted.) This is a potentially dangerous operation; use with care.
func (_class VMClass) PowerStateReset(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.power_state_reset"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Stop executing the specified VM without attempting a clean shutdown.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) HardShutdown(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.hard_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Attempt to cleanly shutdown the specified VM (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) CleanReboot(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.clean_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Attempts to first clean shutdown a VM and if it should fail then perform a hard shutdown on it.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Shutdown(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Attempt to cleanly shutdown the specified VM. (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) CleanShutdown(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.clean_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Resume the specified VM. This can only be called when the specified VM is in the Paused state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Unpause(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.unpause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Pause the specified VM. This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class VMClass) Pause(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.pause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Start the specified VM on a particular host.  This function can only be called with the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  BOOTLOADER_FAILED - The bootloader returned an error
//  UNKNOWN_BOOTLOADER - The requested bootloader is unknown
func (_class VMClass) StartOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_err error) {
	_method := "VM.start_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	return
}

// Start the specified VM.  This function can only be called with the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  VM_HVM_REQUIRED - HVM is required for this operation
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  BOOTLOADER_FAILED - The bootloader returned an error
//  UNKNOWN_BOOTLOADER - The requested bootloader is unknown
//  NO_HOSTS_AVAILABLE - There were no hosts available to complete the specified operation.
//  LICENCE_RESTRICTION - This operation is not allowed under your license.  Please contact your support representative.
func (_class VMClass) Start(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_err error) {
	_method := "VM.start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	return
}

// Inspects the disk configuration contained within the VM's other_config, creates VDIs and VBDs and then executes any applicable post-install script.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class VMClass) Provision(sessionID SessionRef, vm VMRef) (_err error) {
	_method := "VM.provision"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _vmArg)
	return
}

// Checkpoints the specified VM, making a new VM. Checkpoint automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write) and saves the memory image as well.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_CHECKPOINT_SUSPEND_FAILED - An error occured while saving the memory image of the specified virtual machine
//  VM_CHECKPOINT_RESUME_FAILED - An error occured while restoring the memory image of the specified virtual machine
func (_class VMClass) Checkpoint(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.checkpoint"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Reverts the specified VM to a previous state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  VM_REVERT_FAILED - An error occured while reverting the specified virtual machine to the specified snapshot
func (_class VMClass) Revert(sessionID SessionRef, snapshot VMRef) (_err error) {
	_method := "VM.revert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_snapshotArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot"), snapshot)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _snapshotArg)
	return
}

// Copied the specified VM, making a new VM. Unlike clone, copy does not exploits the capabilities of the underlying storage repository in which the VM's disk images are stored. Instead, copy guarantees that the disk images of the newly created VM will be 'full disks' - i.e. not part of a CoW chain.  This function can only be called when the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class VMClass) Copy(sessionID SessionRef, vm VMRef, newName string, sr SRRef) (_retval VMRef, _err error) {
	_method := "VM.copy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg, _srArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Clones the specified VM, making a new VM. Clone automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).   This function can only be called when the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class VMClass) Clone(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.clone"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Snapshots the specified VM with quiesce, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_SNAPSHOT_WITH_QUIESCE_FAILED - The quiesced-snapshot operation failed for an unexpected reason
//  VM_SNAPSHOT_WITH_QUIESCE_TIMEOUT - The VSS plug-in has timed out
//  VM_SNAPSHOT_WITH_QUIESCE_PLUGIN_DEOS_NOT_RESPOND - The VSS plug-in cannot be contacted
//  VM_SNAPSHOT_WITH_QUIESCE_NOT_SUPPORTED - The VSS plug-in is not installed on this virtual machine
func (_class VMClass) SnapshotWithQuiesce(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.snapshot_with_quiesce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Snapshots the specified VM, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class VMClass) Snapshot(sessionID SessionRef, vm VMRef, newName string) (_retval VMRef, _err error) {
	_method := "VM.snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the hardware_platform_version field of the given VM.
func (_class VMClass) SetHardwarePlatformVersion(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the suspend_SR field of the given VM.
func (_class VMClass) SetSuspendSR(sessionID SessionRef, self VMRef, value SRRef) (_err error) {
	_method := "VM.set_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Remove the given key and its corresponding value from the blocked_operations field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations) (_err error) {
	_method := "VM.remove_from_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// Add the given key-value pair to the blocked_operations field of the given VM.
func (_class VMClass) AddToBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations, value string) (_err error) {
	_method := "VM.add_to_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the blocked_operations field of the given VM.
func (_class VMClass) SetBlockedOperations(sessionID SessionRef, self VMRef, value map[VMOperations]string) (_err error) {
	_method := "VM.set_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVMOperationsToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Remove the given value from the tags field of the given VM.  If the value is not in that Set, then do nothing.
func (_class VMClass) RemoveTags(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Add the given value to the tags field of the given VM.  If the value is already in that Set, then do nothing.
func (_class VMClass) AddTags(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the tags field of the given VM.
func (_class VMClass) SetTags(sessionID SessionRef, self VMRef, value []string) (_err error) {
	_method := "VM.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Remove the given key and its corresponding value from the xenstore_data field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromXenstoreData(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// Add the given key-value pair to the xenstore_data field of the given VM.
func (_class VMClass) AddToXenstoreData(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the xenstore_data field of the given VM.
func (_class VMClass) SetXenstoreData(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the recommendations field of the given VM.
func (_class VMClass) SetRecommendations(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Remove the given key and its corresponding value from the other_config field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromOtherConfig(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// Add the given key-value pair to the other_config field of the given VM.
func (_class VMClass) AddToOtherConfig(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the other_config field of the given VM.
func (_class VMClass) SetOtherConfig(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the PCI_bus field of the given VM.
func (_class VMClass) SetPCIBus(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Remove the given key and its corresponding value from the platform field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromPlatform(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// Add the given key-value pair to the platform field of the given VM.
func (_class VMClass) AddToPlatform(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the platform field of the given VM.
func (_class VMClass) SetPlatform(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Remove the given key and its corresponding value from the HVM/boot_params field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromHVMBootParams(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// Add the given key-value pair to the HVM/boot_params field of the given VM.
func (_class VMClass) AddToHVMBootParams(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the HVM/boot_params field of the given VM.
func (_class VMClass) SetHVMBootParams(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the HVM/boot_policy field of the given VM.
func (_class VMClass) SetHVMBootPolicy(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the PV/legacy_args field of the given VM.
func (_class VMClass) SetPVLegacyArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the PV/bootloader_args field of the given VM.
func (_class VMClass) SetPVBootloaderArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the PV/args field of the given VM.
func (_class VMClass) SetPVArgs(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the PV/ramdisk field of the given VM.
func (_class VMClass) SetPVRamdisk(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the PV/kernel field of the given VM.
func (_class VMClass) SetPVKernel(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the PV/bootloader field of the given VM.
func (_class VMClass) SetPVBootloader(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the actions/after_crash field of the given VM.
func (_class VMClass) SetActionsAfterCrash(sessionID SessionRef, self VMRef, value OnCrashBehaviour) (_err error) {
	_method := "VM.set_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnCrashBehaviourToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the actions/after_reboot field of the given VM.
func (_class VMClass) SetActionsAfterReboot(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	_method := "VM.set_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the actions/after_shutdown field of the given VM.
func (_class VMClass) SetActionsAfterShutdown(sessionID SessionRef, self VMRef, value OnNormalExit) (_err error) {
	_method := "VM.set_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Remove the given key and its corresponding value from the VCPUs/params field of the given VM.  If the key is not in that Map, then do nothing.
func (_class VMClass) RemoveFromVCPUsParams(sessionID SessionRef, self VMRef, key string) (_err error) {
	_method := "VM.remove_from_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	return
}

// Add the given key-value pair to the VCPUs/params field of the given VM.
func (_class VMClass) AddToVCPUsParams(sessionID SessionRef, self VMRef, key string, value string) (_err error) {
	_method := "VM.add_to_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	return
}

// Set the VCPUs/params field of the given VM.
func (_class VMClass) SetVCPUsParams(sessionID SessionRef, self VMRef, value map[string]string) (_err error) {
	_method := "VM.set_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the affinity field of the given VM.
func (_class VMClass) SetAffinity(sessionID SessionRef, self VMRef, value HostRef) (_err error) {
	_method := "VM.set_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the is_a_template field of the given VM.
func (_class VMClass) SetIsATemplate(sessionID SessionRef, self VMRef, value bool) (_err error) {
	_method := "VM.set_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the user_version field of the given VM.
func (_class VMClass) SetUserVersion(sessionID SessionRef, self VMRef, value int) (_err error) {
	_method := "VM.set_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the name/description field of the given VM.
func (_class VMClass) SetNameDescription(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Set the name/label field of the given VM.
func (_class VMClass) SetNameLabel(sessionID SessionRef, self VMRef, value string) (_err error) {
	_method := "VM.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	return
}

// Get the auto_update_drivers field of the given VM.
func (_class VMClass) GetAutoUpdateDrivers(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// Get the hardware_platform_version field of the given VM.
func (_class VMClass) GetHardwarePlatformVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the generation_id field of the given VM.
func (_class VMClass) GetGenerationID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_generation_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the version field of the given VM.
func (_class VMClass) GetVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the suspend_SR field of the given VM.
func (_class VMClass) GetSuspendSR(sessionID SessionRef, self VMRef) (_retval SRRef, _err error) {
	_method := "VM.get_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertSRRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the attached_PCIs field of the given VM.
func (_class VMClass) GetAttachedPCIs(sessionID SessionRef, self VMRef) (_retval []PCIRef, _err error) {
	_method := "VM.get_attached_PCIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertPCIRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the VGPUs field of the given VM.
func (_class VMClass) GetVGPUs(sessionID SessionRef, self VMRef) (_retval []VGPURef, _err error) {
	_method := "VM.get_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVGPURefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the order field of the given VM.
func (_class VMClass) GetOrder(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the shutdown_delay field of the given VM.
func (_class VMClass) GetShutdownDelay(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the start_delay field of the given VM.
func (_class VMClass) GetStartDelay(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the appliance field of the given VM.
func (_class VMClass) GetAppliance(sessionID SessionRef, self VMRef) (_retval VMApplianceRef, _err error) {
	_method := "VM.get_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMApplianceRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_snapshot_from_vmpp field of the given VM.
func (_class VMClass) GetIsSnapshotFromVmpp(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_snapshot_from_vmpp"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// Get the protection_policy field of the given VM.
func (_class VMClass) GetProtectionPolicy(sessionID SessionRef, self VMRef) (_retval VMPPRef, _err error) {
	_method := "VM.get_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMPPRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the bios_strings field of the given VM.
func (_class VMClass) GetBiosStrings(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the children field of the given VM.
func (_class VMClass) GetChildren(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	_method := "VM.get_children"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the parent field of the given VM.
func (_class VMClass) GetParent(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	_method := "VM.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_metadata field of the given VM.
func (_class VMClass) GetSnapshotMetadata(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_snapshot_metadata"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_info field of the given VM.
func (_class VMClass) GetSnapshotInfo(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_snapshot_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the blocked_operations field of the given VM.
func (_class VMClass) GetBlockedOperations(sessionID SessionRef, self VMRef) (_retval map[VMOperations]string, _err error) {
	_method := "VM.get_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMOperationsToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the tags field of the given VM.
func (_class VMClass) GetTags(sessionID SessionRef, self VMRef) (_retval []string, _err error) {
	_method := "VM.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the blobs field of the given VM.
func (_class VMClass) GetBlobs(sessionID SessionRef, self VMRef) (_retval map[string]BlobRef, _err error) {
	_method := "VM.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToBlobRefMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the transportable_snapshot_id field of the given VM.
func (_class VMClass) GetTransportableSnapshotID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_transportable_snapshot_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_time field of the given VM.
func (_class VMClass) GetSnapshotTime(sessionID SessionRef, self VMRef) (_retval time.Time, _err error) {
	_method := "VM.get_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTimeToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshots field of the given VM.
func (_class VMClass) GetSnapshots(sessionID SessionRef, self VMRef) (_retval []VMRef, _err error) {
	_method := "VM.get_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_of field of the given VM.
func (_class VMClass) GetSnapshotOf(sessionID SessionRef, self VMRef) (_retval VMRef, _err error) {
	_method := "VM.get_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_a_snapshot field of the given VM.
func (_class VMClass) GetIsASnapshot(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// Get the ha_restart_priority field of the given VM.
func (_class VMClass) GetHaRestartPriority(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the ha_always_run field of the given VM.
func (_class VMClass) GetHaAlwaysRun(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// Get the xenstore_data field of the given VM.
func (_class VMClass) GetXenstoreData(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the recommendations field of the given VM.
func (_class VMClass) GetRecommendations(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the last_booted_record field of the given VM.
func (_class VMClass) GetLastBootedRecord(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_last_booted_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the guest_metrics field of the given VM.
func (_class VMClass) GetGuestMetrics(sessionID SessionRef, self VMRef) (_retval VMGuestMetricsRef, _err error) {
	_method := "VM.get_guest_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMGuestMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the metrics field of the given VM.
func (_class VMClass) GetMetrics(sessionID SessionRef, self VMRef) (_retval VMMetricsRef, _err error) {
	_method := "VM.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMMetricsRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_control_domain field of the given VM.
func (_class VMClass) GetIsControlDomain(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_control_domain"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// Get the last_boot_CPU_flags field of the given VM.
func (_class VMClass) GetLastBootCPUFlags(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_last_boot_CPU_flags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the domarch field of the given VM.
func (_class VMClass) GetDomarch(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_domarch"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the domid field of the given VM.
func (_class VMClass) GetDomid(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_domid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the other_config field of the given VM.
func (_class VMClass) GetOtherConfig(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the PCI_bus field of the given VM.
func (_class VMClass) GetPCIBus(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the platform field of the given VM.
func (_class VMClass) GetPlatform(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the HVM/shadow_multiplier field of the given VM.
func (_class VMClass) GetHVMShadowMultiplier(sessionID SessionRef, self VMRef) (_retval float64, _err error) {
	_method := "VM.get_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertFloatToGo(_method + " -> ", _result.Value)
	return
}

// Get the HVM/boot_params field of the given VM.
func (_class VMClass) GetHVMBootParams(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the HVM/boot_policy field of the given VM.
func (_class VMClass) GetHVMBootPolicy(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/legacy_args field of the given VM.
func (_class VMClass) GetPVLegacyArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/bootloader_args field of the given VM.
func (_class VMClass) GetPVBootloaderArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/args field of the given VM.
func (_class VMClass) GetPVArgs(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/ramdisk field of the given VM.
func (_class VMClass) GetPVRamdisk(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/kernel field of the given VM.
func (_class VMClass) GetPVKernel(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/bootloader field of the given VM.
func (_class VMClass) GetPVBootloader(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the VTPMs field of the given VM.
func (_class VMClass) GetVTPMs(sessionID SessionRef, self VMRef) (_retval []VTPMRef, _err error) {
	_method := "VM.get_VTPMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVTPMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the crash_dumps field of the given VM.
func (_class VMClass) GetCrashDumps(sessionID SessionRef, self VMRef) (_retval []CrashdumpRef, _err error) {
	_method := "VM.get_crash_dumps"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertCrashdumpRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the VBDs field of the given VM.
func (_class VMClass) GetVBDs(sessionID SessionRef, self VMRef) (_retval []VBDRef, _err error) {
	_method := "VM.get_VBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVBDRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the VIFs field of the given VM.
func (_class VMClass) GetVIFs(sessionID SessionRef, self VMRef) (_retval []VIFRef, _err error) {
	_method := "VM.get_VIFs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVIFRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the consoles field of the given VM.
func (_class VMClass) GetConsoles(sessionID SessionRef, self VMRef) (_retval []ConsoleRef, _err error) {
	_method := "VM.get_consoles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertConsoleRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the actions/after_crash field of the given VM.
func (_class VMClass) GetActionsAfterCrash(sessionID SessionRef, self VMRef) (_retval OnCrashBehaviour, _err error) {
	_method := "VM.get_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnCrashBehaviourToGo(_method + " -> ", _result.Value)
	return
}

// Get the actions/after_reboot field of the given VM.
func (_class VMClass) GetActionsAfterReboot(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	_method := "VM.get_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnNormalExitToGo(_method + " -> ", _result.Value)
	return
}

// Get the actions/after_shutdown field of the given VM.
func (_class VMClass) GetActionsAfterShutdown(sessionID SessionRef, self VMRef) (_retval OnNormalExit, _err error) {
	_method := "VM.get_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumOnNormalExitToGo(_method + " -> ", _result.Value)
	return
}

// Get the VCPUs/at_startup field of the given VM.
func (_class VMClass) GetVCPUsAtStartup(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the VCPUs/max field of the given VM.
func (_class VMClass) GetVCPUsMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the VCPUs/params field of the given VM.
func (_class VMClass) GetVCPUsParams(sessionID SessionRef, self VMRef) (_retval map[string]string, _err error) {
	_method := "VM.get_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToStringMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/static_min field of the given VM.
func (_class VMClass) GetMemoryStaticMin(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/dynamic_min field of the given VM.
func (_class VMClass) GetMemoryDynamicMin(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/dynamic_max field of the given VM.
func (_class VMClass) GetMemoryDynamicMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/static_max field of the given VM.
func (_class VMClass) GetMemoryStaticMax(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/target field of the given VM.
func (_class VMClass) GetMemoryTarget(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_target"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/overhead field of the given VM.
func (_class VMClass) GetMemoryOverhead(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the affinity field of the given VM.
func (_class VMClass) GetAffinity(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	_method := "VM.get_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the resident_on field of the given VM.
func (_class VMClass) GetResidentOn(sessionID SessionRef, self VMRef) (_retval HostRef, _err error) {
	_method := "VM.get_resident_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertHostRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the suspend_VDI field of the given VM.
func (_class VMClass) GetSuspendVDI(sessionID SessionRef, self VMRef) (_retval VDIRef, _err error) {
	_method := "VM.get_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVDIRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_a_template field of the given VM.
func (_class VMClass) GetIsATemplate(sessionID SessionRef, self VMRef) (_retval bool, _err error) {
	_method := "VM.get_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertBoolToGo(_method + " -> ", _result.Value)
	return
}

// Get the user_version field of the given VM.
func (_class VMClass) GetUserVersion(sessionID SessionRef, self VMRef) (_retval int, _err error) {
	_method := "VM.get_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertIntToGo(_method + " -> ", _result.Value)
	return
}

// Get the name/description field of the given VM.
func (_class VMClass) GetNameDescription(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the name/label field of the given VM.
func (_class VMClass) GetNameLabel(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get the power_state field of the given VM.
func (_class VMClass) GetPowerState(sessionID SessionRef, self VMRef) (_retval VMPowerState, _err error) {
	_method := "VM.get_power_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMPowerStateToGo(_method + " -> ", _result.Value)
	return
}

// Get the current_operations field of the given VM.
func (_class VMClass) GetCurrentOperations(sessionID SessionRef, self VMRef) (_retval map[string]VMOperations, _err error) {
	_method := "VM.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToEnumVMOperationsMapToGo(_method + " -> ", _result.Value)
	return
}

// Get the allowed_operations field of the given VM.
func (_class VMClass) GetAllowedOperations(sessionID SessionRef, self VMRef) (_retval []VMOperations, _err error) {
	_method := "VM.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertEnumVMOperationsSetToGo(_method + " -> ", _result.Value)
	return
}

// Get the uuid field of the given VM.
func (_class VMClass) GetUUID(sessionID SessionRef, self VMRef) (_retval string, _err error) {
	_method := "VM.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertStringToGo(_method + " -> ", _result.Value)
	return
}

// Get all the VM instances with the given label.
func (_class VMClass) GetByNameLabel(sessionID SessionRef, label string) (_retval []VMRef, _err error) {
	_method := "VM.get_by_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefSetToGo(_method + " -> ", _result.Value)
	return
}

// Destroy the specified VM.  The VM is completely removed from the system.  This function can only be called when the VM is in the Halted State.
func (_class VMClass) Destroy(sessionID SessionRef, self VMRef) (_err error) {
	_method := "VM.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_, _err =  _class.client.APICall(_method, _sessionIDArg, _selfArg)
	return
}

// Create a new VM instance, and return its handle.
// The constructor args are: name_label, name_description, user_version*, is_a_template*, affinity*, memory_target, memory_static_max*, memory_dynamic_max*, memory_dynamic_min*, memory_static_min*, VCPUs_params*, VCPUs_max*, VCPUs_at_startup*, actions_after_shutdown*, actions_after_reboot*, actions_after_crash*, PV_bootloader*, PV_kernel*, PV_ramdisk*, PV_args*, PV_bootloader_args*, PV_legacy_args*, HVM_boot_policy*, HVM_boot_params*, HVM_shadow_multiplier, platform*, PCI_bus*, other_config*, recommendations*, xenstore_data, ha_always_run, ha_restart_priority, tags, blocked_operations, protection_policy, is_snapshot_from_vmpp, appliance, start_delay, shutdown_delay, order, suspend_SR, version, generation_id, hardware_platform_version (* = non-optional).
func (_class VMClass) Create(sessionID SessionRef, args VMRecord) (_retval VMRef, _err error) {
	_method := "VM.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Get a reference to the VM instance with the specified UUID.
func (_class VMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval VMRef, _err error) {
	_method := "VM.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRefToGo(_method + " -> ", _result.Value)
	return
}

// Get a record containing the current state of the given VM.
func (_class VMClass) GetRecord(sessionID SessionRef, self VMRef) (_retval VMRecord, _err error) {
	_method := "VM.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertVMRecordToGo(_method + " -> ", _result.Value)
	return
}

// Return a map of VM references to VM records for all VMs known to the system.
func (_class AsyncVMClass) GetAllRecords(sessionID SessionRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_all_records"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Return a list of all the VMs known to the system.
func (_class AsyncVMClass) GetAll(sessionID SessionRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_all"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Start the 'xenprep' process on the VM; the process will remove any tools and drivers for XenServer and then set auto update drivers true.
func (_class AsyncVMClass) XenprepStart(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.xenprep_start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Import an XVA from a URI
func (_class AsyncVMClass) Import(sessionID SessionRef, url string, sr SRRef, fullRestore bool, force bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.import"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_urlArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "url"), url)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_fullRestoreArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "full_restore"), fullRestore)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _urlArg, _srArg, _fullRestoreArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Check if PV auto update can be set on Windows vm
func (_class AsyncVMClass) AssertCanSetAutoUpdateDrivers(sessionID SessionRef, self VMRef, value bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.assert_can_set_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Enable or disable PV auto update on Windows vm
func (_class AsyncVMClass) SetAutoUpdateDrivers(sessionID SessionRef, self VMRef, value bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Call a XenAPI plugin on this vm
func (_class AsyncVMClass) CallPlugin(sessionID SessionRef, vm VMRef, plugin string, fn string, args map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.call_plugin"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_pluginArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "plugin"), plugin)
	if _err != nil {
		return
	}
	_fnArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "fn"), fn)
	if _err != nil {
		return
	}
	_argsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _pluginArg, _fnArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Query the system services advertised by this VM and register them. This can only be applied to a system domain.
func (_class AsyncVMClass) QueryServices(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.query_services"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Assign this VM to an appliance.
func (_class AsyncVMClass) SetAppliance(sessionID SessionRef, self VMRef, value VMApplianceRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMApplianceRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Import using a conversion service.
func (_class AsyncVMClass) ImportConvert(sessionID SessionRef, atype string, username string, password string, sr SRRef, remoteConfig map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.import_convert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_atypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "type"), atype)
	if _err != nil {
		return
	}
	_usernameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "username"), username)
	if _err != nil {
		return
	}
	_passwordArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "password"), password)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_remoteConfigArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "remote_config"), remoteConfig)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _atypeArg, _usernameArg, _passwordArg, _srArg, _remoteConfigArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Recover the VM
func (_class AsyncVMClass) Recover(sessionID SessionRef, self VMRef, sessionTo SessionRef, force bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.recover"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// List all the SR's that are required for the VM to be recovered
func (_class AsyncVMClass) GetSRsRequiredForRecovery(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_SRs_required_for_recovery"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Assert whether all SRs required to recover this VM are available.
//
// Errors:
//  VM_IS_PART_OF_AN_APPLIANCE - This operation is not allowed as the VM is part of an appliance.
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
func (_class AsyncVMClass) AssertCanBeRecovered(sessionID SessionRef, self VMRef, sessionTo SessionRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.assert_can_be_recovered"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_sessionToArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_to"), sessionTo)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _sessionToArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set this VM's suspend VDI, which must be indentical to its current one
func (_class AsyncVMClass) SetSuspendVDI(sessionID SessionRef, self VMRef, value VDIRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVDIRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set this VM's boot order
func (_class AsyncVMClass) SetOrder(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set this VM's shutdown delay in seconds
func (_class AsyncVMClass) SetShutdownDelay(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set this VM's start delay in seconds
func (_class AsyncVMClass) SetStartDelay(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the protection_policy field
func (_class AsyncVMClass) SetProtectionPolicy(sessionID SessionRef, self VMRef, value VMPPRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertVMPPRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Copy the BIOS strings from the given host to this VM
func (_class AsyncVMClass) CopyBiosStrings(sessionID SessionRef, vm VMRef, host HostRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.copy_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns mapping of hosts to ratings, indicating the suitability of starting the VM at that location according to wlb. Rating is replaced with an error if the VM cannot boot there.
func (_class AsyncVMClass) RetrieveWlbRecommendations(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.retrieve_wlb_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns an error if the VM is not considered agile e.g. because it is tied to a resource local to a host
func (_class AsyncVMClass) AssertAgile(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.assert_agile"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Create a placeholder for a named binary blob of data that is associated with this VM
func (_class AsyncVMClass) CreateNewBlob(sessionID SessionRef, vm VMRef, name string, mimeType string, public bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.create_new_blob"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_nameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "name"), name)
	if _err != nil {
		return
	}
	_mimeTypeArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "mime_type"), mimeType)
	if _err != nil {
		return
	}
	_publicArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "public"), public)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _nameArg, _mimeTypeArg, _publicArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns an error if the VM could not boot on this host for some reason
//
// Errors:
//  HOST_NOT_ENOUGH_FREE_MEMORY - Not enough host memory is available to perform this operation
//  VM_REQUIRES_SR - You attempted to run a VM on a host which doesn't have access to an SR needed by the VM. The VM has at least one VBD attached to a VDI in the SR.
//  VM_HOST_INCOMPATIBLE_VERSION - This VM operation cannot be performed on an older-versioned host during an upgrade.
//  VM_HOST_INCOMPATIBLE_VIRTUAL_HARDWARE_PLATFORM_VERSION - You attempted to run a VM on a host that cannot provide the VM's required Virtual Hardware Platform version.
func (_class AsyncVMClass) AssertCanBootHere(sessionID SessionRef, self VMRef, host HostRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.assert_can_boot_here"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _hostArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Return the list of hosts on which this VM may run.
func (_class AsyncVMClass) GetPossibleHosts(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_possible_hosts"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns a list of the allowed values that a VIF device field can take
func (_class AsyncVMClass) GetAllowedVIFDevices(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_allowed_VIF_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns a list of the allowed values that a VBD device field can take
func (_class AsyncVMClass) GetAllowedVBDDevices(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_allowed_VBD_devices"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Recomputes the list of acceptable operations
func (_class AsyncVMClass) UpdateAllowedOperations(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.update_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Check to see whether this operation is acceptable in the current state of the system, raising an error if the operation is invalid for some reason
func (_class AsyncVMClass) AssertOperationValid(sessionID SessionRef, self VMRef, op VMOperations) (_retval TaskRef, _err error) {
	_method := "Async.VM.assert_operation_valid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_opArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "op"), op)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _opArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Forget the recorded statistics related to the specified data source
func (_class AsyncVMClass) ForgetDataSourceArchives(sessionID SessionRef, self VMRef, dataSource string) (_retval TaskRef, _err error) {
	_method := "Async.VM.forget_data_source_archives"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Query the latest value of the specified data source
func (_class AsyncVMClass) QueryDataSource(sessionID SessionRef, self VMRef, dataSource string) (_retval TaskRef, _err error) {
	_method := "Async.VM.query_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Start recording the specified data source
func (_class AsyncVMClass) RecordDataSource(sessionID SessionRef, self VMRef, dataSource string) (_retval TaskRef, _err error) {
	_method := "Async.VM.record_data_source"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_dataSourceArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "data_source"), dataSource)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _dataSourceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// 
func (_class AsyncVMClass) GetDataSources(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_data_sources"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns a record describing the VM's dynamic state, initialised when the VM boots and updated to reflect runtime configuration changes e.g. CPU hotplug
func (_class AsyncVMClass) GetBootRecord(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_boot_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Assert whether a VM can be migrated to the specified destination.
func (_class AsyncVMClass) AssertCanMigrate(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.assert_can_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Migrate the VM to another host.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class AsyncVMClass) MigrateSend(sessionID SessionRef, vm VMRef, dest map[string]string, live bool, vdiMap map[VDIRef]SRRef, vifMap map[VIFRef]NetworkRef, options map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.migrate_send"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_destArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "dest"), dest)
	if _err != nil {
		return
	}
	_liveArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "live"), live)
	if _err != nil {
		return
	}
	_vdiMapArg, _err := convertVDIRefToSRRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vdi_map"), vdiMap)
	if _err != nil {
		return
	}
	_vifMapArg, _err := convertVIFRefToNetworkRefMapToXen(fmt.Sprintf("%s(%s)", _method, "vif_map"), vifMap)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _destArg, _liveArg, _vdiMapArg, _vifMapArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Returns the maximum amount of guest memory which will fit, together with overheads, in the supplied amount of physical memory. If 'exact' is true then an exact calculation is performed using the VM's current settings. If 'exact' is false then a more conservative approximation is used
func (_class AsyncVMClass) MaximiseMemory(sessionID SessionRef, self VMRef, total int, approximate bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.maximise_memory"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_totalArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "total"), total)
	if _err != nil {
		return
	}
	_approximateArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "approximate"), approximate)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _totalArg, _approximateArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Send the named trigger to this VM.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class AsyncVMClass) SendTrigger(sessionID SessionRef, vm VMRef, trigger string) (_retval TaskRef, _err error) {
	_method := "Async.VM.send_trigger"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_triggerArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "trigger"), trigger)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _triggerArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Send the given key as a sysrq to this VM.  The key is specified as a single character (a String of length 1).  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
func (_class AsyncVMClass) SendSysrq(sessionID SessionRef, vm VMRef, key string) (_retval TaskRef, _err error) {
	_method := "Async.VM.send_sysrq"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _keyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the number of startup VCPUs for a halted VM
func (_class AsyncVMClass) SetVCPUsAtStartup(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the maximum number of VCPUs for a halted VM
func (_class AsyncVMClass) SetVCPUsMax(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the shadow memory multiplier on a running VM
func (_class AsyncVMClass) SetShadowMultiplierLive(sessionID SessionRef, self VMRef, multiplier float64) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_shadow_multiplier_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_multiplierArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "multiplier"), multiplier)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _multiplierArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the shadow memory multiplier on a halted VM
func (_class AsyncVMClass) SetHVMShadowMultiplier(sessionID SessionRef, self VMRef, value float64) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertFloatToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Return true if the VM is currently 'co-operative' i.e. is expected to reach a balloon target and actually has done
func (_class AsyncVMClass) GetCooperative(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_cooperative"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Wait for a running VM to reach its current memory target
func (_class AsyncVMClass) WaitMemoryTargetLive(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.wait_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the memory target for a running VM
func (_class AsyncVMClass) SetMemoryTargetLive(sessionID SessionRef, self VMRef, target int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_target_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_targetArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "target"), target)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _targetArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the memory limits of this VM.
func (_class AsyncVMClass) SetMemoryLimits(sessionID SessionRef, self VMRef, staticMin int, staticMax int, dynamicMin int, dynamicMax int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_limits"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_staticMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_min"), staticMin)
	if _err != nil {
		return
	}
	_staticMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "static_max"), staticMax)
	if _err != nil {
		return
	}
	_dynamicMinArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_min"), dynamicMin)
	if _err != nil {
		return
	}
	_dynamicMaxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "dynamic_max"), dynamicMax)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _staticMinArg, _staticMaxArg, _dynamicMinArg, _dynamicMaxArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the static (ie boot-time) range of virtual memory that the VM is allowed to use.
func (_class AsyncVMClass) SetMemoryStaticRange(sessionID SessionRef, self VMRef, min int, max int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_static_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the memory_static_min field
func (_class AsyncVMClass) SetMemoryStaticMin(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the memory_static_max field
//
// Errors:
//  HA_OPERATION_WOULD_BREAK_FAILOVER_PLAN - This operation cannot be performed because it would invalidate VM failover planning such that the system would be unable to guarantee to restart protected VMs after a Host failure.
func (_class AsyncVMClass) SetMemoryStaticMax(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the minimum and maximum amounts of physical memory the VM is allowed to use.
func (_class AsyncVMClass) SetMemoryDynamicRange(sessionID SessionRef, self VMRef, min int, max int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_dynamic_range"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_minArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "min"), min)
	if _err != nil {
		return
	}
	_maxArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "max"), max)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _minArg, _maxArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the memory_dynamic_min field
func (_class AsyncVMClass) SetMemoryDynamicMin(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the memory_dynamic_max field
func (_class AsyncVMClass) SetMemoryDynamicMax(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Computes the virtualization memory overhead of a VM.
func (_class AsyncVMClass) ComputeMemoryOverhead(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.compute_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the ha_always_run
func (_class AsyncVMClass) SetHaAlwaysRun(sessionID SessionRef, self VMRef, value bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the value of the ha_restart_priority field
func (_class AsyncVMClass) SetHaRestartPriority(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given key-value pair to VM.VCPUs_params, and apply that value on the running VM
func (_class AsyncVMClass) AddToVCPUsParamsLive(sessionID SessionRef, self VMRef, key string, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_to_VCPUs_params_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the number of VCPUs for a running VM
func (_class AsyncVMClass) SetVCPUsNumberLive(sessionID SessionRef, self VMRef, nvcpu int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_VCPUs_number_live"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_nvcpuArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "nvcpu"), nvcpu)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _nvcpuArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Migrate a VM to another Host.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_MIGRATE_FAILED - An error occurred during the migration process.
func (_class AsyncVMClass) PoolMigrate(sessionID SessionRef, vm VMRef, host HostRef, options map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.pool_migrate"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_optionsArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "options"), options)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _optionsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Awaken the specified VM and resume it on a particular Host.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) ResumeOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.resume_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Awaken the specified VM and resume it.  This can only be called when the specified VM is in the Suspended state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) Resume(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.resume"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Suspend the specified VM to disk.  This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) Suspend(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.suspend"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Stop executing the specified VM without attempting a clean shutdown and immediately restart the VM.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) HardReboot(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.hard_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Reset the power-state of the VM to halted in the database only. (Used to recover from slave failures in pooling scenarios by resetting the power-states of VMs running on dead slaves to halted.) This is a potentially dangerous operation; use with care.
func (_class AsyncVMClass) PowerStateReset(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.power_state_reset"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Stop executing the specified VM without attempting a clean shutdown.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) HardShutdown(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.hard_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Attempt to cleanly shutdown the specified VM (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) CleanReboot(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.clean_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Attempts to first clean shutdown a VM and if it should fail then perform a hard shutdown on it.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) Shutdown(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Attempt to cleanly shutdown the specified VM. (Note: this may not be supported---e.g. if a guest agent is not installed). This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) CleanShutdown(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.clean_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Resume the specified VM. This can only be called when the specified VM is in the Paused state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) Unpause(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.unpause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Pause the specified VM. This can only be called when the specified VM is in the Running state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
func (_class AsyncVMClass) Pause(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.pause"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Start the specified VM on a particular host.  This function can only be called with the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  BOOTLOADER_FAILED - The bootloader returned an error
//  UNKNOWN_BOOTLOADER - The requested bootloader is unknown
func (_class AsyncVMClass) StartOn(sessionID SessionRef, vm VMRef, host HostRef, startPaused bool, force bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.start_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_hostArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "host"), host)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _hostArg, _startPausedArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Start the specified VM.  This function can only be called with the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  VM_HVM_REQUIRED - HVM is required for this operation
//  VM_IS_TEMPLATE - The operation attempted is not valid for a template VM
//  OTHER_OPERATION_IN_PROGRESS - Another operation involving the object is currently in progress
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  BOOTLOADER_FAILED - The bootloader returned an error
//  UNKNOWN_BOOTLOADER - The requested bootloader is unknown
//  NO_HOSTS_AVAILABLE - There were no hosts available to complete the specified operation.
//  LICENCE_RESTRICTION - This operation is not allowed under your license.  Please contact your support representative.
func (_class AsyncVMClass) Start(sessionID SessionRef, vm VMRef, startPaused bool, force bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.start"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_startPausedArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "start_paused"), startPaused)
	if _err != nil {
		return
	}
	_forceArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "force"), force)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _startPausedArg, _forceArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Inspects the disk configuration contained within the VM's other_config, creates VDIs and VBDs and then executes any applicable post-install script.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class AsyncVMClass) Provision(sessionID SessionRef, vm VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.provision"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Checkpoints the specified VM, making a new VM. Checkpoint automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write) and saves the memory image as well.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_CHECKPOINT_SUSPEND_FAILED - An error occured while saving the memory image of the specified virtual machine
//  VM_CHECKPOINT_RESUME_FAILED - An error occured while restoring the memory image of the specified virtual machine
func (_class AsyncVMClass) Checkpoint(sessionID SessionRef, vm VMRef, newName string) (_retval TaskRef, _err error) {
	_method := "Async.VM.checkpoint"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Reverts the specified VM to a previous state.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  VM_REVERT_FAILED - An error occured while reverting the specified virtual machine to the specified snapshot
func (_class AsyncVMClass) Revert(sessionID SessionRef, snapshot VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.revert"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_snapshotArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "snapshot"), snapshot)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _snapshotArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Copied the specified VM, making a new VM. Unlike clone, copy does not exploits the capabilities of the underlying storage repository in which the VM's disk images are stored. Instead, copy guarantees that the disk images of the newly created VM will be 'full disks' - i.e. not part of a CoW chain.  This function can only be called when the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class AsyncVMClass) Copy(sessionID SessionRef, vm VMRef, newName string, sr SRRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.copy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_srArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "sr"), sr)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg, _srArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Clones the specified VM, making a new VM. Clone automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).   This function can only be called when the VM is in the Halted State.
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class AsyncVMClass) Clone(sessionID SessionRef, vm VMRef, newName string) (_retval TaskRef, _err error) {
	_method := "Async.VM.clone"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Snapshots the specified VM with quiesce, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
//  VM_SNAPSHOT_WITH_QUIESCE_FAILED - The quiesced-snapshot operation failed for an unexpected reason
//  VM_SNAPSHOT_WITH_QUIESCE_TIMEOUT - The VSS plug-in has timed out
//  VM_SNAPSHOT_WITH_QUIESCE_PLUGIN_DEOS_NOT_RESPOND - The VSS plug-in cannot be contacted
//  VM_SNAPSHOT_WITH_QUIESCE_NOT_SUPPORTED - The VSS plug-in is not installed on this virtual machine
func (_class AsyncVMClass) SnapshotWithQuiesce(sessionID SessionRef, vm VMRef, newName string) (_retval TaskRef, _err error) {
	_method := "Async.VM.snapshot_with_quiesce"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Snapshots the specified VM, making a new VM. Snapshot automatically exploits the capabilities of the underlying storage repository in which the VM's disk images are stored (e.g. Copy on Write).
//
// Errors:
//  VM_BAD_POWER_STATE - You attempted an operation on a VM that was not in an appropriate power state at the time; for example, you attempted to start a VM that was already running.  The parameters returned are the VM's handle, and the expected and actual VM state at the time of the call.
//  SR_FULL - The SR is full. Requested new size exceeds the maximum size
//  OPERATION_NOT_ALLOWED - You attempted an operation that was not allowed.
func (_class AsyncVMClass) Snapshot(sessionID SessionRef, vm VMRef, newName string) (_retval TaskRef, _err error) {
	_method := "Async.VM.snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_vmArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "vm"), vm)
	if _err != nil {
		return
	}
	_newNameArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "new_name"), newName)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _vmArg, _newNameArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the hardware_platform_version field of the given VM.
func (_class AsyncVMClass) SetHardwarePlatformVersion(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the suspend_SR field of the given VM.
func (_class AsyncVMClass) SetSuspendSR(sessionID SessionRef, self VMRef, value SRRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertSRRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given key and its corresponding value from the blocked_operations field of the given VM.  If the key is not in that Map, then do nothing.
func (_class AsyncVMClass) RemoveFromBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations) (_retval TaskRef, _err error) {
	_method := "Async.VM.remove_from_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given key-value pair to the blocked_operations field of the given VM.
func (_class AsyncVMClass) AddToBlockedOperations(sessionID SessionRef, self VMRef, key VMOperations, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_to_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertEnumVMOperationsToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the blocked_operations field of the given VM.
func (_class AsyncVMClass) SetBlockedOperations(sessionID SessionRef, self VMRef, value map[VMOperations]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumVMOperationsToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given value from the tags field of the given VM.  If the value is not in that Set, then do nothing.
func (_class AsyncVMClass) RemoveTags(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.remove_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given value to the tags field of the given VM.  If the value is already in that Set, then do nothing.
func (_class AsyncVMClass) AddTags(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the tags field of the given VM.
func (_class AsyncVMClass) SetTags(sessionID SessionRef, self VMRef, value []string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringSetToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given key and its corresponding value from the xenstore_data field of the given VM.  If the key is not in that Map, then do nothing.
func (_class AsyncVMClass) RemoveFromXenstoreData(sessionID SessionRef, self VMRef, key string) (_retval TaskRef, _err error) {
	_method := "Async.VM.remove_from_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given key-value pair to the xenstore_data field of the given VM.
func (_class AsyncVMClass) AddToXenstoreData(sessionID SessionRef, self VMRef, key string, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_to_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the xenstore_data field of the given VM.
func (_class AsyncVMClass) SetXenstoreData(sessionID SessionRef, self VMRef, value map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the recommendations field of the given VM.
func (_class AsyncVMClass) SetRecommendations(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given key and its corresponding value from the other_config field of the given VM.  If the key is not in that Map, then do nothing.
func (_class AsyncVMClass) RemoveFromOtherConfig(sessionID SessionRef, self VMRef, key string) (_retval TaskRef, _err error) {
	_method := "Async.VM.remove_from_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given key-value pair to the other_config field of the given VM.
func (_class AsyncVMClass) AddToOtherConfig(sessionID SessionRef, self VMRef, key string, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_to_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the other_config field of the given VM.
func (_class AsyncVMClass) SetOtherConfig(sessionID SessionRef, self VMRef, value map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the PCI_bus field of the given VM.
func (_class AsyncVMClass) SetPCIBus(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given key and its corresponding value from the platform field of the given VM.  If the key is not in that Map, then do nothing.
func (_class AsyncVMClass) RemoveFromPlatform(sessionID SessionRef, self VMRef, key string) (_retval TaskRef, _err error) {
	_method := "Async.VM.remove_from_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given key-value pair to the platform field of the given VM.
func (_class AsyncVMClass) AddToPlatform(sessionID SessionRef, self VMRef, key string, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_to_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the platform field of the given VM.
func (_class AsyncVMClass) SetPlatform(sessionID SessionRef, self VMRef, value map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given key and its corresponding value from the HVM/boot_params field of the given VM.  If the key is not in that Map, then do nothing.
func (_class AsyncVMClass) RemoveFromHVMBootParams(sessionID SessionRef, self VMRef, key string) (_retval TaskRef, _err error) {
	_method := "Async.VM.remove_from_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given key-value pair to the HVM/boot_params field of the given VM.
func (_class AsyncVMClass) AddToHVMBootParams(sessionID SessionRef, self VMRef, key string, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_to_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the HVM/boot_params field of the given VM.
func (_class AsyncVMClass) SetHVMBootParams(sessionID SessionRef, self VMRef, value map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the HVM/boot_policy field of the given VM.
func (_class AsyncVMClass) SetHVMBootPolicy(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the PV/legacy_args field of the given VM.
func (_class AsyncVMClass) SetPVLegacyArgs(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the PV/bootloader_args field of the given VM.
func (_class AsyncVMClass) SetPVBootloaderArgs(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the PV/args field of the given VM.
func (_class AsyncVMClass) SetPVArgs(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the PV/ramdisk field of the given VM.
func (_class AsyncVMClass) SetPVRamdisk(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the PV/kernel field of the given VM.
func (_class AsyncVMClass) SetPVKernel(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the PV/bootloader field of the given VM.
func (_class AsyncVMClass) SetPVBootloader(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the actions/after_crash field of the given VM.
func (_class AsyncVMClass) SetActionsAfterCrash(sessionID SessionRef, self VMRef, value OnCrashBehaviour) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnCrashBehaviourToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the actions/after_reboot field of the given VM.
func (_class AsyncVMClass) SetActionsAfterReboot(sessionID SessionRef, self VMRef, value OnNormalExit) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the actions/after_shutdown field of the given VM.
func (_class AsyncVMClass) SetActionsAfterShutdown(sessionID SessionRef, self VMRef, value OnNormalExit) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertEnumOnNormalExitToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Remove the given key and its corresponding value from the VCPUs/params field of the given VM.  If the key is not in that Map, then do nothing.
func (_class AsyncVMClass) RemoveFromVCPUsParams(sessionID SessionRef, self VMRef, key string) (_retval TaskRef, _err error) {
	_method := "Async.VM.remove_from_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Add the given key-value pair to the VCPUs/params field of the given VM.
func (_class AsyncVMClass) AddToVCPUsParams(sessionID SessionRef, self VMRef, key string, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.add_to_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_keyArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "key"), key)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _keyArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the VCPUs/params field of the given VM.
func (_class AsyncVMClass) SetVCPUsParams(sessionID SessionRef, self VMRef, value map[string]string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToStringMapToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the affinity field of the given VM.
func (_class AsyncVMClass) SetAffinity(sessionID SessionRef, self VMRef, value HostRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertHostRefToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the is_a_template field of the given VM.
func (_class AsyncVMClass) SetIsATemplate(sessionID SessionRef, self VMRef, value bool) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertBoolToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the user_version field of the given VM.
func (_class AsyncVMClass) SetUserVersion(sessionID SessionRef, self VMRef, value int) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertIntToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the name/description field of the given VM.
func (_class AsyncVMClass) SetNameDescription(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Set the name/label field of the given VM.
func (_class AsyncVMClass) SetNameLabel(sessionID SessionRef, self VMRef, value string) (_retval TaskRef, _err error) {
	_method := "Async.VM.set_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_valueArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "value"), value)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg, _valueArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the auto_update_drivers field of the given VM.
func (_class AsyncVMClass) GetAutoUpdateDrivers(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_auto_update_drivers"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the hardware_platform_version field of the given VM.
func (_class AsyncVMClass) GetHardwarePlatformVersion(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_hardware_platform_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the generation_id field of the given VM.
func (_class AsyncVMClass) GetGenerationID(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_generation_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the version field of the given VM.
func (_class AsyncVMClass) GetVersion(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the suspend_SR field of the given VM.
func (_class AsyncVMClass) GetSuspendSR(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_suspend_SR"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the attached_PCIs field of the given VM.
func (_class AsyncVMClass) GetAttachedPCIs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_attached_PCIs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the VGPUs field of the given VM.
func (_class AsyncVMClass) GetVGPUs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_VGPUs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the order field of the given VM.
func (_class AsyncVMClass) GetOrder(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_order"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the shutdown_delay field of the given VM.
func (_class AsyncVMClass) GetShutdownDelay(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_shutdown_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the start_delay field of the given VM.
func (_class AsyncVMClass) GetStartDelay(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_start_delay"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the appliance field of the given VM.
func (_class AsyncVMClass) GetAppliance(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_appliance"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_snapshot_from_vmpp field of the given VM.
func (_class AsyncVMClass) GetIsSnapshotFromVmpp(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_is_snapshot_from_vmpp"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the protection_policy field of the given VM.
func (_class AsyncVMClass) GetProtectionPolicy(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_protection_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the bios_strings field of the given VM.
func (_class AsyncVMClass) GetBiosStrings(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_bios_strings"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the children field of the given VM.
func (_class AsyncVMClass) GetChildren(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_children"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the parent field of the given VM.
func (_class AsyncVMClass) GetParent(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_parent"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_metadata field of the given VM.
func (_class AsyncVMClass) GetSnapshotMetadata(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_snapshot_metadata"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_info field of the given VM.
func (_class AsyncVMClass) GetSnapshotInfo(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_snapshot_info"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the blocked_operations field of the given VM.
func (_class AsyncVMClass) GetBlockedOperations(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_blocked_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the tags field of the given VM.
func (_class AsyncVMClass) GetTags(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_tags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the blobs field of the given VM.
func (_class AsyncVMClass) GetBlobs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_blobs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the transportable_snapshot_id field of the given VM.
func (_class AsyncVMClass) GetTransportableSnapshotID(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_transportable_snapshot_id"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_time field of the given VM.
func (_class AsyncVMClass) GetSnapshotTime(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_snapshot_time"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshots field of the given VM.
func (_class AsyncVMClass) GetSnapshots(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_snapshots"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the snapshot_of field of the given VM.
func (_class AsyncVMClass) GetSnapshotOf(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_snapshot_of"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_a_snapshot field of the given VM.
func (_class AsyncVMClass) GetIsASnapshot(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_is_a_snapshot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the ha_restart_priority field of the given VM.
func (_class AsyncVMClass) GetHaRestartPriority(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_ha_restart_priority"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the ha_always_run field of the given VM.
func (_class AsyncVMClass) GetHaAlwaysRun(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_ha_always_run"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the xenstore_data field of the given VM.
func (_class AsyncVMClass) GetXenstoreData(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_xenstore_data"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the recommendations field of the given VM.
func (_class AsyncVMClass) GetRecommendations(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_recommendations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the last_booted_record field of the given VM.
func (_class AsyncVMClass) GetLastBootedRecord(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_last_booted_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the guest_metrics field of the given VM.
func (_class AsyncVMClass) GetGuestMetrics(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_guest_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the metrics field of the given VM.
func (_class AsyncVMClass) GetMetrics(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_metrics"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_control_domain field of the given VM.
func (_class AsyncVMClass) GetIsControlDomain(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_is_control_domain"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the last_boot_CPU_flags field of the given VM.
func (_class AsyncVMClass) GetLastBootCPUFlags(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_last_boot_CPU_flags"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the domarch field of the given VM.
func (_class AsyncVMClass) GetDomarch(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_domarch"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the domid field of the given VM.
func (_class AsyncVMClass) GetDomid(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_domid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the other_config field of the given VM.
func (_class AsyncVMClass) GetOtherConfig(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_other_config"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the PCI_bus field of the given VM.
func (_class AsyncVMClass) GetPCIBus(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_PCI_bus"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the platform field of the given VM.
func (_class AsyncVMClass) GetPlatform(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_platform"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the HVM/shadow_multiplier field of the given VM.
func (_class AsyncVMClass) GetHVMShadowMultiplier(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_HVM_shadow_multiplier"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the HVM/boot_params field of the given VM.
func (_class AsyncVMClass) GetHVMBootParams(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_HVM_boot_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the HVM/boot_policy field of the given VM.
func (_class AsyncVMClass) GetHVMBootPolicy(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_HVM_boot_policy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/legacy_args field of the given VM.
func (_class AsyncVMClass) GetPVLegacyArgs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_PV_legacy_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/bootloader_args field of the given VM.
func (_class AsyncVMClass) GetPVBootloaderArgs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_PV_bootloader_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/args field of the given VM.
func (_class AsyncVMClass) GetPVArgs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_PV_args"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/ramdisk field of the given VM.
func (_class AsyncVMClass) GetPVRamdisk(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_PV_ramdisk"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/kernel field of the given VM.
func (_class AsyncVMClass) GetPVKernel(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_PV_kernel"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the PV/bootloader field of the given VM.
func (_class AsyncVMClass) GetPVBootloader(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_PV_bootloader"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the VTPMs field of the given VM.
func (_class AsyncVMClass) GetVTPMs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_VTPMs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the crash_dumps field of the given VM.
func (_class AsyncVMClass) GetCrashDumps(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_crash_dumps"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the VBDs field of the given VM.
func (_class AsyncVMClass) GetVBDs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_VBDs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the VIFs field of the given VM.
func (_class AsyncVMClass) GetVIFs(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_VIFs"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the consoles field of the given VM.
func (_class AsyncVMClass) GetConsoles(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_consoles"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the actions/after_crash field of the given VM.
func (_class AsyncVMClass) GetActionsAfterCrash(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_actions_after_crash"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the actions/after_reboot field of the given VM.
func (_class AsyncVMClass) GetActionsAfterReboot(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_actions_after_reboot"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the actions/after_shutdown field of the given VM.
func (_class AsyncVMClass) GetActionsAfterShutdown(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_actions_after_shutdown"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the VCPUs/at_startup field of the given VM.
func (_class AsyncVMClass) GetVCPUsAtStartup(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_VCPUs_at_startup"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the VCPUs/max field of the given VM.
func (_class AsyncVMClass) GetVCPUsMax(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_VCPUs_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the VCPUs/params field of the given VM.
func (_class AsyncVMClass) GetVCPUsParams(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_VCPUs_params"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/static_min field of the given VM.
func (_class AsyncVMClass) GetMemoryStaticMin(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_memory_static_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/dynamic_min field of the given VM.
func (_class AsyncVMClass) GetMemoryDynamicMin(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_memory_dynamic_min"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/dynamic_max field of the given VM.
func (_class AsyncVMClass) GetMemoryDynamicMax(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_memory_dynamic_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/static_max field of the given VM.
func (_class AsyncVMClass) GetMemoryStaticMax(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_memory_static_max"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/target field of the given VM.
func (_class AsyncVMClass) GetMemoryTarget(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_memory_target"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the memory/overhead field of the given VM.
func (_class AsyncVMClass) GetMemoryOverhead(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_memory_overhead"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the affinity field of the given VM.
func (_class AsyncVMClass) GetAffinity(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_affinity"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the resident_on field of the given VM.
func (_class AsyncVMClass) GetResidentOn(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_resident_on"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the suspend_VDI field of the given VM.
func (_class AsyncVMClass) GetSuspendVDI(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_suspend_VDI"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the is_a_template field of the given VM.
func (_class AsyncVMClass) GetIsATemplate(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_is_a_template"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the user_version field of the given VM.
func (_class AsyncVMClass) GetUserVersion(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_user_version"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the name/description field of the given VM.
func (_class AsyncVMClass) GetNameDescription(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_name_description"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the name/label field of the given VM.
func (_class AsyncVMClass) GetNameLabel(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the power_state field of the given VM.
func (_class AsyncVMClass) GetPowerState(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_power_state"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the current_operations field of the given VM.
func (_class AsyncVMClass) GetCurrentOperations(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_current_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the allowed_operations field of the given VM.
func (_class AsyncVMClass) GetAllowedOperations(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_allowed_operations"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get the uuid field of the given VM.
func (_class AsyncVMClass) GetUUID(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get all the VM instances with the given label.
func (_class AsyncVMClass) GetByNameLabel(sessionID SessionRef, label string) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_by_name_label"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_labelArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "label"), label)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _labelArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Destroy the specified VM.  The VM is completely removed from the system.  This function can only be called when the VM is in the Halted State.
func (_class AsyncVMClass) Destroy(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.destroy"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Create a new VM instance, and return its handle.
// The constructor args are: name_label, name_description, user_version*, is_a_template*, affinity*, memory_target, memory_static_max*, memory_dynamic_max*, memory_dynamic_min*, memory_static_min*, VCPUs_params*, VCPUs_max*, VCPUs_at_startup*, actions_after_shutdown*, actions_after_reboot*, actions_after_crash*, PV_bootloader*, PV_kernel*, PV_ramdisk*, PV_args*, PV_bootloader_args*, PV_legacy_args*, HVM_boot_policy*, HVM_boot_params*, HVM_shadow_multiplier, platform*, PCI_bus*, other_config*, recommendations*, xenstore_data, ha_always_run, ha_restart_priority, tags, blocked_operations, protection_policy, is_snapshot_from_vmpp, appliance, start_delay, shutdown_delay, order, suspend_SR, version, generation_id, hardware_platform_version (* = non-optional).
func (_class AsyncVMClass) Create(sessionID SessionRef, args VMRecord) (_retval TaskRef, _err error) {
	_method := "Async.VM.create"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_argsArg, _err := convertVMRecordToXen(fmt.Sprintf("%s(%s)", _method, "args"), args)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _argsArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get a reference to the VM instance with the specified UUID.
func (_class AsyncVMClass) GetByUUID(sessionID SessionRef, uuid string) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_by_uuid"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_uuidArg, _err := convertStringToXen(fmt.Sprintf("%s(%s)", _method, "uuid"), uuid)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _uuidArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}

// Get a record containing the current state of the given VM.
func (_class AsyncVMClass) GetRecord(sessionID SessionRef, self VMRef) (_retval TaskRef, _err error) {
	_method := "Async.VM.get_record"
	_sessionIDArg, _err := convertSessionRefToXen(fmt.Sprintf("%s(%s)", _method, "session_id"), sessionID)
	if _err != nil {
		return
	}
	_selfArg, _err := convertVMRefToXen(fmt.Sprintf("%s(%s)", _method, "self"), self)
	if _err != nil {
		return
	}
	_result, _err := _class.client.APICall(_method, _sessionIDArg, _selfArg)
	if _err != nil {
		return
	}
	_retval, _err = convertTaskRefToGo(_method + " -> ", _result.Value)
	return
}
