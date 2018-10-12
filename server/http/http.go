package http

import (
	"fmt"
	"github.com/datakube/datakube/rpc"
	"github.com/datakube/datakube/server/http/handlers"
	"github.com/datakube/datakube/server/http/handlers/api"
	"github.com/datakube/datakube/server/http/handlers/rpc"
	_ "github.com/datakube/datakube/statik"
	"github.com/datakube/datakube/storage"
	"github.com/datakube/datakube/store"
	"github.com/datakube/datakube/store/target"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
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

	statikFS, err := fs.New()

	if err != nil {
		fmt.Printf(err.Error())
	}

	r := gin.Default()
	r.GET("/ping", handlers.GetPing)
	fileRoutes := r.Group("/files/download/")
	fileRoutes.GET("/:name/", handlers.GetFile(h.store, h.storage))
	fileRoutes.GET("/:name/latest", handlers.GetLatestFile(h.store, h.storage))

	apiRouters := r.Group("/api")
	apiRouters.GET("/targets/", api.GetTargets(h.t))
	apiRouters.GET("/jobs/", api.GetJobs(h.store))
	apiRouters.GET("/dumps/", api.GetFiles(h.store))

	r.GET("/", func(c *gin.Context) {
		c.Request.URL.Path = "/dashboard/"
		r.HandleContext(c)
	})


	r.StaticFS("/dashboard/", statikFS)

	datakubeServer := datakube.NewDatakubeServer(rpc.New(h.store, h.t, h.store, h.storage), nil)

	r.POST(datakube.DatakubePathPrefix+"*action", gin.WrapH(datakubeServer))

	r.Run(h.addr)
}
