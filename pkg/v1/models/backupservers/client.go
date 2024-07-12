package backupservers

import (
	"time"
)

type Client struct {
	BackupType       string    `json:"backupType,omitempty"`
	CristieProtected bool      `json:"cristieProtected,omitempty"`
	DisrecDate       time.Time `json:"disrecDate,omitempty"`
	Domain           string    `json:"domain,omitempty"`
	HostId           string    `json:"hostId,omitempty"`
	Id               int32     `json:"id,omitempty"`
	Name             string    `json:"name,omitempty"`
	Platform         string    `json:"platform,omitempty"`
	PointInTimeDates []Pit     `json:"pointInTimeDates,omitempty"`
	SessionSecurity  string    `json:"sessionSecurity,omitempty"`
	SlaDomains       []string  `json:"slaDomains,omitempty"`
	TenantId         string    `json:"tenantId,omitempty"`
}
