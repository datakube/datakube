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


// DumperConfiguration struct to hold Dumper config
type DumperConfiguration struct {
	LogLevel string `mapstructure:"logLevel"`
	Token 	 string `mapstructure:"token"`
	Targets  []Target `mapstructure:"target"`
}

// Target struct to hold Target config
type Target struct {
	TargetType string `mapstructure:"type"`
	Name string `mapstructure:"name"`
	StartImmediately bool `mapstructure:"start_immediately"`
	DBConfig   DatabaseConfiguration `mapstructure:"db"`
	Schedule    ScheduleConfiguration `mapstructure:"schedule"`
}

// DatabaseConfiguration struct to hold Database Configuration
type DatabaseConfiguration struct {
	DatabaseHost     string `description:"Database Host" mapstructure:"host"`
	DatabaseName     string `description:"Database Name" mapstructure:"name"`
	DatabaseUserName string `description:"Database User Name" mapstructure:"user"`
	DatabasePassword string `description:"Database Password" mapstructure:"password"`
	DatabasePort     string `description:"Database Port" mapstructure:"port"`
	DatabaseType     string `description:"Database Type" mapstructure:"type"`
	SQL SQLDatabaseConfiguration `description:"Database Type" mapstructure:"sql"`
}

// SQLDatabaseConfiguration struct to hold SQL Database specific Configuration
type SQLDatabaseConfiguration struct {
	TempDir string `description:"TempDir" mapstructure:"tempdir"`
}

// ScheduleConfiguration struct to hold Schedule Configuration
type ScheduleConfiguration struct {
	Interval string `description:"Interval to run at i.e. weekly" mapstructure:"interval" `
	Day      string `description:"Number of day in month i.e: Monday for week interval and Number (3) for Monthly Interval" mapstructure:"day" `
	At       string `description:"When to start the Dump, i.e. 03:00" mapstructure:"at"`
	StartNow bool   `description:"Interval to run at i.e. 8h or 5m" mapstructure:"db_name"`
}
