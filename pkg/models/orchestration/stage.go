package models

type OrchestrationStage struct {
	Blocks  []OrchestrationBlock `json:"blocks,omitempty"`
	Enabled bool                 `json:"enabled,omitempty"`
	Id      int32                `json:"id,omitempty"`
	JobId   int32                `json:"jobId,omitempty"`
	Name    string               `json:"name,omitempty"`
	Status  string               `json:"status,omitempty"`
}
