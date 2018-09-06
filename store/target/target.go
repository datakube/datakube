package target

import (
	"github.com/SantoDE/datahamster/types"
)

type Store struct {
	targets []types.Target
	identifier string
}

func (s *Store) Subscribe(targetChan chan types.ConfigTargets) {
	for {
		select {
		case targets := <-targetChan:
			s.targets = targets.Targets
		}
	}
}

func (s *Store) ListTargets() []types.Target {
	return s.targets
}

func (s *Store) GetOneTargetByName(targetName string) types.Target {

	for _, target := range s.targets {
		if target.Name == targetName {
			return target
		}
	}

	return *new(types.Target)
}
