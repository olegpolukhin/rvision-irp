package pgsql

import (
	"fmt"
	"github.com/olegpolukhin/rvision-irp/cmd/server"
	"github.com/olegpolukhin/rvision-irp/domain/model"
	"github.com/olegpolukhin/rvision-irp/pkg/datasource"
)

type IncidentRepository struct {
	db *datasource.Postgres
}

func NewIncidentRepository(app *server.App) (helper *IncidentRepository) {
	return &IncidentRepository{app.Postgres}
}

func (r *IncidentRepository) Get(id string) (incident model.IncidentDB, err error) {
	if err := r.db.Connect(); err != nil {
		return model.IncidentDB{}, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.IncidentGetModel(id, &incident); err != nil {
		return model.IncidentDB{}, fmt.Errorf("IncidentGetModel error %s", err.Error())
	}

	if len(incident.ID) == 0 {
		return model.IncidentDB{}, fmt.Errorf("IncidentGetModel error. No Items")
	}

	return
}

func (r *IncidentRepository) Create(id string) (incident model.IncidentDB, err error) {
	if err := r.db.Connect(); err != nil {
		return model.IncidentDB{}, fmt.Errorf("connect to DB error %s", err.Error())
	}

	defer r.db.Close(&err)

	if err := r.db.IncidentGetModel(id, &incident); err != nil {
		return model.IncidentDB{}, fmt.Errorf("IncidentGetModel error %s", err.Error())
	}

	if len(incident.ID) == 0 {
		return model.IncidentDB{}, fmt.Errorf("IncidentGetModel error. No Items")
	}

	return
}
