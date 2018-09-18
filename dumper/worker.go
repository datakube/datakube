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
	"time"
)

//StartWorker function to start the Worker
func StartWorker(c *configuration.DumperConfiguration) {

	client := datakube.NewDatakubeProtobufClient(c.Server, &http.Client{})

	// Attach the headers to a context
	ctx := context.Background()

	ticker := time.NewTicker(time.Second * time.Duration(c.Interval))
	for range ticker.C {
		jobs, err := client.ListJobs(ctx, &datakube.ListJobsRequest{Status: types.STATUS_QUEUED})

		if err != nil {
			log.Error("Errror getting jobs ", err.Error())
			os.Exit(15)
		}

		for _, job := range jobs.Jobs {

			job.State = types.STATUS_IN_PROGRESS
			updateRequest := datakube.UpdateJobRequest{
				Job:job,
			}
			client.UpdateJob(ctx, &updateRequest)

			adapter, err := adapter.CreateNewAdapter(job.Target.Credentials.Host, job.Target.Credentials.Port, job.Target.Credentials.Database, job.Target.Credentials.User, job.Target.Credentials.Password, job.Target.Type)

			if err != nil {
				log.Errorf("Cant execute job for target %s with error => %s", job.Target, err.Error())
				continue;
			}

			res := Run(job.Target.Name, adapter)

			if res.Success == false {
				log.Debug("Something failed in job %i - returning to queue", job.Id)
				job.State = types.STATUS_QUEUED
				updateRequest := datakube.UpdateJobRequest{
					Job:job,
				}
				client.UpdateJob(ctx, &updateRequest)
			}

			data, err := ioutil.ReadFile(res.TemporaryFile)

			if err != nil {
				log.Debugf("Error reading temporary file to send %s", err.Error())
			}

			req := datakube.SaveDumpFileRequest{
				Targetname: res.TargetName,
				Data:       data,
				JobId: job.Id,
			}

			saveresult, err := client.SaveDumpFileForJob(ctx, &req)

			if err != nil {
				log.Debugf("Error sending file to server %s", err.Error())
			}

			if saveresult.Success != true {
				log.Debugf("Transfered dump to Server not successful")
			}

			log.Debugf("Transfered dump to Server successfuly - acknowling")
		}
	}
}
