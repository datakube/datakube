package dumper

import (
	"context"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/dumper/jobs"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/proto"
	"github.com/SantoDE/datahamster/types"
	"io/ioutil"
	"os"
	"net/http"
	"github.com/twitchtv/twirp"
)

//StartWorker function to start the Worker
func StartWorker(c *configuration.DumperConfiguration) {

	conClient := dumper.NewDumperServiceProtobufClient("http://127.0.0.1:8080", &http.Client{})
	fileClient := dumper.NewFileServiceProtobufClient("http://127.0.0.1:8080", &http.Client{})

	// Given some headers ...
	header := make(http.Header)
	header.Set("Datakube-Dumper-Token", "uDRlDxQYbFVXarBvmTncBoWKcZKqrZTY")

	// Attach the headers to a context
	ctx := context.Background()
	ctx, err := twirp.WithHTTPRequestHeaders(ctx, header)

	request := new(dumper.RegisterRequest)
	request.Auth = new(dumper.Authorization)
	request.Auth.Token = c.Token

	var targets []*dumper.Target

	for _, target := range c.Targets {
		requestTarget := dumper.Target{
			Name: target.Name,
		}

		targets = append(targets, &requestTarget)
	}

	request.Targets = targets
	resp, err := conClient.RegisterDumper(ctx, request)

	if err != nil {
		log.Debugf("Error Connecting %s", err.Error())
		os.Exit(15)
	}

	if resp.Success != true {
		log.Debugf("Register was not correct - wrong token maybe?")
		os.Exit(15)
	}

	scheduler := NewScheduler()

	dumps := make(chan types.DumpResult)

	for _, target := range c.Targets {
		j := jobs.NewDumpJob(&target, dumps)

		if target.StartImmediately {
			go j.Run()
		}

		scheduler.Schedule(&target.Schedule, j)
	}

	scheduler.Cron.Start()

	for {
		select {
		case dump := <- dumps:
			data, err := ioutil.ReadFile(dump.TemporaryFile)

			if err != nil {
				log.Debugf("Error reading temporary file to send %s", err.Error())
			}

			req := dumper.SaveDumpFileRequest{
				Targetname: dump.TargetName,
				Data: data,
			}

			res, err := fileClient.SaveDumpFile(ctx, &req)

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