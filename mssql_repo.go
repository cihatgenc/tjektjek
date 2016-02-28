package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	//"golang.org/x/sys/windows/svc/mgr"
	"log"
	//"os"
)

// GetSQLServiceNames - Return all sql service names
func GetSQLServiceNames() ([]string, []string, error) {
	//fmt.Printf("Executing GetSQLServiceNames\n")
	var servicenames []string
	var instancenames []string

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		//fmt.Printf("No installed SQL Service found\n")
		//log.Warn(err)
		return servicenames, instancenames, nil
	}
	defer k.Close()

	stat, err := GetKeyStats(k)
	if err != nil {
		log.Fatal(err)
	}

	names, err := GetKeyNames(k, int(stat.ValueCount))
	if err != nil {
		log.Fatal(err)
	}

	host, err := GetHostName()
	if err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		if name != "MSSQLSERVER" {
			servicenames = append(servicenames, fmt.Sprintf("%s%s", "MSSQL$", name))
		} else {
			servicenames = append(servicenames, name)
		}

		if name == "MSSQLSERVER" {
			instancenames = append(instancenames, host)
		} else {
			instancenames = append(instancenames, fmt.Sprintf("%s\\%s", host, name))
		}
	}

	return servicenames, instancenames, nil
}

// GetSQLAgentNames - Return all sql agent names
func GetSQLAgentNames() {
	fmt.Printf("Executing GetSQLAgentNames")

}

// GetSQLRSNames - Return all sql agent names
func GetSQLRSNames() {
	fmt.Printf("Executing GetSQLRSNames")

}

// GetSQLASNames - Return all sql agent names
func GetSQLASNames() {
	fmt.Printf("Executing GetSQLASNames")

}

// GetSQLISNames - Return all sql agent names
func GetSQLISNames() {
	fmt.Printf("Executing GetSQLISNames")

}
