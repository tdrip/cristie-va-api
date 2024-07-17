package consts

// source targets
const (
	SourceTarget_Type_RBMR           = "RBMR"
	SourceTarget_VMHostType_Standard = "STANDARD"
)

// backup types
const (
	Source_BackupType_FileBased  = "FILE_BASED"  // Fileset
	Source_BackupType_BlockBased = "BLOCK_BASED" // Volume Group
)

const (
	BackupSource_Default_Name = "Unknown Type"
)

const (
	Network_DHCP_Name = "DHCP"
)

const (
	TaskType_Recovery = 0
)

const (
	Action_Leave  = "LEAVE"
	Action_Reboot = "REBOOT"
)
