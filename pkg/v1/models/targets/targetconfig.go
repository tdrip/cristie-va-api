package targets

import (
	disks "github.com/tdrip/cristie-va-api/pkg/v1/models/disks"
	scripts "github.com/tdrip/cristie-va-api/pkg/v1/models/scripts"
)

type TargetConfig struct {
	Arch                     string                       `json:"arch,omitempty"`
	CpuCount                 int                          `json:"cpuCount,omitempty"`
	DisableMultipath         bool                         `json:"disableMultipath,omitempty"`
	Efi                      bool                         `json:"efi,omitempty"`
	EnhancedTesting          bool                         `json:"enhancedTesting,omitempty"`
	FirmwareOpt              string                       `json:"firmwareOpt,omitempty"`
	HardDisksInGB            []disks.DiskType             `json:"hardDisksInGB"` // null required
	Hostname                 string                       `json:"hostname,omitempty"`
	Id                       int                          `json:"id,omitempty"`
	ImportDHCP               bool                         `json:"importDHCP"` //missing from swagger
	LocalDisksOnly           bool                         `json:"localDisksOnly,omitempty"`
	MapDisk                  bool                         `json:"mapDisk,omitempty"`
	DiskMapList              []string                     `json:"diskMapList"`                //missing from swagger
	OriginalHostname         string                       `json:"originalHostname,omitempty"` //missing from swagger
	Os                       string                       `json:"os,omitempty"`
	Osfamily                 string                       `json:"osfamily,omitempty"`
	PostRecoveryCredentialId int                          `json:"postRecoveryCredentialId,omitempty"`
	PostRecoveryPassword     string                       `json:"postRecoveryPassword,omitempty"`
	PostRecoverySSHKey       string                       `json:"postRecoverySSHKey,omitempty"`
	PostRecoveryUsername     string                       `json:"postRecoveryUsername,omitempty"`
	PostScriptCmd            *scripts.OrchestrationScript `json:"postScriptCmd"` // null required
	RamSizeInMB              int64                        `json:"ramSizeInMB,omitempty"`
	RecoveryFailureAction    string                       `json:"recoveryFailureAction,omitempty"`
	RecoverySuccessAction    string                       `json:"recoverySuccessAction,omitempty"`
	SkipMultipath            bool                         `json:"skipMultipath,omitempty"`
	SysServices              string                       `json:"sysServices,omitempty"`
	VerboseLogging           bool                         `json:"verboseLogging,omitempty"`
	VmName                   string                       `json:"vmName,omitempty"`
}
