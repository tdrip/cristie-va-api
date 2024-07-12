package recovery

import (
	"time"

	disks "github.com/tdrip/cristie-va-api/pkg/v1/models/disks"
)

type Status struct {
	Client           string           `json:"client,omitempty"`
	CurrentFilespace string           `json:"currentFilespace,omitempty"`
	Disks            []disks.DiskType `json:"disks,omitempty"`
	DurationInSec    int              `json:"durationInSec,omitempty"`
	EventUuid        string           `json:"eventUuid,omitempty"`
	FilesDone        int              `json:"filesDone,omitempty"`
	Id               int              `json:"id,omitempty"`
	IpAddress        string           `json:"ipAddress,omitempty"`
	IsOrchestration  bool             `json:"isOrchestration,omitempty"`
	Iso              string           `json:"iso,omitempty"`
	IsoVersion       string           `json:"isoVersion,omitempty"`
	LogName          string           `json:"logName,omitempty"`
	OperationStatus  string           `json:"operationStatus,omitempty"`
	Percentage       int              `json:"percentage,omitempty"`
	Platform         string           `json:"platform,omitempty"`
	Product          string           `json:"product,omitempty"`
	RebootStatus     string           `json:"rebootStatus,omitempty"`
	RecStatus        string           `json:"recStatus,omitempty"`
	RecoveryId       int              `json:"recoveryId,omitempty"`
	RestoreEndTime   time.Time        `json:"restoreEndTime,omitempty"`
	RestoreStartTime time.Time        `json:"restoreStartTime,omitempty"`
	Server           string           `json:"server,omitempty"`
	SizeDoneMB       int              `json:"sizeDoneMB,omitempty"`
	SpeedMBPS        float64          `json:"speedMBPS,omitempty"`
	Status           string           `json:"status,omitempty"`
	TotalFiles       int              `json:"totalFiles,omitempty"`
	TotalFilespaces  int              `json:"totalFilespaces,omitempty"`
	TotalSizeMB      int              `json:"totalSizeMB,omitempty"`
	VmName           string           `json:"vmName,omitempty"`
}
