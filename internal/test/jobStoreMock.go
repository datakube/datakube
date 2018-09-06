package test

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/mock"
)

type MockedJobStore struct {
	mock.Mock
}

func (m *MockedJobStore) GetLatestJobByTargetName(targetName string) (types.Job, error) {
	args := m.Called(targetName)
	return args.Get(0).(types.Job), args.Error(1)
}

func (m *MockedJobStore) AllJobsByTargetName(targetName string) ([]types.Job, error) {
	args := m.Called(targetName)
	return nil, args.Error(1)
}

func (m *MockedJobStore) SaveJob(j types.Job) (types.Job, error) {
	args := m.Called(j)
	return j, args.Error(1)
}

func (m *MockedJobStore) AllQueued() ([]types.Job, error) {
	args := m.Called()
	return args.Get(0).([]types.Job), args.Error(1)
}
