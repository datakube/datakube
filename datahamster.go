package datahamster

import (
	"github.com/SantoDE/datahamster/services"
)

// DataStore defines the interface to manage the data.
type DataStore interface {
	Open() error
	Close() error
	MigrateData() error
}

//Services Type to expose Services to RPC and HTTP
type Services struct {
	DumperService services.DumperService
}
