package server

import (
	"github.com/SantoDE/datahamster/server/handlers"
	"github.com/SantoDE/datahamster/services"
	"github.com/SantoDE/datahamster/server/rpc"
	"github.com/SantoDE/datahamster/proto"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/urfave/negroni"
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

	mux := mux.NewRouter()

	mux.HandleFunc("/ping", h.Handler.PingHandler.GET).Methods("GET")
	mux.HandleFunc("/files/download/{targetId}", h.Handler.FileHandler.GET).Methods("GET")
	mux.HandleFunc("/dumper", h.Handler.DumperHandler.POST).Methods("POST")

	dumperHandler := dumper.NewDumperServiceServer(h.services.DumperService, nil)
	fileHandler := dumper.NewFileServiceServer(h.services.FileHandleService, nil)

	mux.Handle(dumper.FileServicePathPrefix, fileHandler)
	mux.PathPrefix(dumper.DumperServicePathPrefix).Handler(dumperHandler)

	n := negroni.New()

	n.Use(negroni.NewLogger());
	n.UseHandler(mux)

	srv := &http.Server{
		Addr: h.addr,
		Handler: n,
	}

	srv.ListenAndServe()
}
