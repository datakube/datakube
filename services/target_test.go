package services_test

import (
	"github.com/SantoDE/datahamster/services"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/types"
	"errors"
)

var _ store.TargetStore = (*TestTargetDataStore)(nil)

type TestTargetDataStore struct{}

func (t *TestTargetDataStore) OneById(tokiden string) (types.DumpTarget, error) {
	var target types.DumpTarget

	target = types.DumpTarget{Name: "existing Target", Schedule: "weekly"}

	if !noHit {
		return target, nil
	}

	return types.DumpTarget{}, errors.New("no dumper with token")
}

func (t *TestTargetDataStore) OneByName(token string) (types.DumpTarget, error) {
	var targets []types.DumpTarget

	targets = append(targets, types.DumpTarget{Name: "existing Target", Schedule: "weekly"})

	if !noHit {
		return types.DumpTarget{
			Name:    "Testdumper",
		}, nil
	}

	return types.DumpTarget{}, errors.New("no dumper with token")
}

func (t *TestTargetDataStore) SaveTarget(dumper types.DumpTarget) (types.DumpTarget, error) {
	return dumper, nil
}


func TestRegisterTargetOK(t *testing.T) {
	testStore := new(TestTargetDataStore)
	ds := services.NewTargetService(testStore)

	res, err := ds.RegisterTarget("testtarget", "weekly")

	assert.Nil(t, err)
	assert.Equal(t, res.Name, "testtarget")
	assert.Equal(t, res.Schedule, "weekly")
}

func TestSaveTargetFileOK(t *testing.T) {
	testStore := new(TestTargetDataStore)
	ts := services.NewTargetService(testStore)

	res, err := ts.SaveTargetFile("existing Target", "testfile", "/tmp/test123.txt")

	assert.Nil(t, err)
	assert.Equal(t, res.ID, 0)
	assert.Equal(t, res.File.Path, "/tmp/test123.txt")
	assert.Equal(t, res.File.Name, "testfile")
}


