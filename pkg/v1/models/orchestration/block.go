package orchestration

import (
	"time"
)

type BlockReportConfig struct {
	BlockId    int32    `json:"blockId,omitempty"`
	Id         int32    `json:"id,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
}

type Block struct {
	Description  string             `json:"description,omitempty"`
	Finished     time.Time          `json:"finished,omitempty"`
	Id           int32              `json:"id,omitempty"`
	JobId        int32              `json:"jobId,omitempty"`
	Name         string             `json:"name,omitempty"`
	ReportConfig *BlockReportConfig `json:"reportConfig,omitempty"`
	Reusable     bool               `json:"reusable,omitempty"`
	RunReport    bool               `json:"runReport,omitempty"`
	StageId      int32              `json:"stageId,omitempty"`
	Started      time.Time          `json:"started,omitempty"`
	Status       string             `json:"status,omitempty"`
	Type_        string             `json:"type,omitempty"`
}
