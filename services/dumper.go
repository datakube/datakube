package services

import (
	"crypto/rand"
	"fmt"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/types"
	"github.com/SantoDE/datahamster/log"
)

//Service struct to hold Dumper Service Information
type DumperService struct {
	datastore store.DumperStore
}

//NewDumperService creates a new Dumper Service
func NewDumperService(d store.DumperStore) DumperService {
	s := DumperService{
		datastore: d,
	}

	return s
}

//Validate checks if the given token is valide
func (s *DumperService) Validate(token string) (bool, error) {
	_, err := s.datastore.One(token)

	if err != nil {
		log.Errorf("Error Validating the token %s => %s", token, err.Error())
		return false, err
	}

	return true, err
}

//Create creates a new Dumper with a random token
func (s *DumperService) Create(name string) (types.Dumper, error) {
	token := randToken()

	Dumper := new(types.Dumper)
	Dumper.Token = token
	Dumper.Name = name

	newDumper, err := s.datastore.Save(*Dumper)

	if err != nil {
		fmt.Print("Error Creating Dumper %s", err.Error())
		return *new(types.Dumper), err
	}

	return newDumper, nil
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
