package test

import (
	"errors"
	"github.com/stretchr/testify/mock"
)

type MysqlDumpMock struct {
	mock.Mock
	Success bool
}

func (m MysqlDumpMock) Dump(host string, port string, database string, user string, password string) ([]byte, error) {

	if m.Success {
		return []byte("Hello World"), nil
	}

	return []byte(""), errors.New("Test Dump Error")
}
