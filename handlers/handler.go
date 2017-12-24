package handlers

import (
	"github.com/SantoDE/datahamster/bolt"
)

type Handler interface {

}

type BaseHandler struct {
	agentService *bolt.AgentService
}