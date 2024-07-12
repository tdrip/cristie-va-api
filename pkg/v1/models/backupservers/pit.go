package backupserver

import (
	"time"
)

type Pit struct {
	Date  time.Time `json:"date,omitempty"`
	Type  string    `json:"type,omitempty"`
}
