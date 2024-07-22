package orchestration

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	models "github.com/tdrip/cristie-va-api/pkg/v1/models"
	orch "github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
)

const (
	UriRecoveryOrchestration       = "v1/recovery/orchestration/"
	UriRecoveryOrchestrationJobs   = "v1/recovery/orchestration/jobs"
	UriRecoveryOrchestrationJob    = "v1/recovery/orchestration/jobs/%d"
	UriRecoveryOrchestrationJobRun = "v1/recovery/orchestration/jobs/%d/run"

	FailureOption = "PAUSE"
)

func CreateJob(crs *cls.Client, name string) (models.Event, error) {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("job creation failed - token missin session")
	}

	request := orch.Job{Name: name, FailureOption: FailureOption}
	bytes, res, err := crs.Session.PostBody(UriRecoveryOrchestration, &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("job creation failed - errors: %v %v %v", err, error_body, nerr)
}

func GetJobs(crs *cls.Client) ([]orch.Job, error) {
	result := []orch.Job{}
	if !crs.Session.HasToken() {
		return result, errors.New("get jobs failed - token missin session")
	}

	bytes, res, err := crs.Session.Get(UriRecoveryOrchestrationJobs)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("get jobs failed - errors: %v %v %v", err, error_body, nerr)
}

func GetJob(crs *cls.Client, jobid int) (orch.Job, error) {
	result := orch.Job{}
	if !crs.Session.HasToken() {
		return result, errors.New("get job failed - token missin session")
	}

	bytes, res, err := crs.Session.Get(fmt.Sprintf(UriRecoveryOrchestrationJob, jobid))

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("get job failed - errors: %v %v %v", err, error_body, nerr)
}

func RunJob(crs *cls.Client, jobid int, stagid int, blockid int) (models.Event, error) {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("run job failed - token missin session")
	}

	uri := fmt.Sprintf(UriRecoveryOrchestrationJobRun, jobid)

	if stagid > 0 || blockid > 0 {
		params := url.Values{}
		if stagid > 0 {
			params.Add("stageId", fmt.Sprintf("%d", stagid))
		}
		if blockid > 0 {
			params.Add("blockId", fmt.Sprintf("%d", blockid))
		}
		uri = fmt.Sprintf("%s?%s", uri, params.Encode())
	}

	bytes, res, err := crs.Session.PostBody(uri, nil)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("run job failed - errors: %v %v %v", err, error_body, nerr)
}
