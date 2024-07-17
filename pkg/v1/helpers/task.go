package helpers

import (
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
)

func CreateTask(jobid int, stagid int, name string, ttype int) orchestration.Block {
	task := orchestration.Block{}
	task.Name = name
	task.Type = ttype
	task.JobId = jobid
	task.StageId = stagid
	return task
}

func CreateRecoveryTask(jobid int, stagid int, name string) orchestration.Block {
	return CreateTask(jobid, stagid, name, consts.TaskType_Recovery)
}
