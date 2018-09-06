package handlers

import (
	"github.com/gin-gonic/gin"
)

//DumperHandler struct to hold DumperHandler specific information
type FileHandler struct {
	BaseHandler
}

var _ Handler = (*FileHandler)(nil)

//NewDumperHandler function to create a new handler Dumper
func NewFileHandler() *FileHandler {
	fh := new(FileHandler)
	return fh
}

//POST function to create a new handler
func (h *FileHandler) GET(c *gin.Context) {
	/*
	targetId, err := strconv.Atoi(c.Param("targetId"))

	provider, err := h.TargetService.GetTargetById(targetId)

	@TODO MAKE BETTER ERROR HANDLING
	if err != nil {

	}

	filename := provider.Files[0].File.Path

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename )
	c.Header("Content-Type", "application/octet-stream")
	c.File(filename)
	*/
}
