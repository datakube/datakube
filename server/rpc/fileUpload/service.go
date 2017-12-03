package fileUpload

import (
	"google.golang.org/grpc"
	pb "github.com/SantoDE/datahamster/services/fileUpload"
	"github.com/SantoDE/datahamster/server/rpc"
	"golang.org/x/net/context"
	"net"
	"github.com/SantoDE/datahamster/log"
)

type Service struct {
	rpc.BaseRpcService
}

// SayHello implements helloworld.GreeterServer
func (f *Service) UploadFile(ctx context.Context, in *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	log.Debugf("In UploadFile Babe!")
	return &pb.FileUploadResponse{Keyword: "hi"}, nil
}

func (s *Service) startServer(lis net.Listener) {
	log.Debugf("Starting FileUpload RPC")
	server := grpc.NewServer()
	log.Debugf("Registering FileUpload RPC")
	pb.RegisterFileUploadServer(server, &Service{})
	log.Debugf("Start Serve FileUpload RPC")
	err := server.Serve(lis)
	if err != nil {
		log.Debugf("Error Starting GRPC %s", err.Error())
	}
}

func (s *Service) Serve(lis net.Listener) {
	log.Debugf("Serving from FileUpload RPC")
	s.startServer(lis)
}