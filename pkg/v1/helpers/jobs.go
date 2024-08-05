package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	client "github.com/tdrip/apiclient/pkg/v1/client"
	config "github.com/tdrip/cristie-va-api/pkg/config"
	bsapi "github.com/tdrip/cristie-va-api/pkg/v1/api/backupservers"
	estate "github.com/tdrip/cristie-va-api/pkg/v1/api/estate"
	orcapi "github.com/tdrip/cristie-va-api/pkg/v1/api/orchestration"
	bs "github.com/tdrip/cristie-va-api/pkg/v1/models/backupservers"
	orch "github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
	targets "github.com/tdrip/cristie-va-api/pkg/v1/models/targets"
)

func FindJobs(crs *client.Client, mac string, buid string) ([]orch.Job, error) {
	jobs, err := orcapi.GetJobs(crs)

	if err != nil {
		return jobs, err
	}

	fds := []orch.Job{}
	for _, j := range jobs {
		jobdetail, err := orcapi.GetJob(crs, j.Id)
		if err != nil {
			continue
		}
		for _, stg := range jobdetail.Stages {
			for _, blk := range stg.Blocks {
				for _, st := range blk.SourceTargetList {
					lmac := *st.TargetMAC
					lbuid := *st.TargetBiosUuid
					if strings.EqualFold(lmac, mac) || strings.EqualFold(lbuid, buid) {
						fds = append(fds, jobdetail)
					}
				}
			}

		}
	}

	return fds, nil
}

func CreateJob(crs *client.Client, updateui UpdateUI, naming NamingConvention, backuptype string, trg config.RecoveryTarget, pausetimeout int, startjob bool) error {
	if updateui == nil {
		updateui = DefaultUpdateUI
	}

	if naming == nil {
		naming = DefaultNamingConvention
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
		return fmt.Errorf("no backup server found of type %s ", backuptype)
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
			if len(trg.OS) == 0 {
				updateui(fmt.Sprintf("target is empty - setting platform to: %s", client.Platform))
				trg.OS = res.Platform
			} else {
				if strings.EqualFold(trg.OS, client.Platform) {
					updateui(fmt.Sprintf("platform %s matched target %s ", client.Platform, trg.OS))
				} else {
					updateui(fmt.Sprintf("platform %s does not target %s ", client.Platform, trg.OS))
				}
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
			jobname := naming("job-", fd, t)

			updateui(fmt.Sprintf("Creating job with name %s", jobname))
			jobevent, err := orcapi.CreateJob(crs, jobname)
			if err != nil {
				return err
			}
			jobid, _ := strconv.Atoi(jobevent.ScheduleReference)

			updateui(fmt.Sprintf("Created job %s with id %d", jobname, jobid))

			stgname := naming("stg-", fd, t)

			updateui(fmt.Sprintf("Creating stage with name %s", stgname))
			stg, err := orcapi.CreateStage(crs, stgname, jobid)
			if err != nil {
				return err
			}

			updateui(fmt.Sprintf("Created stage %s with id %d", stgname, stg.Id))

			src := NewRubrikSource(fd, latestpit, backupserverid)
			taskname := naming("blk-", fd, t)

			updateui(fmt.Sprintf("Creating task with name %s", taskname))
			if len(trg.OS) == 0 {
				trg.OS = fd.Platform
			}

			task := CreateRecoveryTask(jobid, stg.Id, taskname)
			task.SourceTargetList = NewSourceTargets(trg.MacAddress, trg.BiosUuid, trg.OS, NewTargetConfig(fd, cfg), src)

			if pausetimeout >= 0 {
				updateui(fmt.Sprintf("Adding %d seconds validation to task %s", pausetimeout, taskname))
				task = AddValidationPauseToTargets(task, pausetimeout)
			}
			block, err := orcapi.CreateBlock(crs, task)
			if err != nil {
				return err
			}
			updateui(fmt.Sprintf("Created task %s with id %d", taskname, block.Id))

			if startjob {
				updateui(fmt.Sprintf("Running job with id %d", jobid))

				_, err = orcapi.RunJob(crs, jobid, -1, -1)
				if err != nil {
					return err
				}
				updateui(fmt.Sprintf("Sucessfully running job with id %d", jobid))
			}

			return nil
		}

	}

	return nil
}
n