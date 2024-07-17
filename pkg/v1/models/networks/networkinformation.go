package networks

type NetworkInformation struct {
	Description   string `json:"description,omitempty"`
	Dhcp          bool   `json:"dhcp"`   //always required
	DhcpVA        bool   `json:"dhcpVA"` //always required
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
	UserAdded     bool   `json:"userAdded"` //always required
	Virtual       bool   `json:"virtual"`   //always required
	VlanId        int    `json:"vlanId,omitempty"`
}
