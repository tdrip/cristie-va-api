package helpers

import (
	"crypto/tls"
	"net/http"
	"time"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	sess "github.com/tdrip/apiclient/pkg/v1/session"
	config "github.com/tdrip/cristie-va-api/pkg/config"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
)

func BuildCristieClient() *http.Client {
	transport := &http.Transport{Proxy: http.ProxyFromEnvironment, TLSClientConfig: &tls.Config{Renegotiation: tls.RenegotiateOnceAsClient, InsecureSkipVerify: true}, TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{}}
	return &http.Client{
		Timeout:   time.Second * 360,
		Transport: transport,
	}
}

func CreateClient(clint *http.Client, cnt config.VAConnection, logger sess.SessionLog) (*cls.Client, error) {
	crs := client.NewClient(cnt.Server, clint, logger, cnt.Debug, cnt.DumpRequest, cnt.DumpResponses)

	err := client.Login(crs, cnt.User, cnt.PWord)
	if err != nil {
		return nil, err
	}

	return crs, nil
}
