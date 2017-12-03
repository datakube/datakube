package configuration

import "github.com/SantoDE/datahamster/storage"

type GlobalConfiguration struct {
	Server	 ServerConfiguration
	LogLevel string
	Dumps	 []DumpConfiguration
}

type ServerConfiguration struct {
	Address	string
}

type DumpConfiguration struct {
	Identifier 	string
	StorageType	string
	Storage		storage.StorageConfiguration
}