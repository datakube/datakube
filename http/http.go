package http

import (
	"github.com/SantoDE/datahamster/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/SantoDE/datahamster/services"
)

//Server struct to hold HTTP Server Information
type Server struct {
	Handler *Handlers
}

//Handlers struct to hold different Handlers
type Handlers struct {
	PingHandler   *handlers.PingHandler
	DumperHandler *handlers.DumperHandler
}

//NewServer to create a new HTTP Server and wire handlers
func NewServer(services *services.Services) *Server {

	server := new(Server)
	server.Handler = new(Handlers)

	pingHandler := handlers.NewPingHandler()
	DumperHander := handlers.NewDumperHandler(services.DumperService)

	server.Handler.PingHandler = pingHandler
	server.Handler.DumperHandler = DumperHander

	return server
}

//Start HTTP Server
func (h *Server) Start() {
	r := gin.Default()
	r.GET("/ping", h.Handler.PingHandler.GET)
	r.POST("/dumper", h.Handler.DumperHandler.POST)
	r.Run(":8080")
}
