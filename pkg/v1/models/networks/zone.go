package networks

type Zone struct {
	Allowed   []Port        `json:"allowed,omitempty"`
	Cidr      *Address      `json:"cidr,omitempty"`
	Devices   []int         `json:"devices,omitempty"`
	Dhcp      *DhcpConfig   `json:"dhcp,omitempty"`
	Forwards  []ForwardRule `json:"forwards,omitempty"`
	GatewayId int           `json:"gatewayId,omitempty"`
	Id        int           `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
}
