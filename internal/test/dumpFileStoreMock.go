package test

import (
	"errors"
	"github.com/datakube/datakube/types"
	"github.com/stretchr/testify/mock"
)

type DumpFileStoreMock struct {
	mock.Mock
	Success bool
}

func (m DumpFileStoreMock) LoadOneDumpFileByTarget(targetName string) (types.DumpFile, error) {
	args := m.Called(targetName)
	return args.Get(0).(types.DumpFile), args.Error(1)
}

func (m DumpFileStoreMock) LoadOneDumpFileByName(targetName string) (types.DumpFile, error) {
	args := m.Called(targetName)
	return args.Get(0).(types.DumpFile), args.Error(1)
}

func (m DumpFileStoreMock) ListAllDumpFiles() ([]types.DumpFile, error) {

	if !m.Success {
		return nil, errors.New("Test Error")
	}

	dfs1 := types.DumpFile{
		ID: 1337,
		Target: "testTarget",
		File: types.File{
			Name: "testfile",
		},
	}

	dfs2 := types.DumpFile{
		ID: 1337,
		Target: "testTarget",
		File: types.File{
			Name: "testfile",
		},
	}

	return []types.DumpFile{dfs1, dfs2}, nil
}
