package configuration

// GlobalConfiguration struct to hold Application config
type GlobalConfiguration struct {
	Schedule ScheduleConfiguration
	Database DatabaseConfiguration
	Storage  StorageConfiguration
	LogLevel string
}

// DatabaseConfiguration struct to hold Database Configuration
type DatabaseConfiguration struct {
	DatabaseName     string `description:"Database Name"`
	DatabaseUserName string `description:"Database User Name"`
	DatabasePassword string `description:"Database Password"`
	DatabaseHost     string `description:"Database Host"`
	DatabasePort     string `description:"Database Port"`
	DatabaseType     string `description:"Database Type"`
	SQL              SQLDatabaseConfiguration
}

// SQLDatabaseConfiguration struct to hold SQL Database specific Configuration
type SQLDatabaseConfiguration struct {
	TempDir string
}

// ScheduleConfiguration struct to hold Schedule Configuration
type ScheduleConfiguration struct {
	Interval string `description:"Interval to run at i.e. 8h or 5m"`
	StartNow bool   `description:"Interval to run at i.e. 8h or 5m"`
}

// StorageConfiguration struct to hold Storage Configuration
type StorageConfiguration struct {
	Type string `description:"Persistent File Storage"`
	File FileStorageConfiguration
}

// FileStorageConfiguration struct to hold File Storage specific Configuration
type FileStorageConfiguration struct {
	Dir string `description:"Persistent File Storage"`
}
