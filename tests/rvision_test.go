package tests

import (
	"github.com/olegpolukhin/rvision-irp/cmd/server"
	"github.com/olegpolukhin/rvision-irp/config"
	"github.com/olegpolukhin/rvision-irp/usecase"
	"testing"
)

func serverConfigInit() *server.App {
	return server.NewServer(&config.AppConfig{
		Postgres: config.PostgresConf{
			Host:     "127.0.0.1",
			Port:     5432,
			Database: "chimney1",
			Username: "postgres",
			Password: "12345678",
		},
		Auth: config.Auth{},
		URL:  "http://10.10.80.25:5015",
	})
}

func TestGetIncidentList(t *testing.T) {
	revision := serverConfigInit()

	incident := usecase.NewIncidentUsecase(revision)

	list, err := incident.RvisionIncidentList()

	if err != nil {
		t.Errorf("error %v\n", err)
	}

	if len(list.Incidents) == 0 {
		t.Errorf("return no list\n")
		return
	}
}

func TestGetIncident(t *testing.T) {
	revision := serverConfigInit()

	incident := usecase.NewIncidentUsecase(revision)

	list, err := incident.RevisionIncident("1")
	if err != nil {
		t.Errorf("error %v\n", err)
	}

	if len(list.Incidents) == 0 {
		t.Errorf("return no list\n")
		return
	}
}

//func TestGetIncidentDB(t *testing.T) {
//	revision := server.NewServer()
//
//	incident := repository.NewIncidentUsecase(revision)
//
//	list, err := incident.GetIncidentDB("e9de2c10-b813-4d05-bd2b-95d223802775")
//	if err != nil {
//		t.Errorf("error %v\n", err)
//	}
//
//	if len(list.ID) == 0 {
//		t.Errorf("return no list\n")
//		return
//	}
//
//	log.Println(list)
//}
