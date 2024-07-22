package helpers

import (
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
	trg "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func CreateTask(jobid int, stagid int, name string, ttype int) orchestration.CBlock {
	task := orchestration.CBlock{}
	task.Name = name
	task.Type = ttype
	task.JobId = jobid
	task.StageId = stagid
	return task
}

func CreateRecoveryTask(jobid int, stagid int, name string) orchestration.CBlock {
	return CreateTask(jobid, stagid, name, consts.TaskType_Recovery)
}

func AddValidationPauseToTargets(task orchestration.CBlock, timeout int) orchestration.CBlock {
	stargets := []trg.SourceTarget{}
	for _, v := range task.SourceTargetList {
		v1 := AddValidationPause(v, timeout)
		stargets = append(stargets, v1)
	}
	task.SourceTargetList = stargets
	return task
}
