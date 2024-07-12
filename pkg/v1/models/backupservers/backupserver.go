package backupservers

import (
	"time"
)

type BackupServer struct {
	BackupClients []Client  `json:"backupClients,omitempty"`
	Host          string    `json:"host,omitempty"`
	Id            int       `json:"id,omitempty"`
	LastUpdate    time.Time `json:"lastUpdate,omitempty"`
	Legacy        bool      `json:"legacy,omitempty"`
	Name          string    `json:"name,omitempty"`
	OtpSecret     string    `json:"otpSecret,omitempty"`
	Password      string    `json:"password,omitempty"`
	ServerVersion string    `json:"serverVersion,omitempty"`
	Type          string    `json:"type,omitempty"`
	UserName      string    `json:"userName,omitempty"`
}
