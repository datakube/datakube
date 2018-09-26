package test

import (
	"github.com/datakube/datakube/types"
	"github.com/stretchr/testify/mock"
)

type StorageMock struct {
	mock.Mock
	Success bool
}

func (m StorageMock) ReadFile(path string) ([]byte, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func (m StorageMock) SaveFile(file types.File) (types.File, error) {
	return file, nil
}

func (m StorageMock) DeleteFile() {}

func (m StorageMock) ListFiles() {}
