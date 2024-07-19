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
	UriBackupServers = "v1/recovery/backup-servers"
)

func GetServers(crs *cls.Client) (bs.BackupServers, error) {
	results := bs.BackupServers{}
	if !crs.Session.HasToken() {
		return results, errors.New("get servers failed - token missing session")
	}

	bytes, res, err := crs.Session.Get(UriBackupServers)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &results)
		}
		return results, err
	}
	error_body, nerr := client.GetError(bytes)
	return results, fmt.Errorf("get servers failed - errors: %v %v %v", err, error_body, nerr)
}
