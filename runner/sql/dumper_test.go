package sql

import (
	"github.com/SantoDE/datahamster/configuration"
	"reflect"
	"testing"
)

func TestNewSqlDumperOk(t *testing.T) {

	dir := "/tmp/test"

	dbConfig := configuration.DatabaseConfiguration{
		DatabasePassword: "test",
		DatabaseUserName: "test",
		DatabaseName:     "test",
		DatabaseHost:     "test",
		DatabaseType:     "test",
		SQL: configuration.SQLDatabaseConfiguration{
			TempDir: dir,
		},
	}

	config := configuration.Target{
		TargetType: "mysql",
		Name: "testtarget",
		Schedule: *new(configuration.ScheduleConfiguration),
		DBConfig: dbConfig,
	}

	adapter := NewSQLDumper(config)

	if !reflect.DeepEqual(config.DBConfig, adapter.Target.DBConfig) {
		t.Fatalf("Error reading database config: expected %+v, got %+v", config, adapter.Target.DBConfig)
	}

	if !reflect.DeepEqual(dir, adapter.Dir) {
		t.Fatalf("Error reading dir: expected %s, got %s", dir, adapter.Dir)
	}
}
