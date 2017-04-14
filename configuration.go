package main

type GlobalConfiguration struct {
	Schedule ScheduleConfiguration
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	DatabaseName            string		`description:"Database Name"`
	DatabaseUserName        string		`description:"Database User Name"`
	DatabasePassword	string		`description:"Database Password"`
	DatabaseHost         	string		`description:"Database Host"`
	DatabaseType          	string		`description:"Database Type"`
}

type ScheduleConfiguration struct {
	Interval       	  string		`description:"Interval to run at i.e. 8h or 5m"`
}