package dumper

import (
	"github.com/SantoDE/datahamster/adapter"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/types"
)

func Run(targetName string, adapter adapter.DumpAdapter) types.DumpResult{
	log.Debug("Running Dump")
	res, err := adapter.Dump(targetName)
	log.Debug("Dump done")
	if err != nil {
		log.Debug("Error during Dump ", err.Error())
	}

	return res
}
