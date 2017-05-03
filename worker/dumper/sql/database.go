package sql

import (
	"database/sql"
	"fmt"
	"github.com/santode/datahamster/worker/configuration"
	"github.com/santode/datahamster/worker/log"
	"github.com/ziutek/mymysql/godrv"
)

func (d *Dumper) connect(config configuration.DatabaseConfiguration) {
	// Register the mymysql driver
	godrv.Register("SET NAMES utf8")

	connectionString := fmt.Sprintf("tcp:%s:%s*%s/%s/%s", config.DatabaseHost, config.DatabasePort, config.DatabaseName, config.DatabaseUserName, config.DatabasePassword)

	db, err := sql.Open("mymysql", connectionString)

	if err != nil {
		log.Errorf("Error connecting to MySql Datavase %s", err)
	}

	d.Database = *db
}
