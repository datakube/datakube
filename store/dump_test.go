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

func TestDataStore_LoadOneDumpFileByTarget(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	df := types.DumpFile{
		File: types.File{
			Name: "testfile",
			Path: "/test/file",
		},
		Target: "testtarget",
	}

	store.SaveDumpFile(df)

	df, err := store.LoadOneDumpFileByTarget("testtarget")
	assert.Nil(t, err)
	assert.NotNil(t, df.ID)
	assert.Equal(t, df.Target, "testtarget")
	assert.Equal(t, df.File.Path, "/test/file")
	assert.Equal(t, df.File.Name, "testfile")

	df2, err := store.LoadOneDumpFileByTarget("aaaa")
	assert.NotNil(t, err)
	assert.Equal(t,"not found", err.Error())
	assert.Empty(t, df2)
}