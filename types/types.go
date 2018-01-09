package types

import "time"

//Dumper struct to hold Dumper Information
type Dumper struct {
	ID      int          `storm:"id,increment"`
	Token   string       `json:"token"`
	Name    string       `json:"name"`
	Targets []DumpTarget `json:"targets"`
}

type DumpTarget struct {
	ID       int    `storm:"id,increment"`
	Name     string `json:"name"`
	Schedule string `json:"schedule"`
	Files 	 []DumpFile
}

type DumpFile struct {
	ID        int    `storm:"id,increment"`
	CreatedAt time.Time
	File 	  File
}

// File struct which holds the file to save
type File struct {
	Name string `json:"name"`
	Path string
	Data []byte
}

// DumpResult struct to hold a general result for a dump
type DumpResult struct {
	Success bool
	TemporaryFile string
	TargetName string
}