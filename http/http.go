package http

import (
	"github.com/SantoDE/datahamster/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/SantoDE/datahamster/services"
	"github.com/SantoDE/datahamster/http/rpc"
	"github.com/SantoDE/datahamster/proto"
)

//Server struct to hold HTTP Server Information
type Server struct {
	Handler *Handlers
	addr string
	services *rpc.Services
}

//Handlers struct to hold different Handlers
type Handlers struct {
	PingHandler   *handlers.PingHandler
	DumperHandler *handlers.DumperHandler
	FileHandler   *handlers.FileHandler
}

//NewServer to create a new HTTP Server and wire handlers
func NewServer(addr string, dir string, services *services.Services) *Server {

	server := new(Server)
	server.Handler = new(Handlers)

	server.addr = addr

	pingHandler := handlers.NewPingHandler()
	dumperHander := handlers.NewDumperHandler(services.DumperService)
	fileHandler := handlers.NewFileHandler(services.TargetService)

	server.Handler.PingHandler = pingHandler
	server.Handler.DumperHandler = dumperHander
	server.Handler.FileHandler = fileHandler

	server.services = new(rpc.Services)
	server.services.DumperService = rpc.NewDumperService(services.DumperService, services.TargetService)
	server.services.FileHandleService = rpc.NewFileHandleService(services.TargetService, dir)

	return server
}

//Start HTTP Server
func (h *Server) Start() {
	r := gin.Default()
	r.GET("/ping", h.Handler.PingHandler.GET)
	r.POST("/dumper", h.Handler.DumperHandler.POST)
	r.GET("/files/download/:targetId/", h.Handler.FileHandler.GET)

	dumperHandler := dumper.NewDumperServiceServer(h.services.DumperService, nil)
	fileHandler := dumper.NewFileServiceServer(h.services.FileHandleService, nil)

	r.POST(dumper.DumperServicePathPrefix+"*action", gin.WrapH(dumperHandler))
	r.POST(dumper.FileServicePathPrefix+"*action", gin.WrapH(fileHandler))

	r.Run(h.addr)
}
