package handlers

import (
	"github.com/SantoDE/datahamster/services"
	"github.com/SantoDE/datahamster/types"
	"net/http"
	"encoding/json"
	"io/ioutil"
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
func (h *DumperHandler) POST(w http.ResponseWriter, r *http.Request) {
	var newDumper types.Dumper
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

	}

	// This will infer what binder to use depending on the content-type header.
	if err := json.Unmarshal(body, &newDumper); err == nil {
		Dumper, err := h.DumperService.Create(newDumper.Name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError);
		}

		jsonMessage, err := json.Marshal(Dumper)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonMessage)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError);
	}
}