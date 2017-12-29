package handlers

import (
	"github.com/SantoDE/datahamster/services"
	"github.com/SantoDE/datahamster/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

//DumperHandler struct to hold DumperHandler specific information
type DumperHandler struct {
	BaseHandler
	DumperService services.DumperService
}

var _ Handler = (*DumperHandler)(nil)

//NewDumperHandler function to create a new handler Dumper
func NewDumperHandler(as services.DumperService) *DumperHandler {
	ah := new(DumperHandler)
	ah.DumperService = as
	return ah
}

//POST function to create a new handler
func (h *DumperHandler) POST(c *gin.Context) {
	var newDumper types.Dumper
	// This will infer what binder to use depending on the content-type header.
	if err := c.BindJSON(&newDumper); err == nil {
		Dumper, err := h.DumperService.Create(newDumper.Name)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"success": "true", "Dumper": Dumper})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
