package dumper

import (
	"github.com/datakube/datakube/adapter"
	"github.com/datakube/datakube/log"
	"github.com/datakube/datakube/types"
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
