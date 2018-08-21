package rpc

import (
	"github.com/SantoDE/datahamster/proto"
	"github.com/SantoDE/datahamster/services"
	"context"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/storage/file"
	"github.com/SantoDE/datahamster/types"
	"github.com/SantoDE/datahamster/log"
)

//DumperService struct to hold RPC DumperService definition
type FileHandleService struct {
	TargetService services.TargetService
	FileStorage storage.Storage
}

//NewDumperService to create a new RPC Dumper Service
func NewFileHandleService(tas services.TargetService, dir string) *FileHandleService {
	fh := new(FileHandleService)
	fh.TargetService = tas
	fh.FileStorage = file.NewFileStorage(dir)

	return fh
}

//ConnectDumper function which gets called when an Dumper connected
func (f *FileHandleService) SaveDumpFile(ctx context.Context, in *dumper.SaveDumpFileRequest) (*dumper.SaveDumpFileResponse, error) {

	log.Debugf("Received RPC Request to save file with filename %s for token %s", in.Filename, in.Auth.Token)
	savefile := types.File{
		Name: in.Targetname,
		Data: in.Data,
	}

	log.Debugf("Saving file %s", in.Filename)
	savedFile, err := f.FileStorage.SaveFile(savefile)

	if err != nil {
		log.Errorf("Error while saving file %s", savedFile.Name)
	}

	log.Debugf("Persist Saved Dump Record in Database")

	_, err = f.TargetService.SaveTargetFile(in.Targetname, savedFile.Name, savedFile.Path)

	if err != nil {
		return &dumper.SaveDumpFileResponse{
			Success:false,
		}, err
	}

	return &dumper.SaveDumpFileResponse{
		Success:true,
	}, nil

}
