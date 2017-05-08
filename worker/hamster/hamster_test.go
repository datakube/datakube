package hamster

import (
	"github.com/SantoDE/datahamster/storage/file"
	"github.com/SantoDE/datahamster/worker/configuration"
	"github.com/SantoDE/datahamster/worker/dumper/sql"
	"reflect"
	"testing"
)

func TestNewHamsterOk(t *testing.T) {

	dir := "/tmp/test"

	dBConfig := configuration.DatabaseConfiguration{
		DatabasePassword: "test",
		DatabaseUserName: "test",
		DatabaseName:     "test",
		DatabaseHost:     "test",
		DatabaseType:     "test",
		SQL: configuration.SQLDatabaseConfiguration{
			TempDir: dir,
		},
	}

	storageConfig := configuration.StorageConfiguration{
		Type: "file",
		File: configuration.FileStorageConfiguration{
			Dir: dir,
		},
	}

	hamster := NewHamster(dBConfig, storageConfig)

	expectedDumper := sql.NewSQLDumper(dBConfig)
	expectedStorage := file.NewFileStorage(dir)

	if !reflect.DeepEqual(expectedStorage, hamster.Storage) {
		t.Fatalf("Error reading storage: expected %+v, got %+v", expectedStorage, hamster.Storage)
	}

	if !reflect.DeepEqual(expectedDumper, hamster.Dumper) {
		t.Fatalf("Error reading Dumper: expected %+v, got %+v", expectedDumper, hamster.Dumper)
	}
}
