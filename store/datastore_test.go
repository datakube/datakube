package store_test

import (
	"github.com/datakube/datakube/store"
	"io/ioutil"
	"os"
)

type TestDataStore struct {
	*store.DataStore
}

// NewTestDB returns a TestDB using a temporary path.
func NewTestDataStore() *TestDataStore {
	// Retrieve a temporary path.
	f, _ := ioutil.TempFile("", "")
	path := f.Name()
	f.Close()
	os.Remove(path)
	// Open the database.
	db, _ := store.NewStore(f.Name())
	db.Open()
	// Return wrapped type.
	return &TestDataStore{db}
}

// Close and delete Bolt database.
func (db *TestDataStore) Close() {
	defer os.Remove(db.Path)
	db.DataStore.Close()
}
