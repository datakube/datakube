package adapter

import (
	"fmt"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/types"
	"io/ioutil"
	"os/exec"
)

type mysqlDump interface {
	Dump(host string, port string, database string, user string, password string) ([]byte, error)
}

type mysqlDumpCli struct {}

type Sql struct {
	creds  Credentials
	cli mysqlDump
}

type Credentials struct {
	host     string
	port     string
	database string
	user     string
	password string
}

func newSqlAdapter(host string, port string, database string, user string, password string) Sql {
	return Sql{
		creds: Credentials{
			host:     host,
			port:     port,
			database: database,
			user:     user,
			password: password,
		},
		cli: new(mysqlDumpCli),
	}
}

func (s Sql) Dump(targetName string) (types.DumpResult, error) {

	data, err := s.cli.Dump(s.creds.host, s.creds.port, s.creds.database, s.creds.user, s.creds.password)

	if err != nil {
		log.Errorf("Error Dumping MySql Dump: %s", err)
		return types.DumpResult{Success: false}, err
	}

	tempFile, err := ioutil.TempFile("", "")
	ioutil.WriteFile(tempFile.Name(), data, 0755)

	result := types.DumpResult{
		Success:       true,
		TemporaryFile: tempFile.Name(),
		TargetName:    targetName,
	}

	return result, nil
}

func (m mysqlDumpCli) Dump(host string, port string, database string, user string, password string) ([]byte, error) {
	cmd := exec.Command("mysqldump ", createSqlCommandString(host, port, database, user, password))
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}

	return bytes, nil
}

func createSqlCommandString(host string, port string, database string, user string, password string) string {
	cmd := fmt.Sprintf("%s %s %s %s %s", fmt.Sprintf("-P%s", port),fmt.Sprintf("-h%s", host), fmt.Sprintf("-u%s", user), fmt.Sprintf("-p%s", password), fmt.Sprintf("%s", database))
	return cmd
}