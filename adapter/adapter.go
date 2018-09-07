package adapter

import "github.com/SantoDE/datahamster/types"

type DumpAdapter interface {
	Dump(string) (types.DumpResult, error)
}

func CreateNewAdapter(host string, port string, database string, user string, password string, dbtype string) DumpAdapter {
	switch dbtype {
	case "mysql":
		da := newSqlAdapter(host, port, database, user, password)
		da.connect()
		return da
	}
	return nil
}
