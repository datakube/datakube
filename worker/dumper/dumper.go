package dumper

import (
	"github.com/santode/datahamster/worker/configuration"
	"github.com/santode/datahamster/worker/types"
)

// BaseDumper struct which holds basic configuration for all dumpers
type BaseDumper struct {
	Config configuration.DatabaseConfiguration
	Dir    string
}

// Dumper Interface for all dumpers
type Dumper interface {
	Dump() (*types.DumpResult, error)
}
