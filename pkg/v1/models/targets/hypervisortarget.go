package targets

type HypervisorTarget struct {
	HostId       int32  `json:"hostId,omitempty"`
	Id           int32  `json:"id,omitempty"`
	Iso          string `json:"iso,omitempty"`
	IsoDatastore string `json:"isoDatastore,omitempty"`
	Nic          string `json:"nic,omitempty"`
	Type         string `json:"type,omitempty"`
}
