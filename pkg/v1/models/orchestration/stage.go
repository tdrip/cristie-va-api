package orchestration

type Stage struct {
	Blocks  []RBlock `json:"blocks,omitempty"`
	Enabled bool     `json:"enabled,omitempty"`
	Id      int      `json:"id,omitempty"`
	JobId   int      `json:"jobId,omitempty"`
	Name    string   `json:"name,omitempty"`
	Status  string   `json:"status,omitempty"`
}
