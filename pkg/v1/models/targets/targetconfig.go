package targets

import (
	disks "github.com/tdrip/cristie-va-api/pkg/v1/models/disks"
	scripts "github.com/tdrip/cristie-va-api/pkg/v1/models/scripts"
)

type TargetConfig struct {
	Arch                     string                    `json:"arch,omitempty"`
	CpuCount                 int                       `json:"cpuCount,omitempty"`
	DisableMultipath         bool                      `json:"disableMultipath,omitempty"`
	Efi                      bool                      `json:"efi,omitempty"`
	EnhancedTesting          bool                      `json:"enhancedTesting,omitempty"`
	FirmwareOpt              string                    `json:"firmwareOpt,omitempty"`
	HardDisksInGB            []disks.DiskType          `json:"hardDisksInGB,omitempty"`
	Hostname                 string                    `json:"hostname,omitempty"`
	Id                       int                       `json:"id,omitempty"`
	LocalDisksOnly           bool                      `json:"localDisksOnly,omitempty"`
	MapDisk                  bool                      `json:"mapDisk,omitempty"`
	Os                       string                    `json:"os,omitempty"`
	Osfamily                 string                    `json:"osfamily,omitempty"`
	PostRecoveryCredentialId int                       `json:"postRecoveryCredentialId,omitempty"`
	PostRecoveryPassword     string                    `json:"postRecoveryPassword,omitempty"`
	PostRecoverySSHKey       string                    `json:"postRecoverySSHKey,omitempty"`
	PostRecoveryUsername     string                    `json:"postRecoveryUsername,omitempty"`
	PostScriptCmd            *scripts.OrchestrationScript `json:"postScriptCmd,omitempty"`
	RamSizeInMB              int                       `json:"ramSizeInMB,omitempty"`
	RecoveryFailureAction    string                    `json:"recoveryFailureAction,omitempty"`
	RecoverySuccessAction    string                    `json:"recoverySuccessAction,omitempty"`
	SkipMultipath            bool                      `json:"skipMultipath,omitempty"`
	SysServices              string                    `json:"sysServices,omitempty"`
	VerboseLogging           bool                      `json:"verboseLogging,omitempty"`
	VmName                   string                    `json:"vmName,omitempty"`
}
