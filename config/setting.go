package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	DbDriver   string
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
	DbName     string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Fail to read file: ", err)
	}
	AppMode = cfg.Section("server").Key("AppMode").String()
	HttpPort = cfg.Section("server").Key("HttpPort").String()
	DbDriver = cfg.Section("database").Key("DbDriver").String()
	DbHost = cfg.Section("database").Key("DbHost").String()
	DbPort = cfg.Section("database").Key("DbPort").String()
	DbUsername = cfg.Section("database").Key("DbUsername").String()
	DbPassword = cfg.Section("database").Key("DbPassword").String()
	DbName = cfg.Section("database").Key("DbName").String()
}
