package main

import (
	"github.com/SantoDE/datahamster/worker/configuration"
	"github.com/SantoDE/datahamster/worker/dumper"
	"github.com/SantoDE/datahamster/worker/dumper/sql"
	"github.com/SantoDE/datahamster/worker/log"
	"os"
)

// Hamster Structs which knows the DB Configuration and the dumper needed for that Config
type Hamster struct {
	Database *configuration.DatabaseConfiguration
	Dumper   dumper.Dumper
}

// NewHamster Create New Hamster with the given DB Config
func NewHamster(configuration configuration.DatabaseConfiguration) *Hamster {
	hamster := new(Hamster)
	hamster.Dumper = sql.NewSQLDumper(configuration, os.TempDir())
	return hamster
}

func (hamster *Hamster) run() {
	result, err := hamster.Dumper.Dump()

	if err != nil {
		log.Errorf("Error connecting to MySql Datavase %s", err)
	}

	log.Debugf("Result %s", result.Success)
}
