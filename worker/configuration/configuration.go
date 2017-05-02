package configuration

// GlobalConfiguration struct to hold Application config
type GlobalConfiguration struct {
	Schedule ScheduleConfiguration
	Database DatabaseConfiguration
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
}

// ScheduleConfiguration struct to hold Schedule Configuration
type ScheduleConfiguration struct {
	Interval string `description:"Interval to run at i.e. 8h or 5m"`
	StartNow bool   `description:"Interval to run at i.e. 8h or 5m"`
}
