package handlers

import (
	"context"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/types"
	"time"
)

type jobStore interface {
	ListJobsByStatus(status string) ([]types.Job, error)
	DeleteJob(types.Job) error
}

type dumpfileStore interface {
	SaveDumpFile(types.DumpFile) (types.DumpFile, error)
}

type targetStore interface {
	GetOneTargetByName(targetName string) (types.Target)
}

//DumperHandler struct to hold DumperHandler specific information
type RpcHandler struct {
	BaseHandler
	jobStore
	dumpfileStore
	targetStore
	storage.Storage
}

func NewRpcHandler(js jobStore, ts targetStore, df dumpfileStore, storage storage.Storage) RpcHandler {

	h := new(RpcHandler)
	h.jobStore = js
	h.Storage = storage
	h.targetStore = ts
	h.dumpfileStore = df

	return *h
}

//ConnectDumper function which gets called when an Dumper connected
func (h *RpcHandler) SaveDumpFile(ctx context.Context, in *datakube.SaveDumpFileRequest) (*datakube.SaveDumpFileResponse, error) {

	log.Debugf("Received RPC Request to save file with filename %s", in.Filename)
	file := types.File{
		Name: in.Targetname,
		Data: in.Data,
	}

	log.Debugf("Saving file %s", in.Filename)
	savedFile, err := h.Storage.SaveFile(file)

	if err != nil {
		log.Errorf("Error while saving file %s", savedFile.Name)
		return &datakube.SaveDumpFileResponse{
			Success: false,
		}, err
	}

	log.Debugf("Persist Saved Dump Record in Database")
	_, err = h.dumpfileStore.SaveDumpFile(types.DumpFile{
		CreatedAt: time.Now(),
		Target:    in.Targetname,
		File: types.File{
			Name: savedFile.Name,
			Path: savedFile.Path,
		},
	})

	if err != nil {
		return &datakube.SaveDumpFileResponse{
			Success: false,
		}, err
	}

	return &datakube.SaveDumpFileResponse{
		Success: true,
	}, nil
}

func (h *RpcHandler) ListJobs(ctx context.Context, in *datakube.ListJobsRequest) (*datakube.ListJobsResponse, error) {

	jobs, err := h.ListJobsByStatus(in.Status)

	var rpcJobs []*datakube.Job

	if err != nil {
		log.Error("Error fetching all queued Jobs for RPC Call : %s", err)
	}

	for _, job := range jobs {

		target := h.targetStore.GetOneTargetByName(job.Target)

		if target == *new(types.Target){
			h.jobStore.DeleteJob(job)
			continue
		}

		j := new(datakube.Job)
		j.Target = &datakube.Target{
			Name:  job.Target,
			Type: target.DBConfig.DatabaseType,
			Credentials: &datakube.Credentials{
				User: target.DBConfig.DatabaseUserName,
				Host: target.DBConfig.DatabaseHost,
				Database: target.DBConfig.DatabaseName,
				Password: target.DBConfig.DatabasePassword,
				Port: target.DBConfig.DatabasePort,
			},
		}

		j.State = job.Status
		rpcJobs = append(rpcJobs, j)
	}
	return &datakube.ListJobsResponse{
		Jobs: rpcJobs,
	}, err
}
