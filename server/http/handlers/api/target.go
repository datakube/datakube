package api

import "github.com/gin-gonic/gin"

//GET Ping
func GetTargets(ts targetStore) func(*gin.Context) {
	return func(c *gin.Context) {
		targets := ts.ListTargets()

		c.JSON(200, gin.H{
			"targets": targets,
		})
	}
}
