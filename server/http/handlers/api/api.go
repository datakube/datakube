package api

import (
	"github.com/SantoDE/datahamster/types"
)

type targetStore interface {
	ListTargets() []types.Target
}

type jobStore interface {
	ListAllJobs() ([]types.Job, error)
}

//PingHandler to hold Pinghandler information
type ApiHandler struct {
	targetStore
	jobStore
}

//NewPingHandler to create a new Pinghandler
func NewApiHandler(targetStore targetStore, jobStore jobStore) *ApiHandler {
	a := new(ApiHandler)
	a.targetStore = targetStore
	a.jobStore = jobStore
	return a
}
