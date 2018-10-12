package handlers

import (
	"github.com/datakube/datakube/storage"
	"github.com/datakube/datakube/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type dumpFileStore interface {
	LoadOneDumpFileByTarget(targetName string) (types.DumpFile, error)
	LoadOneDumpFileByName(fileName string) (types.DumpFile, error)
}

//POST function to create a new handler
func GetFile(dfs dumpFileStore, storage storage.Storage) func(*gin.Context) {
	return func(c *gin.Context) {
		fileName := c.Param("name")

		if fileName == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No fileName provided",
			})
			return
		}

		dumpFile, error := dfs.LoadOneDumpFileByName(fileName)

		if dumpFile.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No Filename by that name " + fileName + " found",
			})
			return
		}

		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": error.Error(),
			})
			return
		}

		data, err := storage.ReadFile(dumpFile.File.Path)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Type", "application/octet-stream")
		c.Data(http.StatusOK, "application/octet-stream", data)
	}
}

func GetLatestFile(dfs dumpFileStore, storage storage.Storage) func(*gin.Context) {
	return func(c *gin.Context) {
		targetName := c.Param("name")

		if targetName == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No fileName provided",
			})
			return
		}

		dumpFile, error := dfs.LoadOneDumpFileByTarget(targetName)

		if dumpFile.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "No Dumps for target " + targetName + " found",
			})
			return
		}

		if error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": error.Error(),
			})
			return
		}

		fileName := dumpFile.File.Name
		data, err := storage.ReadFile(dumpFile.File.Path)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Type", "application/octet-stream")
		c.Data(http.StatusOK, "application/octet-stream", data)
	}
}
