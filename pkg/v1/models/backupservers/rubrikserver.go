package backupservers

import (
	"time"
)

type RubrikServer struct {
	CloudAddress  string         `json:"cloudAddress,omitempty"`
	CloudId       int32          `json:"cloudId,omitempty"`
	CloudUuid     string         `json:"cloudUuid,omitempty"`
	SmbPassword   string         `json:"smbPassword,omitempty"`
	SmbUsername   string         `json:"smbUsername,omitempty"`
	BackupClients []Client       `json:"backupClients,omitempty"`
	Host          string         `json:"host,omitempty"`
	Id            int32          `json:"id,omitempty"`
	LastUpdate    time.Time      `json:"lastUpdate,omitempty"`
	Legacy        bool           `json:"legacy,omitempty"`
	Name          string         `json:"name,omitempty"`
	OtpSecret     string         `json:"otpSecret,omitempty"`
	Password      string         `json:"password,omitempty"`
	ServerVersion string         `json:"serverVersion,omitempty"`
	Type          string         `json:"type,omitempty"`
	UserName      string         `json:"userName,omitempty"`
}
