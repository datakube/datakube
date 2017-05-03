package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/santode/datahamster/worker/configuration"
	"github.com/santode/datahamster/worker/types"
	"gopkg.in/ory-am/dockertest.v3"
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

	config := configuration.DatabaseConfiguration{
		DatabasePassword: "test",
		DatabaseUserName: "test",
		DatabaseName:     "testdb",
		DatabasePort:     resource.GetPort("3306/tcp"),
		DatabaseHost:     "localhost",
	}

	adapter := NewSQLDumper(config, "/tmp")

	err = adapter.register(config)

	if err != nil {
		err = pool.Purge(resource)
		t.Fatalf("Error while registering: %s", err)
	}

	result, err := adapter.Dump()

	if err != nil {
		err = pool.Purge(resource)
		t.Fatalf("Error while dumping: %s", err)
	}

	expectedResult := &types.DumpResult{
		Success: true,
	}

	if !reflect.DeepEqual(expectedResult, result) {
		pool.Purge(resource)
		t.Fatalf("Error dumping: expected %+v, got %+v", expectedResult, result)
	}

	err = pool.Purge(resource)
}
