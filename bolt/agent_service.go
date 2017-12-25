package bolt

import (
	"crypto/rand"
	"fmt"
	"github.com/SantoDE/datahamster/types"
)

// AgentService represents a service for saving agent details
type AgentService struct {
	datastore *Datastore
}

//NewAgentService creates a new Agent Service
func NewAgentService(d *Datastore) AgentService {
	s := AgentService{
		datastore: d,
	}

	return s
}

//Validate checks if the given token is valide
func (a *AgentService) Validate(token string) (bool, error) {
	var agent types.Agent
	err := a.datastore.db.One("Token", token, &agent)

	if err != nil {
		fmt.Print("Error fetching %s", err)
		return false, err
	}

	return true, err
}


//Create creates a new agent with a random token
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
