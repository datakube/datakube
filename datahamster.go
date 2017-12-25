package datahamster

import (
	"github.com/SantoDE/datahamster/bolt"
)

//Services Type to expose Services to RPC and HTTP
type Services struct {
	AgentService *bolt.AgentService
}
