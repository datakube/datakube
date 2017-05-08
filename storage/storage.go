package storage

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
	ReadFile()
	DeleteFile()
}
