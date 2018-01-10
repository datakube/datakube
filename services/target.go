package services

import (
	"github.com/SantoDE/datahamster/store"
	"fmt"
	"github.com/SantoDE/datahamster/types"
	"time"
)

//Service struct to hold Dumper Service Information
type TargetService struct {
	datastore store.TargetStore
}

//NewDumperService creates a new Dumper Service
func NewTargetService(d store.TargetStore) TargetService {
	s := TargetService{
		datastore: d,
	}

	return s
}


//Register function registers targets with the given dumper
func (s *TargetService) RegisterTarget(name string, schedule string) (types.DumpTarget, error) {

	target := new(types.DumpTarget)
	target.Name = name
	target.Schedule = schedule

	savedTarget, err := s.datastore.SaveTarget(*target)

	if err != nil {
		fmt.Print("Error Registering Target %s", err)
		return *target, err
	}

	return savedTarget, nil
}

//SaveTargetFile adds a saved file to the target
func (s *TargetService) SaveTargetFile(targetname string, filename string, storageKey string) (types.DumpFile, error){

	target, err := s.datastore.OneByName(targetname)

	if err != nil {
		fmt.Print("Error Saving Target File - Target with name %s not found", targetname)
		return *new(types.DumpFile), err
	}

	var file types.DumpFile

	file = types.DumpFile{
		CreatedAt: time.Now(),
		File: types.File {
			Name: filename,
			Path: storageKey,
		},
	}

	target.Files = append(target.Files, file)

	s.datastore.SaveTarget(target)
	return file, nil
}


func (s *TargetService) GetTargetById(targetId string) (types.DumpTarget, error) {
	target, err := s.datastore.OneById(targetId)

	if err != nil {
		fmt.Print("Error fetching %s", err)
		return target, err
	}

	return target, err
}
