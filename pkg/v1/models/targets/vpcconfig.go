package targets

import (
	networks "github.com/tdrip/cristie-va-api/pkg/v1/models/networks"
)

type VpcConfig struct {
	DetachOnReboot bool             `json:"detachOnReboot,omitempty"`
	Device         *networks.Device `json:"device,omitempty"`
	Id             int32            `json:"id,omitempty"`
	Zone           *Znetworks.Zone  `json:"zone,omitempty"`
}