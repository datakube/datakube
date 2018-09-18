package file

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNewFileStorageOkay(t *testing.T) {

	dir := "/tmp/test"

	storage := NewFileStorage(dir)

	if !reflect.DeepEqual(dir, storage.dir) {
		t.Fatalf("Error reading dir: expected %s, got %s", dir, storage.dir)
	}
}

func TestSaveResultOk(t *testing.T) {

	f := new(types.File)
	f.Name = "test_file"
	f.Data = []byte("TEST DATA")

	tempdir, _ := ioutil.TempDir("", "")

	defer os.RemoveAll(tempdir)

	storage := NewFileStorage(tempdir)
	savedFile, err := storage.SaveFile(*f)

	assert.Nil(t, err)

	file, err := os.Open(savedFile.Path)

	if err != nil {
		t.Fatalf("Error while opening the file: %s", err)
	}

	info, _ := file.Stat()

	if info.Size() <= 0 {
		t.Fatalf("Error Dumping: Got an empty file - no data saved")
	}
}

func TestStorage_ReadFile(t *testing.T) {
	f := new(types.File)
	f.Name = "test_file"
	f.Data = []byte("TEST DATA")

	tempDir, _ := ioutil.TempDir("", "")

	defer os.RemoveAll(tempDir)

	storage := NewFileStorage(tempDir)
	savedFile, _ := storage.SaveFile(*f)

	data, err := storage.ReadFile(savedFile.Path)

	assert.Nil(t, err)
	assert.Equal(t, string(data[:]), "TEST DATA")

}
