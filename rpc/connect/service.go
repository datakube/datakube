package connect

import (
	"golang.org/x/net/context"
	"github.com/SantoDE/datahamster/log"
)

type AgentService struct {

}

// SayHello implements helloworld.GreeterServer
func (f *AgentService) ConnectAgent(ctx context.Context, in *ConnectRequest) (*ConnectResponse, error) {
	log.Debugf("In Connect Babe!")
	return &ConnectResponse{Success: true}, nil
}
