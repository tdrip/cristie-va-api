package orchestration

import (
	"encoding/json"
	"errors"
	"fmt"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	models "github.com/tdrip/cristie-va-api/pkg/v1/models"
	orch "github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
)

const (
	UriRecoveryOrchestration = "v1/recovery/orchestration/"
	FailureOption            = "PAUSE"
)

func CreateJob(crs *cls.Client, name string) (models.Event, error) {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("job creation failed - token missin session")
	}

	request := orch.Job{Name: name, FailureOption: FailureOption}
	bytes, res, err := crs.Session.PostBody(UriRecoveryOrchestration, &request)

	if err == nil && res != nil && utils.RequestISsUCCESSFUL(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return res, err
	}

	return result, fmt.Errorf("job creation failed with errors: %v", err)

}
