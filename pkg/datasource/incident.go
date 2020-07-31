package datasource

import (
	"github.com/olegpolukhin/rvision-irp/domain/model"
	udb "upper.io/db.v3"
)

func (pg *Postgres) IncidentCreate(incident model.IncidentDB) (newID string, err error) {
	inserter := pg.tx.InsertInto("incidents").Values(incident)
	newID, err = executeInsertAndGetNewID(inserter)
	return newID, PanicOnError(err)
}

func (pg *Postgres) IncidentList(incident *[]model.IncidentDB) error {
	selector := pg.tx.SelectFrom("incidents")
	err := selector.Where(udb.And(
		udb.Cond{"archived": false},
	)).All(incident)
	return PanicOnError(err)
}

func (pg *Postgres) IncidentGet(id string, incident *model.IncidentDB) error {
	selector := pg.tx.SelectFrom("incidents")
	err := selector.Where(udb.And(
		udb.Cond{"incidents.id": id},
		udb.Cond{"archived": false},
	)).One(incident)
	return PanicOnError(err)
}

func (pg *Postgres) IncidentArchive(id string) error {
	_, err := pg.tx.Update("incidents").
		Set("archived=true").
		Where(udb.And(
			udb.Cond{"id": id},
		)).Exec()
	return PanicOnError(err)
}
