package helper

import (
	"context"
	"github.com/olegpolukhin/rvision-irp/cmd/server"
	"github.com/olegpolukhin/rvision-irp/domain/model"
	"github.com/olegpolukhin/rvision-irp/repository/pgsql"
	"github.com/olegpolukhin/rvision-irp/usecase"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	DefaultUUID = "00000000-0000-0000-0000-000000000000"
)

// TODO НА РЕАЛИЗАЦИИ...

// TODO необходимо провести рефакторинг
// TODO от апи рвизион ожидается получить все заполненные данные, но некеоторые поля остаются пустыми, решить вопрос
// TODO в данном событии работу с БД поместить всю процедуру работы в отдельный поток

// EventIncidentListToDB событие получить данные индцидентов и записать в таблицу
func EventIncidentListToDB(ctx context.Context) {
	var err error

	defer handlerError(&err)

	newServer := server.NewServer(nil)
	incidentUcase := usecase.NewIncidentUsecase(newServer)

	revisionList, err := incidentUcase.RvisionIncidentList()
	if err != nil {
		log.Error("listServ Error", err)
		return
	}

	revisionIncidentMap := make(map[int]model.Incident)

	if len(revisionList.Incidents) == 0 {
		log.Error("GetRevisionIncidentList error. List is empty")
		return
	}

	for i := range revisionList.Incidents {
		revisionIncidentMap[revisionList.Incidents[i].ID] = revisionList.Incidents[i]
	}

	common := pgsql.NewCommonRepository(newServer)

	userList, err := common.UserList()
	if err != nil {
		log.Error("UserList error %v\n", err)
	}

	userListMap := make(map[string]string)
	for i := range userList {
		userListMap[userList[i].Name] = userList[i].ID
	}

	organizationList, err := common.OrganizationList()
	if err != nil {
		log.Error("OrganizationList error %v\n", err)
	}

	//organizationListMap := make(map[string]string)
	//for i := range organizationList {
	//	organizationListMap[organizationList[i].Name] = organizationList[i].ID
	//}

	networkList, err := common.NetworkSegmentList()
	if err != nil {
		log.Error("NetworkSegmentList error %v\n", err)
	}

	szisList, err := common.SziList()
	if err != nil {
		log.Error("SziList error %v\n", err)
	}

	categoryList, err := common.IncidentCategoryList()
	if err != nil {
		log.Error("IncidentCategoryList error %v\n", err)
	}

	categoryListMap := make(map[string]string)
	for i := range categoryList {
		categoryListMap[categoryList[i].Name] = categoryList[i].ID
	}

	priorityList, err := common.IncidentPriorityList()
	if err != nil {
		log.Error("IncidentPriorityList error %v\n", err)
	}

	priorityListMap := make(map[string]string)
	for i := range priorityList {
		priorityListMap[priorityList[i].Name] = priorityList[i].ID
	}

	statusList, err := common.IncidentStatusList()
	if err != nil {
		log.Error("IncidentStatusList error %v\n", err)
	}

	statusListMap := make(map[string]string)
	for i := range statusList {
		statusListMap[statusList[i].Name] = statusList[i].ID
	}

	incidentTypeList, err := common.IncidentTypeList()
	if err != nil {
		log.Error("IncidentTypeList error %v\n", err)
	}

	incidentTypeListMap := make(map[string]string)
	for i := range incidentTypeList {
		incidentTypeListMap[incidentTypeList[i].Name] = incidentTypeList[i].ID
	}

	incident := pgsql.NewIncidentRepository(newServer)

	incidentList, err := incident.List()
	if err != nil {
		log.Error("incident.List error %v\n", err)
		return
	}

	if len(incidentList) == 0 {
		log.Error("incident.List no list")
		return
	}

	incidentMap := make(map[string]model.IncidentDB)
	for i := range incidentList {
		incidentMap[incidentList[i].Identifier] = incidentList[i]
	}

	for i := range revisionList.Incidents {
		rvision := revisionList.Incidents[i]

		if _, ok := userListMap[rvision.Owner]; !ok {
			userListMap[rvision.Owner] = userList[0].ID
		}

		if _, ok := userListMap[rvision.Responsible]; !ok {
			userListMap[rvision.Responsible] = userList[0].ID
		}

		if _, ok := categoryListMap[rvision.IncidentType]; !ok {
			categoryListMap[rvision.IncidentType] = categoryList[0].ID
		}

		if _, ok := priorityListMap[rvision.Level]; !ok {
			priorityListMap[rvision.Level] = priorityList[0].ID
		}

		if _, ok := statusListMap[rvision.Status]; !ok {
			statusListMap[rvision.Status] = statusList[0].ID
		}

		if _, ok := incidentTypeListMap[rvision.Category]; !ok {
			incidentTypeListMap[rvision.Category] = incidentTypeList[0].ID
		}

		if _, ok := incidentMap[rvision.Identifier]; ok {
			entityTry := false
			if rvision.Description != incidentMap[rvision.Identifier].Description {
				entityTry = true
			}

			if entityTry {
				_, err = incident.Create(model.IncidentDB{
					UserID:           userListMap[rvision.Owner],
					ResponsibleID:    userListMap[rvision.Responsible],
					OrganizationID:   organizationList[0].ID,
					NetworkSegmentID: networkList[0].ID,
					SziID:            szisList[0].ID,
					CategoryID:       categoryListMap[rvision.IncidentType],
					PriorityID:       priorityListMap[rvision.Level],
					StatusID:         statusListMap[rvision.Status],
					TypeID:           incidentTypeListMap[rvision.Category],
					Identifier:       rvision.Identifier,
					DetectName:       "-",
					Description:      rvision.Description,
					SavzResult:       "-",
					Path:             "-",
					CreatedDate:      time.Now(),
					CustomFields:     nil,
					PermissionLevel:  1,
					Archived:         false,
				})
				if err != nil {
					log.Error("incident create old %v\n", err)
					return
				}

				err = incident.Archive(incidentMap[rvision.Identifier].ID)
			}
		} else {
			_, err = incident.Create(model.IncidentDB{
				UserID:           userListMap[rvision.Owner],
				ResponsibleID:    userListMap[rvision.Responsible],
				OrganizationID:   organizationList[0].ID,
				NetworkSegmentID: networkList[0].ID,
				SziID:            szisList[0].ID,
				CategoryID:       categoryListMap[rvision.IncidentType],
				PriorityID:       priorityListMap[rvision.Level],
				StatusID:         statusListMap[rvision.Status],
				TypeID:           incidentTypeListMap[rvision.Category],
				Identifier:       rvision.Identifier,
				DetectName:       "-",
				Description:      rvision.Description,
				SavzResult:       "-",
				Path:             "-",
				CreatedDate:      time.Now(),
				CustomFields:     nil,
				PermissionLevel:  1,
				Archived:         false,
			})
			if err != nil {
				log.Error("incident create new %v\n", err)
				return
			}
		}
	}
}
