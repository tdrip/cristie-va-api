package helpers

import (
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	trg "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func NewSourceTargets(targetmac string, targetbuuid string, targetos string, trargetconfig *trg.TargetConfig, source *trg.BackupSource) []trg.SourceTarget {
	target := NewSourceTarget(targetmac, targetbuuid, targetos, trargetconfig, source)
	targets := []trg.SourceTarget{}
	targets = append(targets, target)
	return targets
}

func NewSourceTarget(targetmac string, targetbuuid string, targetos string, trargetconfig *trg.TargetConfig, source *trg.BackupSource) trg.SourceTarget {
	target := trg.NewSourceTarget()
	target.VmHostType = consts.SourceTarget_VMHostType_Standard
	target.TargetConfig = trargetconfig
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

func AddValidationPause(target trg.SourceTarget, timeout int) trg.SourceTarget {
	target.PauseForValidation = true
	target.ValidationPauseDuration = timeout
	return target
}
