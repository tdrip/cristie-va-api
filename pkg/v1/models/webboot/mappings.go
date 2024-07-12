package webboot

type Mapping struct {
	BiosUuid     string   `json:"bios_uuid,omitempty"`
	ConfigId     int      `json:"configId,omitempty"`
	DnsAddrs     []string `json:"dns_addrs,omitempty"`
	DrNetworking bool     `json:"drNetworking,omitempty"`
	Gateway      string   `json:"gateway,omitempty"`
	Id           int      `json:"id,omitempty"`
	Ip           string   `json:"ip,omitempty"`
	Mac          string   `json:"mac,omitempty"`
	Netmask      string   `json:"netmask,omitempty"`
	Os           string   `json:"os,omitempty"`
	Product      string   `json:"product,omitempty"`
}

type Mappings struct {
	Bootmap     []Mapping   `json:"bootmap,omitempty"`
}
