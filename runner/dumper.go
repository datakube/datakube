package runner

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/SantoDE/datahamster/configuration"
)

const TYPE_MSQL string = "mysql"

// Dumper Interface for all dumpers
type DumpRunner interface {
	Dump() (*types.DumpResult, error)
}

// BaseDumper struct which holds basic configuration for all dumpers
type BaseDumper struct {
	Target configuration.Target
	Dir    string
}