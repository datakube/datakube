package handlers

import (
	"net/http"
	"encoding/json"
)

//PingHandler to hold Pinghandler information
type PingHandler struct {
	BaseHandler
}

type PingMessage struct {
	Message string `json:"message"`
}

var _ Handler = (*PingHandler)(nil)

//NewPingHandler to create a new Pinghandler
func NewPingHandler() *PingHandler {
	h := new(PingHandler)

	return h
}

//GET Ping
func (h *PingHandler) GET(w http.ResponseWriter, r *http.Request) {
	message := PingMessage{"pong"};

	jsonMessage, err := json.Marshal(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonMessage)
}
