package runner

import (
	"github.com/SantoDE/datahamster/types"
)

const TYPE_MSQL string = "mysql"

// Dumper Interface for all dumpers
type DumpRunner interface {
	Dump() (*types.DumpResult, error)
}
