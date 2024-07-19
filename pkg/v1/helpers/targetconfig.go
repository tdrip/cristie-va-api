package helpers

import (
	"fmt"
	"strings"

	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	disks "github.com/tdrip/cristie-va-api/pkg/v1/models/disks"
	trg "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func NewTargetConfig(client bs.Client, cfg trg.VmConfiguration) *trg.TargetConfig {
	cnf := trg.TargetConfig{}
	cnf.Id = 0
	cnf.HardDisksInGB = []disks.DiskType{}
	cnf.SkipMultipath = true
	cnf.DiskMapList = []string{}
	cnf.ImportDHCP = false
	//
	cnf.Os = cfg.Os // "Windows X7"

	arch := strings.ToUpper(cfg.Arch)
	if arch == "AMD64" {
		arch = "X64"
	}

	cnf.Arch = arch                   // WINDOWS <= expected  <= Windows from cfg
	cnf.RamSizeInMB = cfg.RamSizeInMB //2048
	cnf.CpuCount = cfg.CpuCount       // 2
	cnf.FirmwareOpt = cfg.EfiSys      // "BIOS"
	cnf.Hostname = client.Name
	cnf.VmName = fmt.Sprintf("RS-%s", client.Name)
	cnf.OriginalHostname = client.Name
	cnf.LocalDisksOnly = true
	cnf.RecoveryFailureAction = consts.Action_Leave
	cnf.RecoverySuccessAction = consts.Action_Reboot
	cnf.PostScriptCmd = nil
	return &cnf
}
