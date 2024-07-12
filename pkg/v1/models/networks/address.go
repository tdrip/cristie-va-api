package networks

type Address struct {
	Address string `json:"address,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Id      int    `json:"id,omitempty"`
	Prefix  int    `json:"prefix,omitempty"`
	Type    string `json:"type,omitempty"`
}
