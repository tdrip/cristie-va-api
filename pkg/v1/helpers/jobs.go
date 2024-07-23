package helpers

import (
	"strings"

	client "github.com/tdrip/apiclient/pkg/v1/client"
	orchapi "github.com/tdrip/cristie-va-api/pkg/v1/api/orchestration"
	orch "github.com/tdrip/cristie-va-api/pkg/v1/models/orchestration"
)

func FindJobs(mac string, buid string, crs *client.Client) ([]orch.Job, error) {
	jobs, err := orchapi.GetJobs(crs)

	if err != nil {
		return jobs, err
	}

	fds := []orch.Job{}
	for _, j := range jobs {
		jobdetail, err := orchapi.GetJob(crs, j.Id)
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
