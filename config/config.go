package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const LogFile = "/tmp/logrus.log"

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

// InitLogRus init logger logrus
func InitLogRus() {
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	//logger.SetOutput(os.Stdout)
	file, err := os.OpenFile(LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logger.Fatal(err)
	}
	//defer file.Close()
	logger.SetOutput(file)
}

// Init config
func Init() {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/appjoin/")  // path to look for the config file in
	viper.AddConfigPath("/tmp/appjoin/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appjoin") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory

	if err := viper.ReadInConfig(); err != nil {
		log.Error(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
}
