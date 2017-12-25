package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/SantoDE/datahamster/bolt"
)


type PingHandler struct {
	BaseHandler
}

var _ Handler = (*PingHandler)(nil)

func NewPingHandler(as *bolt.AgentService) *PingHandler {
	h := new(PingHandler)
	h.setupHandler(as)

	return h
}


func (h *PingHandler) GET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
