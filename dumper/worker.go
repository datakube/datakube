package dumper

import (
	"context"
	"fmt"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/dumper/jobs"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc/proto"
	"github.com/SantoDE/datahamster/types"
	"io/ioutil"
	"os"
)

//StartWorker function to start the Worker
func StartWorker(c *configuration.DumperConfiguration) {

	conn := Connect()

	conClient := dumper.NewDumperServiceClient(conn)

	ctx := context.Background()
	request := new(dumper.RegisterRequest)
	request.Token = "12345"
	resp, err := conClient.RegisterDumper(ctx, request)

	if err != nil {
		log.Debugf("Error Connecting s%s", err.Error())
		os.Exit(500)
	}

	scheduler := NewScheduler()

	dumps := make(chan types.DumpResult)

	for _, target := range c.Targets {


		j := jobs.NewDumpJob(&target, dumps)
		scheduler.Schedule(&target.Schedule, j)
	}

	fileClient := dumper.NewFileServiceClient(conn)

	select {
		case dump := <- dumps:
			data, err := ioutil.ReadFile(dump.TemporaryFile)

			if err != nil {

			}

			req := dumper.SaveDumpFileRequest{
				Targetname: dump.TargetName,
				Data: data,
			}

			res, err := fileClient.SaveDumpFile(ctx, &req)

			if err != nil {

			}

			fmt.Printf("Save grpc %s", res.Success)
		default:
			fmt.Println("no message sent")
	}

	fmt.Printf("response %s", resp.Success)
	defer conn.Close()
}