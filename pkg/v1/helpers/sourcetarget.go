package helpers

import (
	"github.com/tdrip/cristie-va-api/pkg/v1/api/backupservers"
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func NewSourceTargets(targetmac string, targetbuuid string, targetos string, client backupservers.Client, cfg targets.VmConfiguration, source *targets.BackupSource) []targets.SourceTarget {
	target := NewSourceTarget()
	targets := []targets.SourceTarget{}
	targets = append(targets, target)
	return targets
}

func NewSourceTarget(targetmac string, targetbuuid string, targetos string, client backupservers.Client, cfg targets.VmConfiguration, source *targets.BackupSource) targets.SourceTarget {
	target := targets.NewSourceTarget()
	target.VmHostType = consts.SourceTarget_VMHostType_Standard
	target.TargetConfig = NewTargetConfig(source, cfg)
	target.SourceMacAddress = targetmac
	target.Source = source
	if len(targetmac) > 0 {
		target.TargetMAC = &targetmac
	} else {
		target.TargetMAC = nil
	}
	target.Type = consts.SourceTarget_Type_RBMR
	target.IsoTrigger = true // tells iso to start work on boot automatically!
	if len(targetbuuid) > 0 {
		target.TargetBiosUuid = &targetbuuid
	} else {
		target.TargetBiosUuid = nil
	}
	target.WebbootMapping = NewRubrikWebbootMapping(targetmac, targetbuuid, targetos)
	return target
}
