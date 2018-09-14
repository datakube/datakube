package rpc

import (
	"context"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/rpc"
	"github.com/SantoDE/datahamster/types"
	"time"
)

func (h *rpcHandler) SaveDumpFile(ctx context.Context, in *datakube.SaveDumpFileRequest) (*datakube.SaveDumpFileResponse, error) {

	log.Debugf("Received RPC Request to save file with filename %s", in.Filename)
	file := types.File{
		Name: in.Targetname,
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

	return &datakube.SaveDumpFileResponse{
		Success: true,
	}, nil
}