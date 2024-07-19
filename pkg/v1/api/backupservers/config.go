package backupservers

import (
	"encoding/json"
	"errors"
	"fmt"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	tgts "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

const (
	UriRecoveryBackupServerConfig = "v1/recovery/backup-servers/getconfig"
)

func GetConfigDetails(crs *cls.Client, backupserverid int, clientname string, backuptype string) (tgts.VmConfiguration, error) {
	result := tgts.VmConfiguration{}
	if !crs.Session.HasToken() {
		return result, errors.New("backup server - get config details failed - token missing session")
	}

	request := bs.ConfigRequest{}
	request.ServerId = backupserverid
	request.ClientName = clientname
	request.BackupType = backuptype
	request.Password = "null"                // must be provided in this format - this is weird
	request.Username = "null"                // must be provided in this format - this is weird
	request.ClientDomain = "null"            // must be provided in this format - this is weird
	request.EncryptionPasswords = []string{} // needs these to be empty array

	bytes, res, err := crs.Session.PostBody(UriRecoveryBackupServerConfig, &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		if len(bytes) > 0 {
			err = json.Unmarshal(bytes, &result)
		}
		return result, err
	}
	error_body, nerr := client.GetError(bytes)
	return result, fmt.Errorf("backup server - get config details failed- errors: %v %v %v", err, error_body, nerr)
}
