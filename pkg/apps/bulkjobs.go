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
	estate "github.com/tdrip/cristie-va-api/pkg/v1/api/estate"
	orcapi "github.com/tdrip/cristie-va-api/pkg/v1/api/orchestration"
	client "github.com/tdrip/cristie-va-api/pkg/v1/client"
	helpers "github.com/tdrip/cristie-va-api/pkg/v1/helpers"
	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	targets "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

const layout = "01-02-2006-15-04-05"

func RubrikRecoveryJobCreator(clint *http.Client, cnt config.VAConnection, logger sess.SessionLog, updateui UpdateUI, backuptype string, trg config.RecoveryTarget, pausetimeout int) error {
	if updateui == nil {
		updateui = DefaultUpdateUI
	}

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
		if strings.EqualFold(trg.Hostname, client.Name) {
			res, err := bsapi.GetClientDetails(crs, int(client.Id))
			if err != nil {
				return err
			}
			fulldetails = append(fulldetails, res)
			if strings.EqualFold(trg.OS, client.Platform) {
				updateui(fmt.Sprintf("platform %s matched target %s ", client.Platform, trg.OS))
			} else {
				updateui(fmt.Sprintf("platform %s does not target %s ", client.Platform, trg.OS))
			}
		}
	}

	earliest := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	latestpit := bs.Pit{Date: earliest}

	// better clear the estate
	if len(fulldetails) > 0 {
		machines, err := estate.GetMachinesByMac(crs, trg.MacAddress)
		if err != nil {
			return err
		}

		updateui(fmt.Sprintf("Found %d machines matching %s", len(machines.Machines), trg.MacAddress))

		for _, system := range machines.Machines {
			updateui(fmt.Sprintf("Deleting machine %s", system.Uuid))
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
					updateui(fmt.Sprintf("Tried %d times to get configuration for machine %s returning error", tries, fd.Name))
					return err
				}
				if err == nil {
					updateui(fmt.Sprintf("On attempt %d of %d got configuration for machine %s", i, tries, fd.Name))
					cfg = cfg1
					break
				}
			}

			t := time.Now()
			jobname := fmt.Sprintf("%s-%s", fd.Name, t.Format(layout))

			updateui(fmt.Sprintf("Creating job with name %s", jobname))
			jobevent, err := orcapi.CreateJob(crs, jobname)
			if err != nil {
				return err
			}
			jobid, _ := strconv.Atoi(jobevent.ScheduleReference)

			updateui(fmt.Sprintf("Created job %s with id %d", jobname, jobid))

			stgname := fmt.Sprintf("stg-%s-%s", fd.Name, t.Format(layout))

			updateui(fmt.Sprintf("Creating stage with name %s", stgname))
			stg, err := orcapi.CreateStage(crs, stgname, jobid)
			if err != nil {
				return err
			}

			updateui(fmt.Sprintf("Created stage %s with id %d", stgname, stg.Id))

			src := helpers.NewRubrikSource(fd, latestpit, backupserverid)
			taskname := fmt.Sprintf("blk-%s-%s", fd.Name, t.Format(layout))

			updateui(fmt.Sprintf("Creating task with name %s", taskname))
			if len(trg.OS) == 0 {
				trg.OS = fd.Platform
			}

			task := helpers.CreateRecoveryTask(jobid, stg.Id, taskname)
			task.SourceTargetList = helpers.NewSourceTargets(trg.MacAddress, trg.BiosUuid, trg.OS, fd, cfg, src)

			if pausetimeout >= 0 {
				updateui(fmt.Sprintf("Adding %d seconds validation to task %s", pausetimeout, taskname))
				task = helpers.AddValidationPauseToTargets(task, pausetimeout)
			}
			block, err := orcapi.CreateBlock(crs, task)
			if err != nil {
				return err
			}
			updateui(fmt.Sprintf("Created task %s with id %d", taskname, block.Id))
			updateui(fmt.Sprintf("Running job with id %d", jobid))

			_, err = orcapi.RunJob(crs, jobid, -1, -1)
			if err != nil {
				return err
			}
			updateui(fmt.Sprintf("Sucessfully running job with id %d", jobid))
			return nil
		}

	}

	return nil
}
