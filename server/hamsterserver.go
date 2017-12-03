package server

import (
	"github.com/SantoDE/datahamster/server/rpc"
	"github.com/SantoDE/datahamster/server/web"
	"github.com/SantoDE/datahamster/server/configuration"
	"github.com/SantoDE/datahamster/server/rpc/fileUpload"
	"net"
	"github.com/SantoDE/datahamster/log"
)

type HamsterServer struct {
	Configuration 	configuration.ServerConfiguration
	webServer *web.Server
	RpcServices []rpc.RpcService
}

func (s *HamsterServer) Start() {



	webServer := web.NewServer(s.Configuration)
	webServer.Start()

	lis, err := net.Listen("tcp", ":8010")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Debugf("LIstener oppened on %s", lis.Addr().String())

	s.RpcServices = append(s.RpcServices, new(fileUpload.Service))
	for _, r := range s.RpcServices {
		r.Serve(lis)
	}
}