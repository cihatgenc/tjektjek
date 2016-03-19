package main

import (
	// "fmt"
	"github.com/codegangsta/cli"
	"os"
	//"strings"
	//"log"
)

func letsStart() {
	app := cli.NewApp()
	app.Name = "tjektjek"
	app.Usage = "Checks your application"
	app.Version = versionNumber

	app.Commands = []cli.Command{
		{
			Name:  "mssql_services",
			Usage: "Check the SQL Server windows services",
			Action: func(c *cli.Context) {
				statusSQLServices()
			},
		},
		{
			Name:  "mssql_version",
			Usage: "Check the SQL Server version",
			Action: func(c *cli.Context) {
				statusSQLVersion()
			},
		},
	}
	app.Run(os.Args)
}
