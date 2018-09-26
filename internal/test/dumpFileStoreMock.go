package test

import (
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
