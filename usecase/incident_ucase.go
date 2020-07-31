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

// GetRvisionIncidentList returns all incidents api revision you can read.
func (r *IncidentUsecase) RvisionIncidentList() (list model.IncidentList, err error) {
	if err := r.app.Post("/get_incidents", nil, &list); err != nil {
		return list, err
	}
	return list, nil
}

// RevisionIncident returns a incident api revision which has specified id.
func (r *IncidentUsecase) RevisionIncident(id string) (incident model.IncidentList, err error) {
	values := map[string]string{"id": id}
	err = r.app.Post("/get_incidents", values, &incident)
	return
}
