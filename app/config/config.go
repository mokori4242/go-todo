package config

import (
	"log"
	"gopkg.in/ini.v1"
	"app/utils"
)

type ConfigList struct {
    AppPort string
	DBUser string
    DBPassword string
    DBHost string
    DBPort string
	DBName string
	LogFile string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSetting(Config.LogFile)
	InitDB()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		AppPort: cfg.Section("app").Key("port").MustString("8080"),
		DBUser: cfg.Section("db").Key("user").MustString("posgre"),
		DBPassword: cfg.Section("db").Key("password").MustString("password"),
		DBHost: cfg.Section("db").Key("host").MustString("db"),
		DBPort: cfg.Section("db").Key("port").MustString("5432"),
		DBName: cfg.Section("db").Key("name").MustString("posgre"),
		LogFile: cfg.Section("app").Key("logfile").MustString("app.log"),
	}
}