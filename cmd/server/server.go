package server

import (
	"github.com/olegpolukhin/rvision-irp/config"
	"github.com/olegpolukhin/rvision-irp/pkg/client"
	"github.com/olegpolukhin/rvision-irp/pkg/datasource"
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

func NewServer() *App {
	configT := config.AppConfig{}

	configT.Postgres.Host = "127.0.0.1"
	configT.Postgres.Database = "ex"
	configT.Postgres.Username = "ex"
	configT.Postgres.Password = "ex"
	configT.Postgres.Port = 5432

	configT.Auth = config.Auth{}
	configT.URL = "ex"

	return initApp(&configT)
}
