package server

import (
	"net"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/connect"
	"google.golang.org/grpc"
)

type RpcServer struct {
	application *Application
}

func NewRpcServer(app *Application) *RpcServer {

	server := new(RpcServer)
	server.application = app

	return server
}

func (r *RpcServer) Start() {
	lis, err := net.Listen("tcp", ":8010")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	log.Debugf("Registering FileUpload RPC")
	as := connect.NewAgentService(&r.application.AgentService)
	connect.RegisterAgentConnectServer(server, as)
	log.Debugf("Start Serve FileUpload RPC")
	err = server.Serve(lis)
	if err != nil {
		log.Debugf("Error Starting GRPC %s", err.Error())
	}
}