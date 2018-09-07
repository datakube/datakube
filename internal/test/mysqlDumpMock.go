package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
)

type MysqlDumpMock struct {
	mock.Mock
	Success bool
}

func (m *MysqlDumpMock) Dump() (string, error) {

	if m.Success {
		return "/test/file", nil
	}

	return "", errors.New("Test Dump Error")
}
