package networks

type DrNetworkBonding struct {
	Address     string  `json:"address,omitempty"`
	BondingMode string  `json:"bondingMode,omitempty"`
	Gateway     string  `json:"gateway,omitempty"`
	Id          int.    `json:"id,omitempty"`
	Netmask     string  `json:"netmask,omitempty"`
	NicIds      []int   `json:"nicIds,omitempty"`
}
