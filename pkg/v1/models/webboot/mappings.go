package webboot

type Mapping struct {
	BiosUuid     string   `json:"bios_uuid"`
	ConfigId     *int     `json:"configId,omitempty"`
	DHCP         bool     `json:"dhcp,omitempty"`
	DnsAddrs     string   `json:"dns_addrs,omitempty"`
	DrNetworking bool     `json:"drNetworking,omitempty"`
	Gateway      string   `json:"gateway,omitempty"`
	Id           int      `json:"id,omitempty"`
	InsecureTls  bool     `json:"insecure_tls,omitempty"`
	InsecureSSL  bool     `json:"insecure_ssl,omitempty"` //missing from swagger
	Ip           string   `json:"ip,omitempty"`
	Mac          string   `json:"mac"` 
	Netmask      string   `json:"netmask,omitempty"`
	Os           string   `json:"os,omitempty"`
	Product      string   `json:"product,omitempty"`
}

type Mappings struct {
	Bootmap     []Mapping   `json:"bootmap,omitempty"`
}
