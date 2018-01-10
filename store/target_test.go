package store_test

import (
	"testing"
	"github.com/SantoDE/datahamster/store"
	"github.com/stretchr/testify/assert"
	"github.com/SantoDE/datahamster/types"
)

func TestSaveTargetOK(t *testing.T) {
	store, err := store.NewStore("/tmp/test.db")
	defer store.Close()

	assert.Nil(t, err)

	err = store.Open()
	assert.Nil(t, err)

	target := types.DumpTarget{
		Name:     "testtarget",
		Schedule: "weekly",
	}

	savedTarget, err := store.SaveTarget(target)
	assert.Nil(t, err)
	assert.NotNil(t, savedTarget.ID)
	assert.Equal(t, savedTarget.Name, "testtarget")
	assert.Equal(t, savedTarget.Schedule, "weekly")
}

