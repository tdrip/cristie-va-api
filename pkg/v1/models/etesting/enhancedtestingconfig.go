package etesting

type EnhancedTestingConfig struct {
	Direct            bool                      `json:"direct,omitempty"`
	Hypervisor        *EnhancedTestingDiscovery `json:"hypervisor,omitempty"`
	HypervisorId      int                       `json:"hypervisorId,omitempty"`
	Id                int                       `json:"id,omitempty"`
	Iso               string                    `json:"iso,omitempty"`
	SnapshotNetworkId int                       `json:"snapshotNetworkId,omitempty"`
	Testing           bool                      `json:"testing,omitempty"`
	Valid             bool                      `json:"valid,omitempty"`
	Vlan              int                       `json:"vlan,omitempty"`
}