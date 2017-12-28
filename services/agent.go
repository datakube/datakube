package services

import (
	"crypto/rand"
	"fmt"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/types"
)

var _ AgentService = (*Service)(nil)

type AgentService interface {
	Validate(string) (bool, error)
	Create(string) (*types.Agent, error)
}

type Service struct {
	datastore *store.Datastore
}

//NewAgentService creates a new Agent Service
func NewAgentService(d *store.Datastore) Service {
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

//Create creates a new agent with a random token
func (s *Service) Create(name string) (*types.Agent, error) {
	token := randToken()

	agent := new(types.Agent)
	agent.Token = token
	agent.Name = name

	newAgent, err := s.datastore.Save(*agent)

	if err != nil {
		fmt.Print("Error Creating agent %s", err)
		return new(types.Agent), err
	}

	return &newAgent, nil
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
