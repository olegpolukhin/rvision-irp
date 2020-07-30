package usecase

import (
	"github.com/olegpolukhin/rvision-irp/cmd/server"
	"github.com/olegpolukhin/rvision-irp/domain/model"
	"github.com/olegpolukhin/rvision-irp/pkg/client"
)

type IncidentUsecase struct {
	app *client.RevisionApp
}

func NewIncidentUsecase(app *server.App) (helper *IncidentUsecase) {
	return &IncidentUsecase{app.RevisionClient}
}

// GetIncidentList returns all incidents you can read.
func (r *IncidentUsecase) GetIncidentList() (list model.IncidentList, err error) {
	if err := r.app.Post("/get_incidents", nil, &list); err != nil {
		return list, err
	}
	return list, nil
}

// GetIncident returns a incident which has specified id.
func (r *IncidentUsecase) GetIncident(id string) (incident model.IncidentList, err error) {
	values := map[string]string{"id": id}
	err = r.app.Post("/get_incidents", values, &incident)
	return
}
