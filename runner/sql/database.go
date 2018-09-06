package sql

import (
	"database/sql"
	"fmt"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/ziutek/mymysql/godrv"
)

func (d *Dumper) connect(creds datakube.Credentials) {
	// Register the mymysql driver
	godrv.Register("SET NAMES utf8")

	connectionString := fmt.Sprintf("tcp:%s:%s*%s/%s/%s", creds.Host, creds.Port, creds.Database, creds.User, creds.Password)

	log.Debugf("Trying to connect with %s", connectionString)

	db, err := sql.Open("mymysql", connectionString)

	if err != nil {
		log.Errorf("Error connecting to MySql Datavase %s", err)
	}

	d.Database = *db
}
