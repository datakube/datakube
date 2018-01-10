package handlers

import (
	"github.com/SantoDE/datahamster/services"
	"github.com/gin-gonic/gin"
	"fmt"
	"io"
)

//DumperHandler struct to hold DumperHandler specific information
type FileHandler struct {
	BaseHandler
	TargetService services.TargetService
}

var _ Handler = (*FileHandler)(nil)

//NewDumperHandler function to create a new handler Dumper
func NewFileHandler(ts services.TargetService) *FileHandler {
	fh := new(FileHandler)
	fh.TargetService = ts
	return fh
}

//POST function to create a new handler
func (h *FileHandler) GET(c *gin.Context) {
	targetId := c.Param("targetId")

	target, err := h.TargetService.GetTargetById(targetId)

	if err != nil {

	}

	filename := target.Files[0].File.Path

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename )
	c.Header("Content-Type", "application/octet-stream")

	c.File(filename)
}
