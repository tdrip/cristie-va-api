package orchestration

type Job struct {
	CurrentStage  string  `json:"currentStage,omitempty"`
	FailureOption string  `json:"failureOption,omitempty"`
	Id            int     `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	Stages        []Stage `json:"stages,omitempty"`
	Status        string  `json:"status,omitempty"`
}
