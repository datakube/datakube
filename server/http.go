package server

import (
	"github.com/gin-gonic/gin"
	"github.com/SantoDE/datahamster/handlers"
)

type HttpServer struct {
	application *Application
	agentHandler *handlers.AgentHandler
	pingHandler *handlers.PingHandler
}

func NewHttpServer(app *Application) *HttpServer {

	server := new(HttpServer)
	server.application = app
	server.agentHandler = handlers.NewAgentHandler(&app.AgentService)

	return server
}

func (h *HttpServer) Start() {
	r := gin.Default()
	r.GET("/ping", h.pingHandler.GET)
	r.POST("/agent", h.agentHandler.POST)
	r.Run(":8080")
}
