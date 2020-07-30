package datasource

import (
	"github.com/olegpolukhin/rvision-irp/domain/model"
	udb "upper.io/db.v3"
)

func (pg *Postgres) IncidentGetModel(id string, incident *model.IncidentDB) error {
	selector := pg.tx.SelectFrom("incidents")
	err := selector.Where(udb.And(
		udb.Cond{"incidents.id": id},
		udb.Cond{"archived": false},
	)).One(incident)
	return PanicOnError(err)
}
