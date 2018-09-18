package rpc

import (
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/types"
)

type dumpfileStore interface {
	SaveDumpFile(types.DumpFile) (types.DumpFile, error)
}

type targetStore interface {
	GetOneTargetByName(targetName string) (types.Target, error)
}

type jobStore interface {
	ListJobsByStatus(string) ([]types.Job, error)
	DeleteJob(types.Job) error
	GetJobById(id int) (types.Job, error)
	SaveJob(job types.Job) (types.Job, error)
}

func New(dfs dumpfileStore, ts targetStore, js jobStore, stg storage.Storage) *rpcHandler {
	return &rpcHandler{js: js, dfs: dfs, ts: ts, stg: stg}
}

type rpcHandler struct {
	dfs dumpfileStore
	ts  targetStore
	js  jobStore
	stg storage.Storage
}
