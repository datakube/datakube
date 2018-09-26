package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFiles(dfs dumpfileStore) func(*gin.Context) {
	return func(c *gin.Context) {
		dumpFiles, err := dfs.ListAllDumpFiles()

		code := http.StatusOK
		msg := ""

		if err != nil {
			code = http.StatusInternalServerError
			msg = err.Error()
		}

		c.JSON(code, gin.H{
			"dumps":  dumpFiles,
			"error": msg,
		})
	}
}

