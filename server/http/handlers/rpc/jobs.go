package rpc

import (
	"context"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/types"
)

func (h *rpcHandler) ListJobs(ctx context.Context, in *datakube.ListJobsRequest) (*datakube.ListJobsResponse, error) {

	jobs, err := h.js.ListJobsByStatus(in.Status)

	var rpcJobs []*datakube.Job

	if err != nil {
		log.Error("Error fetching all queued Jobs for RPC Call", err)
	}

	for _, job := range jobs {

		target, _ := h.ts.GetOneTargetByName(job.Target)

		if target == *new(types.Target) {
			h.js.DeleteJob(job)
			continue
		}

		j := new(datakube.Job)
		j.Target = &datakube.Target{
			Name: job.Target,
			Type: target.DBConfig.DatabaseType,
			Credentials: &datakube.Credentials{
				User:     target.DBConfig.DatabaseUserName,
				Host:     target.DBConfig.DatabaseHost,
				Database: target.DBConfig.DatabaseName,
				Password: target.DBConfig.DatabasePassword,
				Port:     target.DBConfig.DatabasePort,
			},
		}

		j.State = job.Status
		rpcJobs = append(rpcJobs, j)
	}
	return &datakube.ListJobsResponse{
		Jobs: rpcJobs,
	}, err
}
