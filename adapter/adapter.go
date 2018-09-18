package adapter

import "github.com/datakube/datakube/types"

type DumpAdapter interface {
	Dump(string) (types.DumpResult, error)
}

func CreateNewAdapter(host string, port string, database string, user string, password string, dbtype string) (DumpAdapter, error) {

	var err error
	var adapter DumpAdapter

	switch dbtype {
	case "mysql":
		da := newSqlAdapter(host, port, database, user, password)
		adapter = da
	}
	return adapter, err
}
