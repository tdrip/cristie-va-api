package networks

type ForwardRule struct {
	ConfigId int    `json:"configId,omitempty"`
	Id       int    `json:"id,omitempty"`
	Iport    int    `json:"iport,omitempty"`
	Oaddress string `json:"oaddress,omitempty"`
	Oport    int    `json:"oport,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}
