package rpc

import (
	pb "github.com/SantoDE/datahamster/rpc/internal/connect"
	"github.com/SantoDE/datahamster/services"
	"golang.org/x/net/context"
)

//AgentService struct to hold RPC AgentService definition
type AgentService struct {
	boltAgentService services.AgentService
}

//NewAgentService to create a new RPC Agent Service
func NewAgentService(bas services.AgentService) *AgentService {
	as := new(AgentService)
	as.boltAgentService = bas

	return as
}

//ConnectAgent function which gets called when an agent connected
func (f *AgentService) ConnectAgent(ctx context.Context, in *pb.ConnectRequest) (*pb.ConnectResponse, error) {
	res, err := f.boltAgentService.Validate(in.Token)

	if err != nil {
		return &pb.ConnectResponse{Success: false}, err
	}

	return &pb.ConnectResponse{Success: res}, nil
}
