package handlers

import (
	"github.com/SantoDE/datahamster/bolt"
)

type Handler interface {
}

type BaseHandler struct {
	agentService *bolt.AgentService
}

func (b *BaseHandler) setupHandler(as *bolt.AgentService) {
	b.agentService = as
}
