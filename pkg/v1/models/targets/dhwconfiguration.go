package targets

type DhwConfiguration struct {
	Id       int   `json:"id,omitempty"`
	Password string `json:"password,omitempty"`
	Path     string `json:"path,omitempty"`
	UserName string `json:"userName,omitempty"`
}