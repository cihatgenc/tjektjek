package main

import (
	// "fmt"
	"os"

	"github.com/codegangsta/cli"
	//"strings"
	//"log"
)

func letsStart() {
	app := cli.NewApp()
	app.Name = "tjektjek"
	app.Usage = "Checks your application"
	app.Version = VERSION

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
		{
			Name:  "mssql_dbstates",
			Usage: "Check the SQL Server database states",
			Action: func(c *cli.Context) {
				statusSQLDbState()
			},
		},
		{
			Name:  "mssql_connection",
			Usage: "Check the SQL Server connection",
			Action: func(c *cli.Context) {
				statusSQLConnection()
			},
		},
		{
			Name:  "mssql_new",
			Usage: "Check A new feature",
			Action: func(c *cli.Context) {
				statusSQLNew()
			},
		},
	}
	app.Run(os.Args)
}
