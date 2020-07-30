package tests

import (
	"github.com/olegpolukhin/rvision-irp/cmd/server"
	"github.com/olegpolukhin/rvision-irp/usecase"
	"testing"
)

func TestGetIncidentList(t *testing.T) {
	revision := server.NewServer()

	incident := usecase.NewIncidentUsecase(revision)

	list, err := incident.GetIncidentList()

	if err != nil {
		t.Errorf("error %v\n", err)
	}

	if len(list.Incidents) == 0 {
		t.Errorf("return no list\n")
		return
	}
}

func TestGetIncident(t *testing.T) {
	revision := server.NewServer()

	incident := usecase.NewIncidentUsecase(revision)

	list, err := incident.GetIncident("1")
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
