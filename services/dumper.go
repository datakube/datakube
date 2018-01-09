package services

import (
	"crypto/rand"
	"fmt"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/types"
	"time"
)

var _ DumperService = (*Service)(nil)

//DumperService Interface to expose the public API
type DumperService interface {
	Validate(string) (bool, error)
	Create(string) (types.Dumper, error)
	RegisterTarget(string, string, string) (types.DumpTarget, error)
	SaveTargetFile(string, string, string, string) (types.DumpFile, error)
}

//Service struct to hold Dumper Service Information
type Service struct {
	datastore store.DumperStore
}

//NewDumperService creates a new Dumper Service
func NewDumperService(d store.DumperStore) Service {
	s := Service{
		datastore: d,
	}

	return s
}

//Validate checks if the given token is valide
func (s *Service) Validate(token string) (bool, error) {
	_, err := s.datastore.One(token)

	if err != nil {
		fmt.Print("Error fetching %s", err)
		return false, err
	}

	return true, err
}

//Register function registers targets with the given dumper
func (s *Service) RegisterTarget(token string, name string, schedule string) (types.DumpTarget, error) {
	existing, err := s.datastore.One(token)

	if err != nil {
		fmt.Print("Error Registering Target - Dumper with token %s not found", token)
		return *new(types.DumpTarget), err
	}

	target := new(types.DumpTarget)
	target.Name = name
	target.Schedule = schedule

	existing.Targets = append(existing.Targets, *target)

	s.datastore.Save(existing)

	if err != nil {
		fmt.Print("Error Registering Target %s", err)
		return *target, err
	}

	return *target, nil
}

//SaveTargetFile adds a saved file to the target
func (s *Service) SaveTargetFile(token string, targetname string, filename string, storageKey string) (types.DumpFile, error){

	dumper, err := s.datastore.One(token)

	if err != nil {
		fmt.Print("Error Saving Target File - Dumper with token %s not found", token)
		return *new(types.DumpFile), err
	}

	var file types.DumpFile

	for _, target := range dumper.Targets {
		if target.Name == targetname {
			file = types.DumpFile{
				CreatedAt: time.Now(),
				File: types.File {
					Name: filename,
					Path: storageKey,
				},
			}

			target.Files = append(target.Files, file)
		}
	}

	s.datastore.Save(dumper)
	return file, nil
}

//Create creates a new Dumper with a random token
func (s *Service) Create(name string) (types.Dumper, error) {
	token := randToken()

	Dumper := new(types.Dumper)
	Dumper.Token = token
	Dumper.Name = name

	newDumper, err := s.datastore.Save(*Dumper)

	if err != nil {
		fmt.Print("Error Creating Dumper %s", err)
		return *new(types.Dumper), err
	}

	return newDumper, nil
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
