package targets

import (
	disks "github.com/tdrip/cristie-va-api/pkg/v1/models/disks"
	networks "github.com/tdrip/cristie-va-api/pkg/v1/models/networks"
)

type VmConfiguration struct {
	Arch            string                        `json:"arch,omitempty"`
	CpuCount        int                           `json:"cpuCount,omitempty"`
	DiskNo          int                           `json:"diskNo,omitempty"`
	Disks           []disks.StandardDisk          `json:"disks,omitempty"`
	Efi             bool                          `json:"efi,omitempty"`
	EfiSys          string                        `json:"efiSys,omitempty"`
	Hostname        string                        `json:"hostname,omitempty"`
	IsHourlyBilling bool                          `json:"isHourlyBilling,omitempty"`
	NetWorkInfo     []networks.NetworkInformation `json:"netWorkInfo,omitempty"`
	Os              string                        `json:"os,omitempty"`
	OsFamily        string                        `json:"osFamily,omitempty"`
	RamSizeInMB     int64                         `json:"ramSizeInMB,omitempty"`
	StoragePools    []disks.StoragePool           `json:"storagePools,omitempty"`
}
