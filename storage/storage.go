package storage

// StorageConfiguration struct to hold Storage Configuration
type StorageConfiguration struct {
	Type string `description:"Persistent File Storage"`
	File FileStorageConfiguration
}

// FileStorageConfiguration struct to hold File Storage specific Configuration
type FileStorageConfiguration struct {
	Dir string `description:"Persistent File Storage"`
}

// BaseStorage struct which holds basic configuration for all dumpers
type BaseStorage struct {
}

// File struct which holds the file to save
type File struct {
	Name string
	Path string
}

// Storage Interface for all dumpers
type Storage interface {
	SaveFile(file File) error
	ListFiles()
	DeleteFile()
}
