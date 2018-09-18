package target

import (
	"errors"
	"github.com/datakube/datakube/types"
	"sync"
)

type Store struct {
	targets []types.Target
	sync.RWMutex
}

func (s *Store) Subscribe(targetChan chan types.ConfigTargets) {
	for {
		select {
		case targets := <-targetChan:
			s.RWMutex.Lock()
			s.targets = targets.Targets
			s.RWMutex.Unlock()
		}
	}
}

func (s *Store) ListTargets() []types.Target {
	return s.targets
}

func (s *Store) GetOneTargetByName(targetName string) (types.Target, error) {
	s.RLock()
	defer s.RUnlock()
	for _, target := range s.targets {
		if target.Name == targetName {
			return target, nil
		}
	}

	return *new(types.Target), errors.New("No target found")
}
