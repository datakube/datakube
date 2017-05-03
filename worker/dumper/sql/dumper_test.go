package sql

import (
	"github.com/santode/datahamster/worker/configuration"
	"reflect"
	"testing"
)

func TestNewSqlDumperOk(t *testing.T) {

	config := configuration.DatabaseConfiguration{
		DatabasePassword: "test",
		DatabaseUserName: "test",
		DatabaseName:     "test",
		DatabaseHost:     "test",
		DatabaseType:     "test",
	}

	dir := "/tmp/test"

	adapter := NewSQLDumper(config, "/tmp/test")

	if !reflect.DeepEqual(config, adapter.Config) {
		t.Fatalf("Error reading database config: expected %+v, got %+v", config, adapter.Config)
	}

	if !reflect.DeepEqual(dir, adapter.Dir) {
		t.Fatalf("Error reading dir: expected %s, got %s", dir, adapter.Dir)
	}
}
