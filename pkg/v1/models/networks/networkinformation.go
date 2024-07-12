package networks

type NetworkInformation struct {
	Description   string `json:"description,omitempty"`
	Dhcp          bool   `json:"dhcp,omitempty"`
	DhcpVA        bool   `json:"dhcpVA,omitempty"`
	Gateway       string `json:"gateway,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	Id            int    `json:"id,omitempty"`
	InternalIndex int    `json:"internalIndex,omitempty"`
	IpAddr        string `json:"ipAddr,omitempty"`
	MacAddr       string `json:"macAddr,omitempty"`
	Name          string `json:"name,omitempty"`
	NameServer    string `json:"nameServer,omitempty"`
	NetMask       string `json:"netMask,omitempty"`
	Network       string `json:"network,omitempty"`
	NicId         int    `json:"nicId,omitempty"`
	UserAdded     bool   `json:"userAdded,omitempty"`
	Virtual       bool   `json:"virtual,omitempty"`
	VlanId        int    `json:"vlanId,omitempty"`
}
