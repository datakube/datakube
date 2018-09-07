package http

import (
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/server/http/handlers"
	"github.com/SantoDE/datahamster/server/http/handlers/api"
	"github.com/SantoDE/datahamster/server/http/handlers/rpc"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/store/target"
	"github.com/gin-gonic/gin"
)

//Server struct to hold HTTP Server Information
type Server struct {
	addr    string
	storage storage.Storage
	store   *store.DataStore
	t       *target.Store
}

//NewServer to create a new HTTP Server and wire handlers
func NewServer(addr string) *Server {

	server := new(Server)
	server.addr = addr

	return server
}

func (h *Server) Init(storage storage.Storage, store *store.DataStore, t *target.Store) {
	h.t = t
	h.store = store
	h.storage = storage
}

//Start HTTP Server
func (h *Server) Start() {
	r := gin.Default()
	r.GET("/ping", handlers.GetPing)
	r.GET("/files/download/:targetId/", handlers.GetFile)
	r.GET("/targets/", api.GetTargets(h.t))
	r.GET("/jobs/", api.GetJobs(h.store))

	datakubeServer := datakube.NewDatakubeServer(rpc.New(h.store, h.t, h.store, h.storage), nil)

	r.POST(datakube.DatakubePathPrefix+"*action", gin.WrapH(datakubeServer))

	r.Run(h.addr)
}
