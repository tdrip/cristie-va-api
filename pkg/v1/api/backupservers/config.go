package backupservers

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
	UriRecoveryBackupServerConfig = "v1/recovery/backup-servers/getconfig"
)

func GetConfigDetails(crs *cls.Client,backupserverid int, clientname string, backuptype string) (models.Event, error) {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("backup server - get config details failed - token missing session")
	}

	request := orch.Job{}
	bytes, res, err := crs.Session.PostBody(UriRecoveryBackupServerConfig, &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("backup server - get config details failed- errors: %v %v %v", err, error_body, nerr)
}