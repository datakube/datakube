package rpc

import (
	"github.com/SantoDE/datahamster/log"
	pb "github.com/SantoDE/datahamster/rpc/proto"
	"google.golang.org/grpc"
	"net"
	"github.com/SantoDE/datahamster/services"
)

//Server struct to hold RPC Server Information
type Server struct {
	services *Services
}

//Services struct to hold RPC Services Information
type Services struct {
	DumperService *DumperService
}

//NewServer function to create a new RPC Server
func NewServer(services *services.Services) *Server {
	server := new(Server)
	server.services = new(Services)
	server.services.DumperService = NewDumperService(services.DumperService)
	return server
}

//Start function to create a new RPC Server
func (r *Server) Start() {
	lis, err := net.Listen("tcp", ":8010")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	log.Debugf("Registering FileUpload RPC")
	pb.RegisterDumperServiceServer(server, r.services.DumperService)
	log.Debugf("Start Serve FileUpload RPC")
	err = server.Serve(lis)
	if err != nil {
		log.Debugf("Error Starting GRPC %s", err.Error())
	}
}
