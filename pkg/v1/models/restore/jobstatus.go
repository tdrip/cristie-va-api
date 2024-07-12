package restore

import (
	"time"
)

type JobStatus struct {
	Address       string    `json:"address,omitempty"`
	BackupId      int.      `json:"backupId,omitempty"`
	BackupVersion time.Time `json:"backupVersion,omitempty"`
	Details       string    `json:"details,omitempty"`
	EndTime       time.Time `json:"end_time,omitempty"`
	Id            int       `json:"id,omitempty"`
	Platform      string    `json:"platform,omitempty"`
	Product       string    `json:"product,omitempty"`
	Progress      float64   `json:"progress,omitempty"`
	ProxyUuid     string    `json:"proxyUuid,omitempty"`
	Source        string    `json:"source,omitempty"`
	SourceName    string    `json:"sourceName,omitempty"`
	StartTime     time.Time `json:"start_time,omitempty"`
	Status        string    `json:"status,omitempty"`
	Target        string    `json:"target,omitempty"`
}
