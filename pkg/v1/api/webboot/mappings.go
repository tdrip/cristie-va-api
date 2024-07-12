package webboot

import (
	"encoding/json"
	"errors"
	"fmt"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	models "github.com/tdrip/cristie-va-api/pkg/v1/models"
	orch "github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
)

const (
	UriWebbootMap = "v1/webboot/map/"
)

func CreateDHCPMapping(crs *cls.Client, product string) (models.Event, error) {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("webboot mapping failed - token missing session")
	}

	request := orch.Job{Name: product}
	bytes, res, err := crs.Session.PostBody(UriWebbootMap, &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
	error_body, nerr := client.GetError(bytes, res)
	return result, fmt.Errorf("webboot mapping  - errors: %v %v %v", err, error_body, nerr)
}
