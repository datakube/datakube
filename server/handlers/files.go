package handlers

import (
	"github.com/SantoDE/datahamster/services"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
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
func (h *FileHandler) GET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	targetId, err := strconv.Atoi(vars["targetId"])

	target, err := h.TargetService.GetTargetById(targetId)

	//@TODO MAKE BETTER ERROR HANDLING
	if err != nil {

	}

	filename := target.Files[0].File.Path

	w.Header().Set("Content-Description", "File Transfer")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename )
	w.Header().Set("Content-Type", "application/octet-stream")

	http.ServeFile(w, r, filename)
}
