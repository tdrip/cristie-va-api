package networks

type Binding struct {
	Address  string `json:"address,omitempty"`
	ConfigId int    `json:"configId,omitempty"`
	Id       int    `json:"id,omitempty"`
	Mac      string `json:"mac,omitempty"`
	ZoneId   int    `json:"zoneId,omitempty"`
}