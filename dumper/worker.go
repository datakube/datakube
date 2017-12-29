package dumper

import (
	"context"
	"fmt"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/connect"
)

//StartWorker function to start the Worker
func StartWorker() {

	conn := Connect()
	defer conn.Close()

	request := new(connect.ConnectRequest)
	request.Token = "12345"
	conClient := connect.NewDumperConnectClient(conn)
	resp, err := conClient.ConnectDumper(context.Background(), request)

	if err != nil {
		log.Debugf("Error Connecting s%s", err.Error())
	}

	fmt.Printf("response %s", resp.Success)
}
