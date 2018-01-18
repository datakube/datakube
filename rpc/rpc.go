package rpc

import (
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/proto"
	"google.golang.org/grpc"
	"net"
	"github.com/SantoDE/datahamster/services"
	"fmt"
)

//Server struct to hold RPC Server Information
type Server struct {
	services *Services
}

//Services struct to hold RPC Services Information
type Services struct {
	DumperService *DumperService
	FileHandleService *FileHandleService
}

//NewServer function to create a new RPC Server
func NewServer(services *services.Services) *Server {
	server := new(Server)
	server.services = new(Services)
	server.services.DumperService = NewDumperService(services.DumperService, services.TargetService)
	server.services.FileHandleService = NewFileHandleService(services.TargetService)
	return server
}

//Start function to create a new RPC Server
func (r *Server) Start() {
	lis, err := net.Listen("tcp", ":8010")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	grpc.MaxCallRecvMsgSize(17179869184)
	log.Debugf("Registering Dumper RPC")

	dumper.RegisterDumperServiceServer(server, r.services.DumperService)
	log.Debugf("Registering FileService RPC")
	dumper.RegisterFileServiceServer(server, r.services.FileHandleService)
	log.Debugf("Start RPC socket")

	tmp := server.GetServiceInfo()

	for _, info := range tmp {
		fmt.Printf("Data %s", info.Metadata)
	}

	err = server.Serve(lis)
	if err != nil {
		log.Debugf("Error Starting GRPC %s", err.Error())
	}
}
