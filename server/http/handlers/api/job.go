package api

import "github.com/gin-gonic/gin"

func (a *ApiHandler) GetJobs(c *gin.Context) {
	jobs, err := a.jobStore.ListAllJobs()

	if err != nil {

	}

	c.JSON(200, gin.H{
		"jobs": jobs,
	})
}
