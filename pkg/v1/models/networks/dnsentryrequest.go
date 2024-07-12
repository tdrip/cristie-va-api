package networks

type DnsEntryRequest struct {
	AddPTR           bool   `json:"addPTR,omitempty"`
	ConfigId         int    `json:"configId,omitempty"`
	Id               int    `json:"id,omitempty"`
	NewHostname      string `json:"newHostname,omitempty"`
	NewIP            string `json:"newIP,omitempty"`
	OriginalHostname string `json:"originalHostname,omitempty"`
	OriginalIP       string `json:"originalIP,omitempty"`
	PtrZone          string `json:"ptrZone,omitempty"`
	Reverse          bool   `json:"reverse,omitempty"`
	ServerId         int    `json:"serverId,omitempty"`
	Valid            bool   `json:"valid,omitempty"`
	ZoneName         string `json:"zoneName,omitempty"`
}
