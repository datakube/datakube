package rpc

import (
	dumper "github.com/SantoDE/datahamster/rpc/proto"
	"github.com/SantoDE/datahamster/services"
	"golang.org/x/net/context"
)

//DumperService struct to hold RPC DumperService definition
type DumperService struct {
	DumperService services.DumperService
}

//NewDumperService to create a new RPC Dumper Service
func NewDumperService(bas services.DumperService) *DumperService {
	as := new(DumperService)
	as.DumperService = bas

	return as
}

//ConnectDumper function which gets called when an Dumper connected
func (f *DumperService) RegisterDumper(ctx context.Context, in *dumper.RegisterRequest) (*dumper.RegisterResponse, error) {
	res, err := f.DumperService.Validate(in.Token)

	if err != nil {
		return &dumper.RegisterResponse{Success: false}, err
	}

	for _, target := range in.Targets {
		f.DumperService.RegisterTarget(in.Token, target.Name, target.Schedule)
	}

	return &dumper.RegisterResponse{Success: res}, nil
}
