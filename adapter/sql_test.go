package adapter

import (
	"github.com/SantoDE/datahamster/internal/test"
	"github.com/SantoDE/datahamster/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNewSqlAdapter(t *testing.T) {
	a := newSqlAdapter("h", "p", "d", "u", "p")
	assert.Equal(t, a.creds.host, "h")
	assert.Equal(t, a.creds.port, "p")
	assert.Equal(t, a.creds.database, "d")
	assert.Equal(t, a.creds.user, "u")
	assert.Equal(t, a.creds.port, "p")

}

func TestSql_DumpOk(t *testing.T) {

	mysqldumpMock := test.MysqlDumpMock{}
	mysqldumpMock.On("Dump", "").Return(types.DumpResult{Success: true}, nil)

	sqlAdapter := Sql{
		dumper: &mysqldumpMock,
	}

	mysqldumpMock.Success = true

	res, _ := sqlAdapter.Dump("abc")

	assert.Equal(t, res.Success, true)
	assert.Equal(t, res.TargetName, "abc")
	assert.Equal(t, res.TemporaryFile, "/test/file")

	mysqldumpMock.Success = false

	res, err := sqlAdapter.Dump("abc")

	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.TargetName, "")
	assert.Equal(t, res.TemporaryFile, "")
	assert.Equal(t, err.Error(), "Test Dump Error")
}