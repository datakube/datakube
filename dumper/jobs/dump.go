package jobs

import (
	"github.com/robfig/cron"
	"fmt"
	"github.com/SantoDE/datahamster/runner"
)

var _ cron.Job = (*DumpJob)(nil)

type DumpJob struct {
	runner runner.DumpRunner
}

func (p *DumpJob) Run() {
	fmt.Printf("I'm running babe")
}
