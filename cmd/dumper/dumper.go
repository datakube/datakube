package main

import (
	"fmt"
	"github.com/SantoDE/datahamster/configuration"
	"github.com/SantoDE/datahamster/dumper"
	"github.com/SantoDE/datahamster/log"
	"github.com/SantoDE/datahamster/storage"
	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "database.name",
			Usage:  "The Database Name to Backup",
			EnvVar: "DATABASE_NAME",
		},
		cli.StringFlag{
			Name:   "database.host",
			Usage:  "The Database Host to Connect to",
			EnvVar: "DATABASE_HOST",
		},
		cli.StringFlag{
			Name:   "database.user",
			Usage:  "The User to use for the connection",
			EnvVar: "DATABASE_USER",
		},
		cli.StringFlag{
			Name:   "database.password",
			Usage:  "The Password to use for the connection",
			EnvVar: "DATABASE_PASSWORD",
		},
		cli.StringFlag{
			Name:   "database.port",
			Usage:  "The Password to use for the connection",
			EnvVar: "DATABASE_PORT",
		},
		cli.StringFlag{
			Name:   "database.type",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "DATABASE_TYPE",
		},
		cli.StringFlag{
			Name:   "database.sql.tempdir",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "DATABASE_TYPE",
		},
		cli.StringFlag{
			Name:   "schedule-duration",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
		cli.StringFlag{
			Name:   "start-now",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
		cli.StringFlag{
			Name:   "log-level",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
		cli.StringFlag{
			Name:   "storage.type",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
		cli.StringFlag{
			Name:   "storage.file.dir",
			Usage:  "The Database Type to connect to (currently SQL is supported only)",
			EnvVar: "SCHEDULE_DURATION",
		},
	}

	app.Action = func(c *cli.Context) error {
		DumperConfiguration := initConfig(c)
		level, err := logrus.ParseLevel(strings.ToLower(DumperConfiguration.LogLevel))
		if err != nil {
			log.Error("Error getting level", err)
		}
		log.SetLevel(level)

		//exit := make(chan struct{}, 1)

		dumper.StartWorker(&DumperConfiguration)

		return nil
	}

	app.Name = "Datahamster - Worker"
	app.Usage = "Worker to automatically get databse dumps and forward them to the server"

	app.Run(os.Args)
}

func initConfig(c *cli.Context) configuration.DumperConfiguration {

	var dataBaseName = c.String("database.name")
	var databaseHost = c.String("database.host")
	var databaseUser = c.String("database.user")
	var databaseType = c.String("database.type")
	var databasePort = c.String("database.port")
	var databasePassword = c.String("database.password")
	var logLevel = c.String("log-level")

	storageConfig := new(storage.Configuration)

	switch storageType := c.String("storage.type"); storageType {

	case "file":
		var storageDir = c.String("storage.file.dir")
		storageConfig.File = storage.FileConfiguration{
			Dir: storageDir,
		}

	default:
		storageConfig.Type = storageType
		fmt.Printf("Default")
	}

	dbConfig := configuration.DatabaseConfiguration{
		DatabaseName:     dataBaseName,
		DatabaseType:     databaseType,
		DatabasePassword: databasePassword,
		DatabaseUserName: databaseUser,
		DatabaseHost:     databaseHost,
		DatabasePort:     databasePort,
	}

	switch databaseType := c.String("database.type"); databaseType {

	case "SQL":
		var tempDir = c.String("database.sql.tempdir")
		dbConfig.SQL = configuration.SQLDatabaseConfiguration{
			TempDir: tempDir,
		}
	default:
		dbConfig.DatabaseType = databaseType
		fmt.Printf("Default")
	}

	config := configuration.DumperConfiguration{
		LogLevel: logLevel,
		Targets:  make([]configuration.Target, 1),
	}

	return config
}
