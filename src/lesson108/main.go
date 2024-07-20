package main

import (
	"gopkg.in/go-ini/ini.v1"
)

// ini
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(8080),            // 8080 is default value
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"), // example.sql is default value
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	println(Config.Port)
	println(Config.DbName)
	println(Config.SQLDriver)
}
