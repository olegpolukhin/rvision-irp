package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type PostgresConf struct {
	Host     string
	Port     uint
	Database string
	Username string
	Password string
}

type Auth struct {
	Username string
	APIToken string
}

type AppConfig struct {
	Postgres PostgresConf
	Auth     Auth
	URL      string
}

// Init config
func Init() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appjoin/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appjoin") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		log.Println(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
}
