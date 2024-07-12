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
	UriRecoveryOrchestrationJobsStagesBlocks = "v1/recovery/orchestration/jobs/%d/stages/%d/blocks"
)

func CreateBlock(crs *cls.Client, name string, jobid int, stagid int) (orch.Block, error) {
	result := orch.Block{}
	if !crs.Session.HasToken() {
		return result, errors.New("task creation failed - token missin session")
	}

	request := orch.Block{Name: name, JobId: jobid, StageId: stagid}
	bytes, res, err := crs.Session.PostBody(fmt.Sprintf(UriRecoveryOrchestrationJobsStagesBlocks, jobid, stagid), &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("task creation failed - errors: %v %v %v", err, error_body, nerr)
}