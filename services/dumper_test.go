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
			Targets: targets,
		}, nil
	}

	return types.Dumper{}, errors.New("no dumper with token")
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

func TestRegisterTargetNOK(t *testing.T) {
	noHit = true
	testStore := new(TestDataStore)
	ds := services.NewDumperService(testStore)

	res, err := ds.RegisterTarget("12345", "testtarget", "weekly")

	assert.NotNil(t, err)
	assert.Equal(t, res.ID, 0)
	noHit = false
}

func TestSaveTargetFileOK(t *testing.T) {
	testStore := new(TestDataStore)
	ds := services.NewDumperService(testStore)

	res, err := ds.SaveTargetFile("12345", "existing Target", "testfile", "/tmp/test123.txt")

	assert.Nil(t, err)
	assert.Equal(t, res.ID, 0)
	assert.Equal(t, res.File.Path, "/tmp/test123.txt")
	assert.Equal(t, res.File.Name, "testfile")
}

