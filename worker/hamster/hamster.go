package hamster

import (
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/storage/file"
	"github.com/SantoDE/datahamster/worker/configuration"
	"github.com/SantoDE/datahamster/worker/dumper"
	"github.com/SantoDE/datahamster/worker/dumper/sql"
	"os"
)

// Hamster Structs which knows the DB Configuration and the dumper needed for that Config
type Hamster struct {
	Dumper  dumper.Dumper
	Storage storage.Storage
}

// NewHamster Create New Hamster with the given DB Config
func NewHamster(dbConfiguration configuration.DatabaseConfiguration, storage configuration.StorageConfiguration) *Hamster {
	hamster := new(Hamster)
	hamster.Dumper = sql.NewSQLDumper(dbConfiguration)
	hamster.Storage = file.NewFileStorage(storage.File.Dir)
	return hamster
}

func (hamster *Hamster) run() (*storage.File, error) {
	result, err := hamster.Dumper.Dump()

	if err != nil {
		log.Errorf("Error connecting to MySql Database %s", err)
		return nil, err
	}

	log.Debugf("Dump Succesfull - going to save it")

	fileInfo, _ := os.Stat(result.Path)

	name := fileInfo.Name()

	f := new(storage.File)
	f.Path = result.Path
	f.Name = name

	hamster.Storage.SaveFile(*f)

	log.Debugf("Saved File with new Name %s", f.Name)
	return f, nil
}
