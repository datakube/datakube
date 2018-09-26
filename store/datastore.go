package store

import (
	"github.com/asdine/storm"
	"github.com/datakube/datakube/types"
	"go.etcd.io/bbolt"
	"time"
)

//Datastore struct to hold Datastore Information
type DataStore struct {
	Path string
	db   *storm.DB
}

//NewStore creates a new Datastore
func NewStore(path string) (*DataStore, error) {
	s := &DataStore{
		Path: path,
	}

	return s, nil
}

//Open opens a database connection
func (s *DataStore) Open() error {
	db, err := storm.Open(s.Path, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))

	if err != nil {
		return err
	}
	s.db = db

	s.setup()
	return nil
}

//Close closes a database connection
func (s *DataStore) Close() error {
	err := s.db.Close()

	if err != nil {
		return err
	}

	return nil
}

func (s *DataStore) setup() error {

	_ = s.db.Init(&types.DumpFile{})
	_ = s.db.Init(&types.Job{})

	return nil
}
