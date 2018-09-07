package test

import (
	"errors"
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/mock"
)

type JobStoreMock struct {
	mock.Mock
	Success bool
}

func (m JobStoreMock) GetLatestJobByTargetName(targetName string) (types.Job, error) {
	args := m.Called(targetName)
	return args.Get(0).(types.Job), args.Error(1)
}

func (m JobStoreMock) AllJobsByTargetName(targetName string) ([]types.Job, error) {
	args := m.Called(targetName)
	return nil, args.Error(1)
}

func (m JobStoreMock) SaveJob(j types.Job) (types.Job, error) {
	args := m.Called(j)
	return j, args.Error(1)
}

func (m JobStoreMock) AllQueued() ([]types.Job, error) {
	args := m.Called()
	return args.Get(0).([]types.Job), args.Error(1)
}

func (m JobStoreMock) ListAllJobs() ([]types.Job, error) {

	var jobs []types.Job

	if !m.Success {
		return jobs, errors.New("ListAllJobs error")
	}

	jobs = append(jobs, types.Job{ID: 1})
	jobs = append(jobs, types.Job{ID: 2})

	return jobs, nil
}
