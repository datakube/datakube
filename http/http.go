package http

import (
	"github.com/SantoDE/datahamster"
	"github.com/SantoDE/datahamster/http/handlers"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Handler *Handlers
}

type Handlers struct {
	PingHandler  *handlers.PingHandler
	AgentHandler *handlers.AgentHandler
}

func NewHttpServer(services *datahamster.Services) *HttpServer {

	server := new(HttpServer)
	server.Handler = new(Handlers)

	pingHandler := handlers.NewPingHandler(services.AgentService)
	agentHander := handlers.NewAgentHandler(services.AgentService)

	server.Handler.PingHandler = pingHandler
	server.Handler.AgentHandler = agentHander

	return server
}

func (h *HttpServer) Start() {
	r := gin.Default()
	r.GET("/ping", h.Handler.PingHandler.GET)
	r.POST("/agent", h.Handler.AgentHandler.POST)
	r.Run(":8080")
}
