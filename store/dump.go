package store

import (
	"errors"
	"github.com/asdine/storm/q"
	"github.com/datakube/datakube/types"
)

//Save function to save the given Dumper
func (d *DataStore) SaveDumpFile(file types.DumpFile) (types.DumpFile, error) {
	err := d.db.Save(&file)

	if err != nil {
		return *new(types.DumpFile), err
	}

	return file, nil
}

func (d *DataStore) LoadOneDumpFileByTarget(targetName string) (types.DumpFile, error) {

	var df types.DumpFile
	query := d.db.Select(q.Eq("Target", targetName)).OrderBy("CreatedAt").Reverse()
	err := query.First(&df)

	if err != nil {
		return *new(types.DumpFile), err
	}

	return df, nil
}

func (d *DataStore) LoadOneDumpFileByName(fileName string) (types.DumpFile, error) {

	var dfs []types.DumpFile

	err := d.db.Select().Find(&dfs)

	if err != nil {
		return *new(types.DumpFile), err
	}

	for _ , df := range dfs {
		if df.File.Name == fileName {
			return df, nil
		}
	}

	return types.DumpFile{}, errors.New("No file found by given filename")
}

func (d *DataStore) ListAllDumpFiles() ([]types.DumpFile, error) {

	var dfs []types.DumpFile
	err := d.db.All(&dfs)

	if err != nil {
		return []types.DumpFile{}, err
	}

	return dfs, nil
}