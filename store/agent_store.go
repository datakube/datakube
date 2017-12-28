package store

import (
	"fmt"
	"github.com/SantoDE/datahamster/types"
)

var _ AgentStore = (*Datastore)(nil)

type AgentStore interface {
	One(string) (types.Agent, error)
	Save(types.Agent) (types.Agent, error)
}

func (ds *Datastore) One(token string) (types.Agent, error){
	var agent types.Agent
	err := ds.db.One("Token", token, &agent)

	if err != nil {
		fmt.Print("Error fetching %s", err)
		return *new(types.Agent), err
	}

	return agent, nil
}

func (ds *Datastore) Save(agent types.Agent) (types.Agent, error){
	err := ds.db.Save(&agent)

	if err != nil {
		return *new(types.Agent), err
	}

	return agent, nil
}