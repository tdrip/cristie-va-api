package webboot

import (
	"encoding/json"
	"errors"
	"fmt"

	cls "github.com/tdrip/apiclient/pkg/v1/client"
	utils "github.com/tdrip/apiclient/pkg/v1/utils"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	wb "github.com/tdrip/cristie-va-api/pkg/v1/models/webboot"
)

const (
	UriWebbootMap = "v1/webboot/map/"
)

func CreateDHCPMapping(crs *cls.Client, product string, os string, mac string, buuid string) (wb.Mappings, error) {
	result := wb.Mappings{}
	if !crs.Session.HasToken() {
		return result, errors.New("webboot DHCP mapping failed - token missing session")
	}

	request := wb.Mapping{}
	request.DHCP = true
	request.InsecureTls = false
	request.ConfigId = nil
	request.Mac = mac
	request.BiosUuid = buuid
	request.Ip = "DHCP"
	request.Netmask = "DHCP"
	request.Gateway = "DHCP"
	request.DnsAddrs = ""
	request.OS = os
	request.Product = product
	return CreateMapping(crs, request)
}

func CreateMapping(crs *cls.Client, request wb.Mapping) (wb.Mappings, error) {
	result := wb.Mappings{}
	if !crs.Session.HasToken() {
		return result, errors.New("webboot mapping failed - token missing session")
	}

	bytes, res, err := crs.Session.PostBody(UriWebbootMap, &request)

	if err == nil && res != nil && utils.RequestIsSuccessful(res.StatusCode) {
		err = json.Unmarshal(bytes, &result)
		return result, err
	}
	error_body, nerr := client.GetError(bytes, res)
	return result, fmt.Errorf("webboot mapping  - errors: %v %v %v", err, error_body, nerr)
}
