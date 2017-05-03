package sql

import (
	"database/sql"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/SantoDE/datahamster/worker/configuration"
	"github.com/SantoDE/datahamster/worker/dumper"
	"github.com/SantoDE/datahamster/worker/log"
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
func NewSQLDumper(config configuration.DatabaseConfiguration, dir string) *Dumper {
	d := new(Dumper)
	d.Config = config
	d.Dir = dir
	return d
}

func (d *Dumper) register(config configuration.DatabaseConfiguration) error {

	d.connect(config)

	dumper, err := mysqldump.Register(&d.Database, d.Dir, time.ANSIC)

	if err != nil {
		log.Errorf("Error Registering MySql Dump: %s", err)
		return err
	}

	d.dumper = dumper

	return nil
}

// Dump Method to really ceate the Database Dump
func (d *Dumper) Dump() (*types.DumpResult, error) {

	err := d.dumper.Dump()

	if err != nil {
		log.Errorf("Error Dumping MySql Dump: %s", err)
		return &types.DumpResult{Success: false}, err
	}

	result := types.DumpResult{
		Success: true,
	}

	return &result, nil
}
