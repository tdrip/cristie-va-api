package apps

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	sess "github.com/tdrip/apiclient/pkg/v1/session"
	config "github.com/tdrip/cristie-va-api/pkg/config"
	bsapi "github.com/tdrip/cristie-va-api/pkg/v1/api/backupservers"
	"github.com/tdrip/cristie-va-api/pkg/v1/api/estate"
	orcapi "github.com/tdrip/cristie-va-api/pkg/v1/api/orchestration"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	helpers "github.com/tdrip/cristie-va-api/pkg/v1/helpers"
	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	targets "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

const layout = "01-02-2006-15-04-05"

func RubrikRecoveryJobCreator(clint *http.Client, cnt config.VAConnection, logger sess.SessionLog, backuptype string, trg config.RecoveryTarget) error {
	crs := client.NewClient(cnt.Server, clint, logger, cnt.Debug, cnt.DumpRequest, cnt.DumpResponses)

	err := client.Login(crs, cnt.User, cnt.PWord)
	if err != nil {
		return err
	}

	bservers, err := bsapi.GetServers(crs)
	if err != nil {
		return err
	}

	backupserverid := -1
	for _, v := range bservers.Servers {
		if strings.EqualFold(v.Type, backuptype) {
			backupserverid = v.Id
		}

	}

	if backupserverid < 0 {
		return fmt.Errorf("no backup server found of type %s found on %s", backuptype, cnt.Server)
	}

	res, err := bsapi.GetClients(crs, backupserverid)
	if err != nil {
		return err
	}

	fulldetails := []bs.Client{}
	for _, client := range res {
		res, err := bsapi.GetClientDetails(crs, int(client.Id))
		if err != nil {
			return err
		}
		fulldetails = append(fulldetails, res)
	}

	earliest := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	latestpit := bs.Pit{Date: earliest}

	// better clear the estate
	if len(fulldetails) > 0 {
		machines, err := estate.GetMachinesByMac(crs, trg.MacAddress)
		if err != nil {
			return err
		}

		for _, system := range machines.Machines {
			_, err := estate.DeleteMachine(crs, system.Uuid)
			if err != nil {
				return err
			}
		}
	}

	for _, fd := range fulldetails {

		for _, pit := range fd.PointInTimeDates {
			if pit.Date.After(latestpit.Date) {
				latestpit = pit
			}

		}

		if latestpit.Date.After(earliest) {
			cfg := targets.VmConfiguration{}
			tries := 3

			for i := 0; i < tries; i++ {
				cfg1, err := bsapi.GetConfigDetails(crs, backupserverid, fd.Name, latestpit.Type)
				if err != nil && i == tries {
					return err
				}
				if err == nil {
					cfg = cfg1
				}
			}

			t := time.Now()
			jobname := fmt.Sprintf("%s-%s", fd.Name, t.Format(layout))
			job, err := orcapi.CreateJob(crs, jobname)
			if err != nil {
				return err
			}
			stgname := fmt.Sprintf("stg-%s-%s", fd.Name, t.Format(layout))
			jobid, _ := strconv.Atoi(job.ScheduleReference)

			stg, err := orcapi.CreateStage(crs, stgname, jobid)
			if err != nil {
				return err
			}

			src := helpers.NewRubrikSource(fd, latestpit, backupserverid)
			task := helpers.CreateRecoveryTask(jobid, stg.Id, fmt.Sprintf("blk-%s-%s", fd.Name, t.Format(layout)))
			task.SourceTargetList = helpers.NewSourceTargets(trg.MacAddress, trg.BiosUuid, trg.OS, fd, cfg, src)

			_, err = orcapi.CreateBlock(crs, task)
			if err != nil {
				return err
			}

			_, err = orcapi.RunJob(crs, jobid, -1, -1)
			return err
		}

	}

	return nil
}
