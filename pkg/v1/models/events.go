package models

import (
	"time"
)

type Event struct {
	Abortable         bool      `json:"abortable,omitempty"`         // false
	BlockId           int32     `json:"blockId,omitempty"`           // 0
	Completed         bool      `json:"completed,omitempty"`         // false
	Details           string    `json:"details,omitempty"`           // Details of what is currently happening - "Something was created"
	Error             int32     `json:"error,omitempty"`             // 0
	Failed            int32     `json:"failed,omitempty"`            // 0
	FinishTime        time.Time `json:"finishTime,omitempty"`        // "2024-07-01T22:18:14.636481742Z"
	Finish_Time       time.Time `json:"finish_time,omitempty"`       // "2024-07-01T22:18:14.636481742Z"
	Id                int32     `json:"id,omitempty"`                // ID of the event - 323
	Identifier        string    `json:"identifier,omitempty"`        // "ochestration"
	Logged            bool      `json:"logged,omitempty"`            // false
	ModifiedTimestamp int64     `json:"modifiedTimestamp,omitempty"` // 0
	NodeUuid          string    `json:"nodeUuid,omitempty"`          // null
	Percent           int32     `json:"percent,omitempty"`           // 0
	Replayable        bool      `json:"replayable,omitempty"`        // false
	Result            string    `json:"result,omitempty"`            // null
	Schedule          string    `json:"schedule,omitempty"`          // ""
	ScheduleReference string    `json:"scheduleReference,omitempty"` // "96"
	Source            string    `json:"source,omitempty"`            //  Event source (if applicable) - "<>"
	StartTime         time.Time `json:"startTime,omitempty"`         // "2024-07-01T22:18:14.636481742Z"
	Start_Time        time.Time `json:"start_time,omitempty"`        // "2024-07-01T22:18:14.636481742Z"
	Status            string    `json:"status,omitempty"`            // The current status of this event - "Completed"
	Target            string    `json:"target,omitempty"`            // Event target (if applicable) - ""
	Task              string    `json:"task,omitempty"`              // The task being performed by this event - "DR Orchestration"
	Type              int32     `json:"type,omitempty"`              // 1
	User              string    `json:"user,omitempty"`              // "loggedin:user"
	Uuid              string    `json:"uuid,omitempty"`              // UUID of the event (deprecated) - "f3456789-e678-e678-e678-f3456789"
}
