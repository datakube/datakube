package rpc

import (
	"context"
	"github.com/datakube/datakube/log"
	"github.com/datakube/datakube/rpc"
	"github.com/datakube/datakube/types"
)

func (h *rpcHandler) ListJobs(ctx context.Context, in *datakube.ListJobsRequest) (*datakube.ListJobsResponse, error) {

	jobs, err := h.js.ListJobsByStatus(in.Status)

	var rpcJobs []*datakube.Job

	if err != nil {
		log.Error("Error fetching all queued Jobs for RPC Call", err)
		if err.Error() == "not found" {
			err = nil
		}
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
		j.Id = int32(job.ID)

		j.State = job.Status
		rpcJobs = append(rpcJobs, j)
	}
	return &datakube.ListJobsResponse{
		Jobs: rpcJobs,
	}, err
}

func (h *rpcHandler) UpdateJob(ctx context.Context, in *datakube.UpdateJobRequest) (*datakube.UpdateJobResponse, error) {

	job, err := h.js.GetJobById(int(in.Job.Id))

	if err != nil {
		return &datakube.UpdateJobResponse{
			Job: in.Job,
			Success: false,
		}, err
	}

	job.Status = in.Job.State
	job.Message = in.Message

	h.js.SaveJob(job)

	return &datakube.UpdateJobResponse{
		Job: in.Job,
		Success: true,
	}, err
}
