package dumper

import (
	"context"
	"fmt"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/dumper/jobs"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/proto"
)

//StartWorker function to start the Worker
func StartWorker(c *configuration.DumperConfiguration) {

	conn := Connect()
	defer conn.Close()

	conClient := dumper.NewDumperServiceClient(conn)

	request := new(dumper.RegisterRequest)
	request.Token = "12345"
	resp, err := conClient.RegisterDumper(context.Background(), request)

	if err != nil {
		log.Debugf("Error Connecting s%s", err.Error())
	}

	scheduler := NewScheduler()

	for _, target := range c.Targets {
		j := new(jobs.DumpJob)
		scheduler.Schedule(&target.Schedule, j)
	}

	fmt.Printf("response %s", resp.Success)
}