package api

import "github.com/gin-gonic/gin"

func GetJobs(js jobStore) func(*gin.Context) {
	return func(c *gin.Context) {
		jobs, err := js.ListAllJobs()

		if err != nil {

		}

		c.JSON(200, gin.H{
			"jobs": jobs,
		})
	}
}
