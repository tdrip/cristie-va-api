package orchestration

import (
	"time"

	targets "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

type BlockReportConfig struct {
	BlockId    int      `json:"blockId,omitempty"`
	Id         int      `json:"id,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
}

type Block struct {
	Description      string                 `json:"description,omitempty"`
	Finished         *time.Time             `json:"finished,omitempty"`
	Id               *int                   `json:"id"` // always required
	JobId            int                    `json:"jobId,omitempty"`
	Name             string                 `json:"name,omitempty"`
	ReportConfig     *BlockReportConfig     `json:"reportConfig,omitempty"`
	Reusable         bool                   `json:"reusable"`      // always required
	RunReport        bool                   `json:"runReport"`     // always required
	SelectedStage    *int                   `json:"selectedStage"` // always required
	SourceTargetList []targets.SourceTarget `json:"sourceTargetList,omitempty"`
	StageId          int                    `json:"stageId,omitempty"`
	Started          *time.Time             `json:"started,omitempty"`
	Status           int                    `json:"status,omitempty"` // string and INT?
	Type             int                    `json:"type"`             // always required
}
