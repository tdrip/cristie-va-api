package helpers

import (
	cls "github.com/tdrip/apiclient/pkg/v1/client"
	"github.com/tdrip/cristie-va-api/pkg/v1/api/estate"
)

func ClearEstate(crs *cls.Client, MacAddress string) error {
	// better clear the estate
	machines, err := estate.GetMachinesByMac(crs, MacAddress)
	if err != nil {
		return err
	}

	for _, system := range machines.Machines {
		_, err := estate.DeleteMachine(crs, system.Uuid)
		if err != nil {
			return err
		}
	}

	return nil
}
