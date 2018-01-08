package sql

import (
	"database/sql"
	"github.com/JamesStewy/go-mysqldump"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/types"
	"time"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/runner"
)

// Dumper struct for the SQL Dumper
type Dumper struct {
	runner.BaseDumper
	dumper   *mysqldump.Dumper
	Database sql.DB
}

// NewSQLDumper function to create new dumper
func NewSQLDumper(target configuration.Target) *Dumper {
	d := new(Dumper)
	d.Target = target
	d.Dir = target.DBConfig.SQL.TempDir
	return d
}

func (d *Dumper) register() error {

	d.connect(d.Target.DBConfig)

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
		TemporaryFile: dumpPath,
		TargetName: d.Target.Name,
	}

	return &result, nil
}
