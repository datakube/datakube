package handlers

import "github.com/gin-gonic/gin"

//GET Ping
func GetPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
