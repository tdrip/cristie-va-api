package estate

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	models "github.com/tdrip/cristie-va-api/pkg/v1/models"
	est "github.com/tdrip/cristie-va-api/pkg/v1/models/estate"
)

const (
	UriEstateMachines = "v1/control/estate/machines"
	UriEstateMachine  = "v1/control/estate/machines/%s"
)

func GetMachines(crs *cls.Client) (est.SystemsList, error) {
	result := est.SystemsList{}
	if !crs.Session.HasToken() {
		return result, errors.New("get machines failed - token missing session")
	}

	bytes, res, err := crs.Session.Get(UriEstateMachines)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("get machines failed - errors: %v %v %v", err, error_body, nerr)
}

func GetMachinesByMac(crs *cls.Client, macadd string) (est.SystemsList, error) {
	result := est.SystemsList{}
	if !crs.Session.HasToken() {
		return result, errors.New("get machines by mac failed - token missing session")
	}

	if len(macadd) == 0 {
		return result, errors.New("get machines by mac failed - no mac provided")
	}

	params := url.Values{}
	params.Add("macAddresss", macadd)

	bytes, res, err := crs.Session.Get(fmt.Sprintf("%s?%s", UriEstateMachines, params.Encode()))

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("get machines by mac failed - errors: %v %v %v", err, error_body, nerr)
}

func DeleteMachine(crs *cls.Client, uuid string) (models.Event, error) {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("delete machine failed - token missing session")
	}

	if len(uuid) == 0 {
		return result, errors.New("delete machine failed - no mac provided")
	}

	bytes, res, err := crs.Session.Delete(fmt.Sprintf(UriEstateMachine, url.QueryEscape(uuid)))

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("delete machine failed - errors: %v %v %v", err, error_body, nerr)
}
