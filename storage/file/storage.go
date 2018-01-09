package file

import (
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/storage"
	"os"
	"path"
	"github.com/SantoDE/datahamster/types"
	"io/ioutil"
)

var _ storage.Storage = (*Storage)(nil)

// Storage struct for the File Storage
type Storage struct {
	storage.BaseStorage
	dir string
}

// NewFileStorage function to create new dumper
func NewFileStorage(dir string) *Storage {
	storage := new(Storage)
	storage.dir = dir
	return storage
}

// SaveFile function to save a new file
func (storage *Storage) SaveFile(f types.File) (types.File, error) {

	path := path.Join(storage.dir, f.Name)

	log.Debugf("Saving File from %s on File Storage to Path %s", f.Path, path)

	err := ioutil.WriteFile(path, f.Data, 0644)

	f.Path = path

	if err != nil {
		log.Errorf("Error Moving the file %s to location %s on File Storage: %s", f.Path, path, err)
		return *new(types.File), err
	}

	return f, nil
}

// ReadFile function to read a specific file
func (*Storage) ReadFile() {

}

// ListFiles function to list all files
func (storage *Storage) ListFiles() {

}

// DeleteFile function to delete a file
func (*Storage) DeleteFile() {

}
