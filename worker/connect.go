package worker

import (
	"google.golang.org/grpc"
	"github.com/SantoDE/datahamster/log"
)

func Connect() *grpc.ClientConn {
	conn, err := grpc.Dial("127.0.0.1:8010", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Debugf("Error Dialing RPC %s", err.Error())
	}

	log.Debugf("Dailed successfull")

	return conn
}