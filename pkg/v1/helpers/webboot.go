package helpers

import (
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func NewWebbootMapping(mac string, buuid string, os string, prodct string) *targets.WebbootMapping {
	wb := targets.WebbootMapping{}
	wb.Id = 0         // this will be omitted in JSON
	wb.ConfigId = nil // this needs to be null in JSON

	wb.InsecureTls = false
	wb.InsecureSSL = false
	wb.DrNetworking = false

	if len(mac) > 0 {
		wb.Mac = mac
	}

	if len(buuid) > 0 {
		wb.BiosUuid = buuid
	}

	// Recovery
	wb.Product = prodct
	wb.Os = os

	// default networking
	wb.DHCP = true
	wb.Ip = consts.Network_DHCP_Name
	wb.Netmask = consts.Network_DHCP_Name
	wb.Gateway = consts.Network_DHCP_Name
	wb.DnsAddrs = []string{}
	return &wb
}

func NewRubrikWebbootMapping(mac string, buuid string, os string) *targets.WebbootMapping {
	return NewWebbootMapping(mac, buuid, os, consts.SourceTarget_Type_RBMR)
}
