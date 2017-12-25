package bolt

import (
	"crypto/rand"
	"fmt"
	"github.com/SantoDE/datahamster/types"
)

// DialService represents a service for managing dials.
type AgentService struct {
	datastore *Datastore
}

func NewAgentService(d *Datastore) AgentService {
	s := AgentService{
		datastore: d,
	}

	return s
}

func (a *AgentService) Validate(token string) (bool, error) {
	fmt.Printf("In Validate baby!")
	var agent types.Agent
	err := a.datastore.db.One("Token", token, &agent)

	if err != nil {
		fmt.Print("Error fetching %s", err)
		return false, err
	}

	return true, err
}

func (a *AgentService) Create(name string) *types.Agent {
	token := randToken()

	agent := new(types.Agent)
	agent.Token = token
	agent.Name = name

	err := a.datastore.db.Save(agent)

	if err != nil {

	}

	return agent
}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
