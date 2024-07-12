package client

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
	UriUsersAuthPath      = "v1/users/auth/"
	UriUsersAuthSession   = "v1/users/session"    // <-- 4.9.1+
	UriUsersAuthSessionId = "v1/users/session_id" // <-- 4.8.3
)

func Login(crs *cls.Client, userid string, pwd string) error {
	result := models.Event{}
	if !crs.Session.HasToken() {
		return result, errors.New("job creation failed - token missin session")
	}

	request := orch.Job{Name: userid, FailureOption: userid}
	bytes, res, err := crs.Session.PostBody(UriUsersAuthPath, &request)

	if err == nil && res != nil && utils.RequestISsUCCESSFUL(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return res, err
	}

	return fmt.Errorf("job creation failed with errors: %v", err)

}
