package target_test

import (
	"github.com/SantoDE/datahamster/store/target"
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTargetStoreOk(t *testing.T) {
	store := new(target.Store)

	testChan := make(chan types.ConfigTargets)

	var targets []types.Target

	target1 := types.Target{
		Name: "testtarget",
	}

	targets = append(targets, target1)

	go store.Subscribe(testChan)

	assert.Nil(t, store.ListTargets())
	assert.Equal(t, len(store.ListTargets()), 0)

	target, _ := store.GetOneTargetByName("testtarget")
	assert.Equal(t, target, *new(types.Target))

	testChan <- types.ConfigTargets{
		Targets: targets,
	}

	time.Sleep(200 * time.Millisecond)

	target, _ = store.GetOneTargetByName("testtarget")
	assert.NotNil(t, store.ListTargets())
	assert.Equal(t, len(store.ListTargets()), 1)
	assert.Equal(t, target.Name, target1.Name)
}
