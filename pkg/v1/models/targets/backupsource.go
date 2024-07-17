package targets

import (
	"time"
)

type BackupSource struct {
	BackupType   string    `json:"backupType,omitempty"` //missing from swagger
	Hostname     string    `json:"hostName,omitempty"`   //missing from swagger
	HostId       string    `json:"hostId,omitempty"`     //missing from swagger
	Id           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"` //missing from swagger
	Pit          bool      `json:"pit,omitempty"`
	PitTimestamp time.Time `json:"pitTimestamp,omitempty"`
	ServerId     int       `json:"serverId,omitempty"`
	Type         string    `json:"type,omitempty"`
}
