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
	var servicenames []string
	var instancenames []string

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
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

// GetSQLInstanceNames - Return all sql instance names
func GetSQLInstallNames() ([]string, []string, error) {
	var installnames []string
	var instancenames []string

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		return installnames, instancenames, nil
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
		installname, err := GetKeyStrValues(k, name)
		if err != nil {
			log.Fatal(err)
		}
		installnames = append(installnames, installname)

		if name == "MSSQLSERVER" {
			instancenames = append(instancenames, host)
		} else {
			instancenames = append(instancenames, fmt.Sprintf("%s\\%s", host, name))
		}
	}

	return installnames, instancenames, nil
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

func GetSQLVersion(instance string) (string, error) {
	var regkey string
	var version string
	searchkey := "PatchLevel"

	regkey = fmt.Sprintf("%s%s%s", `SOFTWARE\Microsoft\Microsoft SQL Server\`, instance, `\Setup`)

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, regkey, registry.QUERY_VALUE)
	if err != nil {
		return version, nil
	}
	defer k.Close()

	version, err = GetKeyStrValues(k, searchkey)
	if err != nil {
		log.Fatal(err)
	}

	return version, nil
}

func GetSQLEdition(instance string) (string, error) {
	var regkey string
	var edition string
	searchkey := "Edition"

	regkey = fmt.Sprintf("%s%s%s", `SOFTWARE\Microsoft\Microsoft SQL Server\`, instance, `\Setup`)

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, regkey, registry.QUERY_VALUE)
	if err != nil {
		return edition, nil
	}
	defer k.Close()

	edition, err = GetKeyStrValues(k, searchkey)
	if err != nil {
		log.Fatal(err)
	}

	return edition, nil
}
