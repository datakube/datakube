package file

import (
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
