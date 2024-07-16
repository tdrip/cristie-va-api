package etesting

import (
	"time"
)

type EnhancedTestingDiscovery struct {
	Address    string    `json:"address,omitempty"`
	Id         int       `json:"id,omitempty"`
	LastUpdate time.Time `json:"lastUpdate,omitempty"`
	MacAddress string    `json:"macAddress,omitempty"`
}
