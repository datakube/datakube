package web

import (
	"github.com/SantoDE/datahamster/server/configuration"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Manager struct which holds the application configuration
type Server struct {
	stopChan        chan bool
	configuration 	configuration.ServerConfiguration
}

// NewManager with the application config
func NewServer(configuration configuration.ServerConfiguration) *Server {
	server := new(Server)
	server.configuration = configuration

	return server
}

func(server *Server) Start() {
	router := server.buildDefaultHTTPRouter()

	go router.Run()
}

func (server *Server) buildDefaultHTTPRouter() gin.Engine {
	router := gin.Default()

	// This handler will match /user/john but will not match neither /user/ or /user
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello")
	})


	return *router
}

// Wait blocks until server is shutted down.
func (server *Server) Wait() {
	<-server.stopChan
}