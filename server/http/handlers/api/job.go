package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetJobs(js jobStore) func(*gin.Context) {
	return func(c *gin.Context) {
		jobs, err := js.ListAllJobs()

		code := http.StatusOK
		msg := ""

		if err != nil {
			code = http.StatusInternalServerError
			msg = err.Error()
		}

		c.JSON(code, gin.H{
			"jobs":  jobs,
			"error": msg,
		})
	}
}
