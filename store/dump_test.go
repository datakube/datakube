package store_test

import (
	"github.com/datakube/datakube/types"
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
	assert.Equal(t, "not found", err.Error())
	assert.Empty(t, df2)
}

func TestDataStore_ListAllDumpFiles(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	df1 := types.DumpFile{
		File: types.File{
			Name: "testfile",
			Path: "/test/file",
		},
		Target: "testtarget",
	}
	df2 := types.DumpFile{
		File: types.File{
			Name: "testfile",
			Path: "/test/file",
		},
		Target: "testtarget2",
	}

	store.SaveDumpFile(df1)
	store.SaveDumpFile(df2)

	dfs, err := store.ListAllDumpFiles()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(dfs))
	assert.NotNil(t, dfs[0].ID)
	assert.Equal(t, dfs[0].Target, "testtarget")
	assert.Equal(t, dfs[0].File.Path, "/test/file")
	assert.Equal(t, dfs[0].File.Name, "testfile")
	assert.NotNil(t, dfs[1].ID)
	assert.Equal(t, dfs[1].Target, "testtarget2")
	assert.Equal(t, dfs[1].File.Path, "/test/file")
	assert.Equal(t, dfs[1].File.Name, "testfile")
}

func TestDataStore_LoadOneDumpFileByName(t *testing.T) {
	store := NewTestDataStore()
	defer store.Close()

	df1 := types.DumpFile{
		File: types.File{
			Name: "testfile",
			Path: "/test/file",
		},
		Target: "testtarget",
	}

	store.SaveDumpFile(df1)

	testFile, err := store.LoadOneDumpFileByName("testfile")
	assert.Nil(t, err)
	assert.NotNil(t, testFile.ID)
	assert.Equal(t,  "testtarget" ,testFile.Target,)
	assert.Equal(t, "/test/file", testFile.File.Path)
	assert.Equal(t, "testfile", testFile.File.Name)

}