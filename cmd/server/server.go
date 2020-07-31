package server

import (
	"github.com/olegpolukhin/rvision-irp/config"
	"github.com/olegpolukhin/rvision-irp/pkg/client"
	"github.com/olegpolukhin/rvision-irp/pkg/datasource"
	"github.com/spf13/viper"
	"net/http"
	"upper.io/db.v3/postgresql"
)

type App struct {
	Config         *config.AppConfig
	Postgres       *datasource.Postgres
	RevisionClient *client.RevisionApp
}

func initApp(config *config.AppConfig) *App {
	app := App{}

	app.Config = config

	pgConnUri := postgresql.ConnectionURL{
		Host:     config.Postgres.Host,
		Database: config.Postgres.Database,
		User:     config.Postgres.Username,
		Password: config.Postgres.Password,
	}
	app.Postgres = datasource.NewPgDatasource(pgConnUri)

	app.RevisionClient = &client.RevisionApp{
		Auth:    &config.Auth,
		BaseURL: config.URL,
		Client:  http.DefaultClient,
	}

	return &app
}

func NewServer(baseConfig *config.AppConfig) *App {
	configTemp := config.AppConfig{}

	if baseConfig == nil {
		configTemp.Postgres.Host = viper.GetString("db_host")
		configTemp.Postgres.Database = viper.GetString("db_name")
		configTemp.Postgres.Username = viper.GetString("db_username")
		configTemp.Postgres.Password = viper.GetString("db_password")
		configTemp.Postgres.Port = viper.GetSizeInBytes("db_port")

		configTemp.Auth = config.Auth{}
		configTemp.URL = viper.GetString("api_url")
	} else {
		configTemp = *baseConfig
	}

	return initApp(&configTemp)
}
