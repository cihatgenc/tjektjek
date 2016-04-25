package main

type mssqlinfo struct {
	ServiceName  string `json:"servicename"`
	InstallName  string `json:"installname"`
	InstanceName string `json:"instancename"`
	Port         string `json:"port"`
}

type mssqlinfos []mssqlinfo
