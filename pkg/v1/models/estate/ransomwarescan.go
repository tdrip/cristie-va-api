package estate

import (
	"time"
)

type RansomwareScan struct {
	Date    *time.Time         `json:"date,omitempty"`
	Id      int                `json:"id,omitempty"`
	Metric  float64            `json:"metric,omitempty"`
	Metrics []RansomwareMetric `json:"metrics,omitempty"`
}

type RansomwareMetric struct {
	Cause      string  `json:"cause,omitempty"`
	Identifier string  `json:"identifier,omitempty"`
	Metric     float64 `json:"metric,omitempty"`
	Type       string  `json:"type,omitempty"`
}
