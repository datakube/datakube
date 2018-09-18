package api

import (
	"github.com/datakube/datakube/types"
)

type targetStore interface {
	ListTargets() []types.Target
}

type jobStore interface {
	ListAllJobs() ([]types.Job, error)
}
