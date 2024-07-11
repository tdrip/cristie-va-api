package models

import (
	"time"
)

type Event struct {
	Abortable         bool      `json:"abortable,omitempty"`
	BlockId           int32     `json:"blockId,omitempty"`
	Completed         bool      `json:"completed,omitempty"`
	Details           string    `json:"details,omitempty"`
	Error             int32     `json:"error,omitempty"`
	Failed            int32     `json:"failed,omitempty"`
	FinishTime        time.Time `json:"finishTime,omitempty"`
	Finish_Time       time.Time `json:"finish_time,omitempty"`
	Id                int32     `json:"id,omitempty"`
	Identifier        string    `json:"identifier,omitempty"`
	Logged            bool      `json:"logged,omitempty"`
	ModifiedTimestamp int64     `json:"modifiedTimestamp,omitempty"`
	NodeUuid          string    `json:"nodeUuid,omitempty"`
	Percent           int32     `json:"percent,omitempty"`
	Replayable        bool      `json:"replayable,omitempty"`
	Result            string    `json:"result,omitempty"`
	Schedule          string    `json:"schedule,omitempty"`
	ScheduleReference string    `json:"scheduleReference,omitempty"`
	Source            string    `json:"source,omitempty"`
	StartTime         time.Time `json:"startTime,omitempty"`
	Start_Time        time.Time `json:"start_time,omitempty"`
	Status            string    `json:"status,omitempty"`
	Target            string    `json:"target,omitempty"`
	Task              string    `json:"task,omitempty"`
	Type              int32     `json:"type,omitempty"`
	User              string    `json:"user,omitempty"`
	Uuid              string    `json:"uuid,omitempty"`
}
