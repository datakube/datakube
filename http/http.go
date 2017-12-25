package http

import (
	"github.com/SantoDE/datahamster"
	"github.com/SantoDE/datahamster/http/handlers"
	"github.com/gin-gonic/gin"
)

//Server struct to hold HTTP Server Information
type Server struct {
	Handler *Handlers
}

//Handlers struct to hold different Handlers
type Handlers struct {
	PingHandler  *handlers.PingHandler
	AgentHandler *handlers.AgentHandler
}

//NewServer to create a new HTTP Server and wire handlers
func NewServer(services *datahamster.Services) *Server {

	server := new(Server)
	server.Handler = new(Handlers)

	pingHandler := handlers.NewPingHandler(services.AgentService)
	agentHander := handlers.NewAgentHandler(services.AgentService)

	server.Handler.PingHandler = pingHandler
	server.Handler.AgentHandler = agentHander

	return server
}

//Start HTTP Server
func (h *Server) Start() {
	r := gin.Default()
	r.GET("/ping", h.Handler.PingHandler.GET)
	r.POST("/agent", h.Handler.AgentHandler.POST)
	r.Run(":8080")
}
