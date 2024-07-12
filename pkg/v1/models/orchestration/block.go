package orchestration

import (
	"time"
)

type BlockReportConfig struct {
	BlockId    int      `json:"blockId,omitempty"`
	Id         int      `json:"id,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
}

type Block struct {
	Description  string             `json:"description,omitempty"`
	Finished     time.Time          `json:"finished,omitempty"`
	Id           int                `json:"id,omitempty"`
	JobId        int                `json:"jobId,omitempty"`
	Name         string             `json:"name,omitempty"`
	ReportConfig *BlockReportConfig `json:"reportConfig,omitempty"`
	Reusable     bool               `json:"reusable,omitempty"`
	RunReport    bool               `json:"runReport,omitempty"`
	StageId      int                `json:"stageId,omitempty"`
	Started      time.Time          `json:"started,omitempty"`
	Status       string             `json:"status,omitempty"`
	Type         string             `json:"type,omitempty"`
}
