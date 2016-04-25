package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	//"strings"
)

const (
	sensuOk       = 0
	sensuWarning  = 1
	sensuCritical = 2
	sensuUnknown  = 3

	sensuOktext       = "OK - "
	sensuWarningtext  = "WARNING - "
	sensuCriticaltext = "CRITICAL - "
	sensuUnknowntext  = "UNKNOWN - "
)

// statusSQLServices - Check SQL Server Services
func statusSQLServices() {
	var message string
	var allok = true

	servicenames, instancenames, err := GetSQLServiceNames()
	if err != nil {
		log.Fatal(err)
	}

	for i, servicename := range servicenames {

		starttype, err := GetServiceStartType(servicename)
		if err != nil {
			log.Fatal(err)
		}

		state, err := GetServiceStatus(servicename)
		if err != nil {
			log.Fatal(err)
			allok = false
			message += fmt.Sprintf("%s=%s,", instancenames[i], "unknown")
			continue
		}

		if state != "running" && starttype == "automatic" {
			allok = false
		}
		message += fmt.Sprintf("%s=%s,", instancenames[i], state)
	}

	if message == "" {
		message = fmt.Sprintf("%s%s,", sensuUnknowntext, "No installed SQL Service found")
		fmt.Println(message)
		os.Exit(sensuUnknown)
	}

	message = message[0:(len(message) - 1)]
	if allok == false {
		message = fmt.Sprintf("%s%s", sensuCriticaltext, message)
		fmt.Println(message)
		os.Exit(sensuCritical)
	} else {
		message = fmt.Sprintf("%s%s", sensuOktext, message)
		fmt.Println(message)
		os.Exit(sensuOk)
	}
}

// statusSQLVersion - Check SQL Server Version and Edition
func statusSQLVersion() {
	var message string

	installnames, instancenames, err := GetSQLInstallNames()
	if err != nil {
		log.Fatal(err)
	}

	for i, installname := range installnames {
		version, err := GetSQLVersion(installname)
		if err != nil {
			log.Fatal(err)
			message += fmt.Sprintf("%s=%s,", instancenames[i], "unknown")
			continue
		}

		edition, err := GetSQLEdition(installname)
		if err != nil {
			log.Fatal(err)
			message += fmt.Sprintf("%s=%s,", instancenames[i], "unknown")
			continue
		}

		message += fmt.Sprintf("%s=%s %s,", instancenames[i], edition, version)
	}

	if message == "" {
		message = fmt.Sprintf("%s%s,", sensuUnknowntext, "No installed SQL Service found")
		fmt.Println(message)
		os.Exit(sensuUnknown)
	}

	message = message[0:(len(message) - 1)]
	message = fmt.Sprintf("%s%s", sensuOktext, message)
	fmt.Println(message)
	os.Exit(sensuOk)
}

// statusSQLConnection - Check SQL Server connection
func statusSQLConnection() {
	var message string
	var allok = true

	sqlservers, err := GetSQLServerInfo()
	if err != nil {
		log.Fatal(err)
	}

	for _, sqlserver := range sqlservers {
		//fmt.Println(sqlserver)

		db, err := GetSQLHandle(sqlserver.InstanceName, "master", sqlserver.Port)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		result := db.Ping()
		if result != nil {
			allok = false
			message += fmt.Sprintf("%s=%s,", sqlserver.InstanceName, "Failed")
			continue
		}

		message += fmt.Sprintf("%s=%s,", sqlserver.InstanceName, "Connected")
	}

	message = message[0:(len(message) - 1)]
	if allok == false {
		message = fmt.Sprintf("%s%s", sensuCriticaltext, message)
		fmt.Println(message)
		os.Exit(sensuCritical)
	} else {
		message = fmt.Sprintf("%s%s", sensuOktext, message)
		fmt.Println(message)
		os.Exit(sensuOk)
	}
}

// statusSQLDbState - Check SQL Server database states
func statusSQLDbState() {
	var instancename = "SBPLT182"
	//fmt.Println("Gonna do sql open")
	db, err := GetSQLHandle(instancename, "master", "1433")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Gonna do ping sql")
	blaat := db.Ping()
	if blaat != nil {
		//fmt.Println(blaat)
		os.Exit(2)
	}
	fmt.Println("It worked")
	os.Exit(0)
}

func statusSQLNew() {
	var message string
	var allok = true
	var sqlcmd = "select name from sys.sysdatabases where dbid = 4"

	sqlservers, err := GetSQLServerInfo()
	if err != nil {
		log.Fatal(err)
	}

	for _, sqlserver := range sqlservers {
		//fmt.Println(sqlserver)

		db, err := GetSQLHandle(sqlserver.InstanceName, "master", sqlserver.Port)
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		//		result := db.Ping()
		//		if result != nil {
		//			allok = false
		//			message += fmt.Sprintf("%s=%s,", sqlserver.InstanceName, "Failed")
		//			continue
		//		}

		//		message += fmt.Sprintf("%s=%s,", sqlserver.InstanceName, "Connected")

		stmt, err := db.Prepare(sqlcmd)
		fmt.Println(stmt)
		if err != nil {
			log.Fatal("Prepare failed:", err.Error())
		}
		defer stmt.Close()

		row := stmt.QueryRow()
		fmt.Println(row)
		var somechars string
		err = row.Scan(&somechars)
		if err != nil {
			log.Fatal("Scan failed:", err.Error())
		}

		fmt.Printf("somechars:%s\n", somechars)

	}

	message = message[0:(len(message) - 1)]
	if allok == false {
		message = fmt.Sprintf("%s%s", sensuCriticaltext, message)
		fmt.Println(message)
		os.Exit(sensuCritical)
	} else {
		message = fmt.Sprintf("%s%s", sensuOktext, message)
		fmt.Println(message)
		os.Exit(sensuOk)
	}
}
