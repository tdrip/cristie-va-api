package orchestration

import (
	"encoding/json"
	"errors"
	"fmt"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	orch "github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
)

const (
	UriRecoveryOrchestrationJobsStages = "v1/recovery/orchestration/jobs/%d/stages"
)

func CreateStage(crs *cls.Client, name string, jobid int) (orch.Stage, error) {
	result := orch.Stage{}
	if !crs.Session.HasToken() {
		return result, errors.New("stage creation failed - token missin session")
	}

	request := orch.Stage{Name: name}
	bytes, res, err := crs.Session.PostBody(fmt.Sprintf(UriRecoveryOrchestrationJobsStages, jobid), &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
	error_body, nerr := client.GetError(bytes, res)
	return result, fmt.Errorf("stage creation failed - errors: %v %v %v", err, error_body, nerr)
}
