package worker

import (
	"github.com/SantoDE/datahamster/rpc/connect"
	"context"
	"github.com/SantoDE/datahamster/log"
	"fmt"
)

func StartWorker() {

	conn := Connect()
	defer conn.Close()

	request := new(connect.ConnectRequest)
	request.Type = "sql"
	conClient := connect.NewAgentConnectClient(conn)
	resp, err := conClient.ConnectAgent(context.Background(), request)

	if err != nil {
		log.Debugf("Error Connecting s%s", err.Error())
	}

	fmt.Printf("response %s", resp.Success)
}