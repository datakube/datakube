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
	}
	app.Name = "Datahamster - Worker"
	app.Usage = "Worker to automatically get databse dumps and forward them to the server"

	app.Run(os.Args);
}

