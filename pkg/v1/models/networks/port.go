package networks

type Port struct {
	Id       int  `json:"id,omitempty"`
	Port     int  `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}
