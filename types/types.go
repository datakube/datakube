package types

import (
	"time"
)

const STATUS_QUEUED = "queued"
const STATUS_SUCCESS = "success"

type ConfigTargets struct {
	Targets []Target `toml:"target"`
}

// Target struct to hold Target config
type Target struct {
	Name             string   `toml:"name"`
	StartImmediately bool     `toml:"start_immediately"`
	DBConfig         Database `toml:"db"`
	Schedule         Schedule `toml:"schedule"`
}

// ScheduleConfiguration struct to hold Schedule Configuration
type Schedule struct {
	Interval string `description:"Interval to run at i.e. weekly" toml:"interval" `
	Day      string `description:"Number of day in month i.e: Monday for week interval and Number (3) for Monthly Interval" toml:"day" `
	At       string `description:"When to start the Dump, i.e. 03:00" toml:"at"`
	StartNow bool   `description:"Interval to run at i.e. 8h or 5m" toml:"db_name"`
}

// DatabaseConfiguration struct to hold Database Configuration
type Database struct {
	DatabaseHost     string      `description:"Database Host" toml:"host"`
	DatabaseName     string      `description:"Database Name" toml:"name"`
	DatabaseUserName string      `description:"Database User Name" toml:"user"`
	DatabasePassword string      `description:"Database Password" toml:"password"`
	DatabasePort     string      `description:"Database Port" toml:"port"`
	DatabaseType     string      `description:"Database Type" toml:"type"`
	SQL              SQLDatabase `description:"Database Type" toml:"sql"`
}

// SQLDatabaseConfiguration struct to hold SQL Database specific Configuration
type SQLDatabase struct {
	TempDir string `description:"TempDir" mapstructure:"tempdir"`
}

type DumpFile struct {
	ID        int `storm:"id,increment"`
	CreatedAt time.Time
	File      File
	Target    string
}

// File struct which holds the file to save
type File struct {
	Name string `json:"name"`
	Path string
	Data []byte
}

// DumpResult struct to hold a general result for a dump
type DumpResult struct {
	Success       bool
	TemporaryFile string
	TargetName    string
}

// Job struct to hold information about what to dump when and be the provider for the polling
type Job struct {
	ID     int       `storm:"id,increment"`
	Status  string    `json:"state"`
	Target string    `json:"provider"`
	RunAt  time.Time `json:"runAt"`
}
