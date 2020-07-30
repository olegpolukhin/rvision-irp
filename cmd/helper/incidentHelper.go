package helper

import (
	"context"
	"github.com/olegpolukhin/rvision-irp/cmd/server"
	"github.com/olegpolukhin/rvision-irp/repository/pgsql"
	"github.com/olegpolukhin/rvision-irp/usecase"
	"log"
)

// EventIncidentListToDB событие получить данные индцидентов и записать в таблицу
func EventIncidentListToDB(ctx context.Context) {
	var err error

	defer handlerError(&err)

	newServer := server.NewServer()
	incidentUcase := usecase.NewIncidentUsecase(newServer)

	listServ, err := incidentUcase.GetIncidentList()
	if err != nil {
		log.Println("listServ Error", err)
		return
	}

	log.Println(listServ)

	incidentDB := pgsql.NewIncidentRepository(newServer)
	list, err := incidentDB.Get("e9de2c10-b813-4d05-bd2b-95d223802775")
	if err != nil {
		log.Println("GetIncidentDB error %v\n", err)
		return
	}

	if len(list.ID) == 0 {
		log.Println("GetIncidentDB no list\n")
		return
	}

	log.Println(list)
}
