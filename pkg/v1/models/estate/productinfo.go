package estate

import (
	"time"
)

type ProductInfo struct {
	ActivationCode       string    `json:"activationCode,omitempty"`
	Bmr                  bool      `json:"bmr,omitempty"`
	ContractId           int       `json:"contractId,omitempty"`
	LicExpiry            time.Time `json:"licExpiry,omitempty"`
	LicType              string    `json:"licType,omitempty"`
	Licensed             bool      `json:"licensed,omitempty"`
	NextAvailableVersion string    `json:"nextAvailableVersion,omitempty"`
	RentalSync           bool      `json:"rentalSync,omitempty"`
	Signature            string    `json:"signature,omitempty"`
	Tokens               []int     `json:"tokens,omitempty"`
	Type_                string    `json:"type,omitempty"`
	Version              string    `json:"version,omitempty"`
}
