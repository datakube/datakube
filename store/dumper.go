package store

import (
	"fmt"
	"github.com/SantoDE/datahamster/types"
)

var _ DumperStore = (*Store)(nil)

//DumperStore Interface to expose Dumper Stores API
type DumperStore interface {
	One(string) (types.Dumper, error)
	Save(types.Dumper) (types.Dumper, error)
}

//One function to retrieve one Dumper by the given token
func (s *Store) One(token string) (types.Dumper, error) {
	var Dumper types.Dumper
	err := s.db.One("Token", token, &Dumper)

	if err != nil {
		fmt.Print("Error fetching %s", err)
		return *new(types.Dumper), err
	}

	return Dumper, nil
}

//Save function to save the given Dumper
func (s *Store) Save(Dumper types.Dumper) (types.Dumper, error) {
	err := s.db.Save(&Dumper)

	if err != nil {
		return *new(types.Dumper), err
	}

	return Dumper, nil
}
