package orchestration

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
	UriRecoveryOrchestrationJobsStagesBlocks = "v1/recovery/orchestration/jobs/%d/stages/%d/blocks"
)

func CreateBlock(crs *cls.Client, request orch.CBlock) (models.Event, error) {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("task creation failed - token missin session")
	}

	bytes, res, err := crs.Session.PostBody(fmt.Sprintf(UriRecoveryOrchestrationJobsStagesBlocks, request.JobId, request.StageId), &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("task creation failed - errors: %v %v %v", err, error_body, nerr)
}
