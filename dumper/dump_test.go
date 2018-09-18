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

	tda := new(TestDumpAdapter)

	res := Run("TestTarget", tda)

	assert.Equal(t, res.Success, true)
	assert.Equal(t, res.TargetName, "TestTarget")
	assert.Equal(t, res.TemporaryFile, "/tmp/testfile")

}
