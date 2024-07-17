package helpers

import (
	"github.com/tdrip/cristie-va-api/pkg/v1/consts"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	"github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func NewBackupSource(client backupservers.Client, pit backupservers.Pit, backupserverid int, stype string, name string) *targets.BackupSource {
	bs := targets.BackupSource{}
	bs.Type = stype
	bs.Name = name

	// pit details
	bs.PitTimestamp = &pit.Date
	bs.BackupType = pit.Type

	// backup server id
	bs.ServerId = backupserverid

	// client details
	bs.Hostname = client.Name
	bs.HostId = client.HostId

	return &bs
}

func NewRubrikSource(client backupservers.Client, pit backupservers.Pit, backupserverid int) *targets.BackupSource {
	return NewBackupSource(client, pit, backupserverid, consts.SourceTarget_Type_RBMR, consts.BackupSource_Default_Name)
}
