package store

import (
"github.com/SantoDE/datahamster/types"
)

var _ TargetStore = (*Store)(nil)

//DumperStore Interface to expose Dumper Stores API
type TargetStore interface {
	SaveTarget(target types.DumpTarget) (types.DumpTarget, error)
	OneById(string) (types.DumpTarget, error)
	OneByName(string) (types.DumpTarget, error)
}

//Save function to save the given Dumper
func (s *Store) SaveTarget(target types.DumpTarget) (types.DumpTarget, error) {
	err := s.db.Save(&target)

	if err != nil {
		return *new(types.DumpTarget), err
	}

	return target, nil
}

//TargetById fetches a target by id
func (s *Store) OneById(id string) (types.DumpTarget, error) {
	var target types.DumpTarget

	err :=  s.db.One("Id", id, &target)

	if err != nil {
		return *new(types.DumpTarget), err
	}

	return target, nil
}

//TargetById fetches a target by id
func (s *Store) OneByName(name string) (types.DumpTarget, error) {
	var target types.DumpTarget

	err :=  s.db.One("Name", name, &target)

	if err != nil {
		return *new(types.DumpTarget), err
	}

	return target, nil
}