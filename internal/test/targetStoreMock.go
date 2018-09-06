package test

import (
	"github.com/stretchr/testify/mock"
	"github.com/SantoDE/datahamster/types"
)

type TargetStoreMock struct{
	mock.Mock
}

func (m *TargetStoreMock) OneByName(name string) (types.Target, error) {
	args := m.Called(name)
	return args.Get(0).(types.Target), args.Error(1)
}

func (m *TargetStoreMock) OneById(id int) (types.Target, error) {
	args := m.Called(id)
	return args.Get(0).(types.Target) , args.Error(1)
}

func (m *TargetStoreMock) Save(t types.Target) (types.Target, error) {
	args := m.Called(t)
	return t,  args.Error(1)
}