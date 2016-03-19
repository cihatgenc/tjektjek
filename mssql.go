package main

import (
	"fmt"
	"log"
	"os"
	//"golang.org/x/sys/windows/registry"
	//"golang.org/x/sys/windows/svc/mgr"
)

var sensuOk = 0
var sensuWarning = 1
var sensuCritical = 2
var sensuUnknown = 3

var sensuOktext = "OK - "
var sensuWarningtext = "WARNING - "
var sensuCriticaltext = "CRITICAL - "
var sensuUnknowntext = "UNKNOWN - "

// statusSQLServices - Check SQL Server Services
func statusSQLServices() {
	//fmt.Printf("Executing checkSQLServices\n")
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
	fmt.Println(message)
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
