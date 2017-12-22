package server

import (
	"net"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/connect"
	"google.golang.org/grpc"
)

func StartRpc() {
	lis, err := net.Listen("tcp", ":8010")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	log.Debugf("Registering FileUpload RPC")
	connect.RegisterAgentConnectServer(server, &connect.AgentService{})
	log.Debugf("Start Serve FileUpload RPC")
	err = server.Serve(lis)
	if err != nil {
		log.Debugf("Error Starting GRPC %s", err.Error())
	}
}