package store

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/asdine/storm"
	"github.com/coreos/bbolt"
	"time"
)

// DataStore defines the interface to manage the data.
type DataStore interface {
	Open() error
}

//Datastore struct to hold Datastore Information
type Store struct {
	Path string
	db   *storm.DB
}

//NewStore creates a new Datastore
func NewStore(path string) (*Store, error) {
	s := &Store{
		Path: path,
	}

	return s, nil
}

//Open opens a database connection
func (s *Store) Open() error {
	db, err := storm.Open(s.Path, storm.BoltOptions(0600, &bolt.Options{Timeout: 1 * time.Second}))

	if err != nil {
		return err
	}
	s.db = db

	s.setup()
	return nil
}

//Close closes a database connection
func (s *Store) Close() error {
	err := s.db.Close()

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) setup() error {
	err := s.db.Init(&types.Dumper{})

	if err != nil {
		return err
	}

	return nil
}
