package api

import "github.com/gin-gonic/gin"

//GET Ping
func (a *ApiHandler) GetTargets(c *gin.Context) {
	c.JSON(200, gin.H{
		"targets": a.targetStore.ListTargets(),
	})
}
