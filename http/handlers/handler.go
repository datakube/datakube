package handlers

import (
	"github.com/SantoDE/datahamster/bolt"
)

//Handler Interface
type Handler interface {
}

//BaseHandler Interface
type BaseHandler struct {
	agentService *bolt.AgentService
}

func (b *BaseHandler) setupHandler(as *bolt.AgentService) {
	b.agentService = as
}
