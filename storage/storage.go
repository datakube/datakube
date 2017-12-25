package storage

// Configuration struct to hold Storage Configuration
type Configuration struct {
	Type string `description:"Persistent File Storage"`
	File FileConfiguration
}

// FileConfiguration struct to hold File Storage specific Configuration
type FileConfiguration struct {
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
