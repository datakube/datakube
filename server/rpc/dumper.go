package rpc

import (
	"github.com/SantoDE/datahamster/proto"
	"github.com/SantoDE/datahamster/services"
	"context"
	"github.com/SantoDE/datahamster/log"
)

//DumperService struct to hold RPC DumperService definition
type DumperService struct {
	DumperService services.DumperService
	TargetService services.TargetService
}

//NewDumperService to create a new RPC Dumper Service
func NewDumperService(bas services.DumperService, tas services.TargetService) *DumperService {
	as := new(DumperService)
	as.DumperService = bas
	as.TargetService = tas

	return as
}

//ConnectDumper function which gets called when an Dumper connected
func (f *DumperService) RegisterDumper(ctx context.Context, in *dumper.RegisterRequest) (*dumper.RegisterResponse, error) {
	res, err := f.DumperService.Validate(in.Auth.Token)

	if err != nil {
		return &dumper.RegisterResponse{Success: false}, err
	}

	for _, target := range in.Targets {
		_, err := f.TargetService.RegisterTarget(target.Name, target.Schedule)

		if err != nil {
			log.Debug("Error registering target ", err.Error())
		}
	}

	return &dumper.RegisterResponse{Success: res}, nil
}
