package handlers

import (
	"github.com/SantoDE/datahamster/services"
	"github.com/SantoDE/datahamster/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AgentHandler struct to hold AgentHandler specific information
type AgentHandler struct {
	BaseHandler
	agentService services.AgentService
}

var _ Handler = (*AgentHandler)(nil)

//NewAgentHandler function to create a new handler agent
func NewAgentHandler(as services.AgentService) *AgentHandler {
	ah := new(AgentHandler)
	ah.agentService = as
	return ah
}

//POST function to create a new handler
func (h *AgentHandler) POST(c *gin.Context) {
	var newAgent types.Agent
	// This will infer what binder to use depending on the content-type header.
	if err := c.BindJSON(&newAgent); err == nil {
		agent, err := h.agentService.Create(newAgent.Name)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"success": "true", "agent": agent})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
