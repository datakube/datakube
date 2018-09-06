package file

import (
	"github.com/SantoDE/datahamster/types"
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

	storage := NewFileStorage("/tmp")
	storage.SaveFile(*f)

	file, err := os.Open("/tmp/test_file")

	if err != nil {
		t.Fatalf("Error while opening the file: %s", err)
	}

	info, _ := file.Stat()

	if info.Size() <= 0 {
		t.Fatalf("Error Dumping: Got an empty file - no data saved")
	}
}
