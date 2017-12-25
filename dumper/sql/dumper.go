package sql

import (
	"database/sql"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/SantoDE/datahamster/dumper"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/worker/configuration"
	"github.com/SantoDE/datahamster/worker/types"
	"time"
)

var _ dumper.Dumper = (*Dumper)(nil)

// Dumper struct for the SQL Dumper
type Dumper struct {
	dumper.BaseDumper
	dumper   *mysqldump.Dumper
	Database sql.DB
}

// NewSQLDumper function to create new dumper
func NewSQLDumper(config configuration.DatabaseConfiguration) *Dumper {
	d := new(Dumper)
	d.Config = config
	d.Dir = config.SQL.TempDir
	return d
}

func (d *Dumper) register() error {

	d.connect(d.Config)

	dumper, err := mysqldump.Register(&d.Database, d.Dir, time.RFC3339)

	if err != nil {
		log.Errorf("Error Registering MySql Dump: %s", err)
		return err
	}

	d.dumper = dumper

	return nil
}

// Dump Method to really ceate the Database Dump
func (d *Dumper) Dump() (*types.DumpResult, error) {

	err := d.register()

	if err != nil {
		log.Errorf("Error registering to create MySql Dump: %s", err)
		return &types.DumpResult{Success: false}, err
	}

	dumpPath, err := d.dumper.Dump()

	if err != nil {
		log.Errorf("Error Dumping MySql Dump: %s", err)
		return &types.DumpResult{Success: false}, err
	}

	result := types.DumpResult{
		Success: true,
		Path:    dumpPath,
	}

	return &result, nil
}
