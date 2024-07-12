package targets

import (
	"time"
)

type BackupSource struct {
	Id           int       `json:"id,omitempty"`
	Pit          bool      `json:"pit,omitempty"`
	PitTimestamp time.Time `json:"pitTimestamp,omitempty"`
	ServerId     int       `json:"serverId,omitempty"`
	Type         string    `json:"type,omitempty"`
}
