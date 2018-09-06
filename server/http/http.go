package http

import (
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/server/http/handlers"
	"github.com/SantoDE/datahamster/server/http/handlers/api"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/store/target"
	"github.com/gin-gonic/gin"
)

//Server struct to hold HTTP Server Information
type Server struct {
	Handler *Handlers
	addr    string
}

//Handlers struct to hold different Handlers
type Handlers struct {
	PingHandler *handlers.PingHandler
	FileHandler *handlers.FileHandler
	RpcHandler  *handlers.RpcHandler
	ApiHandler  *api.ApiHandler
}

//NewServer to create a new HTTP Server and wire handlers
func NewServer(addr string) *Server {

	server := new(Server)
	server.Handler = new(Handlers)

	server.addr = addr

	return server
}

func (h *Server) Init(storage storage.Storage, store *store.DataStore, t *target.Store) {

	pingHandler := handlers.NewPingHandler()
	fileHandler := handlers.NewFileHandler()
	rpcHandler := handlers.NewRpcHandler(store, t, store, storage)
	apiHandler := api.NewApiHandler(t, store)

	h.Handler.PingHandler = pingHandler
	h.Handler.FileHandler = fileHandler
	h.Handler.RpcHandler = &rpcHandler
	h.Handler.ApiHandler = apiHandler
}

//Start HTTP Server
func (h *Server) Start() {
	r := gin.Default()
	r.GET("/ping", h.Handler.PingHandler.GET)
	r.GET("/files/download/:targetId/", h.Handler.FileHandler.GET)
	r.GET("/targets/", h.Handler.ApiHandler.GetTargets)
	r.GET("/jobs/", h.Handler.ApiHandler.GetJobs)

	datakubeServer := datakube.NewDatakubeServer(h.Handler.RpcHandler, nil)

	r.POST(datakube.DatakubePathPrefix+"*action", gin.WrapH(datakubeServer))

	r.Run(h.addr)
}
