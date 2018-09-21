package target

import (
	"errors"
	"github.com/datakube/datakube/types"
	"sync"
)

type Store struct {
	targets map[string][]types.Target
	sync.RWMutex
}

func (s *Store) Subscribe(targetChan chan types.ConfigTargets) {

	s.RWMutex.Lock()
	s.targets = make(map[string][]types.Target)
	s.RWMutex.Unlock()

	for {
		select {
		case targets := <-targetChan:
			s.RWMutex.Lock()
			s.targets[targets.Provider] = targets.Targets
			s.RWMutex.Unlock()
		}
	}
}

func (s *Store) ListTargets() []types.Target {
	s.RLock()
	defer s.RUnlock()

	var targetList []types.Target

	for _, targets := range s.targets {
		for _, target := range targets {
			targetList = append(targetList, target)
		}
	}
	return targetList
}

func (s *Store) GetOneTargetByName(targetName string) (types.Target, error) {
	s.RLock()
	defer s.RUnlock()
	for _, targets := range s.targets {
		for _, target := range targets {
			if target.Name == targetName {
				return target, nil
			}
		}
	}
	return *new(types.Target), errors.New("No target found")
}
