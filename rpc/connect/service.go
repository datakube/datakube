package connect

import (
	"golang.org/x/net/context"
	"github.com/SantoDE/datahamster/services"
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
func (f *AgentService) ConnectAgent(ctx context.Context, in *ConnectRequest) (*ConnectResponse, error) {
	res, err := f.boltAgentService.Validate(in.Token)

	if err != nil {
		return &ConnectResponse{Success: false}, err
	}

	return &ConnectResponse{Success: res}, nil
}
