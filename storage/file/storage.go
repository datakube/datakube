package file

import (
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/types"
	"io/ioutil"
	"strings"
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

	s := []string{storage.dir, f.Name}
	path := strings.Join(s, "/")

	log.Debugf("Saving File to path %s", path)
	err := ioutil.WriteFile(path, f.Data, 0644)

	f.Path = path

	if err != nil {
		log.Errorf("Error Saving the file %s to location %s on File Storage: %s", f.Path, path, err)
		return *new(types.File), err
	}

	return f, nil
}

// ReadFile function to read a specific file
func (*Storage) ReadFile(path string) ([]byte, error){
	dat, err := ioutil.ReadFile(path)

	return dat, err
}

// ListFiles function to list all files
func (storage *Storage) ListFiles() {

}

// DeleteFile function to delete a file
func (*Storage) DeleteFile() {

}
