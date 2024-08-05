package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tdrip/cristie-va-api/pkg/config"
	"github.com/tdrip/cristie-va-api/pkg/v1/helpers"
)

func main() {
	tm := flag.String("tm", "", "target mac address: must be colon format aa:bb:cc:dd:ee:ff")

	u := flag.String("u", "", "username - example: jobcreatoruser")
	p := flag.String("p", "", "password for the account in --u ")
	s := flag.String("s", "", "cristie va server which we need to connect to  - FQDN ")

	flag.Parse()

	targetmacaddress := *tm

	if len(targetmacaddress) == 0 {
		uiupdate("Missing parameter - BOISUUID or MAC missing")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(targetmacaddress) > 0 {
		targetmacaddress = strings.ToUpper(targetmacaddress)
	}

	vaserver := *(s)
	user := *(u)
	pword := *(p)

	cnt := config.VAConnection{User: user, Server: vaserver, PWord: pword, Debug: true, DumpRequest: true, DumpResponses: true}
	trg := config.RecoveryTarget{MacAddress: targetmacaddress}

	clnt, err := helpers.CreateClient(helpers.BuildCristieClient(), cnt, dumpapidata)
	if err != nil {
		uiupdate(err.Error())
		os.Exit(1)
	}

	err = helpers.ClearEstate(clnt, trg.MacAddress)

	if err != nil {
		uiupdate(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
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
