package webboot

type Mapping struct {
	BiosUuid     string `json:"bios_uuid,omitempty"`
	ConfigId     *int   `json:"configId,omitempty"`
	DHCP         bool   `json:"dhcp"`
	DnsAddrs     string `json:"dns_addrs"`
	DrNetworking bool   `json:"drNetworking"`
	Gateway      string `json:"gateway,omitempty"`
	Id           int    `json:"id,omitempty"`
	InsecureTls  bool   `json:"insecure_tls"`
	InsecureSSL  bool   `json:"insecure_ssl"` //missing from swagger
	Ip           string `json:"ip,omitempty"`
	Mac          string `json:"mac"`
	Netmask      string `json:"netmask,omitempty"`
	Os           string `json:"os,omitempty"`
	Product      string `json:"product,omitempty"`
}

type Mappings struct {
	Bootmap []Mapping `json:"bootmap,omitempty"`
}
