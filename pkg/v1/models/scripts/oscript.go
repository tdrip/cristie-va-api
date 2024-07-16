package scripts

type OrchestrationScript struct {
	Contents string `json:"contents,omitempty"`
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Platform string `json:"platform,omitempty"`
	Type     string `json:"type,omitempty"`
}
