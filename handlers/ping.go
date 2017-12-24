package handlers

import "github.com/gin-gonic/gin"

type PingHandler struct {
	BaseHandler
}

func (h *PingHandler)GET(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
