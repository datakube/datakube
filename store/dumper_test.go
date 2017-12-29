package store_test

import (
	"testing"
	"github.com/SantoDE/datahamster/store"
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
)

func TestSaveOK(t *testing.T) {
	store, err := store.NewStore("/tmp/test.db")

	assert.Nil(t, err)

	err = store.Open()
	assert.Nil(t, err)

	var targets []types.DumpTarget

	targets = append(targets, types.DumpTarget{
		Name: "testtarget",
		Schedule: "weekly",
	})

	dumper := types.Dumper{
		Token: "1234",
		Name: "Testdumper",
		Targets: targets,
	}

	savedDumper, err := store.Save(dumper)
	assert.Nil(t, err)
	assert.NotNil(t, savedDumper.ID)
	assert.NotNil(t, savedDumper.Targets)
	assert.Equal(t, savedDumper.Name, "Testdumper")
	assert.Equal(t, savedDumper.Targets[0].Name, "testtarget")
}
