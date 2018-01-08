package jobs_test

import (
	"github.com/SantoDE/datahamster/dumper/jobs"
	"testing"
	"github.com/SantoDE/datahamster/types"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/stretchr/testify/assert"
)


type TestDumpRunner struct {}

func (tdr *TestDumpRunner) Dump() (*types.DumpResult, error){
	result := new(types.DumpResult)

	result.Success = true
	result.TemporaryFile = "/tmp/testfile"
	result.TargetName="TestTarget"

	return result, nil
}

func TestDumpJobRunt(t *testing.T) {

	cfg := configuration.Target{
		TargetType: "mysql",
		Name: "TestTarget",
		Schedule: *new(configuration.ScheduleConfiguration),
		DBConfig: *new(configuration.DatabaseConfiguration),

	}
	events := make(chan types.DumpResult)
	tdr := new(TestDumpRunner)
	job := jobs.NewDumpJob(&cfg, events)
	job.Runner = tdr

	go job.Run()

	var eventHit bool

	select {
	case dump := <- events:
		eventHit = true
		assert.Equal(t, dump.Success, true)
		assert.Equal(t, dump.TargetName, "TestTarget")
		assert.Equal(t, dump.TemporaryFile, "/tmp/testfile")
	}

	assert.Equal(t, true, eventHit)
}
