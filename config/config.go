package config

import (
	"log"

	"github.com/go-ini/ini"
)

type ConfigT struct {
	Server       string `ini:"server"`
	ClientSecret string `ini:"client_secret"`
	ClientID     string `ini:"client_id"`
	SessionKey   string `ini:"session_key"`
	Host         string `ini:"host"`
}

var Config ConfigT

func LoadConfig() {
	confile, err := ini.Load("config.ini")
	if err != nil {
		log.Println("Failed to read config")
		log.Panic(err)
	}
	consection := confile.Section("")
	if err := consection.MapTo(&Config); err != nil {
		log.Println("Failed to read config")
		log.Panic(err)
	}
}
