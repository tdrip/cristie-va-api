package models

type ErrorInformation struct {
	Description    string   `json:"description,omitempty"`
	Field          string   `json:"field,omitempty"`
	Format         string   `json:"format,omitempty"`
	HttpErrorCode  int      `json:"httpErrorCode,omitempty"`
	Message        string   `json:"message,omitempty"`
	PossibleValues []string `json:"possibleValues,omitempty"`
	Resource       string   `json:"resource,omitempty"`
	Value          string   `json:"value,omitempty"`
}

type Exception struct {
	ErrorCode        int32  `json:"errorCode,omitempty"`
	Information      string `json:"information,omitempty"`
	LocalizedMessage string `json:"localizedMessage,omitempty"`
	Message          string `json:"message,omitempty"`
}
