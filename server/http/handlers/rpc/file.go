package rpc

import (
	"context"
	"errors"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/types"
	"strconv"
	"time"
)

func (h *rpcHandler) SaveDumpFileForJob(ctx context.Context, in *datakube.SaveDumpFileRequest) (*datakube.SaveDumpFileResponse, error) {

	log.Debugf("Received RPC Request to save file with TargetName %s for job %d", in.Targetname, in.JobId)

	job, err := h.js.GetJobById(int(in.JobId))

	if err != nil {
		log.Debugf("No Job with Id %d found - denying save request", in.JobId)
		return &datakube.SaveDumpFileResponse{
			Success: false,
		}, err
	}

	if in.Targetname == "" {
		err := errors.New("No proper Target passed - not accepting the request to save the file")
		log.Error(err.Error())
		return &datakube.SaveDumpFileResponse{
			Success: false,
		}, err
	}

	now :=  strconv.FormatInt(time.Now().UTC().Unix(), 10)

	file := types.File{
		Name: in.Targetname + "_" + now,
		Data: in.Data,
	}

	log.Debugf("Saving file %s", in.Filename)
	savedFile, err := h.stg.SaveFile(file)

	if err != nil {
		log.Errorf("Error while saving file %s", savedFile.Name)
		return &datakube.SaveDumpFileResponse{
			Success: false,
		}, err
	}

	log.Debugf("Persist Saved Dump Record in Database")
	_, err = h.dfs.SaveDumpFile(types.DumpFile{
		CreatedAt: time.Now(),
		Target:    in.Targetname,
		File: types.File{
			Name: savedFile.Name,
			Path: savedFile.Path,
		},
	})

	if err != nil {
		return &datakube.SaveDumpFileResponse{
			Success: false,
		}, err
	}

	job.Status = types.STATUS_SUCCESS
	h.js.SaveJob(job)

	return &datakube.SaveDumpFileResponse{
		Success: true,
	}, nil
}
