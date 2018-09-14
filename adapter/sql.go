package adapter

import (
	"database/sql"
	"fmt"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/types"
	"github.com/ziutek/mymysql/godrv"
	"io/ioutil"
	"time"
)

type mysqlDump interface {
	Dump() (string, error)
}

type Sql struct {
	creds  credentials
	dumper mysqlDump
}

type credentials struct {
	host     string
	port     string
	database string
	user     string
	password string
}

func newSqlAdapter(host string, port string, database string, user string, password string) Sql {
	return Sql{
		creds: credentials{
			host:     host,
			port:     port,
			database: database,
			user:     user,
			password: password,
		},
	}
}

func (s Sql) Dump(targetName string) (types.DumpResult, error) {

	dumpPath, err := s.dumper.Dump()

	if err != nil {
		log.Errorf("Error Dumping MySql Dump: %s", err)
		return types.DumpResult{Success: false}, err
	}

	result := types.DumpResult{
		Success:       true,
		TemporaryFile: dumpPath,
		TargetName:    targetName,
	}

	return result, nil
}

func (s *Sql) connect() {
	// Register the mymysql driver
	godrv.Register("SET NAMES utf8")

	connectionString := fmt.Sprintf("tcp:%s:%s*%s/%s/%s", s.creds.host, s.creds.port, s.creds.database, s.creds.user, s.creds.password)

	log.Debugf("Trying to connect with %s", connectionString)

	db, err := sql.Open("mymysql", connectionString)

	if err != nil {
		log.Errorf("Error connecting to MySql Datavase %s", err)
	}

	dir, err := ioutil.TempDir("", "dump")

	if err != nil {

	}

	dumper, err := mysqldump.Register(db, dir, time.RFC3339)

	if err != nil {

	}

	s.dumper = dumper
}
