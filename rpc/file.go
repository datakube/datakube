package rpc

import (
	"github.com/SantoDE/datahamster/rpc/proto"
	"github.com/SantoDE/datahamster/services"
	"golang.org/x/net/context"
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
func NewFileHandleService(tas services.TargetService) *FileHandleService {
	fh := new(FileHandleService)
	fh.TargetService = tas
	//@TODO make storage location configurable
	fh.FileStorage = file.NewFileStorage("/tmp/dumps")

	return fh
}

//ConnectDumper function which gets called when an Dumper connected
func (f *FileHandleService) SaveDumpFile(ctx context.Context, in *dumper.SaveDumpFileRequest) (*dumper.SaveDumpFileResponse, error) {

	log.Debug("Received RPC Request to save file with filename %s for token %s", in.Filename, in.Token)
	file := types.File{
		Name: in.Targetname,
		Data: in.Data,
	}

	log.Debug("Saving file %s", in.Filename)
	savedFile, err := f.FileStorage.SaveFile(file)

	if err != nil {
		log.Debug("Error while saving file %s", savedFile.Name)
	}

	log.Debug("Persist Saved Dump Record in Database")

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
