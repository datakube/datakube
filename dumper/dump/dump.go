package dump

import (
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/runner"
	"github.com/SantoDE/datahamster/runner/sql"
	"github.com/SantoDE/datahamster/types"
	"github.com/robfig/cron"
)

var _ cron.Job = (*DumpJob)(nil)

type DumpJob struct {
	Runner runner.DumpRunner
	events chan<- types.DumpResult
}

func NewDumpJob(target *datakube.Target, dumpResults chan<- types.DumpResult) *DumpJob {

	dj := new(DumpJob)
	var runner runner.DumpRunner

	switch target.Type {
	case "mysql":
		runner = sql.NewSQLDumper(target)
	}

	dj.events = dumpResults
	dj.Runner = runner

	return dj
}

func (p *DumpJob) Run() {
	log.Debug("Running Dump")
	res, err := p.Runner.Dump()
	log.Debug("Dump done")
	if err != nil {
		log.Debug("Error during Dump %s", err.Error())
	}

	p.events <- *res
}
