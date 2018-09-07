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
