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

		fmt.Printf("State for %s is %s\n", instancenames[i], state)

		if state != "running" && starttype == "automatic" {
			allok = false
		}

		message += fmt.Sprintf("%s=%s,", instancenames[i], state)
	}

	if message == "" {
		message = "No installed SQL Service found"
		fmt.Println(message)
		os.Exit(sensuUnknown)
	}

	message = message[0:(len(message) - 1)]
	if allok == false {
		fmt.Println(message)
		os.Exit(sensuCritical)
	} else {
		fmt.Println(message)
		os.Exit(sensuOk)
	}
}
