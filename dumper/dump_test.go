package dumper

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestDumpAdapter struct{}

func (t TestDumpAdapter) Dump(targetName string) (types.DumpResult, error) {

	result := *new(types.DumpResult)

	result.Success = true
	result.TemporaryFile = "/tmp/testfile"
	result.TargetName = "TestTarget"

	return result, nil
}

func TestDumpJobRunt(t *testing.T) {

	events := make(chan types.DumpResult)
	tda := new(TestDumpAdapter)

	go Run("TestTarget", tda, events)

	var eventHit bool

	select {
	case dump := <-events:
		eventHit = true
		assert.Equal(t, dump.Success, true)
		assert.Equal(t, dump.TargetName, "TestTarget")
		assert.Equal(t, dump.TemporaryFile, "/tmp/testfile")
	}

	assert.Equal(t, true, eventHit)
}
