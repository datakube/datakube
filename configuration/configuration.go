package configuration

// GlobalConfiguration struct to hold Application config
type GlobalConfiguration struct {
	Server   ServerConfiguration
	LogLevel string
}

// ServerConfiguration struct to hold Server config
type ServerConfiguration struct {
	Address string
}

// Target struct to hold Target config
type Target struct {
	TargetType string
	Schedule   ScheduleConfiguration
	DBConfig   DatabaseConfiguration
}

// DumperConfiguration struct to hold Dumper config
type DumperConfiguration struct {
	Targets  []Target
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
	Interval string `description:"Interval to run at i.e. weekly"`
	Day      string `description:"Number of day in month i.e: Monday for week interval and Number (3) for Monthly Interval"`
	At       string `description:"When to start the Dump, i.e. 03:00"`
	StartNow bool   `description:"Interval to run at i.e. 8h or 5m"`
}
