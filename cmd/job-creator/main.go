package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tdrip/cristie-va-api/pkg/config"
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/helpers"
)

func main() {
	tm := flag.String("tm", "", "target mac address: must be colon format aa:bb:cc:dd:ee:ff")
	tb := flag.String("tb", "", "target bios uuid: ABCDEFGH-1234-1234-1234-ABCDEFGHFFFF")
	to := flag.String("to", "", "OS Type: linux or pwindows (can be null)")
	th := flag.String("th", "", "target hostname - example: somehost.domain.com")

	u := flag.String("u", "", "username - example: jobcreatoruser")
	p := flag.String("p", "", "password for the account in --u ")
	s := flag.String("s", "", "cristie va server which we need to connect to  - FQDN ")

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
		targetbuuid = strings.ToUpper(targetbuuid)
	}

	if len(targetmacaddress) > 0 {
		targetmacaddress = strings.ToUpper(targetmacaddress)
	}

	vaserver := *(s)
	user := *(u)
	pword := *(p)

	cnt := config.VAConnection{User: user, Server: vaserver, PWord: pword, Debug: true, DumpRequest: true, DumpResponses: true}
	trg := config.RecoveryTarget{Hostname: targethost, MacAddress: targetmacaddress, BiosUuid: targetbuuid, OS: targetos}

	clnt, err := helpers.CreateClient(helpers.BuildCristieClient(), cnt, dumpapidata)
	if err != nil {
		uiupdate(err.Error())
		os.Exit(1)
	}

	fds, err := helpers.FindJobs(clnt, targetmacaddress, targetbuuid)
	if err != nil {
		uiupdate(err.Error())
		os.Exit(1)
	}
	if len(fds) == 0 {
		err = helpers.CreateJob(clnt, uiupdate, nil, consts.SourceTarget_Type_RBMR, trg, 300, false)

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

func dumpapidata(msg string, data string, err error) {
	if err != nil {
		uiupdate(fmt.Sprintf("Message (Error) %s\n  Data:  %s\n   Err:  %s\n", msg, data, err.Error()))
	} else {
		uiupdate(fmt.Sprintf("Message (Debug):%s\n  Data:  %s\n", msg, data))
	}
}
