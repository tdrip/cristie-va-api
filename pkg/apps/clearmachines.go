package apps

import (
	"net/http"

	sess "github.com/tdrip/apiclient/pkg/v1/session"
	config "github.com/tdrip/cristie-va-api/pkg/config"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	helpers "github.com/tdrip/cristie-va-api/pkg/v1/helpers"
)

func ClearEstate(clint *http.Client, cnt config.VAConnection, logger sess.SessionLog, backuptype string, trg config.RecoveryTarget) error {
	crs := client.NewClient(cnt.Server, clint, logger, cnt.Debug, cnt.DumpRequest, cnt.DumpResponses)

	err := client.Login(crs, cnt.User, cnt.PWord)
	if err != nil {
		return err
	}
	return helpers.ClearEstate(crs, trg.MacAddress)
}
