package main

import (
	"github.com/urfave/cli"
	"os"
)


func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "database-name",
			Usage: "The Database Name to Backup",
			EnvVar: "DATABASE_NAME",
		},
		cli.StringFlag{
			Name: "database-host",
			Usage: "The Database Host to Connect to",
			EnvVar: "DATABASE_HOST",
		},
		cli.StringFlag{
			Name: "database-user",
			Usage: "The User to use for the connection",
			EnvVar: "DATABASE_USER",
		},
		cli.StringFlag{
			Name: "database-password",
			Usage: "The Password to use for the connection",
			EnvVar: "DATABASE_PASSWORD",
		},
		cli.StringFlag{
			Name: "database-type",
			Usage: "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "DATABASE_TYPE",
		},
		cli.StringFlag{
			Name: "schedule-duration",
			Usage: "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
	}

	app.Action = func(c *cli.Context) error {

		globalConfiguration := initConfig(c)

		exit := make(chan struct{}, 1)
		manager := NewManager(globalConfiguration)
		manager.run(exit)

		return nil;
	}

	app.Name = "Datahamster - Worker"
	app.Usage = "Worker to automatically get databse dumps and forward them to the server"

	app.Run(os.Args);
}

func initConfig(c *cli.Context) GlobalConfiguration{

	var dataBaseName 		= c.String("database-name")
	var databaseHost		= c.String("database-host")
	var databaseUser		= c.String("database-user")
	var databaseType		= c.String("database-type")
	var databasePassword		= c.String("database-type")
	var scheduleDuration		= c.String("schedule-duration")

	config := GlobalConfiguration {
		Schedule: ScheduleConfiguration{
			Interval: scheduleDuration,
		},
		Database: DatabaseConfiguration{
			DatabaseName: dataBaseName,
			DatabaseType: databaseType,
			DatabasePassword: databasePassword,
			DatabaseUserName: databaseUser,
			DatabaseHost: databaseHost,
		},
	}

	return config
}