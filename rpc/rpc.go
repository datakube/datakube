package rpc

import (
	"net"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/connect"
	"google.golang.org/grpc"
	"github.com/SantoDE/datahamster"
)

type RpcServer struct {
	services *RpcServices
}

type RpcServices struct {
	AgentService *connect.AgentService
}

func NewRpcServer(services *datahamster.Services) *RpcServer {
	server := new(RpcServer)
	server.services = new(RpcServices)
	server.services.AgentService = connect.NewAgentService(services.AgentService)
	return server
}

func (r *RpcServer) Start() {
	lis, err := net.Listen("tcp", ":8010")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	log.Debugf("Registering FileUpload RPC")
	connect.RegisterAgentConnectServer(server, r.services.AgentService)
	log.Debugf("Start Serve FileUpload RPC")
	err = server.Serve(lis)
	if err != nil {
		log.Debugf("Error Starting GRPC %s", err.Error())
	}
}