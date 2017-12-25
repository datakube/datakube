package bolt

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/asdine/storm"
	"github.com/coreos/bbolt"
	"time"
)

type Datastore struct {
	Path string
	db   *storm.DB
}

func NewStore(path string) (*Datastore, error) {
	s := &Datastore{
		Path: path,
	}

	return s, nil
}

func (d *Datastore) Open() error {
	db, err := storm.Open(d.Path, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))

	if err != nil {
		return err
	}
	d.db = db

	d.setup()
	return nil
}

func (d *Datastore) setup() error {
	err := d.db.Init(&types.Agent{})

	if err != nil {
		return err
	}

	return nil
}
