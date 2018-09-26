package dumper

import (
	"github.com/datakube/datakube/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testDumpAdapter struct {
	Success bool
}

func (t testDumpAdapter) Dump(targetName string) (types.DumpResult, error) {

	result := *new(types.DumpResult)

	result.Success = t.Success

	if t.Success {
		result.TemporaryFile = "/tmp/testfile"
	}

	result.TargetName = "TestTarget"

	return result, nil
}

func TestDumpJobRunt(t *testing.T) {

	tda := new(testDumpAdapter)

	tda.Success = true
	res := Run("TestTarget", tda)

	assert.Equal(t, true, res.Success)
	assert.Equal(t, "TestTarget", res.TargetName)
	assert.Equal(t, "/tmp/testfile", res.TemporaryFile)

	tda.Success = false
	res = Run("TestTarget", tda)

	assert.Equal(t, false, res.Success)
	assert.Equal(t, "TestTarget", res.TargetName)
	assert.Equal(t, "", res.TemporaryFile)
}
