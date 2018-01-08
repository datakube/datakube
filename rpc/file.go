package rpc

import (
	dumper "github.com/SantoDE/datahamster/rpc/proto"
	"github.com/SantoDE/datahamster/services"
	"golang.org/x/net/context"
	"github.com/SantoDE/datahamster/storage"
	"github.com/SantoDE/datahamster/storage/file"
	"github.com/SantoDE/datahamster/types"
)

//DumperService struct to hold RPC DumperService definition
type FileHandleService struct {
	DumperService services.DumperService
	FileStorage storage.Storage
}

//NewDumperService to create a new RPC Dumper Service
func NewFileHandleService(bas services.DumperService) *FileHandleService {
	fh := new(FileHandleService)
	fh.DumperService = bas
	fh.FileStorage = file.NewFileStorage("/var/tmp")

	return fh
}

//ConnectDumper function which gets called when an Dumper connected
func (f *FileHandleService) SaveDumpFile(ctx context.Context, in *dumper.SaveDumpFileRequest) (*dumper.SaveDumpFileResponse, error) {

	file := types.File{
		Name: in.Filename,
		Data: in.Data,
	}

	savedFile, err := f.FileStorage.SaveFile(file)

	if err != nil {

	}

	f.DumperService.SaveTargetFile(in.Token, in.Targetname, savedFile.Name, savedFile.Path)

	return &dumper.SaveDumpFileResponse{
		Success:true,
	}, nil
}
