package targets

type HypervisorTarget struct {
	HostId       int    `json:"hostId,omitempty"`
	Id           int    `json:"id,omitempty"`
	Iso          string `json:"iso,omitempty"`
	IsoDatastore string `json:"isoDatastore,omitempty"`
	Nic          string `json:"nic,omitempty"`
	Type         string `json:"type,omitempty"`
}
