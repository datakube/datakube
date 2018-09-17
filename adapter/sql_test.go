package adapter

import (
	"fmt"
	"github.com/SantoDE/datahamster/internal/test"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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

	testCreds := Credentials{
		user: "user",
		password: "password",
		host: "host",
		database: "testdb",
		port: "3306",
	}

	mysqldumpMock := test.MysqlDumpMock{}
	mysqldumpMock.On("Dump", testCreds).Return()

	sqlAdapter := Sql{
		cli: &mysqldumpMock,
		creds:testCreds,
	}

	mysqldumpMock.Success = true

	res, _ := sqlAdapter.Dump("abc")

	assert.Equal(t, res.Success, true)
	assert.Equal(t, res.TargetName, "abc")
	assert.NotNil(t, res.TemporaryFile, "")
	data, _ := ioutil.ReadFile(res.TemporaryFile)
	dataString := fmt.Sprintf("%s", data)
	assert.Equal(t,"Hello World", dataString)

	mysqldumpMock.Success = false

	res, err := sqlAdapter.Dump("abc")

	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.TargetName, "")
	assert.NotNil(t, res.TemporaryFile, "")
	assert.Equal(t, err.Error(), "Test Dump Error")
}

func TestCreateSqlCommand(t *testing.T) {
	res := createSqlCommandString("testhost", "testport", "testdb", "testuser", "testpw")

	assert.Equal(t, "-Ptestport -htesthost -utestuser -ptestpw testdb", res)
}