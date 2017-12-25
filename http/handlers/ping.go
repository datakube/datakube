package handlers

import (
	"github.com/SantoDE/datahamster/bolt"
	"github.com/gin-gonic/gin"
)

//PingHandler to hold Pinghandler information
type PingHandler struct {
	BaseHandler
}

var _ Handler = (*PingHandler)(nil)

//NewPingHandler to create a new Pinghandler
func NewPingHandler(as *bolt.AgentService) *PingHandler {
	h := new(PingHandler)
	h.setupHandler(as)

	return h
}

//GET Ping
func (h *PingHandler) GET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
