package handlers

import (
	"context"
	"github.com/SantoDE/datahamster/types"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/storage"
	"time"
)

type jobStore interface {
	ListAllQueuedJobs() ([]types.Job, error)
}

type dumpfileStore interface {
	SaveDumpFile(types.DumpFile) (types.DumpFile, error)
}

//DumperHandler struct to hold DumperHandler specific information
type RpcHandler struct {
	BaseHandler
	jobStore
	dumpfileStore
	storage.Storage
}

func NewRpcHandler(store jobStore, storage storage.Storage) RpcHandler {

	h := new(RpcHandler)
	h.jobStore = store
	h.Storage = storage

	return *h
}

//ConnectDumper function which gets called when an Dumper connected
func (h *RpcHandler) SaveDumpFile(ctx context.Context, in *rpc.SaveDumpFileRequest) (*rpc.SaveDumpFileResponse, error) {

	log.Debugf("Received RPC Request to save file with filename %s for token %s", in.Filename, in.Token)
	file := types.File{
		Name: in.Targetname,
		Data: in.Data,
	}

	log.Debugf("Saving file %s", in.Filename)
	savedFile, err := h.Storage.SaveFile(file)

	if err != nil {
		log.Errorf("Error while saving file %s", savedFile.Name)
	}

	log.Debugf("Persist Saved Dump Record in Database")
	_, err = h.dumpfileStore.SaveDumpFile(types.DumpFile{
		CreatedAt: time.Now(),
		Target: in.Targetname,
		File: types.File{
			Name: savedFile.Name,
			Path: savedFile.Path,
		},
	})

	if err != nil {
		return &rpc.SaveDumpFileResponse{
			Success:false,
		}, err
	}

	return &rpc.SaveDumpFileResponse{
		Success:true,
	}, nil
}

func (h *RpcHandler) ListQueuedJobs(ctx context.Context, in *rpc.ListJobsRequest) (*rpc.ListJobsResponse, error) {

	jobs, err := h.ListAllQueuedJobs()

	var rpcJobs []*rpc.Job

	if err != nil {
		log.Error("Error fetching all queued Jobs for RPC Call : %s", err)
	}

	for _, job := range jobs {
		j := new(rpc.Job)
		j.Target = &rpc.Target{
			Name: job.Target,
			State: job.State,
		}
		rpcJobs = append(rpcJobs, j)
	}
	return &rpc.ListJobsResponse{
		Jobs: rpcJobs,
	}, err
}

