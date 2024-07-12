package backupservers

import (
	"encoding/json"
	"errors"
	"fmt"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
)

const (
	UriBackupServersClient  = "v1/recovery/backup-servers/client/%d"
	UriBackupServersClients = "v1/recovery/backup-servers/%d/clients"
)

func GetClients(crs *cls.Client, backupid int) ([]bs.Client, error) {
	results := []bs.Client{}
	if !crs.Session.HasToken() {
		return results, errors.New("get clients failed - token missing session")
	}

	bytes, res, err := crs.Session.Get(fmt.Sprintf(UriBackupServersClients, backupid))

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &results)
		return results, err
	}
	error_body, nerr := client.GetError(bytes)
	return results, fmt.Errorf("get clients failed - errors: %v %v %v", err, error_body, nerr)
}

func GetClientDetails(crs *cls.Client, clientid int) (bs.Client, error) {
	result := bs.Client{}
	if !crs.Session.HasToken() {
		return result, errors.New("get client details failed - token missing session")
	}

	bytes, res, err := crs.Session.Get(fmt.Sprintf(UriBackupServersClient, clientid))

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("get client details failed - errors: %v %v %v", err, error_body, nerr)
}
