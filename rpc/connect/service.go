package connect

import (
	"golang.org/x/net/context"
	"github.com/SantoDE/datahamster/bolt"
)

type AgentService struct {
	boltAgentService *bolt.AgentService
}

func NewAgentService(bas *bolt.AgentService) *AgentService {
	as := new(AgentService)
	as.boltAgentService = bas

	return as
}

// SayHello implements helloworld.GreeterServer
func (f *AgentService) ConnectAgent(ctx context.Context, in *ConnectRequest) (*ConnectResponse, error) {
	res, err := f.boltAgentService.Validate(in.Token)

	if err != nil {
		return &ConnectResponse{Success: false}, err
	}

	return &ConnectResponse{Success: res}, nil
}
