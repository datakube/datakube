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
	FileHandler   *handlers.FileHandler
}

//NewServer to create a new HTTP Server and wire handlers
func NewServer(services *services.Services) *Server {

	server := new(Server)
	server.Handler = new(Handlers)

	pingHandler := handlers.NewPingHandler()
	dumperHander := handlers.NewDumperHandler(services.DumperService)
	fileHandler := handlers.NewFileHandler(services.TargetService)

	server.Handler.PingHandler = pingHandler
	server.Handler.DumperHandler = dumperHander
	server.Handler.FileHandler = fileHandler

	return server
}

//Start HTTP Server
func (h *Server) Start() {
	r := gin.Default()
	r.GET("/ping", h.Handler.PingHandler.GET)
	r.POST("/dumper", h.Handler.DumperHandler.POST)
	r.GET("/files/download/:targetId/", h.Handler.FileHandler.GET)
	r.Run(":8080")
}
