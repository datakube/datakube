package services_test

import (
	"errors"
	"github.com/SantoDE/datahamster/services"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ store.DumperStore = (*TestDataStore)(nil)
var noHit = false

type TestDataStore struct{}

func (t *TestDataStore) One(token string) (types.Dumper, error) {
	var targets []types.DumpTarget

	targets = append(targets, types.DumpTarget{Name: "existing Target", Schedule: "weekly"})

	if !noHit {
		return types.Dumper{
			Token:   token,
			Name:    "Testdumper",
		}, nil
	}

	return types.Dumper{}, errors.New("no dumper with token")
}

func (t *TestDataStore) Save(dumper types.Dumper) (types.Dumper, error) {
	return dumper, nil
}

func TestValidateDumperOK(t *testing.T) {
	testStore := new(TestDataStore)
	ds := services.NewDumperService(testStore)

	res, err := ds.Validate("12345")

	assert.Nil(t, err)
	assert.Equal(t, res, true)
}


func TestSaveDumperOK(t *testing.T) {
	testStore := new(TestDataStore)
	ts := services.NewDumperService(testStore)

	res, err := ts.Create("TestDumper")

	assert.Nil(t, err)
	assert.Equal(t, res.Name, "TestDumper")
	assert.NotNil(t, res.ID)
	assert.NotNil(t, res.Token)
}
