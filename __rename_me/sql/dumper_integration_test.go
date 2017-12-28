package sql

import (
	"database/sql"
	"fmt"
	"github.com/SantoDE/datahamster/worker/configuration"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ory-am/dockertest.v3"
	"os"
	"reflect"
	"testing"
)

var db *sql.DB

func TestSqlDumpOk(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	pool, err := dockertest.NewPool("")

	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("santode/datahamster-worker-integration-test-db", "latest", []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_DATABASE=testdb", "MYSQL_USER=test", "MYSQL_PASSWORD=test"})
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("test:test@(localhost:%s)/testdb", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		t.Fatalf("Could not connect to docker database container: %s", err)
		err = pool.Purge(resource)

	}

	dir := "/tmp"

	config := configuration.DatabaseConfiguration{
		DatabasePassword: "test",
		DatabaseUserName: "test",
		DatabaseName:     "testdb",
		DatabasePort:     resource.GetPort("3306/tcp"),
		DatabaseHost:     "localhost",
		SQL: configuration.SQLDatabaseConfiguration{
			TempDir: dir,
		},
	}

	adapter := NewSQLDumper(config)

	err = adapter.register()

	if err != nil {
		err = pool.Purge(resource)
		t.Fatalf("Error while registering: %s", err)
	}

	result, err := adapter.Dump()

	if err != nil {
		err = pool.Purge(resource)
		t.Fatalf("Error while dumping: %s", err)
	}

	if reflect.TypeOf(result).String() == "DumpResult" {
		pool.Purge(resource)
		t.Fatalf("Error Dumping: Did not receive a DumpResult, instead received %+v", result)
	}

	if reflect.ValueOf(result.Success).Bool() != true {
		pool.Purge(resource)
		t.Fatalf("Error Dumping: expected result %+v, got %+v", true, reflect.ValueOf(result.Success).Bool())
	}

	path := reflect.ValueOf(result.Path).String()

	if path == "" {
		pool.Purge(resource)
		t.Fatalf("Error Dumping: Got an empty filename - probably no dump created")
	}

	file, err := os.Open(path)

	if err != nil {
		err = pool.Purge(resource)
		t.Fatalf("Error while opening the file: %s", err)
	}

	info, _ := file.Stat()

	if info.Size() <= 0 {
		pool.Purge(resource)
		t.Fatalf("Error Dumping: Got an empty file - no data dumped")
	}

	err = pool.Purge(resource)
}
