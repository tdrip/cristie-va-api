package networks

type Device struct {
	Addresses    []Address `json:"addresses,omitempty"`
	Dhcp         bool      `json:"dhcp,omitempty"`
	FriendlyName string    `json:"friendlyName,omitempty"`
	GatewayId    int       `json:"gatewayId,omitempty"`
	Id           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Parent       string    `json:"parent,omitempty"`
	Type         string    `json:"type,omitempty"`
	VlanId       int       `json:"vlan_id,omitempty"`
}