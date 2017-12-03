package rpc

import (
	"net"
)

type RpcService interface {
	Serve(listener net.Listener)
}

type BaseRpcService struct {

}