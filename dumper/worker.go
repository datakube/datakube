package dumper

import (
	"context"
	"github.com/SantoDE/datahamster/adapter"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/types"
	"io/ioutil"
	"net/http"
	"os"
)

//StartWorker function to start the Worker
func StartWorker(c *configuration.DumperConfiguration) {

	client := datakube.NewDatakubeProtobufClient(c.Server, &http.Client{})

	// Attach the headers to a context
	ctx := context.Background()

	dumps := make(chan types.DumpResult)

	jobs, err := client.ListJobs(ctx, &datakube.ListJobsRequest{Status: types.STATUS_QUEUED})

	if err != nil {
		log.Error("Errror getting jobs ", err.Error())
		os.Exit(15)
	}

	for _, job := range jobs.Jobs {
		adapter, err := adapter.CreateNewAdapter(job.Target.Credentials.Host, job.Target.Credentials.Port, job.Target.Credentials.Database, job.Target.Credentials.User, job.Target.Credentials.Password, job.Target.Type)

		if err != nil {
			log.Errorf("Cant execute job for target %s with error => %s", job.Target, err.Error())
			continue;
		}

		go Run(job.Target.Type, adapter, dumps)
	}

	for {
		select {
		case dump := <-dumps:
			data, err := ioutil.ReadFile(dump.TemporaryFile)

			if err != nil {
				log.Debugf("Error reading temporary file to send %s", err.Error())
			}

			req := datakube.SaveDumpFileRequest{
				Targetname: dump.TargetName,
				Data:       data,
			}

			res, err := client.SaveDumpFile(ctx, &req)

			if err != nil {
				log.Debugf("Error sending file to server %s", err.Error())
			}

			if res.Success != true {
				log.Debugf("Transfered dump to Server not successful")
			}
			log.Debugf("Transfered dump to Server successfuly")
		}
	}
}
