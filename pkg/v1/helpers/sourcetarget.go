package helpers

import (
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	trg "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func NewSourceTargets(targetmac string, targetbuuid string, targetos string, client bs.Client, cfg trg.VmConfiguration, source *trg.BackupSource) []trg.SourceTarget {
	target := NewSourceTarget(targetmac, targetbuuid, targetos, client, cfg, source)
	targets := []trg.SourceTarget{}
	targets = append(targets, target)
	return targets
}

func NewSourceTarget(targetmac string, targetbuuid string, targetos string, client bs.Client, cfg trg.VmConfiguration, source *trg.BackupSource) trg.SourceTarget {
	target := trg.NewSourceTarget()
	target.VmHostType = consts.SourceTarget_VMHostType_Standard
	target.TargetConfig = NewTargetConfig(client, cfg)
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
