package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	cfg "github.com/tdrip/apiclient/pkg/v1/cfg"
	sess "github.com/tdrip/apiclient/pkg/v1/session"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	auth "github.com/tdrip/cristie-va-api/pkg/v1/api/auth"
	models "github.com/tdrip/cristie-va-api/pkg/v1/models"
)

const (
	UriUsersAuthPath      = "v1/users/auth/"
	UriUsersAuthSession   = "v1/users/session"    // <-- 4.9.1+
	UriUsersAuthSessionId = "v1/users/session_id" // <-- 4.8.3
)

func New(server string, logger sess.SessionLog, dbg bool, dumpreq bool, dumpres bool) *cls.Client {
	apis, _ := cfg.NewAPIServer(server, "")
	aser, _ := cfg.NewAuthServer(server, UriUsersAuthPath)
	crsclient := cls.NewTlsSkipCustomLogger(apis, aser, logger)
	crsclient.Session.Debug = dbg
	crsclient.Session.DumpRequest = dumpreq
	crsclient.Session.DumpResponse = dumpres
	return crsclient
}

func Login(crs *cls.Client, userid string, pwd string) error {
	request := auth.OAuthRequest{Username: userid, Password: pwd}
	bytes, res, err := crs.Session.PostBody(UriUsersAuthPath, &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		result := auth.OAuthResponse{}
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			return nil
		}
		err := cls.SetAuthToken(crs, result.AccessToken)
		return err
	}

	return fmt.Errorf("login failed with errors: %v", err)

}

func CheckSession(crs *cls.Client, usesessionid bool) (auth.SessionResponse, error) {
	result := auth.SessionResponse{}
	uri := UriUsersAuthSession
	if usesessionid {
		uri = UriUsersAuthSessionId
	}
	bytes, res, err := crs.Session.Get(uri)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		if err != nil {
			return result, nil
		}
		return result, err
	}

	return result, fmt.Errorf("check session failed with errors: %v", err)

}

func GetError(rawbody []byte, res *http.Response) (models.Exception, error) {
	if len(rawbody) > 0 && res != nil && res.Body != http.NoBody {
		res := models.Exception{}
		err := json.Unmarshal(rawbody, &res)
		return res, err
	}
	return models.Exception{}, errors.New("no data in http boy to parse as error")
}
