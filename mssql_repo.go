package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"golang.org/x/sys/windows/registry"
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

// Get SQL Server information on all instances
func GetSQLServerInfo() (sqlinfo mssqlinfos, err error) {
	var result mssqlinfos

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		return nil, nil
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
		var resultrow mssqlinfo

		installname, err := GetKeyStrValues(k, name)
		if err != nil {
			log.Fatal(err)
		}

		resultrow.InstallName = installname
		resultrow.ServiceName = fmt.Sprintf("%s%s", "MSSQL$", name)

		if name == "MSSQLSERVER" {
			resultrow.InstanceName = host
		} else {
			resultrow.InstanceName = fmt.Sprintf("%s\\%s", host, name)
		}

		resultrow.Port, _ = GetSQLPort(installname)

		result = append(result, resultrow)
	}

	return result, nil
}

// Get SQL Server information on all instances that are active
func GetActiveSQLServerInfo() (sqlinfo mssqlinfos, err error) {
	var result mssqlinfos

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Microsoft SQL Server\Instance Names\SQL`, registry.QUERY_VALUE)
	if err != nil {
		return nil, nil
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
		var resultrow mssqlinfo

		installname, err := GetKeyStrValues(k, name)
		if err != nil {
			log.Fatal(err)
		}

		resultrow.InstallName = installname
		resultrow.ServiceName = fmt.Sprintf("%s%s", "MSSQL$", name)

		if name == "MSSQLSERVER" {
			resultrow.InstanceName = host
		} else {
			resultrow.InstanceName = fmt.Sprintf("%s\\%s", host, name)
		}

		resultrow.Port, _ = GetSQLPort(installname)

		result = append(result, resultrow)
	}

	return result, nil
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

func GetSQLVersion(installname string) (string, error) {
	var regkey string
	var version string
	searchkey := "PatchLevel"

	regkey = fmt.Sprintf("%s%s%s", `SOFTWARE\Microsoft\Microsoft SQL Server\`, installname, `\Setup`)

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

func GetSQLEdition(installname string) (string, error) {
	var regkey string
	var edition string
	searchkey := "Edition"

	regkey = fmt.Sprintf("%s%s%s", `SOFTWARE\Microsoft\Microsoft SQL Server\`, installname, `\Setup`)

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

func GetSQLPort(installname string) (port string, err error) {
	var regkey string
	var portnumber string

	staticport := "TcpPort"
	dynamicport := "TcpDynamicPorts"

	regkey = fmt.Sprintf("%s%s%s", `SOFTWARE\Microsoft\Microsoft SQL Server\`, installname, `\MSSQLServer\SuperSocketNetLib\Tcp\IPAll`)

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, regkey, registry.QUERY_VALUE)
	if err != nil {
		return "", nil
	}
	defer k.Close()

	portnumber, err = GetKeyStrValues(k, staticport)
	if err != nil {
		log.Fatal(err)
	}

	// If port is not static, look for the dynamic port
	if portnumber == "" {
		portnumber, err = GetKeyStrValues(k, dynamicport)
		if err != nil {
			log.Fatal(err)
		}
	}
	return portnumber, nil

}

func GetSQLHandle(instance string, database string, port string) (*sql.DB, error) {
	dsn := fmt.Sprintf("Server=%s;Database=%s;port=%s;", instance, database, port)
	//fmt.Println(dsn)
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println(err)
		fmt.Println("ben ik hier??" + instance)
		return db, err
	}
	// defer db.Close()

	return db, nil
}

func GetDatabaseStates(instance string) ([]string, error) {
	var result []string

	return result, nil
}
