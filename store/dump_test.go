package store_test

import (
	"github.com/SantoDE/datahamster/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveDumpFile(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	df := types.DumpFile{
		File: types.File{
			Name: "testfile",
			Path: "/test/file",
		},
		Target: "testtarget",
	}

	savedDumpFile, err := store.SaveDumpFile(df)
	assert.Nil(t, err)
	assert.NotNil(t, savedDumpFile.ID)
	assert.Equal(t, savedDumpFile.Target, "testtarget")
	assert.Equal(t, savedDumpFile.File.Path, "/test/file")
	assert.Equal(t, savedDumpFile.File.Name, "testfile")
}
