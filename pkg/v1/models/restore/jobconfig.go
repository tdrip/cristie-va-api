package restore

import (
	"time"
)

type JobConfig struct {
	BackupId      int       `json:"backupId,omitempty"`
	BackupVersion time.Time `json:"backupVersion,omitempty"`
	CreateVm      bool      `json:"create_vm,omitempty"`
	EventUuid     string    `json:"eventUuid,omitempty"`
	HypervisorId  int       `json:"hypervisorId,omitempty"`
	Id            int       `json:"id,omitempty"`
	Manual        bool      `json:"manual,omitempty"`
	Offsite       bool      `json:"offsite,omitempty"`
	Platform      string    `json:"platform,omitempty"`
	Product       string    `json:"product,omitempty"`
	ProxyAddr     string    `json:"proxyAddr,omitempty"`
	ProxyUuid     string    `json:"proxyUuid,omitempty"`
	ScheduleUuid  string    `json:"scheduleUuid,omitempty"`
	Source        string    `json:"source,omitempty"`
	VmName        string    `json:"vm_name,omitempty"`
	VmNetworkIds  []string  `json:"vm_network_ids,omitempty"`
	VmParentId    string    `json:"vm_parent_id,omitempty"`
	VmStorageId   string    `json:"vm_storage_id,omitempty"`
}
