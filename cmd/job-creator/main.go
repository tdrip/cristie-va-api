package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tdrip/cristie-va-api/pkg/apps"
	"github.com/tdrip/cristie-va-api/pkg/config"
	"github.com/tdrip/cristie-va-api/pkg/v1/client"
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/helpers"
)

func Dump(msg string, data string, err error) {
	if err != nil {
		uiupdate(fmt.Sprintf("Message (Error) %s\n  Data:  %s\n   Err:  %s\n", msg, data, err.Error()))
	} else {
		uiupdate(fmt.Sprintf("Message (Debug):%s\n  Data:  %s\n", msg, data))
	}
}

func BuildCristieClient() *http.Client {
	transport := &http.Transport{Proxy: http.ProxyFromEnvironment, TLSClientConfig: &tls.Config{Renegotiation: tls.RenegotiateOnceAsClient, InsecureSkipVerify: true}, TLSNextProto: map[string]func(authority string, c *tls.Conn) http.RoundTripper{}}
	return &http.Client{
		Timeout:   time.Second * 360,
		Transport: transport,
	}
}

func main() {
	tm := flag.String("tm", "", "target mac address")
	tb := flag.String("tb", "", "target bios uuid")
	to := flag.String("to", "", "OS Type")
	th := flag.String("th", "", "target hostname - example: ")

	u := flag.String("u", "", "username - example: rackn")
	p := flag.String("p", "", "target hostname - example: ")
	s := flag.String("s", "", "target hostname - example: ")

	flag.Parse()
	targethost := *th
	targetmacaddress := *tm
	targetbuuid := *tb
	targetos := *to

	if len(targetbuuid) == 0 && len(targetmacaddress) == 0 {
		uiupdate("Missing parameter - BOISUUID or MAC missing")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(targethost) == 0 {
		uiupdate("Missing parameter -  Target Hostname missing")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(targetbuuid) > 0 {
		targetbuuid = strings.ToLower(targetbuuid)
	}

	if len(targetmacaddress) > 0 {
		targetmacaddress = strings.ToLower(targetmacaddress)
	}

	vaserver := *(s)
	user := *(u)
	pword := *(p)

	cnt := config.VAConnection{User: user, Server: vaserver, PWord: pword, Debug: true, DumpRequest: true, DumpResponses: true}
	trg := config.RecoveryTarget{Hostname: targethost, MacAddress: targetmacaddress, BiosUuid: targetbuuid, OS: targetos}
	clnt := client.NewClient(cnt.Server, BuildCristieClient(), Dump, cnt.Debug, cnt.DumpRequest, cnt.DumpResponses)

	err := client.Login(clnt, cnt.User, cnt.PWord)
	if err != nil {
		uiupdate(err.Error())
		os.Exit(1)
	}

	fds, err := helpers.FindJobs(targetmacaddress, targetbuuid, clnt)
	if err != nil {
		uiupdate(err.Error())
		os.Exit(1)
	}
	if len(fds) == 0 {
		err = apps.RubrikRecoveryJobCreator(BuildCristieClient(), cnt, Dump, uiupdate, nil, consts.SourceTarget_Type_RBMR, trg, 300, false)

		if err != nil {
			uiupdate(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	} else {
		for _, fd := range fds {
			uiupdate(fmt.Sprintf("matched jobid %d ", fd.Id))
		}
		os.Exit(0)

	}

}

func uiupdate(msg string) {
	fmt.Println(msg)
}
