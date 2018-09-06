package store

import "github.com/SantoDE/datahamster/types"

//Save function to save the given Dumper
func (d *DataStore) SaveDumpFile(file types.DumpFile) (types.DumpFile, error) {
	err := d.db.Save(&file)

	if err != nil {
		return *new(types.DumpFile), err
	}

	return file, nil
}
