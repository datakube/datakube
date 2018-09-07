package test

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/mock"
)

type TargetStoreMock struct {
	mock.Mock
	Success bool
}

func (m TargetStoreMock) OneByName(name string) (types.Target, error) {
	args := m.Called(name)
	return args.Get(0).(types.Target), args.Error(1)
}

func (m TargetStoreMock) OneById(id int) (types.Target, error) {
	args := m.Called(id)
	return args.Get(0).(types.Target), args.Error(1)
}

func (m TargetStoreMock) Save(t types.Target) (types.Target, error) {
	args := m.Called(t)
	return t, args.Error(1)
}

func (m TargetStoreMock) ListTargets() []types.Target {

	var targets []types.Target

	if !m.Success {
		return targets
	}

	targets = append(targets, types.Target{Name: "Test Target 1"})
	targets = append(targets, types.Target{Name: "Test Target 2"})

	return targets
}
