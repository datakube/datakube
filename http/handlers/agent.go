package handlers

import (
	"github.com/SantoDE/datahamster/bolt"
	"github.com/SantoDE/datahamster/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AgentHandler struct {
	BaseHandler
}

var _ Handler = (*AgentHandler)(nil)

func NewAgentHandler(as *bolt.AgentService) *AgentHandler {
	ah := new(AgentHandler)
	ah.setupHandler(as)

	return ah
}

func (h *AgentHandler) POST(c *gin.Context) {
	var newAgent types.Agent
	// This will infer what binder to use depending on the content-type header.
	if err := c.BindJSON(&newAgent); err == nil {
		agent := h.agentService.Create(newAgent.Name)
		c.JSON(http.StatusOK, gin.H{"success": "true", "agent": agent})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
