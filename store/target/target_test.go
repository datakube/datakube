package target_test

import (
	"github.com/SantoDE/datahamster/internal/store"
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveTargetOK(t *testing.T) {
	store := store.NewTestDataStore()
	defer store.Close()

	target := types.Target{
		Name:     "testtarget",
		Schedule: "weekly",
	}

	savedTarget, err := store.SaveTarget(target)
	assert.Nil(t, err)
	assert.NotNil(t, savedTarget.ID)
	assert.Equal(t, savedTarget.Name, "testtarget")
	assert.Equal(t, savedTarget.Schedule, "weekly")
}

func TestTargetOneByName(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	target := types.Target{
		Name:     "testtarget",
		Schedule: "weekly",
	}

	store.SaveTarget(target)

	savedTarget, err := store.OneTargetByName("testtarget")
	assert.Nil(t, err)
	assert.NotNil(t, savedTarget.ID)
	assert.Equal(t, savedTarget.Name, "testtarget")
	assert.Equal(t, savedTarget.Schedule, "weekly")
}

func TestTargetOneById(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	target := types.Target{
		Name:     "testtarget",
		Schedule: "weekly",
	}

	saved, _ := store.SaveTarget(target)

	savedTarget, err := store.OneTargetById(saved.ID)
	assert.Nil(t, err)
	assert.NotNil(t, savedTarget.ID)
	assert.Equal(t, savedTarget.Name, "testtarget")
	assert.Equal(t, savedTarget.Schedule, "weekly")
}
