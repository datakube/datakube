package services_test

import (
	"github.com/SantoDE/datahamster/services"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ store.DumperStore = (*TestDataStore)(nil)

type TestDataStore struct{}

func (t *TestDataStore) One(token string) (types.Dumper, error) {
	var targets []types.DumpTarget

	targets = append(targets, types.DumpTarget{Name: "existing Target", Schedule: "weekly"})

	return types.Dumper{
		Token:   token,
		Name:    "Testdumper",
		Targets: targets,
	}, nil
}

func (t *TestDataStore) Save(dumper types.Dumper) (types.Dumper, error) {
	return dumper, nil
}

func TestRegisterTargetOK(t *testing.T) {
	testStore := new(TestDataStore)
	ds := services.NewDumperService(testStore)

	res, err := ds.RegisterTarget("12345", "testtarget", "weekly")

	assert.Nil(t, err)
	assert.Equal(t, res.Name, "testtarget")
	assert.Equal(t, res.Schedule, "weekly")
}
