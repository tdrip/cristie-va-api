package helpers

import (
	"fmt"
	"strings"

	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/disks"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func NewTargetConfig(client backupservers.Client, cfg targets.VmConfiguration) targets.TargetConfig {
	cnf := targets.TargetConfig{}
	cnf.Id = 0
	cnf.HardDisksInGB = []disks.DiskType{}
	cnf.SkipMultipath = true
	cnf.DiskMapList = []string{}
	cnf.ImportDHCP = false
	//
	cnf.Os = cfg.Os                      // "Windows X7"
	cnf.Arch = strings.ToUpper(cfg.Arch) // WINDOWS <= expected  <= Windows from cfg
	cnf.RamSizeInMB = cfg.RamSizeInMB    //2048
	cnf.CpuCount = cfg.CpuCount          // 2
	cnf.FirmwareOpt = cfg.EfiSys         // "BIOS"
	cnf.Hostname = client.Name
	cnf.VmName = fmt.Sprintf("RS-%s", client.Name)
	cnf.OriginalHostname = client.Name
	cnf.LocalDisksOnly = true
	cnf.RecoveryFailureAction = consts.Action_Leave
	cnf.RecoverySuccessAction = consts.Action_Reboot
	cnf.PostScriptCmd = nil
	return cnf
}
